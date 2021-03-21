package example2

import (
	"testing"
)

// go test -bench=.
func BenchmarkCaseArray1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmployerArray1()
	}
}

func BenchmarkCaseArray2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		getEmployerArray2()
	}
}
