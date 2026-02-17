# Atomic Winner Competition

A Go project demonstrating the **Compare-And-Swap (CAS)** mechanism using the `sync/atomic` package.

## Overview
This program simulates a race between **100 goroutines**. The goal is to perform a "First-Come, First-Served" action without using traditional Mutex locks.

The core logic uses `atomic.CompareAndSwapInt64`, which checks if a variable matches an expected value (`0`) and only then updates it to a new value (`1`). This operation is performed atomically at the CPU level.



## 🛠 Features
* **Lock-free Synchronization**: Uses CAS instead of `sync.Mutex` for better performance in simple state transitions.
* **Concurrency Control**: Manages 100 simultaneous workers using `sync.WaitGroup`.
* **Atomic Boolean Logic**: Demonstrates how to handle success/failure results from atomic operations.
