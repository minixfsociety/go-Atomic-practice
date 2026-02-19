package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var jobProgress int64
	go func() {
		for {
			atomic.AddInt64(&jobProgress, 1)
			time.Sleep(10 * time.Microsecond)
		}
	}()
	for i := 0; i < 5; i++ {
		val := atomic.LoadInt64(&jobProgress)
		fmt.Printf("Monitoring: current progress is %d\n", val)
		time.Sleep(time.Second)
	}
}
