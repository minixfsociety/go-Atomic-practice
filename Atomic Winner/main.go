package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var winner int64
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			if atomic.CompareAndSwapInt64(&winner, 0, 1) {
				fmt.Printf("Goroutine #%d is the winner! \n", id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Competition finished. Final winner state:", winner)
}
