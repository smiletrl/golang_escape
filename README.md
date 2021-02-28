## Golang Escape Analysis

According to [Golang FAQ](https://golang.org/doc/faq#stack_or_heap), it's not very 
clear when a variable will be allocated at heap.

For programming, it's a good idea to keep heap memory as small as possible, since
we want to reduce the garbage collection pressure. This repositry includes golang escapse cases and no-escape cases.

Taking these cases as a guide to write efficient golang code for high performance.

Result of this repository is based on go version: `go version go1.15.8 darwin/amd64`

### To see the escape result for each file, run commands like below

```
cd pkg/escape
go build -gcflags="-m" case1.go
```

Any feedback or new prs are welcome :)
