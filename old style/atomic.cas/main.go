package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var lockIn int64 = 0
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()
			if atomic.CompareAndSwapInt64(&lockIn, 0, 1) {
				fmt.Printf("Goroutine #%d: I am the WINNER! I took leadership\n", id)
			} else {
				fmt.Printf("Goroutine #%d: Too late, the place is already taken\n", id)
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Battle is over")
}
