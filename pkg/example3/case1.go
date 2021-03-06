package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	caseArray()
	fmt.Printf("druation is: %+vs\n", time.Now().Sub(start).Seconds())
}

func caseArray() {
	var titles []string
	emps := getemployerArray()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

func caseSlice() {
	var titles []string
	emps := getemployerSlice()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

type employer struct {
	Name  string
	Age   int
	Title string
}

//go:noinline
func getemployerArray() [1e5]employer {
	var emps [1e5]employer
	for i := 0; i < 1e5; i++ {
		e := employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
	return emps
}

//go:noinline
func getemployerSlice() []employer {
	var emps = make([]employer, 1e5)
	for i := 0; i < 1e5; i++ {
		e := employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
	return emps
}
