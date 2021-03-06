package main

import (
	"fmt"
	"unsafe"
)

func case1Array() {
	var titles []string
	emps := getEmployer1Array()
	for _, emp := range emps {
		titles = append(titles, emp.Title)
	}
}

func case1Array1() {
	var titles []string
	emps := getEmployer1Array2()
	fmt.Printf("emp size is: %+v\n", unsafe.Sizeof(emps))
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
		emps[i] = employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
	}
	return emps
}

//go:noinline
func getEmployer1Array2() [1e6]employer1 {
	var emps [1e6]employer1 //moved to heap: emps
	for i := 0; i < 1e6; i++ {
		emps[i] = employer1{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
	}
	return emps
}
