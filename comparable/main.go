package main

import "fmt"

type values interface {
	string | int32 | float32 | bool
}

type Test[C comparable, V values] struct {
	Key   C
	Value V
}

func compare[V values](m map[string]V, src, dst string) bool {
	return m[src] == m[dst]
}

func main() {
	mi := map[string]int32{
		"Tiago":   32,
		"João":    27,
		"Marta":   25,
		"Marlene": 32,
	}

	fmt.Println(compare(mi, "Tiago", "João"))
	fmt.Println(compare(mi, "Tiago", "Marlene"))

	t1 := Test[string, int32]{Key: "VAI", Value: 2}
	t2 := Test[string, int32]{Key: "VAI", Value: 2}
	t3 := Test[int, bool]{Key: 332, Value: true}

	fmt.Println(t1 == t2)
	fmt.Printf("%+v\n", t3)

}
