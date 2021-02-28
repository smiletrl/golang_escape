package escape

import (
	"fmt"
	"time"
)

// Func pointer argument has pointer field, and this pointer referenced field escapes.

func case3() {
	x := 42
	fmt.Println(x)

	emp := employer3{}
	updateEmployer(&emp)
}

type employer3 struct {
	Name     string
	Age      int
	Sex      *string
	Birthday time.Time
}

//go:noinline
func updateEmployer(emp *employer3) {
	sex := "woman" // moved to heap: sex
	emp.Name = "sara"
	emp.Sex = &sex
}
