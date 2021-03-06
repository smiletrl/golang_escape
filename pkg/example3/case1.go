package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	caseSlice()
	fmt.Printf("druation is: %+vs\n", time.Now().Sub(start).Seconds())
}

func caseArray() {
	for i := 0; i < 100; i++ {
		getemployerArray()
	}
}

func caseSlice() {
	for i := 0; i < 100; i++ {
		getemployerSlice()
	}
}

type employer struct {
	Name  string
	Age   int
	Title string
}

//go:noinline
func getemployerArray() {
	var emps [1e5]employer
	for i := 0; i < 1e5; i++ {
		e := employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
}

//go:noinline
func getemployerSlice() {
	var emps = make([]employer, 1e5)
	for i := 0; i < 1e5; i++ {
		e := employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
}
