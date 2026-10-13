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

// Any function (including anonymous functions) can be run as a goroutine
// by prefixing the function call with the `go` keyword.

// PrintNumbers prints numbers from 0 to 4 with a delay between each.
func PrintNumbers(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(100 * time.Millisecond) // Simulate work
	}
}

// --- 3. THE PROBLEM WITH TIME.SLEEP FOR SYNCHRONIZATION ---

// In the original example, time.Sleep was used to wait for goroutines to finish.
// This is problematic because:
//   a) We don't know exactly how long goroutines will take
//   b) We might wait too long (wasting time) or too short (cutting off work)
//   c) It doesn't scale - what if we have 1000 goroutines?

// Let's demonstrate this issue first:
func basicGoroutineExample() {
	fmt.Println("=== Basic Goroutine Example (with time.Sleep) ===")
	for i := 0; i < 3; i++ {
		go PrintNumbers(i)
	}
	// We guess that 2 seconds is enough... but what if it's not?
	time.Sleep(2 * time.Second)
	fmt.Println("Main function continuing...")
	fmt.Println()
}

// --- 4. PROPER SYNCHRONIZATION WITH WAITGROUP ---

// sync.WaitGroup is the standard way to wait for goroutines to finish.
// It maintains a counter and blocks until the counter reaches zero.

// WorkerWithWaitGroup performs work and signals completion via WaitGroup.
func WorkerWithWaitGroup(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the counter when the function returns
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d: task %d\n", id, i)
		time.Sleep(150 * time.Millisecond)
	}
}

// WaitGroupExample demonstrates proper synchronization.
func WaitGroupExample() {
	fmt.Println("=== WaitGroup Example ===")
	var wg sync.WaitGroup

	// We want to wait for 3 workers
	wg.Add(3)

	// Launch 3 workers as goroutines
	for i := 0; i < 3; i++ {
		go WorkerWithWaitGroup(i, &wg)
	}

	// Block until all workers signal completion
	wg.Wait()
	fmt.Println("All workers completed!")
	fmt.Println()
}

// --- 5. RACE CONDITIONS ---

// When multiple goroutines access shared memory without synchronization,
// we get race conditions which lead to unpredictable results.

// UnsafeCounter is NOT safe for concurrent access.
type UnsafeCounter struct {
	value int
}

// Increment increases the counter (NOT thread-safe).
func (c *UnsafeCounter) Increment() {
	c.value++ // This is not atomic!
}

// Value returns the current counter value.
func (c *UnsafeCounter) Value() int {
	return c.value
}

// RaceConditionExample demonstrates the problem.
func RaceConditionExample() {
	fmt.Println("=== Race Condition Example ===")
	counter := UnsafeCounter{value: 0}
	const numGoroutines = 100
	const incrementsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}(i)
	}

	wg.Wait()
	expected := numGoroutines * incrementsPerGoroutine
	fmt.Printf("Expected value: %d, Actual value: %d\n", expected, counter.value)
	if counter.value != expected {
		fmt.Println("❌ Race condition detected! Value is incorrect.")
	} else {
		fmt.Println("✅ By chance, value is correct (race conditions are non-deterministic).")
	}
	fmt.Println()
}

// --- 6. MUTEX FOR MUTUAL EXCLUSION ---

// sync.Mutex provides mutual exclusion - only one goroutine can hold the lock at a time.

// SafeCounter IS safe for concurrent access.
type SafeCounter struct {
	value int
	mu    sync.Mutex
}

// Increment increases the counter (thread-safe).
func (c *SafeCounter) Increment() {
	c.mu.Lock()   // Acquire lock
	defer c.mu.Unlock() // Release lock when function returns
	c.value++
}

// Value returns the current counter value (thread-safe).
func (c *SafeCounter) Value() int {
	c.mu.Lock()   // Acquire lock
	defer c.mu.Unlock() // Release lock
	return c.value
}

// MutexExample demonstrates proper synchronization.
func MutexExample() {
	fmt.Println("=== Mutex Example ===")
	counter := SafeCounter{value: 0}
	const numGoroutines = 10
	const incrementsPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}(i)
	}

	wg.Wait()
	expected := numGoroutines * incrementsPerGoroutine
	fmt.Printf("Expected value: %d, Actual value: %d\n", expected, counter.Value())
	if counter.Value() == expected {
		fmt.Println("✅ Mutex protected the counter - value is correct!")
	} else {
		fmt.Println("❌ Still incorrect - check your mutex usage.")
	}
	fmt.Println()
}

// --- 7. CHANNELS FOR COMMUNICATION ---

// Channels are Go's mechanism for communicating between goroutines.
// They allow goroutines to send and receive values with built-in synchronization.

// sender sends values to a channel and then closes it.
func sender(ch chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that sender is done
	for i := 0; i < 3; i++ {
		ch <- i // Send value to channel (blocks if buffer is full)
		fmt.Printf("Sender %d sent: %d\n", id, i)
		time.Sleep(time.Duration(id*100) * time.Millisecond)
	}
	close(ch) // Close channel to indicate no more values will be sent
	fmt.Printf("Sender %d done\n", id)
}

// receiver receives values from a channel until it's closed.
func receiver(ch <-chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range ch { // Receives until channel is closed
		fmt.Printf("Receiver %d got: %d\n", id, val)
		time.Sleep(50 * time.Millisecond)
	}
	fmt.Printf("Receiver %d done\n", id)
}

// ChannelExample demonstrates unbuffered channel communication.
func ChannelExample() {
	fmt.Println("=== Channel Communication Example ===")
	// Unbuffered channel - sender blocks until receiver is ready
	messages := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2) // Wait for sender and receiver

	go sender(messages, 1, &wg)
	go receiver(messages, 1, &wg)

	wg.Wait()
	fmt.Println("Channel communication completed")
	fmt.Println()
}

// --- 8. BUFFERED CHANNELS ---

// Buffered channels have a capacity and allow sends to succeed
// until the buffer is full, reducing blocking.

// BufferedChannelExample shows how buffered channels work.
func BufferedChannelExample() {
	fmt.Println("=== Buffered Channel Example ===")
	// Buffered channel with capacity 2
	messages := make(chan string, 2)

	// We can send up to 2 values without blocking
	messages <- "hello"
	messages <- "world"
	fmt.Println("Sent two messages to buffered channel")

	// Now we receive them
	fmt.Printf("Received: %s\n", <-messages)
	fmt.Printf("Received: %s\n", <-messages)

	// Try to send a third - this will block because buffer is full
	go func() {
		time.Sleep(100 * time.Millisecond)
		messages <- "third"
		fmt.Println("Sent third message")
	}()

	// Give the goroutine time to send, then receive
	time.Sleep(200 * time.Millisecond)
	fmt.Printf("Received: %s\n", <-messages)
	fmt.Println()
}

// --- 9. SELECT STATEMENT ---

// select lets a goroutine wait on multiple channel operations.
// It chooses randomly if multiple cases are ready.

// SelectExample demonstrates the select statement.
func SelectExample() {
	fmt.Println("=== Select Statement Example ===")
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		channel1 <- "from channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		channel2 <- "from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-channel1:
			fmt.Println("Received:", msg1)
		case msg2 := <-channel2:
			fmt.Println("Received:", msg2)
		}
	}
	fmt.Println()
}

// --- 10. WORKER POOL PATTERN ---

// A common pattern for limiting concurrent goroutines while processing a queue of jobs.

// Job represents a unit of work.
type Job struct {
	id     int
	random int
}

// Result represents the outcome of processing a job.
type Result struct {
	job     Job
	sum     int
	workerID int
}

// worker processes jobs from the jobs channel and sends results to results channel.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate work: sum of digits
		sum := 0
		n := job.random
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		results <- Result{job: job, sum: sum, workerID: id}
		time.Sleep(50 * time.Millisecond) // Simulate processing time
	}
}

// WorkerPoolExample demonstrates limiting concurrent workers.
func WorkerPoolExample() {
	fmt.Println("=== Worker Pool Example ===")
	const numJobs = 20
	const numWorkers = 3

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{id: j, random: j * 17} // Some pseudo-random work
	}
	close(jobs) // No more jobs will be sent

	// Wait for workers to finish
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Worker %d processed job %d (random=%d), sum=%d\n",
			result.workerID, result.job.id, result.job.random, result.sum)
	}
	fmt.Println("All jobs processed!")
	fmt.Println()
}

// --- 11. CONTEXT FOR CANCELLATION ---

// context.Context allows us to propagate cancellation signals across goroutine boundaries.

// We'll skip context for this introductory example to keep it focused,
// but in real applications, you'd use context for timeouts and cancellation.

// --- 12. BEST PRACTICES SUMMARY ---

func main() {
	fmt.Println("=== Go Goroutines Learning Guide ===\n")

	// Demonstrate each concept
	basicGoroutineExample()
	WaitGroupExample()
	RaceConditionExample()
	MutexExample()
	ChannelExample()
	BufferedChannelExample()
	SelectExample()
	WorkerPoolExample()

	fmt.Println("=== Best Practices Summary ===")
	fmt.Println("✓ Use goroutines for concurrent execution")
	fmt.Println("✓ Never use time.Sleep for synchronization - use WaitGroup or channels")
	fmt.Println("✓ Protect shared memory with mutexes or use channels to communicate")
	fmt.Println("✓ Prefer communicating over sharing memory (Go's motto)")
	fmt.Println("✓ Use buffered channels to reduce blocking when appropriate")
	fmt.Println("✓ Use select for multiplexing channel operations")
	fmt.Println("✓ Consider worker pool pattern for limiting concurrency")
	fmt.Println("✓ Use context.Context for cancellation and timeouts in larger applications")
	fmt.Println("✓ Always check for race conditions with: go run -race")
	fmt.Println()
	fmt.Println("Try running with race detector: go run -race goroutines.go")
}