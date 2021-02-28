package noescape

import (
	"fmt"
	"math/rand"
)

func case1() {
	lenth := getLen()
	s := getSlice(int64(lenth))
	if 1 == 2 {
		fmt.Println(s)
	}
}

//go:noinline
func getLen() int {
	lent := rand.Intn(1e9)
	return lent
}

//go:noinline
func getSlice(len int64) []int {
	j := make([]int, len)
	return j
}
