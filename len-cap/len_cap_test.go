package main

import (
	"testing"
)

var result []string

func BenchmarkLen(b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		r = Len()
	}

	result = r
}

func BenchmarkCap(b *testing.B) {
	var r []string
	for n := 0; n < b.N; n++ {
		r = Cap()
	}

	result = r
}
