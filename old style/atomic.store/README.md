# Project 3: Atomic Store (Server Speed Controller)

## 📌 Overview
This project demonstrates how to **dynamically update** a shared configuration across different goroutines without stopping the program. We simulate a server whose "sleep time" (processing speed) can be changed instantly by an administrator.

## 🔴 The Challenge
Usually, if one goroutine is reading a variable in a loop and another one changes it, you might run into two problems:
1. **Race Condition**: The program might crash or behave unpredictably.
2. **Caching**: The worker goroutine might "remember" the old value in its CPU cache and never see the update made by the main thread.

## ✅ The Solution: `atomic.StoreInt64`
We use `atomic.StoreInt64` to perform a **thread-safe write** operation.
- **Immediate Visibility**: As soon as the "Admin" (main thread) calls `Store`, the "Server" (background goroutine) sees the new value on its next `Load`.
- **No Mutex Needed**: Since we are only updating a single number, Atomics are much faster than locking the entire code block with a Mutex.

## 🛠 How it Works
1. **The Variable**: `sleepTime` starts at `1000` (1 second).
2. **The Background Worker**: Constantly "peeks" at the memory using `atomic.LoadInt64`.
3. **The Update**: After 3 seconds, we use `atomic.StoreInt64(&sleepTime, 100)`.
4. **The Result**: The server suddenly starts working 10 times faster.


