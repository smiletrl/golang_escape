package main

import (
	"testing"
)

// go test -bench=. -benchmem
func BenchmarkSlice1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getemployerSlice1()
	}
}

func BenchmarkSlice2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getemployerSlice2()
	}
}
