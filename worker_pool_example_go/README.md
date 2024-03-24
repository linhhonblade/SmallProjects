# Worker Pool Example Go

## Worker Pool Basics

**Worker Pool** is a software design pattern where a set of worker (pool) are created to concurrently process from a queue of task

This pattern includes 2 parts:

- A queuing system to deliver tasks to workers
- Logic for workers to acquire tasks

### Advantage

- Simplified code
- Resource management
- Improved performance
- Scalability

### Disadvantage

- **Idle thread overhead**: Even when there are no tasks to execute, the worker threads are still running in the background, consuming resources. This overhead can be negligible for CPU-bound tasks but can become significant for I/O-bound tasks.
- **Limited flexibility**:  The worker pool enforces a fixed number of worker threads. This might not be suitable for situations where the number of concurrent tasks varies significantly.
- **Potential for deadlocks**: If tasks within the pool acquire resources in a specific order and then wait for each other, a deadlock can occur

### Tuning

Worker pool size should be tuned to the specific application. Things to consider:

- How many tasks are coming?
- How quickly can tasks come in? Are they bursty, or more steady-state?
- How quickly can it handle the results of those tasks (if applicable)?
- How much of my system resources am I willing to put towards processing these tasks?
- Should the number of workers be dynamic? 


## Things Learned

### The use of `select`

```go
func (p *SimplePool) AddWork(t Task) {
    select {
    case p.tasks <- t:
    case <-p.quit:
    }
}
```
**Purpose**

The `select` statement in this function serves as a non-blocking communication mechanism, enabling the `AddWork` function to efficiently manage two operations:

1. **Adding a task to the work queue**: If the `tasks` channel is ready to receive data (not full), the `case p.tasks <- t:` clause executes, adding the new `Task` (represented by t) to the queue.
2. **Handling pool shutdown**: If the `quit` channel receives a signal (not empty), the `case <-p.quit:` clause executes, indicating that the pool is shutting down. The function exits, preventing further task additions.

**Non-blocking behavior**

- The `select` statement operates in a non-blocking fashion. This means that if neither channel (`tasks` or `quit`) is ready for immediate interaction, the function does not wait indefinitely. Instead, it gracefully continues execution without blocking the thread.

**Key benefit**

- This non-blocking approach is crucial for preventing deadlocks or stalls in the program. If the `AddWork` function were to block while adding a `task` (waiting for the `tasks` channel to become non-full), it could potentially block other goroutines (lightweight concurrent processes in Go) that might need to access the `quit` channel to signal pool shutdown.

### Unbuffered channel

In Go, creating a channel with a size of 0 (using `make(chan T, 0))` results in an **unbuffered channel**. This means that:

1. **No internal buffer**: Unlike channels with a capacity greater than 0, unbuffered channels don't have any built-in storage to hold elements waiting to be received.
2. **Synchronization**: Operations on unbuffered channels, both sending (`<-`) and receiving (`->`), act as synchronization points.
3. **Blocking behavior**: When sending to an unbuffered channel, the sending operation will **block** until a receiving operation is ready to accept the element. Conversely, when receiving from an unbuffered channel, the operation will **block** until an element is sent through the channel.

**Reasons to use a channel with size 0:**

1. **Ensuring synchronization**: Unbuffered channels guarantee that sending and receiving operations happen in a strict order. This is useful when you need to ensure that certain actions occur sequentially or avoid race conditions.
2. **Preventing data loss**: Since unbuffered channels don't have a buffer, they cannot lose data due to being full. This is relevant when dealing with critical or sensitive information that cannot be discarded.
3. **Efficient signaling**: Unbuffered channels are often used for signaling purposes, where the channel's primary function is to indicate events or completion rather than transferring large amounts of data. Sending an empty struct through an unbuffered channel is a common practice for efficient signaling.

**However, unbuffered channels also come with some drawbacks:**

1. **Blocking operations**: As mentioned previously, both sending and receiving operations can block, potentially impacting performance if not used carefully.
2. **Complexity**: Using unbuffered channels can sometimes lead to more complex code due to the need to manage synchronization explicitly.