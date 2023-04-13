package main

import "fmt"

func main() {
	var exit bool
	for !exit {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 7 {
				break
			}
		}
	}
}
