package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulation: An Image Processing Service
// It takes an image URL, compresses it in the background,
// notifies the WaitGroup when done, and sends the finished result into the Channel!
func worker(url string, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // 1. Always decrement counter right before exiting

	time.Sleep(time.Millisecond * 60) // Simulating 60ms image compression
	fmt.Println("image processed ", url)

	// 2. Safely push the finished result into the Channel (Arrow points INTO ch)
	ch <- "PROCESSED: " + url
}

func main() {
	// 1. Declare WaitGroup to track active background workers
	var wg sync.WaitGroup

	// 2. Create a Buffered Channel (Capacity = 2 slots for our 2 images)
	// A buffered channel acts as an in-memory FIFO drop-box
	resultCHan := make(chan string, 2)

	// Start a timer to measure total concurrent execution time
	startime := time.Now()

	// 3. Set WaitGroup counter to 2 (Waiting for 2 background workers)
	wg.Add(2)

	// 4. Launch both workers concurrently!
	// Pass &wg as a pointer (shared counter) and resultCHan as the delivery pipe
	go worker("image_12.png", &wg, resultCHan)
	go worker("image_13.png", &wg, resultCHan)

	// 5. Meanwhile, main() could do other tasks here if needed!
	// fmt.Println("Main is doing other tasks while images compress...")

	// 6. Wait for both background workers to finish their jobs (Counter reaches 0)
	wg.Wait()

	// 7. Pull the finished results out of the Channel (Arrow points OUT of channel)
	res1 := <-resultCHan
	res2 := <-resultCHan

	fmt.Println("\n--- FINAL RESULTS RECEIVED IN MAIN ---")
	fmt.Println("Result 1:", res1)
	fmt.Println("Result 2:", res2)

	fmt.Printf("\nit took %s\n", time.Since(startime))
}

// -----------------------------------------------------------------------------
// CORE ARCHITECTURAL TAKEAWAYS:
// -----------------------------------------------------------------------------
// 1. sync.WaitGroup is the ALARM CLOCK:
//    It tells main() WHEN all background tasks are done so the program doesn't exit early.
//
// 2. The Channel is the SECURE DELIVERY PIPE:
//    It is a thread-safe FIFO queue that transports the data from background
//    workers back to main() with zero race conditions!
//
// 3. True Parallelism:
//    Both 60ms image tasks ran simultaneously on separate CPU cores,
//    completing in ~60ms total instead of 120ms sequential time!
