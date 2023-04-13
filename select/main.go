package main

import (
	"fmt"
	"math/rand"
)

func sum(total chan int, exit chan bool) {
	valor := rand.Intn(20)
	for {
		select {
		case total <- valor:
			valor += rand.Intn(20)

		case <-exit:
			fmt.Println("exit :)")
			return
		}
	}
}

func main() {
	total := make(chan int)
	exit := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-total)
		}

		exit <- true
	}()

	sum(total, exit)
}
