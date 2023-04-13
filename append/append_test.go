package main

import (
	"testing"
)

func BenchmarkAppend1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append1()
	}
}

func BenchmarkAppend2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append2()
	}
}

func BenchmarkAppend3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append3()
	}
}
