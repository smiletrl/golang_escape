package example2

type employer struct {
	Name  string
	Age   int
	Title string
}

//go:noinline
func getEmployerArray1() [1e5]employer {
	var emps [1e5]employer
	for i := 0; i < 1e5; i++ {
		emps[i] = employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
	}
	return emps
}

//go:noinline
func getEmployerArray2() [1e6]employer {
	var emps [1e6]employer //moved to heap: emps
	for i := 0; i < 1e6; i++ {
		emps[i] = employer{
			Name:  "adam",
			Age:   23,
			Title: "ceo",
		}
	}
	return emps
}
