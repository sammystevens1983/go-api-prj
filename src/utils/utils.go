package utils

import (
	"bufio"
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
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Use a buffered writer for efficient file writing
	bufferedWriter := bufio.NewWriter(file)

	// Pre-generate a single chunk of data
	chunk := strings.Repeat("A", chunkSize)

	receivedSize := 0 // Tracks the total bytes received and written
	for receivedSize < totalSize {
		remaining := totalSize - receivedSize
		// Write the remaining chunk size or the full chunk, whichever is smaller
		toWrite := chunk
		if remaining < chunkSize {
			toWrite = strings.Repeat("A", remaining)
		}

		_, err := bufferedWriter.WriteString(toWrite)
		if err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			return
		}
		receivedSize += len(toWrite)

		// Print progress less frequently (e.g., every 5%)
		if receivedSize%int(float64(totalSize)*0.05) == 0 || receivedSize == totalSize {
			fmt.Printf("\rReceived %d / %d bytes", receivedSize, totalSize)
		}
	}

	// Flush any remaining data in the buffer
	err = bufferedWriter.Flush()
	if err != nil {
		fmt.Printf("Error flushing buffer: %v\n", err)
		return
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
