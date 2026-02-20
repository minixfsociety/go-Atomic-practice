package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var active atomic.Bool
	active.Store(true)
	go func() {
		for active.Load() {
			fmt.Println("Server is working...")
			time.Sleep(100 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Println("Stopping server...")
	active.Store(false)
	time.Sleep(1 * time.Second)
	fmt.Println("Goodbye")

}
