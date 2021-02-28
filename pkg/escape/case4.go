package escape

import (
	"fmt"
	"time"
)

// If chan's struct value has pointer field value, the referenced variable
// will escape.

func case4() {
	num := 42
	fmt.Println(num)

	var employerChan chan employer4
	getEmployer4(employerChan)
	fmt.Println(<-employerChan)
}

type employer4 struct {
	Name     string
	Age      int
	Sex      *string
	Birthday time.Time
}

//go:noinline
func getEmployer4(emp chan<- employer4) {
	sex := "man" // moved to heap: sex
	bir := time.Now()
	e := employer4{
		Name:     "adam",
		Age:      23,
		Sex:      &sex,
		Birthday: bir,
	}
	emp <- e
}
