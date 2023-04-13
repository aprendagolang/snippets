package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var total int64

	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddInt64(&total, int64(c))
			}
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(total)
}
