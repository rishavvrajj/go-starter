package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Example 1: Creating and Writing Files
func createAndWriteFiles() {
	fmt.Println("=== Creating and Writing Files ===")

	// Create a file and write to it
	file, err := os.Create("write_example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write bytes
	byteSlice := []byte("Hello Go File I/O!\n")
	bytesWritten, err := file.Write(byteSlice)
	if err != nil {
		fmt.Printf("Error writing to file: %v\n", err)
		return
	}
	fmt.Printf("Wrote %d bytes\n", bytesWritten)

	// Write string
	stringWritten, err := file.WriteString("This is written with WriteString\n")
	if err != nil {
		fmt.Printf("Error writing string: %v\n", err)
		return
	}
	fmt.Printf("Wrote %d bytes via WriteString\n", stringWritten)

	// Sync to ensure data is flushed to disk
	file.Sync()
	fmt.Println("File written and synced successfully\n")
}

// Example 2: Reading Files
func readFiles() {
	fmt.Println("=== Reading Files ===")

	// Read entire file at once
	data, err := os.ReadFile("write_example.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Printf("File contents (%d bytes):\n%s\n", len(data), string(data))

	// Open file and read line by line
	file, err := os.Open("write_example.txt")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Reading line by line:")
	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Line %d: %s\n", lineNum, line)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println()
}

// Example 3: Buffered Writing
func bufferedWriting() {
	fmt.Println("=== Buffered Writing ===")

	file, err := os.Create("buffered_example.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Create buffered writer
	writer := bufio.NewWriter(file)

	// Write data to buffer (not yet to disk)
	for i := 0; i < 5; i++ {
		line := fmt.Sprintf("Buffered line %d\n", i+1)
		_, err := writer.WriteString(line)
		if err != nil {
			fmt.Printf("Error writing to buffer: %v\n", err)
			return
		}
	}

	// Flush buffer to disk
	err = writer.Flush()
	if err != nil {
		fmt.Printf("Error flushing buffer: %v\n", err)
		return
	}
	fmt.Println("Buffered write completed and flushed\n")
}

// Example 4: Copying Files
func copyFile(source, destination string) error {
	fmt.Printf("=== Copying %s to %s ===\n", source, destination)

	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer destFile.Close()

	// Copy using io.Copy (handles buffer management)
	bytesCopied, err := io.Copy(destFile, sourceFile)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	fmt.Printf("Successfully copied %d bytes\n", bytesCopied)
	return nil
}

// Example 5: Working with Directories
func directoryOperations() {
	fmt.Println("=== Directory Operations ===")

	// Create a directory for demo
	dirName := "demo_directory"
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}
	defer os.RemoveAll(dirName) // Clean up at end

	// Create some files in the directory
	for i := 1; i <= 3; i++ {
		fileName := filepath.Join(dirName, fmt.Sprintf("file%d.txt", i))
		err := os.WriteFile(fileName, []byte(fmt.Sprintf("Content of file %d\n", i)), 0644)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", fileName, err)
			return
		}
	}
	fmt.Printf("Created directory '%s' with 3 files\n", dirName)

	// Read directory contents
	files, err := os.ReadDir(dirName)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
		return
	}
	fmt.Println("Files in directory:")
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			fmt.Printf("Error getting info for %s: %v\n", file.Name(), err)
			continue
		}
		fmt.Printf("  %s (%d bytes, IsDir: %v)\n", file.Name(), info.Size(), file.IsDir())
	}
	fmt.Println()
}

// Example 6: File Information
func fileInfo() {
	fmt.Println("=== File Information ===")

	// Get file info
	fileInfo, err := os.Stat("write_example.txt")
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}

	fmt.Printf("File: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Permissions: %o\n", fileInfo.Mode().Perm())
	fmt.Printf("Is Directory: %v\n", fileInfo.IsDir())
	fmt.Printf("Modification Time: %v\n", fileInfo.ModTime())
	fmt.Printf("Access Time: (not available on all systems)\n")
	fmt.Println()
}

// Example 7: Temporary Files
func temporaryFiles() {
	fmt.Println("=== Temporary Files ===")

	// Create temporary file
	tempFile, err := os.CreateTemp("", "example_*.txt")
	if err != nil {
		fmt.Printf("Error creating temp file: %v\n", err)
		return
	}
	tempFileName := tempFile.Name()
	fmt.Printf("Created temp file: %s\n", tempFileName)

	// Write to temp file
	tempFile.Write([]byte("Temporary data for demonstration\n"))
	tempFile.Close()

	// Read it back
	data, err := os.ReadFile(tempFileName)
	if err != nil {
		fmt.Printf("Error reading temp file: %v\n", err)
		return
	}
	fmt.Printf("Temp file contents: %s", string(data))

	// Clean up (temp files are usually cleaned by OS, but good practice)
	err = os.Remove(tempFileName)
	if err != nil {
		fmt.Printf("Error removing temp file: %v\n", err)
		return
	}
	fmt.Println("Temp file cleaned up\n")
}

// Example 8: File Existence and Permissions
func filePermissions() {
	fmt.Println("=== File Permissions ===")

	// Check if file exists
	if _, err := os.Stat("write_example.txt"); os.IsNotExist(err) {
		fmt.Println("write_example.txt does not exist")
	} else if err != nil {
		fmt.Printf("Error checking file: %v\n", err)
	} else {
		fmt.Println("write_example.txt exists")
	}

	// Change file permissions
	err := os.Chmod("write_example.txt", 0644) // rw-r--r--
	if err != nil {
		fmt.Printf("Error changing permissions: %v\n", err)
		return
	}
	fmt.Println("Changed permissions of write_example.txt to 0644 (rw-r--r--)")

	// Check current permissions
	info, err := os.Stat("write_example.txt")
	if err != nil {
		fmt.Printf("Error getting file info: %v\n", err)
		return
	}
	fmt.Printf("Current permissions: %o\n", info.Mode().Perm())
	fmt.Println()
}

// Example 9: Renaming and Moving Files
func renameMoveFiles() {
	fmt.Println("=== Renaming and Moving Files ===")

	// Create a file to rename
	err := os.WriteFile("to_rename.txt", []byte("This file will be renamed\n"), 0644)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer os.Remove("renamed.txt") // Cleanup

	// Rename/move file
	err = os.Rename("to_rename.txt", "renamed.txt")
	if err != nil {
		fmt.Printf("Error renaming file: %v\n", err)
		return
	}
	fmt.Println("Renamed 'to_rename.txt' to 'renamed.txt'")

	// Verify the move
	_, err = os.Stat("renamed.txt")
	if err == nil {
		fmt.Println("'renamed.txt' exists after rename")
	} else {
		fmt.Printf("Error: 'renamed.txt' not found: %v\n", err)
	}
	fmt.Println()
}

// Example 10: Reading CSV-like Data (Simple Parsing)
func parseSimpleData() {
	fmt.Println("=== Parsing Simple Data ===")

	// Create sample data file
	data := []string{
		"name,age,city",
		"Alice,25,New York",
		"Bob,30,London",
		"Charlie,35,Tokyo",
	}
	err := os.WriteFile("sample.csv", []byte(data[0]+"\n"+data[1]+"\n"+data[2]+"\n"+data[3]+"\n"), 0644)
	if err != nil {
		fmt.Printf("Error creating sample data: %v\n", err)
		return
	}
	defer os.Remove("sample.csv")

	// Read and parse
	file, err := os.Open("sample.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	fmt.Println("Parsing CSV-like data:")
	scanner := bufio.NewScanner(file)
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineNum++
		if lineNum == 1 {
			fmt.Printf("Header: %s\n", line)
			continue
		}
		// Simple split (in real code, use encoding/csv for proper CSV handling)
		fields := []string{}
		start := 0
		for i, ch := range line {
			if ch == ',' {
				fields = append(fields, line[start:i])
				start = i + 1
			}
		}
		fields = append(fields, line[start:]) // Last field

		if len(fields) >= 3 {
			fmt.Printf("Record %d: Name=%s, Age=%s, City=%s\n", lineNum-1, fields[0], fields[1], fields[2])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println()
}

func main() {
	createAndWriteFiles()
	readFiles()
	bufferedWriting()

	// Copy file example
	err := copyFile("write_example.txt", "copied_example.txt")
	if err != nil {
		fmt.Printf("Copy failed: %v\n", err)
	}
	defer os.Remove("copied_example.txt") // Cleanup

	directoryOperations()
	fileInfo()
	temporaryFiles()
	filePermissions()
	renameMoveFiles()
	parseSimpleData()

	// Final cleanup
	os.Remove("write_example.txt")
	os.Remove("buffered_example.txt")

	fmt.Println("=== All examples completed ===")
}