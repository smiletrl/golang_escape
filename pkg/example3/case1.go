package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	case1Array()
	fmt.Printf("druation is: %+vs\n", time.Now().Sub(start).Seconds())
}

func case1Array() {
	var titles []string
	emps := getEmployer1Array()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

func case1Slice() {
	var titles []string
	emps := getEmployer1Slice()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

func case1Slice2() {
	var titles []string
	emps := getEmployer1Slice2()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

type employer1 struct {
	Name  string
	Age   int
	Title string
}

//go:noinline
func getEmployer1Array() [1e5]employer1 {
	var emps [1e5]employer1
	for i := 0; i < 1e5; i++ {
		e := employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
	return emps
}

//go:noinline
func getEmployer1Array2() [1e6]employer1 {
	var emps [1e6]employer1 //moved to heap: emps
	for i := 0; i < 1e6; i++ {
		e := employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
	return emps
}

//go:noinline
func getEmployer1Slice() []employer1 {
	var emps = make([]employer1, 1e5)
	for i := 0; i < 1e5; i++ {
		e := employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
	return emps
}

//go:noinline
func getEmployer1Slice2() []employer1 {
	var emps = make([]employer1, 1e5)
	for i := 0; i < 1e5; i++ {
		e := employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps = append(emps, e)
	}
	return emps
}
