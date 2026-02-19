# Project 4: Atomic CAS (Compare-and-Swap)

## 📌 Overview
This project demonstrates the most powerful atomic operation: **Compare-And-Swap (CAS)**. It is the core of lock-free programming and leader election algorithms.

## 🛠 How it Works
1. We have a `lockIn` variable initialized to `0` (available).
2. 10 goroutines start simultaneously, each trying to change the value from `0` to `1`.
3. **The CAS Rule**: A goroutine can only update the value if it is **still** `0` at the exact moment of the attempt.
4. Only **one** goroutine succeeds and returns `true`. All others return `false`.

## 🚀 Why Use CAS?
- To implement **Optimistic Locking**.
- To choose a **Leader** in a distributed system.
- To build high-performance data structures without traditional Mutexes.

## How to Run
```bash
go run atomic_cas.go