package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func process(num int, wg *sync.WaitGroup) {
	fmt.Printf("iniciando goroutine num %d \n", num)
	time.Sleep(time.Second * time.Duration(rand.Intn(10)))
	fmt.Printf("finalizando goroutine num %d \n", num)

	wg.Done()
}

func main() {
	limit := rand.Intn(100)
	var wg sync.WaitGroup
	for i := 0; i <= limit; i++ {
		wg.Add(1)
		go process(i, &wg)
	}

	wg.Wait()

	fmt.Printf("%d goroutines foram finalizadas com sucesso! \n", limit)
}
