package main

import (
	"testing"
)

// go test -bench=.
func BenchmarkCase1Array(b *testing.B) {
	for n := 0; n < b.N; n++ {
		case1Array()
	}
}

func BenchmarkCase1Slice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		case1Slice()
	}
}
