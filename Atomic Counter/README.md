***Atomic Counter in Go***
A simple project to demonstrate concurrency and thread-safety in Golang using the sync/atomic package.

***Overview***
This program spawns 1000 goroutines, where each one increments a shared counter.
The primary goal is to illustrate the solution to the Race Condition problem. In a concurrent environment, a simple count++ is not "atomic" (it involves reading, incrementing, and writing). Without proper synchronization, some increments would be lost. This project uses atomic.AddInt64 to ensure every single increment is recorded correctly at the CPU level.

***Features***
Goroutines: For high-performance concurrent execution.
sync.WaitGroup: To synchronize the main function with worker goroutines.
sync/atomic: For low-level, lock-free memory safety.

***Key Takeaways***
Race Conditions: Understanding why shared variables behave unpredictably in parallel loops.
Atomic vs Mutex: Atomic operations are faster than Mutexes for simple numeric updates because they don't require the operating system to "lock" a thread.
Memory Addresses: Learning why atomic functions require a pointer (&count) to modify the value directly in memory.
