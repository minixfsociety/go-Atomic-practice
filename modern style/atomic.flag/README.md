# Project 2: Modern Atomic Bool (Graceful Shutdown)

## 📌 Overview
This project demonstrates the **Modern Style (Go 1.19+)** of handling state flags using `atomic.Bool`. We simulate a background server process that can be safely stopped by a signal from the main thread.

## 🆕 What's New (Modern Style)
In the "Old Style," we had to use `int32` or `int64` (0 or 1) to represent a boolean flag. Now, Go provides a native `atomic.Bool` type which makes the code:
1. **Type-safe**: You can't accidentally store a number; it only accepts `true` or `false`.
2. **Readable**: No more pointers (`&`) needed for methods. You just call `.Load()` and `.Store()`.

## 🛠 Logic Flow
1. **Initialization**: `active.Store(true)` starts the "engine".
2. **The Worker**: A goroutine runs a `for` loop that checks `active.Load()` on every iteration.
3. **The Trigger**: After 2 seconds, the main thread calls `active.Store(false)`.
4. **The Exit**: The worker sees the `false` value, breaks the loop, and shuts down gracefully.

## 🚀 Key Advantage
Using `atomic.Bool` prevents **Race Conditions**. If you used a regular `bool` variable, the CPU might cache the value, and the background goroutine might never "see" that the variable changed to `false`, causing the server to run forever.

## How to Run
```bash
go run main.go