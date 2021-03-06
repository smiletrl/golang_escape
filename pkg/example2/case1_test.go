package example2

import (
	"testing"
)

// go test -bench=.
func BenchmarkCase1Array(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmployer1Array()
	}
}

func BenchmarkCase1Array2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmployer1Array2()
	}
}
