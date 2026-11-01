# Concurrency in Go: The Master Cheatsheet

## 1. Operating System vs Go Concurrency Model

| Metric | OS Thread (C / C++ / Java) | Go Goroutine (`go fn()`) |
| :--- | :--- | :--- |
| **Managed By** | OS Kernel (Ring 0 System Calls) | Go Runtime Scheduler (Ring 3 User Space) |
| **Initial Memory** | **1 MB – 8 MB** (Fixed Stack) | **~2 KB** (Dynamically Resizable) |
| **Creation Cost** | High (Traps to Kernel, sets up TCB) | Negligible (Plain memory allocation) |
| **Switch Time** | ~1,000 – 2,000 ns (CPU cache flush) | ~10 – 20 ns (Only 3 registers swapped: PC, SP, DX) |
| **Capacity** | A few thousand before OOM crash | Hundreds of thousands to millions |

---

## 2. Event Loop (Node.js/Python) vs Go Multi-Core Parallelism

* **Node.js / Python Async**:
  * Single thread running in an event loop (the "juggler").
  * Great for non-blocking I/O (network/file waiting).
  * **Critical weakness:** Heavy CPU calculation (image processing, crypto) freezes the entire event loop for all users.
* **Go Goroutines**:
  * True hardware parallelism: Go's **M:N Scheduler** maps $M$ Goroutines onto $N$ real OS threads across all CPU cores.
  * A heavy calculation on Core 1 does not slow down other Goroutines running on Core 2, 3, or 4.
  * No infectious `async/await` syntax ("function coloring"). Functions remain regular, idiomatic Go functions.

---

## 3. The 3 Phases of `sync.WaitGroup`

`sync.WaitGroup` maintains an internal thread-safe integer counter:

```go
var wg sync.WaitGroup

// 1. ADD: Increment counter by N before launching goroutines
wg.Add(2) // Counter = 2

// 2. DONE: Decrement counter by 1 inside worker (Always use defer!)
go func() {
    defer wg.Done() // Counter = Counter - 1
    // do work...
}()

// 3. WAIT: Pause the caller until counter drops to 0
wg.Wait()
```

### The Golden Rule: Always Pass `*sync.WaitGroup` by Pointer!
```go
// CORRECT: Pointer passes shared state
func worker(url string, wg *sync.WaitGroup) { ... }

// WRONG: Value copy creates an isolated counter; wg.Wait() deadlocks!
func worker(url string, wg sync.WaitGroup) { ... }
```

---

## 4. Go Channels: Thread-Safe FIFO Queues

Go's official motto:
> *"Do not communicate by sharing memory; instead, share memory by communicating."*

### Channel Syntax Reference:

```go
// 1. Unbuffered Channel (Synchronous Handshake: Sender & Receiver must meet)
ch := make(chan string)

// 2. Buffered Channel (FIFO Queue of capacity N)
ch := make(chan string, 10)

// 3. Send data INTO the channel (Worker)
ch <- "my_result"

// 4. Receive data OUT of the channel (Main)
result := <-ch

// 5. Close a channel (Signals to receivers: no more data will be sent)
close(ch) // ALWAYS close from the sender side, never from the receiver!

// 6. Range over channel until closed
for item := range ch {
    fmt.Println(item)
}
```

---

## 5. The `select` Statement (Channel Switchboard & Timeouts)

`select` is like a `switch` statement, but for channels. It blocks until **one of its cases is ready to execute**.

### Syntax: Handling Multiple Channels + Timeout

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from Service 1:", msg1)

case msg2 := <-ch2:
    fmt.Println("Received from Service 2:", msg2)

case <-time.After(2 * time.Second):
    // If neither channel responds within 2 seconds, this case runs!
    fmt.Println("Request timed out! Aborting.")

default:
    // Non-blocking case: runs instantly if all channels above are busy
    fmt.Println("No channels ready, doing other work...")
}
```

---

## 6. `context.Context` (Cancellation & Deadlines)

Used in production to signal background goroutines to **stop working** (e.g., when a user cancels an HTTP request or navigates away).

### Syntax: WithTimeout & WithCancel

```go
// 1. Create a context that automatically cancels after 3 seconds
ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
defer cancel() // Always release resources when main finishes

// 2. Pass ctx to the background goroutine
go func(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            // Received cancellation signal!
            fmt.Println("Task canceled:", ctx.Err())
            return
        default:
            // Do normal work in batches...
            time.Sleep(500 * time.Millisecond)
        }
    }
}(ctx)
```

---

## 7. `sync.Mutex` & `sync.RWMutex` (Thread-Safe Shared State)

When you **must** share a plain variable (like a counter or a map/cache) between goroutines without channels:

### 1. `sync.Mutex` (Mutual Exclusion Lock)
```go
type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()         // Only 1 goroutine can enter this block at a time
    defer c.mu.Unlock() // Always unlock right before exiting!
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

### 2. `sync.RWMutex` (Read-Write Lock: High-Performance Caching)
* Many goroutines can **READ** at the same time (`RLock()` / `RUnlock()`).
* Only one goroutine can **WRITE** at a time (`Lock()` / `Unlock()`).

```go
var rw sync.RWMutex

// For reading:
rw.RLock()
data := cache["key"]
rw.RUnlock()

// For writing:
rw.Lock()
cache["key"] = "new_value"
rw.Unlock()
```

---

## 8. Summary Comparison of Go Concurrency Tools

| Tool | When to Use | Key Mental Model |
| :--- | :--- | :--- |
| **Goroutines (`go fn()`)** | Execute a task asynchronously in the background | Pressing start on the washing machine |
| **`sync.WaitGroup`** | Wait for a batch of background tasks to complete | Counting attendance before the bus departs |
| **Channels (`chan`)** | Safely pass data/results between goroutines | Thread-safe FIFO conveyor belt |
| **`select`** | Listen to multiple channels or enforce a timeout | Switchboard answering the first ringing phone |
| **`context.Context`** | Propagate cancellation signals down a call tree | The "Emergency Stop" button |
| **`sync.Mutex`** | Guard a shared variable or map in memory | The single key to the bathroom door |
