## Golang Escape Analysis

According to [Golang FAQ](https://golang.org/doc/faq#stack_or_heap), it's not very 
clear when a variable will be allocated at heap.

Golang uses [Escape Analysis](https://github.com/golang/go/wiki/CompilerOptimizations#escape-analysis) to decide where a vairable will be allocated.

For programming, it's a good idea to keep heap memory as small as possible, since we want to reduce the garbage collection pressure. This repositry includes golang escapse examples/cases and no-escape example cases.

These examples could serve as a reference to write efficient golang code for high performance, i.e, avoid writing code like escape examples.

Result of this repository is based on go version: `go version go1.15.8 darwin/amd64`

To illustrate the basic idea, [inline optimization](https://github.com/golang/go/wiki/CompilerOptimizations#function-inlining) is mostly disabled in this repository. Feel free to remove line `//go:noinline` from code and run below commands to see the difference.

### To see the escape result for each file, run commands like below

```
cd pkg/escape
go build -gcflags="-m" case1.go
```

An example result is like below:

```
smiletrl@Rulins-MacBook-Pro escape % go build -gcflags="-m" case1.go
# command-line-arguments
./case1.go:13:13: inlining call to fmt.Println
./case1.go:16:13: inlining call to fmt.Println
./case1.go:28:2: moved to heap: title
./case1.go:13:13: num escapes to heap
./case1.go:13:13: []interface {} literal does not escape
./case1.go:16:13: emp escapes to heap
./case1.go:16:13: []interface {} literal does not escape
<autogenerated>:1: .this does not escape
```

Note: `num escapes to heap` means `num` lives out of its own function scope, which doesn't necessarily mean `num` is allocated to heap. `moved to heap: title` means variable `title` is allocated at heap.

For more verbose compile result, use double `-m` as `go build -gcflags="-m -m" case1.go`.

Any feedback or new prs are welcome :)
