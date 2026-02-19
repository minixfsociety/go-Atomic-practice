package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var sleepTime int64 = 1000

	go func() {
		for {
			currentSleep := atomic.LoadInt64(&sleepTime)
			fmt.Printf("Server is working (sleep: %dms)...\n", currentSleep)
			time.Sleep(time.Duration(currentSleep) * time.Millisecond)
		}
	}()
	time.Sleep(3 * time.Second)
	fmt.Println(">>> ADMIN: Changing speed to FAST")
	atomic.StoreInt64(&sleepTime, 100)
	time.Sleep(2 * time.Second)
}
