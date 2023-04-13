package main

import "fmt"

type MyType interface {
	int | string | float32
}

func Reverse[V MyType](m []V) {
	first := 0
	last := len(m) - 1
	for first < last {
		m[first], m[last] = m[last], m[first]
		first++
		last--
	}
}

func main() {
	ints := []int{1, 2, 3, 4, 5, 6}
	strs := []string{"Tiago", "Dani", "Pedro", "Maria"}

	Reverse(ints)
	Reverse(strs)

	fmt.Println("int:", ints)
	fmt.Println("str:", strs)
}
