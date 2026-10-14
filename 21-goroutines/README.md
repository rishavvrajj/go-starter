# Goroutines

## What this topic is
Goroutines are lightweight threads managed by the Go runtime. They enable concurrent execution of functions with minimal overhead (typically starting at just 2KB of stack space). Goroutines are the foundation of concurrency in Go.

## When to use it
- When you need to perform multiple tasks concurrently
- Performing I/O operations that can happen in parallel (like web requests, file reading)
- Background processing or periodic tasks
- When you want to avoid blocking the main thread
- Implementing worker pools or pipelines
- Any situation where you can benefit from concurrent execution

## Minimal syntax example
```go
package main

import (
    "fmt"
    "sync"
    "time"
)

// A simple worker function
func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Signal completion when function returns
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second) // Simulate work
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    // We want to wait for 3 workers
    wg.Add(3)
    
    // Launch 3 workers as goroutines
    for i := 1; i <= 3; i++ {
        go worker(i, &wg)
    }
    
    // Block until all workers signal completion
    wg.Wait()
    fmt.Println("All workers completed!")
}
```

## Common mistakes
- Forgetting to wait for goroutines to finish (main function may exit before they complete)
- Using time.Sleep for synchronization instead of proper synchronization primitives
- Not understanding that goroutine stacks grow and shrink as needed
- Forgetting to call wg.Add() before launching goroutines
- Not calling wg.Done() in each goroutine (causing wg.Wait() to block forever)
- Copying a WaitGroup after it has been used (should pass by pointer)
- Not realizing that goroutines share the same address space (need synchronization for shared data)
- Using goroutines for trivial tasks where the overhead outweighs the benefit
- Not handling panics in goroutines (they can crash the entire program if not recovered)

## Command to run it
```bash
go run goroutines.go
```