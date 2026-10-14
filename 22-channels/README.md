# Channels

## What this topic is
Channels in Go are typed conduits through which you can send and receive values with the channel operator, <-. They provide a way for goroutines to communicate with each other and synchronize their execution. Channels are the primary mechanism for sharing data between goroutines in Go.

## When to use it
- When you need to communicate data between goroutines
- For synchronizing goroutine execution (waiting for signals)
- Implementing worker pools or pipelines
- When you want to avoid sharing memory by communicating instead
- Building concurrent systems like queues, pub/sub, or rate limiters
- When you need to multiplex or select between multiple channel operations

## Minimal syntax example
```go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create an unbuffered channel
    messages := make(chan string)
    
    // Send a value from a goroutine
    go func() { messages <- "ping" }()
    
    // Receive the value in the main goroutine
    msg := <-messages
    fmt.Println(msg) // prints "ping"
    
    // Buffered channel (can hold up to 2 values without blocking)
    buffered := make(chan int, 2)
    buffered <- 1
    buffered <- 2
    fmt.Println(<-buffered) // 1
    fmt.Println(<-buffered) // 2
    
    // Closing and ranging over channels
    jobs := make(chan int, 3)
    go func() {
        for j := 1; j <= 3; j++ {
            jobs <- j
            fmt.Println("sent job", j)
        }
        close(jobs) // signal no more values
    }()
    
    // Receive values until channel is closed
    for j := range jobs {
        fmt.Println("received job", j)
    }
    
    // Select statement (wait on multiple channels)
    c1 := make(chan string)
    c2 := make(chan string)
    go func() { time.Sleep(time.Second); c1 <- "one" }()
    go func() { time.Sleep(time.Second * 2); c2 <- "two" }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
        }
    }
}
```

## Common mistakes
- Forgetting that sending to a nil channel blocks forever
- Not understanding that unbuffered channels block both sender and receiver until the other is ready
- Trying to receive from a closed channel without checking if it's closed (use comma-ok idiom)
- Forgetting to close channels when no more values will be sent (causing range loops to block forever)
- Sending to a closed channel (causes panic)
- Not realizing that channel directions (send-only/receive-only) are only for type safety, they don't change the underlying channel
- Overusing buffered channels when unbuffered would suffice for synchronization
- Not understanding that select chooses randomly when multiple cases are ready
- Using channels for simple synchronization when sync.WaitGroup or other primitives would be clearer
- Forgetting that channels are reference types (passing copies of the channel header, not the underlying data structure)

## Command to run it
```bash
go run channels.go
```