package main

import (
	"fmt"
	"sync"
	"time"
)

// Example 1: Basic Mutex Usage
func basicMutex() {
	fmt.Println("=== Basic Mutex ===")
	var counter int
	var mu sync.Mutex

	// Increment counter from multiple goroutines safely
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

// Example 2: Proper defer pattern (recommended)
func properDeferPattern() {
	fmt.Println("\n=== Proper Defer Pattern ===")
	var counter int
	var mu sync.Mutex

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock() // Unlock happens when function returns
			counter++
			// If we had a return or panic here, unlock would still happen
		}()
	}
	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}

// Example 3: Mutex protecting a struct
type SafeCounter struct {
	mu      sync.Mutex
	Counter map[string]int
}

func (sc *SafeCounter) Inc(key string) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.Counter[key]++
}

func (sc *SafeCounter) Value(key string) int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.Counter[key]
}

func mutexWithStruct() {
	fmt.Println("\n=== Mutex with Struct ===")
	sc := SafeCounter{Counter: make(map[string]int)}
	var wg sync.WaitGroup

	keys := []string{"a", "b", "c", "d", "e"}
	for _, k := range keys {
		wg.Add(1)
		go func(key string) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				sc.Inc(key)
			}
		}(k)
	}
	wg.Wait()

	for _, k := range keys {
		fmt.Printf("Key %s: %d\n", k, sc.Value(k))
	}
}

// Example 4: RWMutex for read-heavy workloads
func rwMutexExample() {
	fmt.Println("\n=== RWMutex Example ===")
	var data map[string]string
	var rwmu sync.RWMutex

	// Writer goroutine
	go func() {
		for i := 0; i < 10; i++ {
			rwmu.Lock()
			if data == nil {
				data = make(map[string]string)
			}
			data[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
			rwmu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Multiple reader goroutines
	var wg sync.WaitGroup
	for r := 0; r < 3; r++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				rwmu.RLock()
				if len(data) > 0 {
					for k, v := range data {
						fmt.Printf("Reader %d: %s=%s\n", readerID, k, v)
					}
				}
				rwmu.RUnlock()
				time.Sleep(50 * time.Millisecond)
			}
		}(r)
	}

	time.Sleep(2 * time.Second) // Let goroutines run
	wg.Wait()
}

// Example 5: sync.Once for initialization
func onceExample() {
	fmt.Println("\n=== sync.Once Example ===")
	var once sync.Once
	var sharedResource string

	initialize := func() {
		fmt.Println("Initializing shared resource...")
		sharedResource = "Initialized Value"
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			once.Do(initialize)
			fmt.Printf("Goroutine %d sees: %s\n", id, sharedResource)
		}(i)
	}
	wg.Wait()
}

// Example 6: Common Mistake - Double Lock (causes deadlock)
func doubleLockMistake() {
	fmt.Println("\n=== Double Lock Mistake (will deadlock) ===")
	var mu sync.Mutex
	fmt.Println("Attempting to double lock...")
	mu.Lock()
	// This will deadlock because we're trying to lock an already locked mutex
	// mu.Lock() // Uncommenting this line would cause deadlock
	mu.Unlock()
	fmt.Println("Successfully unlocked after single lock")
}

// Example 7: Conditional variable (less common but useful)
func condExample() {
	fmt.Println("\n=== Cond Variable Example ===")
	var mu sync.Mutex
	var cond = sync.NewCond(&mu)
	queue := []interface{}{}

	// Consumer
	go func() {
		mu.Lock()
		for len(queue) == 0 {
			cond.Wait() // Releases lock and waits, re-acquires when signaled
		}
		item := queue[0]
		queue = queue[1:]
		fmt.Printf("Consumed: %v\n", item)
		mu.Unlock()
	}()

	// Producer
	mu.Lock()
	queue = append(queue, "item-1")
	fmt.Println("Produced: item-1")
	cond.Signal() // Wake up one waiting goroutine
	mu.Unlock()

	time.Sleep(100 * time.Millisecond) // Let consumer finish
}

// Example 8: Pool for object reuse
func poolExample() {
	fmt.Println("\n=== Pool Example ===")
	var bufferPool = sync.Pool{
		New: func() interface{} {
			return new([]byte) // Returns a pointer to a byte slice
		},
	}

	// Get from pool (may create new or reuse)
	buf := bufferPool.Get().(*[]byte)
	*buf = append(*buf, 'h', 'e', 'l', 'l', 'o')
	fmt.Printf("Buffer content: %s\n", *buf)

	// Return to pool
	bufferPool.Put(buf)

	// Get again (might get the same buffer back)
	buf2 := bufferPool.Get().(*[]byte)
	fmt.Printf("Buffer2 length: %d\n", len(*buf2)) // Should be 0 if cleared
	bufferPool.Put(buf2)
}

func main() {
	basicMutex()
	properDeferPattern()
	mutexWithStruct()
	rwMutexExample()
	onceExample()
	doubleLockMistake()
	condExample()
	poolExample()
}