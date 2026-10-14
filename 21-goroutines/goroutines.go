package main

import (
	"fmt"
	"sync"
	"time"
)

// --- 1. WHAT ARE GOROUTINES? ---

// A goroutine is a lightweight thread managed by the Go runtime.
// Goroutines enable concurrent execution of functions.
// They are much cheaper than OS threads (typically ~2KB stack vs 1MB+).

// --- 2. BASIC GOROUTINE CREATION ---

// Any function can be run as a goroutine by prefixing the function call with the `go` keyword.

// printNumbers prints numbers from 0 to 4 with a delay between each.
func printNumbers(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// --- 3. THE PROBLEM: NOT WAITING FOR GOROUTINES ---

// If we don't wait for goroutines to finish, the main function might exit
// before the goroutines complete their work.

func demoWithoutWaiting() {
	fmt.Println("=== Demo: NOT Waiting for Goroutines ===")
	for i := 0; i < 3; i++ {
		go printNumbers(i)
	}
	// Main function continues immediately - goroutines may still be running
	fmt.Println("Main function continuing without waiting...")
	// Give goroutines a chance to finish (but this is unreliable!)
	time.Sleep(1 * time.Second)
	fmt.Println()
}

// --- 4. THE SOLUTION: SYNC.WAITGROUP ---

// sync.WaitGroup is the standard way to wait for goroutines to finish.
// It maintains a counter and blocks until the counter reaches zero.

// workerWithWaitGroup performs work and signals completion via WaitGroup.
func workerWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the function returns
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: task %d\n", id, i)
		time.Sleep(150 * time.Millisecond)
	}
}

// demoWithWaitGroup demonstrates proper synchronization.
func demoWithWaitGroup() {
	fmt.Println("=== Demo: Waiting with WaitGroup ===")
	var wg sync.WaitGroup

	// We want to wait for 3 workers
	wg.Add(3)

	// Launch 3 workers as goroutines
	for i := 0; i < 3; i++ {
		go workerWithWaitGroup(i, &wg)
	}

	// Block until all workers signal completion
	wg.Wait()
	fmt.Println("All workers completed!")
	fmt.Println()
}

// --- 5. KEY POINTS ABOUT WAITGROUP ---

// 1. Call wg.Add(n) BEFORE launching the goroutines to set the expected count
// 2. Each goroutine should call wg.Done() when it finishes (usually via defer)
// 3. Call wg.Wait() to block until all goroutines have finished
// 4. WaitGroup should NOT be copied after first use

// --- 6. COMPLETE EXAMPLE ---

func main() {
	fmt.Println("=== Go Goroutines and WaitGroup Learning Guide ===")

	// Demonstrate the problem
	demoWithoutWaiting()

	// Demonstrate the solution
	demoWithWaitGroup()

	fmt.Println("=== Key Takeaways ===")
	fmt.Println("✓ Goroutines are lightweight concurrent functions")
	fmt.Println("✓ Use 'go' keyword to launch a goroutine")
	fmt.Println("✓ Never rely on time.Sleep for synchronization")
	fmt.Println("✓ Use sync.WaitGroup to wait for goroutines to finish")
	fmt.Println("✓ Call wg.Add() before launching goroutines")
	fmt.Println("✓ Call wg.Done() in each goroutine when finished (use defer)")
	fmt.Println("✓ Call wg.Wait() to block until all goroutines complete")
	fmt.Println()
	fmt.Println("Try removing the wg.Wait() to see the difference!")
}