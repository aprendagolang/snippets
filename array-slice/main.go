package main

import "fmt"

func bubbleSort(n [6]int) [6]int {
	for i, _ := range n {
		for j := 0; j < len(n)-i-1; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
			}
		}
	}

	return n
}

func main() {
	unsorted := [...]int{1, 7, 8, 5, 3, 9}
	equal := [...]int{1, 7, 8, 5, 3, 9}
	sorted := bubbleSort(unsorted)
	fmt.Println(unsorted)
	fmt.Println(sorted)
	fmt.Println(unsorted == equal)

	slice := []int{1, 2, 3, 4, 5}

	a := slice[:3]
	b := slice[2:]

	slice[2] = 99
	
	fmt.Println(a, b)

}
