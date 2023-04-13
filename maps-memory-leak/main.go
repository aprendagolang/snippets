package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := 1_000_000
	m := make(map[int][256]byte)
	printAlloc()

	for i := 0; i < n; i++ {
		m[i] = [256]byte{}
	}
	printAlloc()

	for i := 0; i < n; i++ {
		delete(m, i)
	}

	// m = make(map[int][128]byte)
	runtime.GC()
	printAlloc()
	runtime.KeepAlive(m)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
