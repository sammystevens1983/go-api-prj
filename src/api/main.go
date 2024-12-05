package main

/*
#cgo CFLAGS: -I/workspaces/go-api-demo/src/c_sqr
#cgo LDFLAGS: -L/workspaces/go-api-demo/src/c_sqr -lc_sqr_lib
#include "/workspaces/go-api-demo/src/c_sqr/mylib.h"
*/
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
    var loadedString *[]byte

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Print Hello, World!")
		fmt.Println("2. Call C Library to Square a Number")
		fmt.Println("3. Simulate Receiving and Reconstructing a Large File")
		fmt.Println("4. Load and Modify String in Memory")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil || (choice < 1 || choice > 5) {
			fmt.Println("Invalid choice. Please enter a valid option.")
			continue
		}

		switch choice {
		case 1:
			fmt.Println("Hello, World!")
		case 2:
			fmt.Print("Enter a number to square: ")
			numInput, _ := reader.ReadString('\n')
			num, err := strconv.Atoi(strings.TrimSpace(numInput))
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer.")
				continue
			}

			// Call the C function
			result := C.square(C.int(num))
			fmt.Printf("The square of %d is %d.\n", num, int(result))
		case 3:
			simulateFileStreaming()
		case 4:
            loadedString = loadAndModifyString(loadedString)
		case 5:
			fmt.Println("Exiting...")
			return
		}
	}
}

func simulateFileStreaming() {
	const chunkSize = 1024 // Simulated chunk size in bytes
	const totalSize = 10 * 1024 * 1024 // Simulated total file size: 10 MB
	filePath := "received_large_file.txt"

	fmt.Println("Simulating receiving a large file...")

	// Open the file for writing
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	receivedSize := 0
	for receivedSize < totalSize {
		// Simulate receiving a chunk
		chunk := strings.Repeat("A", chunkSize)
		_, err := file.WriteString(chunk)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
		receivedSize += chunkSize
		fmt.Printf("\rReceived %d / %d bytes", receivedSize, totalSize)
	}

	fmt.Println("\nFile reconstructed successfully!")
	fmt.Printf("File saved to: %s\n", filePath)
}

func loadAndModifyString(currentBytes *[]byte) *[]byte {
	reader := bufio.NewReader(os.Stdin)

	if currentBytes == nil {
		// Load the string into memory as a byte slice
		str := "BascomHunter"
		bytes := []byte(str)
		currentBytes = &bytes
	}

	fmt.Printf("Current string in memory: %s\n", string(*currentBytes))

	// Display hex values
	fmt.Print("Hex values in memory: ")
	for _, b := range *currentBytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	// Prompt the user to modify a byte
	fmt.Print("Enter index to modify (0-based): ")
	indexInput, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(indexInput))
	if err != nil || index < 0 || index >= len(*currentBytes) {
		fmt.Println("Invalid index.")
		return currentBytes
	}

	fmt.Print("Enter new hex value (e.g., 41 for 'A'): ")
	hexInput, _ := reader.ReadString('\n')
	newByte, err := strconv.ParseUint(strings.TrimSpace(hexInput), 16, 8)
	if err != nil {
		fmt.Println("Invalid hex value.")
		return currentBytes
	}

	// Modify the byte slice
	(*currentBytes)[index] = byte(newByte)

	// Display the modified hex values
	fmt.Print("Modified hex values in memory: ")
	for _, b := range *currentBytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	fmt.Printf("Modified string in memory: %s\n", string(*currentBytes))

	return currentBytes
}

