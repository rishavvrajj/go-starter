# Mutex and Synchronization

## What this topic is
Mutex (mutual exclusion) in Go is a synchronization primitive used to prevent race conditions when multiple goroutines access shared data concurrently. The `sync.Mutex` type ensures that only one goroutine can access a critical section of code at a time. Go also provides `sync.RWMutex` for read-heavy workloads, `sync.Once` for initialization that should happen only once, and other synchronization primitives.

## When to use it
- When you have shared data accessed by multiple goroutines
- To prevent race conditions in concurrent programs
- Implementing thread-safe data structures (caches, counters, etc.)
- When you need to ensure initialization happens only once (sync.Once)
- For read-heavy workloads where you want to allow multiple concurrent readers (sync.RWMutex)
- When you need to coordinate goroutines beyond simple waiting (sync.Cond)
- For object pooling to reuse memory and reduce allocation overhead (sync.Pool)

## Minimal syntax example
```go
package main

import (
    "fmt"
    "sync"
)

// Shared data
var (
    counter int
    mutex   sync.Mutex
)

// Safe increment function using mutex
func increment() {
    mutex.Lock()
    defer mutex.Unlock() // Ensure unlock happens even if we panic
    counter++
}

func main() {
    var wg sync.WaitGroup
    
    // Launch 10 goroutines to increment counter
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for j := 0; j < 100; j++ {
                increment()
            }
        }()
    }
    
    // Wait for all goroutines to finish
    wg.Wait()
    fmt.Printf("Final counter value: %d\n", counter) // Should be 1000
    
    // Example of RWMutex for read-heavy workloads
    var data map[string]string
    var rwmu sync.RWMutex
    
    // Writer goroutine
    go func() {
        rwmu.Lock()
        data = map[string]string{"key": "value"}
        rwmu.Unlock()
    }()
    
    // Multiple reader goroutines
    var readers sync.WaitGroup
    readers.Add(3)
    for i := 0; i < 3; i++ {
        go func(id int) {
            defer readers.Done()
            rwmu.RLock()
            if val := data["key"]; val != "" {
                fmt.Printf("Reader %d: %s\n", id, val)
            }
            rwmu.RUnlock()
        }(i)
    }
    readers.Wait()
    
    // Example of sync.Once for initialization
    var once sync.Once
    var shared string
    
    setup := func() {
        fmt.Println("Initializing...")
        shared = "Ready"
    }
    
    var initWg sync.WaitGroup
    for i := 0; i < 5; i++ {
        initWg.Add(1)
        go func(id int) {
            defer initWg.Done()
            once.Do(setup) // Only runs once
            fmt.Printf("Worker %d sees: %s\n", id, shared)
        }(i)
    }
    initWg.Wait()
}
```

## Common mistakes
- Forgetting to unlock a mutex (especially when there are multiple return paths or potential panics)
- Not using defer to ensure mutex unlock happens (leads to deadlocks)
- Locking a mutex twice from the same goroutine (causes deadlock)
- Using mutex when you don't actually have shared state
- Over-synchronizing (locking too much, reducing concurrency benefits)
- Not understanding the difference between Mutex and RWMutex
- Forgetting that sync.Once can only be used once (can't be reset)
- Using sync.Cond without understanding its complex semantics
- Not realizing that map access is not thread-safe by default
- Expecting mutexes to work with map iteration (they don't - you need to lock the entire iteration)
- Using mutexes for channel synchronization (channels already provide synchronization)

## Command to run it
```bash
go run mutex.go
```