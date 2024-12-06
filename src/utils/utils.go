package utils

import (
	"fmt"     // Provides formatted I/O functions
	"os"      // Provides operating system functions for file handling
	"strings" // Provides utilities for manipulating and repeating strings
)

// SimulateFileStreaming simulates the process of receiving a large file in chunks
// and reconstructing it by writing the chunks to a file.
//
// Parameters:
// - chunkSize: The size of each chunk (in bytes) to be written to the file.
// - totalSize: The total size of the file (in bytes) to be reconstructed.
// - filePath: The path where the reconstructed file will be saved.
//
// This function writes data to the file in chunks, simulating a streaming process.
// It provides a real-time progress indicator to the user.
func SimulateFileStreaming(chunkSize, totalSize int, filePath string) {
	fmt.Println("Simulating receiving a large file...")

	// Open the file for writing, creating it if it doesn't exist
	file, err := os.Create(filePath)
	if err != nil {
		// Handle file creation errors
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	// Ensure the file is closed when the function exits
	defer file.Close()

	receivedSize := 0 // Tracks the total bytes received and written
	for receivedSize < totalSize {
		// Generate a chunk of data to simulate receiving a file
		chunk := strings.Repeat("A", chunkSize) // Repeat 'A' to fill the chunk
		_, err := file.WriteString(chunk)       // Write the chunk to the file
		if err != nil {
			// Handle file writing errors
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
		receivedSize += chunkSize // Update the total bytes received
		// Display progress to the user
		fmt.Printf("\rReceived %d / %d bytes", receivedSize, totalSize)
	}

	fmt.Println("\nFile reconstructed successfully!")
	fmt.Printf("File saved to: %s\n", filePath)
}

// LoadAndModifyString loads a string into memory (if not already loaded) and allows
// modification of its contents directly in memory.
//
// Parameters:
// - currentBytes: A pointer to a byte slice that stores the string in memory.
//
// Returns:
// - A pointer to the modified byte slice.
//
// This function demonstrates how to manipulate a string in memory, displaying
// its hexadecimal representation before and after modification.
func LoadAndModifyString(currentBytes *[]byte) *[]byte {
	if currentBytes == nil {
		// Initialize the string in memory if it's not already loaded
		str := "BascomHunter"    // Default string to load
		bytes := []byte(str)     // Convert the string to a byte slice
		currentBytes = &bytes // Store the byte slice pointer
	}

	// Display the current string in memory
	fmt.Printf("Current string in memory: %s\n", string(*currentBytes))

	// Display the hexadecimal representation of the string before modification
	fmt.Print("Pre Hex values in memory: ")
	for _, b := range *currentBytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	// Modify a specific byte in the string
	index := 6              // Index to modify (0-based)
	newByte := byte('Z')    // New byte value to replace at the specified index
	(*currentBytes)[index] = newByte // Modify the byte slice directly

	// Display the modified string
	fmt.Printf("Modified string in memory: %s\n", string(*currentBytes))

	// Display the hexadecimal representation of the string after modification
	fmt.Print("Post Hex values in memory: ")
	for _, b := range *currentBytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	return currentBytes // Return the pointer to the modified byte slice
}
