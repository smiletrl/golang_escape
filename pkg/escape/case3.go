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
	Title    *string
	Birthday time.Time
}

//go:noinline
func updateEmployer(emp *employer3) {
	title := "ceo" // moved to heap: title
	emp.Name = "sara"
	emp.Title = &title
}
