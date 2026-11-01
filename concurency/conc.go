package main

import (
	"fmt"
	"time"
)

// =============================================================================
// GO CHANNELS SYNTAX & CONCEPT CHEATSHEET
// =============================================================================
// What is a Channel?
// 1. A typed, thread-safe pipe/conduit that connects concurrent Goroutines.
// 2. Operates as a FIFO Queue (First-In, First-Out).
// 3. Avoids Race Conditions: Transfers data ownership without shared RAM locks!
// =============================================================================

// -----------------------------------------------------------------------------
// Pattern 1: Unbuffered Channel (Synchronous Handshake)
// -----------------------------------------------------------------------------
// - make(chan Type) has capacity = 0.
// - Sender BLOCKS until Receiver is ready to read.
// - Receiver BLOCKS until Sender puts data into the channel.
func demoUnbuffered() {
	fmt.Println("--- 1. UNBUFFERED CHANNEL ---")

	// 1. Create an unbuffered channel for strings
	ch := make(chan string)

	// 2. Launch background worker that sends data
	go func() {
		time.Sleep(50 * time.Millisecond) // Simulating work
		// The arrow points INTO the channel: ch <- data
		ch <- "Hello from Goroutine!"
		fmt.Println("Worker: Message sent into channel")
	}()

	// 3. Main receives data (The arrow points OUT of the channel: <-ch)
	// main() BLOCKS here until the worker sends the data!
	msg := <-ch
	fmt.Println("Main received:", msg)
}

// -----------------------------------------------------------------------------
// Pattern 2: Buffered Channel (Asynchronous Queue)
// -----------------------------------------------------------------------------
// - make(chan Type, capacity) has an internal buffer/queue.
// - Sender does NOT block as long as there is room in the buffer.
// - Sender only blocks when the buffer is 100% full.
func demoBuffered() {
	fmt.Println("\n--- 2. BUFFERED CHANNEL (Capacity = 2) ---")

	// Create a channel that can hold up to 2 items without blocking
	ch := make(chan string, 2)

	// We can send 2 items immediately without any background receiver waiting!
	ch <- "Item 1"
	ch <- "Item 2"

	// Reading them back follows FIFO (First-In, First-Out):
	fmt.Println("Popped:", <-ch) // "Item 1"
	fmt.Println("Popped:", <-ch) // "Item 2"
}

// -----------------------------------------------------------------------------
// Pattern 3: Closing Channels & Ranging Over Them
// -----------------------------------------------------------------------------
// - close(ch) signals to receivers: "No more data will ever be sent."
// - 'for item := range ch' keeps reading until the channel is closed.
// - Important rule: ONLY the sender should close a channel, never the receiver!
func demoCloseAndRange() {
	fmt.Println("\n--- 3. CLOSE & RANGE OVER CHANNEL ---")

	ch := make(chan int, 3)

	// Producer Goroutine: Sends numbers and then closes the channel
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch) // Close channel when done producing!
	}()

	// Consumer (main): Reads all items until closed
	for num := range ch {
		fmt.Println("Received number:", num)
	}
	fmt.Println("Channel closed, loop finished cleanly!")
}

// -----------------------------------------------------------------------------
// Pattern 4: Solving the Image Processing Problem with Channels!
// -----------------------------------------------------------------------------
// Returning data from concurrent workers safely without data races:
func processImageWorker(url string, out chan string) {
	time.Sleep(60 * time.Millisecond) // Simulating image compression
	// Send processed result back into the pipe:
	out <- "PROCESSED: " + url
}

func demoImageProcessingWithChannels() {
	fmt.Println("\n--- 4. RETURNING DATA SAFELY WITH CHANNELS ---")

	results := make(chan string, 2)

	// Fire 2 concurrent workers and pass the channel
	go processImageWorker("image_12.png", results)
	go processImageWorker("image_13.png", results)

	// Collect the results safely from the queue:
	res1 := <-results
	res2 := <-results

	fmt.Println("Got result 1:", res1)
	fmt.Println("Got result 2:", res2)
}

func main() {
	demoUnbuffered()
	demoBuffered()
	demoCloseAndRange()
	demoImageProcessingWithChannels()
}
