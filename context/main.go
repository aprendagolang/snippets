package main

import (
	"context"
	"fmt"
	"time"
)

func doSomeHeavyWork(ctx context.Context, msg string) {
	c := make(chan bool)

	go func(c chan bool) {
		time.Sleep(2 * time.Second)
		c <- true
	}(c)

	select {
	case <-ctx.Done():
		fmt.Println("context timeout!")
	case <-c:
		fmt.Println(msg)
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	doSomeHeavyWork(ctx, "work completed!")
}
