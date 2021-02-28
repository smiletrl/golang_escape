## Golang Escape Analysis

For programming, it's a good idea to keep heap memory as small as possible, since
we want to reduce the garbage collection presure. This repositry includes golang escapse cases and no escape cases.

Result of this repository is based on go version 15.8.
`go version go1.15.8 darwin/amd64`

### Command to run to see the escape result for each file

```
cd pkg/escape
go build -gcflags="-m" case1.go
```
