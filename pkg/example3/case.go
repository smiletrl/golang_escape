package main

type employer struct {
	Name  string
	Age   int
	Title string
}

//go:noinline
func getemployerSlice1() {
	var emps = make([]employer, 10)
	for i := 0; i < 10; i++ {
		e := employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
		emps[i] = e
	}
}

//go:noinline
func getemployerSlice2() {
	var emps []employer
	e := employer{
		Name:  "adam",
		Age:   23,
		Title: "ceo",
	}
	emps = append(emps, e)
}

func main() {}
