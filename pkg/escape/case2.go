package escape

import (
	"fmt"
	"time"
)

// Func returned struct pointer will escape.

func case2() {
	num := 42
	fmt.Println(num)

	emp := getEmployer2()
	fmt.Println(emp)
}

type employer2 struct {
	Name     string
	Age      int
	Title    *string
	Birthday time.Time
}

//go:noinline
func getEmployer2() *employer2 {
	title := "man" // moved to heap: title
	bir := time.Now()
	e := employer2{ // // moved to heap: e
		Name:     "adam",
		Age:      23,
		Title:    &title,
		Birthday: bir,
	}
	return &e
}
