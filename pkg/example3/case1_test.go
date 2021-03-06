package main

import (
	"testing"
)

// go test -bench=. benmem
func BenchmarkCase1Array(b *testing.B) {
	for n := 0; n < b.N; n++ {
		caseArray()
	}
}

func BenchmarkCase1Slice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		caseSlice()
	}
}
