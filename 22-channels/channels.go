package main

import (
	"fmt"
	"time"
)

// Example 1: Basic unbuffered channel communication
func basicChannel() {
	fmt.Println("=== Basic Channel ===")
	// Create an unbuffered channel of strings
	msgChan := make(chan string)

	// Send a value to the channel from a goroutine
	go func() {
		msgChan <- "Hello from goroutine"
	}()

	// Receive the value from the channel in main goroutine
	msg := <-msgChan
	fmt.Println(msg)
}

// Example 2: Buffered channel
func bufferedChannel() {
	fmt.Println("\n=== Buffered Channel ===")
	// Create a buffered channel with capacity 3
	ch := make(chan int, 3)

	// Send values without blocking (until buffer is full)
	ch <- 10
	ch <- 20
	ch <- 30

	// Receive values
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// Example 3: Channel directions
func ping(pings chan<- string, msg string) {
	// pings is send-only channel
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	// pings is receive-only, pongs is send-only
	msg := <-pings
	pongs <- msg
}

func channelDirections() {
	fmt.Println("\n=== Channel Directions ===")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

// Example 4: Closing channels and ranging
func channelClosing() {
	fmt.Println("\n=== Channel Closing and Ranging ===")
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Worker goroutine
	go func() {
		for {
			// Receive values until channel is closed
			j, more := <-jobs
			if more {
				fmt.Println("received job:", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// Send jobs and then close the channel
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// Wait for worker to finish
	<-done
}

// Example 5: Range over channels
func rangeOverChannel() {
	fmt.Println("\n=== Range Over Channel ===")
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Range receives values until channel is closed
	for elem := range queue {
		fmt.Println(elem)
	}
}

// Example 6: Select statement
func selectExample() {
	fmt.Println("\n=== Select Statement ===")
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// Example 7: Worker pool
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker %d started job %d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func workerPool() {
	fmt.Println("\n=== Worker Pool ===")
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send 5 jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for a := 1; a <= 5; a++ {
		<-results
	}
}

// Example 8: Rate limiting with channels
func rateLimiting() {
	fmt.Println("\n=== Rate Limiting ===")
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limiter channel receives a value every 200ms
	limiter := time.Tick(time.Millisecond * 200)

	// Process requests with rate limiting
	for req := range requests {
		<-limiter // Wait for limiter tick
		fmt.Println("request", req, "processed at", time.Now())
	}
}

// Example 9: Timeout with select
func timeoutSelect() {
	fmt.Println("\n=== Timeout with Select ===")
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}

func main() {
	basicChannel()
	bufferedChannel()
	channelDirections()
	channelClosing()
	rangeOverChannel()
	selectExample()
	workerPool()
	rateLimiting()
	timeoutSelect()
}