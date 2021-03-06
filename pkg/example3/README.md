## Example Description

This is an example of `large local variable` escaping to heap.

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

We see only line 60 `moved to heap: emps` from compiler's output, but line 46 `var emps [1e5]employer1` is not moved to heap, even they are both local array variables. This comes to Golang's rule of `Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.` from [FAQ](https://golang.org/doc/faq#stack_or_heap).

