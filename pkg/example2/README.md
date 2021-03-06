## Example Description

This is an example of a `large local variable` allocated to heap.

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

We see only line 60 `moved to heap: emps` from compiler's output, but line 46 `var emps [1e5]employer1` is not moved to heap, even though they are both local array variables. This comes to Golang's rule of `Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.` from [FAQ](https://golang.org/doc/faq#stack_or_heap).

But how `large` is large? In this example, `var emps [1e5]employer1` size is `4MB = (40 Byte * 1e5)`, and `var emps [1e6]employer1` size is `40MB = (40 Byte * 1e6)`. So `40MB` is large enough for Golang to move this variable to heap instead of stack. We will not cover the internal implementation of how golang makes the decision, which probably deserves a separate topic to talk.

## How to get the variable size?

In this example, we see the variable size matters for whether it will be allocated to heap. This example provides a few options to get variable size.

To get this variable `var emps [1e5]employer1` size, each struct `employer` contains 2 string fields, and one int field. In my machine (mac amd64), each string is `16 Byte`, and int is `8 Byte`, which gives the total size of one `employer` struct to be 40 Byte.

One thing is this size `16 Byte` for string is simply the [`string header struct`](https://golang.org/pkg/reflect/#StringHeader) size, which is a result of [unsafe.Sizeof](https://golang.org/pkg/unsafe/#Sizeof). See `The size does not include any memory possibly referenced by x` from `Sizeof` doc. `String` in golang is actually a reference type (`Data uintptr` from string header struct is a pointer to the real byte data).

To get the `real` memory (i.e, including the pointed value's memory) allocated to the string variable, we can do

```
stringSize := len(str) + unsafe.Sizeof(str)
```

`len(str)` will return the byte length of the pointed value from the string `str`.

Another option to get the size of a variable is to use the built in benchmark tool

Run below command:

```
smiletrl@Rulins-MacBook-Pro example2 % go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/smiletrl/golang_escape/pkg/example2
BenchmarkCase1Array-16     	    1280	    952224 ns/op	       0 B/op	       0 allocs/op
BenchmarkCase1Array2-16    	      25	  49134560 ns/op	40001616 B/op	       1 allocs/op
```

See [Benchmark result](https://golang.org/pkg/testing/#BenchmarkResult) for meaning. We are interested at `40001616 B/op`, which is returned from [AllocedBytesPerOp](https://golang.org/pkg/testing/#BenchmarkResult.AllocedBytesPerOp). In this case, we see the size of `[1e6]employer1` (maybe plus the allocation of `i`) is `40001616 Byte ~= 40 MB`.



