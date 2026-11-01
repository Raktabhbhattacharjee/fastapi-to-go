package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulation: An Image Processing Service
// It takes an image URL, simulates a heavy task (resizing/compressing),
// prints when done, and notifies the WaitGroup.

func worker(url string, wg *sync.WaitGroup) string {
	defer wg.Done()
	time.Sleep(time.Millisecond * 60)
	fmt.Println("image processed ", url)
	// wg.Done() // Decrements counter: counter = counter - 1
	return url
}

func main() {
	// Declare a WaitGroup from the 'sync' package to track active background workers
	var wg sync.WaitGroup

	// Start a timer to measure how long the concurrent execution takes
	startime := time.Now()

	// 1. Tell WaitGroup we have 2 background jobs to wait for (Counter = 2)
	wg.Add(2)

	// 2. Launch both workers concurrently (Pass &wg as a pointer so state is shared)
	go worker("image_12.png", &wg)
	go worker("image_13.png", &wg)

	// Note: Goroutine return values are ignored by Go.
	// (To send values back to main, we use Go Channels).
	// fmt.Println("result", result)
	// fmt.Println("result1", result1)

	// 3. WaitGroup maintains an internal counter:
	// wg.Wait() pauses main() here until the counter drops back down to 0
	wg.Wait()

	fmt.Printf("it took %s\n", time.Since(startime))
}

// -----------------------------------------------------------------------------
// CORE TAKEAWAYS & COMPARISON WITH SEQUENTIAL CODE:
// -----------------------------------------------------------------------------
// 1. Sequential execution (Standard C/C++ style without threads):
//    main() would call task 1 -> wait 60ms -> call task 2 -> wait 60ms.
//    Total time would be: 60ms + 60ms = 120ms.
//
// 2. Go concurrent execution with 'go':
//    main() spawns both workers in the background and moves to the next line immediately.
//    Without a wait mechanism, main() would exit in 0.001ms before workers could finish!
//
// 3. Why we use sync.WaitGroup:
//    It acts as a synchronized barrier holding the door open until all workers
//    signal completion via wg.Done(). Both tasks run in parallel in ~60ms total.
