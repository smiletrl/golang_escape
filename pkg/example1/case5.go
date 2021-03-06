package examples

import (
	"fmt"
	"time"
)

// content in slice has pointer or content value has pointer, the referenced variable escapes to heap.
// array does not work like slice though.

func case5() {
	num := 42
	fmt.Println(num)

	emp := getEmployer5()
	fmt.Println(emp)
}

type employer5 struct {
	Name     string
	Age      int
	Title    *string
	Birthday time.Time
}

//go:noinline
func getEmployer5() int {
	var names []*string
	name := "adam" // moved to heap: title
	for i := 0; i < 100; i++ {
		names = append(names, &name)
	}

	var employers []employer5
	title := "ceo" // moved to heap: title
	for k := 0; k < 100; k++ {
		employers = append(employers, employer5{
			Title: &title,
		})
	}

	var employers2 []employer5
	for k := 0; k < 100; k++ {
		employers2 = append(employers, employer5{})
	}

	var ages [80]*int
	age := 12 // does not escape !!!
	for j := 0; j < 80; j++ {
		ages[j] = &age
	}
	return len(names) + len(employers) + len(employers2) + len(ages)
}
