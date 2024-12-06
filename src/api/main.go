package main

/*
#cgo CFLAGS: -I/workspaces/go-api-demo/src/c_sqr
#cgo LDFLAGS: -L/workspaces/go-api-demo/src/c_sqr -lc_sqr_lib
#include "/workspaces/go-api-demo/src/c_sqr/mylib.h"
*/
import "C"

import (
	"bufio"                // Used for reading input from the user
	"fmt"                  // Provides formatted I/O functions
	"go-api-prj/src/utils" // Imports utility functions for file streaming and string modification
	"os"                   // Provides operating system functions like file handling
	"strconv"              // Provides functions to convert strings to integers
	"strings"              // Provides string manipulation utilities
)

func main() {
	// `loadedString` is a pointer to a byte slice that stores the string for the "Load and Modify String in Memory" option.
	var loadedString *[]byte

	// `reader` is used to read user input from the console.
	reader := bufio.NewReader(os.Stdin)

	// Main program loop that displays a menu and processes user input.
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
			// `SimulateFileStreaming` takes the chunk size, total size, and file path as arguments.
			utils.SimulateFileStreaming(1024, 10*1024*1024, "received_large_file.txt")
		case 4:
			// Option 4: Load a string into memory and allow modifications
			// This function uses the `loadedString` variable to track the string state.
			loadedString = utils.LoadAndModifyString(loadedString)
		case 5:
			// Option 5: Exit the program
			fmt.Println("Exiting...")
			return
		}
	}
}

// callCSquareFunction handles user input, calls the C library's `square` function, and displays the result.
func callCSquareFunction() {
	// `reader` is used to read the number input from the user.
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
	result := C.square(C.int(num)) // Convert the Go integer to a C integer
	fmt.Printf("The square of %d is %d.\n", num, int(result))
}
