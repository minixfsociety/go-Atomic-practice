# Project 1: Atomic Add (Global Request Counter)

## 📌 Overview
This project demonstrates how to handle a **Shared State** in a concurrent environment using the `sync/atomic` package. We simulate a scenario where 1000 independent goroutines try to increment a single global counter simultaneously.

## 🔴 The Problem: Race Condition
In a standard Go program, an operation like `counter++` is not "atomic". It consists of three steps:
1. Read the current value.
2. Increment the value in memory.
3. Write the new value back.

When multiple goroutines do this at the same time, they might read the same old value, leading to lost updates. This is known as a **Race Condition**.

## ✅ The Solution: Low-Level Atomics
To solve this without the overhead of a Mutex, we use `atomic.AddInt64`. 
- **Efficiency**: It performs the operation at the CPU instruction level.
- **Safety**: It guarantees that the "Read-Modify-Write" cycle is completed without interruption.

## 🛠 Concepts Used
- `sync.WaitGroup`: To ensure the main thread waits for all 1000 goroutines to finish.
- `sync/atomic`: For thread-safe integer manipulation.
- **Pointers (`&`)**: Passing the memory address to the atomic function so it modifies the actual variable.
