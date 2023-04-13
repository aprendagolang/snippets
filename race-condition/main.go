package main

import (
	"fmt"
	"sync"
)

var total = 0

func count(wg *sync.WaitGroup, c chan bool) {
	c <- true
	total++
	<-c
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	c := make(chan bool, 1)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go count(&wg, c)
	}
	wg.Wait()
	fmt.Println("total: ", total)
}
