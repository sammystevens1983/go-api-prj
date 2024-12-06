package main

/*
#cgo CFLAGS: -I/workspaces/go-api-demo/src/c_sqr
#cgo LDFLAGS: -L/workspaces/go-api-demo/src/c_sqr -lc_sqr_lib
#include "/workspaces/go-api-demo/src/c_sqr/mylib.h"
*/
import "C"

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-api-prj/src/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// JSON handler function for writing data to a file
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON payload from the request body
	var payload DataPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON payload.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Write the payload to a file
	file, err := os.Create("data.json")
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create file: %v", err), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // Pretty-print JSON
	if err := encoder.Encode(payload); err != nil {
		http.Error(w, fmt.Sprintf("Failed to write JSON to file: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "JSON payload successfully written to data.json\n")
}

// DataPayload represents the structure of the JSON payload to be received
type DataPayload struct {
	Message string `json:"message"` // Example field
	Number  int    `json:"number"`  // Example field
}


// File upload handler with optional destination directory and filename
func fileUploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Only POST is allowed.", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form (limit to 10MB)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data.", http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form data.", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Retrieve optional destination directory and filename from form data
	destDir := r.FormValue("destDir")
	if destDir == "" {
		destDir = "." // Default to the current directory
	}

	destFilename := r.FormValue("destFilename")
	if destFilename == "" {
		destFilename = handler.Filename // Default to the source filename
	}

	// Ensure the destination directory exists
	err = os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating destination directory: %v", err), http.StatusInternalServerError)
		return
	}

	// Construct the full destination file path
	destFilePath := fmt.Sprintf("%s/%s", strings.TrimRight(destDir, "/"), destFilename)

	// Create the destination file
	destFile, err := os.Create(destFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error creating file on server: %v", err), http.StatusInternalServerError)
		return
	}
	defer destFile.Close()

	// Copy the file contents to the destination file
	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving file: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File successfully uploaded to: %s\n", destFilePath)
}

func main() {
    // Print the PID of the program at startup
	fmt.Printf("Program is running with PID: %d\n", os.Getpid())

    // Start the API server in a separate goroutine
	go func() {
		http.HandleFunc("/square", squareHandler)    // Existing square endpoint
		http.HandleFunc("/save-json", jsonHandler)  // Existing JSON endpoint
		http.HandleFunc("/upload-file", fileUploadHandler) // New file upload endpoint
		fmt.Println("Starting API server on port 5000...")
		if err := http.ListenAndServe(":5000", nil); err != nil {
			fmt.Printf("Error starting API server: %v\n", err)
			os.Exit(1)
		}
	}()

	// CLI menu functionality
	var loadedString *[]byte
	reader := bufio.NewReader(os.Stdin)

	for {
		// Display the menu options
		fmt.Println("\nMenu:")
		fmt.Println("1. Print Hello, World!")
		fmt.Println("2. Call C Library to Square a Number")
		fmt.Println("3. Simulate Receiving and Reconstructing a Large File")
		fmt.Println("4. Load and Modify String in Memory")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		// Read user input and trim any whitespace
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input)) // Convert input to an integer
		if err != nil || (choice < 1 || choice > 5) {
			// Handle invalid input
			fmt.Println("Invalid choice. Please enter a valid option.")
			continue
		}

		// Handle menu options based on user input
		switch choice {
		case 1:
			// Option 1: Print a simple greeting
			fmt.Println("Hello, World!")
		case 2:
			// Option 2: Call the C library function to square a number
			callCSquareFunction()
		case 3:
			// Option 3: Simulate receiving and reconstructing a large file
			utils.SimulateFileStreaming(1024, 100*1024*1024, "received_large_file.txt")
		case 4:
			// Option 4: Load a string into memory and allow modifications
			loadedString = utils.LoadAndModifyString(loadedString)
		case 5:
			// Option 5: Exit the program
			fmt.Println("Exiting...")
			return
		}
	}
}

// CLI handler for squaring a number using the C library
func callCSquareFunction() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number to square: ")

	// Read and parse the input number
	numInput, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(numInput))
	if err != nil {
		// Handle invalid input
		fmt.Println("Invalid input. Please enter an integer.")
		return
	}

	// Call the `square` function from the C library
	result := C.square(C.int(num))
	fmt.Printf("The square of %d is %d.\n", num, int(result))
}

// squareHandler handles HTTP requests for squaring a number.
func squareHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the number from the query parameters
	keys, ok := r.URL.Query()["number"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing 'number' query parameter", http.StatusBadRequest)
		return
	}

	// Convert the number to an integer
	num, err := strconv.Atoi(keys[0])
	if err != nil {
		http.Error(w, "Invalid 'number' query parameter. Must be an integer.", http.StatusBadRequest)
		return
	}

	// Call the C function to square the number
	result := C.square(C.int(num))

	// Respond with the squared result
	fmt.Fprintf(w, "The square of %d is %d\n", num, int(result))
}
