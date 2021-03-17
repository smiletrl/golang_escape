## Example Description

This is an example of a `large local variable` allocated to heap.

Run command, and result is like:

```
smiletrl@Rulins-MacBook-Pro example2 % go build -gcflags="-m -l" case1.go
# command-line-arguments
./case1.go:24:6: moved to heap: emps
```

We see only line 24 `moved to heap: emps` from compiler's output, but line 11 `var emps [1e5]employer1` is not moved to heap, even though they are both local array variables. This comes to Golang's rule of `Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.` from [FAQ](https://golang.org/doc/faq#stack_or_heap).

But how `large` is large? In this example, `var emps [1e5]employer1` size is `4MB = (40 Byte * 1e5)`, and `var emps [1e6]employer1` size is `40MB = (40 Byte * 1e6)`. So `40MB` is large enough for Golang to move this variable to heap instead of stack. For the exact size, check [this article](https://smiletrl.github.io/post/golang-local-large-heap-allocated-variable/).

## How to get the variable memory size?

In this example, we see the variable memory size matters for whether it will be allocated to heap. This example provides a few options to get variable size.

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

See [Benchmark result](https://golang.org/pkg/testing/#BenchmarkResult) for meaning. We are interested at `40001616 B/op`, which is returned from [AllocedBytesPerOp](https://golang.org/pkg/testing/#BenchmarkResult.AllocedBytesPerOp). In this case, we see the size of `[1e6]employer1` is `40001616 Byte ~= 40 MB`. `var emps [1e5]employer1` is not allocated at heap, so its result is empty.

To add [pprof profile](https://golang.org/pkg/runtime/pprof/) support, run command

```
smiletrl@Rulins-MacBook-Pro example2 % go test -bench BenchmarkCase1Array2 -benchmem -memprofile memprofile.out
smiletrl@Rulins-MacBook-Pro example2 % go tool pprof memprofile.out
```

To avoid confusion, the above test command only runs `BenchmarkCase1Array2`.

This profile `memprofile.out` shows memory allocated from `getEmployer1Array2()` to heap, i.e, this variable `var emps [1e6]employer1`, and then we may [play with this output profile](https://blog.golang.org/pprof). This repository includes one png format of this memory profile as `profile001.png`. It shows 38.15 MB. Note, [pprof uses MiB](https://github.com/google/pprof/issues/136) as measurement. To do the conversion, we get `38.15MiB ~= 40 MB`, which is consistent with the above `40MB` result.
