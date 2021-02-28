package escape

import (
	"fmt"
	"time"
)

// Func returned struct value includes a pointer field, and this field has a referenced
// variable, this variable will escape.

func case1() {
	num := 42
	fmt.Println(num)

	emp := getEmployer1()
	fmt.Println(emp)
}

type employer1 struct {
	Name     string
	Age      int
	Sex      *string
	Birthday time.Time
}

//go:noinline
func getEmployer1() employer1 {
	sex := "man" // moved to heap: sex
	bir := time.Now()
	e := employer1{
		Name:     "adam",
		Age:      23,
		Sex:      &sex,
		Birthday: bir,
	}
	return e
}
