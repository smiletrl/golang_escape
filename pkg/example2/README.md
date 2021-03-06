## Example Description

Run command, and result is like:

```
smiletrl@Rulins-MacBook-Pro example2 % go build -gcflags="-m -l"
# github.com/smiletrl/golang_escape/pkg/example2
./case1.go:11:12: ... argument does not escape
./case1.go:11:65: time.Now().Sub(start).Seconds() escapes to heap
./case1.go:74:17: make([]employer1, 100000) escapes to heap
./case1.go:88:17: make([]employer1, 100000) escapes to heap
./case1.go:60:6: moved to heap: emps
```

We see only line 60, `moved to heap: emps`. But in fact, more variables got moved to heap!!!
