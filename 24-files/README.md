# Files

## What this topic is
File I/O in Go is handled through the standard library packages like `os`, `bufio`, `io`, and `path/filepath`. These packages provide functions for creating, reading, writing, and manipulating files and directories.

## When to use it
- When you need to read from or write to files
- Working with configuration files
- Logging to files
- Processing data from files (CSV, JSON, etc.)
- Creating temporary files
- Working with directories and file paths
- File compression and archiving
- Implementing file upload/download functionality

## Minimal syntax example
```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

// Writing to a file
func writeFile() {
    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Printf("Error creating file: %v\n", err)
        return
    }
    defer file.Close()
    
    // Write bytes
    file.Write([]byte("Hello, File!\n"))
    
    // Write string
    file.WriteString("Another line\n")
    
    // Flush to disk
    file.Sync()
}

// Reading from a file
func readFile() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }
    fmt.Println(string(data))
}

// Reading line by line
func readFileLineByLine() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
    
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}

// Copying a file
func copyFile(src, dst string) error {
    source, err := os.Open(src)
    if err != nil {
        return err
    }
    defer source.Close()
    
    destination, err := os.Create(dst)
    if err != nil {
        return err
    }
    defer destination.Close()
    
    _, err = io.Copy(destination, source)
    return err
}

func main() {
    writeFile()
    readFile()
    readFileLineByLine()
}
```

## Common mistakes
- Forgetting to close files (always use defer file.Close())
- Not checking for errors after file operations
- Trying to write to a nil file pointer
- Not understanding the difference between os.Create (truncates) and os.OpenFile (with flags)
- Forgetting to flush buffered writers
- Using os.ReadFile for very large files (loads entire file into memory)
- Not using bufio for efficient line-by-line reading
- Forgetting to handle different path separators on different OS (use path/filepath)
- Not checking if a file exists before trying to open it
- Confusing file permissions (Unix-style octal notation)
- Not cleaning up temporary files
- Using string concatenation for paths instead of filepath.Join

## Command to run it
```bash
go run files.go
```