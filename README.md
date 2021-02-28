## Golang Escape Analysis

According to [Golang FAQ](https://golang.org/doc/faq#stack_or_heap), it's not very 
clear when a variable will be allocated at heap.

Golang uses [Escape Analysis](https://github.com/golang/go/wiki/CompilerOptimizations#escape-analysis) to decide where a vairable should be allocated.

For programming, it's a good idea to keep heap memory as small as possible, since
we want to reduce the garbage collection pressure. This repositry includes golang escapse cases and no-escape cases.

Taking these cases as a guide to write efficient golang code for high performance.

Result of this repository is based on go version: `go version go1.15.8 darwin/amd64`

To illustrate the basci idea, [inline optimization](https://github.com/golang/go/wiki/CompilerOptimizations#function-inlining) is mostly disabled in this repository. Feel free to remove line `//go:noinline` from code and run below commands to see the difference.

### To see the escape result for each file, run commands like below

```
cd pkg/escape
go build -gcflags="-m" case1.go
```

This repository mostly disables inline optimization 

Any feedback or new prs are welcome :)
