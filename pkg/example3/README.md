## Example Description
 
This is an example of [`slice`](https://blog.golang.org/slices) escaping to heap.
 
Looking at the code, both `var emps [1e5]employer` and `var emps = make([]employer, 1e5)` are local variables, and they are not used in func return either. The difference is the first one being array, while the second one being slice.
 
In example 2, we know array size `1e5` for `var emps [1e5]employer` is allocated at stack. This example will see how slice behaves.
 
Run command, and result is like:
 
```
smiletrl@Rulins-MacBook-Pro example3 % go test -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/smiletrl/golang_escape/pkg/example3
BenchmarkCase1Array-16           501       2336946 ns/op           0 B/op          0 allocs/op
BenchmarkCase1Slice-16             4     267056822 ns/op    400589504 B/op       111 allocs/op
```
 
We will see no memory for `caseArray()` is allocated at heap, but `caseSlice()` does allocate variable `var emps = make([]employer, 1e5)` to heap for 101 times during each operation.
 
Slice itself includes [one pointer to the array content](https://golang.org/pkg/reflect/#SliceHeader), kind of similar to the string we discussed at example 2. In this case, slice size should include the backed array's size and itself (24 Bytes for 3 `int` type from the `SliceHeader` struct). Basically, this slice size `var emps = make([]employer, 1e3)` is relatively the same as array `var emps [1e5]employer`.
 
Now try to change the slice length from `1e5` to `1e4` and `1e3` for slice `var emps = make([]employer, 1e5)`.
 
```
//go:noinlinex
func getemployerSlice() {
   var emps = make([]employer, 1e3)
   for i := 0; i < 1e3; i++ {
       e := employer{
           Name:  "adam",
           Age:   23,
           Title: "ceo",
       }
       emps[i] = e
   }
}
```
 
Run above command `go test -bench=. -benchmem` again, and we will see slice gets allocated to stack again with `1e3`.
 
Now try to change the slice value assignment to `append`, and change the slice length to `1` in the loop.
 
```
//go:noinline
func getemployerSlice() {
   var emps = make([]employer, 1)
   for i := 0; i < 1; i++ {
       e := employer{
           Name:  "adam",
           Age:   23,
           Title: "ceo",
       }
       emps = append(emps, e)
   }
}
```
 
Run above command `go test -bench=. -benchmem` again, and this time, we will see slice always gets allocated to heap even with length `1`.
 
## Conclusion
 
1. When a slice is used with `append`, this slice may dynamically grow pretty large, which is not identical to an array with a fixed size. So it's reasonable to allocate the slice variable to [heap - dynamic memory allocation](https://en.wikipedia.org/wiki/Memory_management#DYNAMIC).
 
2. When a slice is not using `append`, it will be determined by its size.
a. If the size is small, then it will be allocated to stack.
b. If the size is large, then it will be allocated to the heap.
c. With the same size as the array, slice has more chance to be allocated to heap. This comes to a similar question in example 2, how `large` is large. We will cover this in another topic.

## Run GC command
 
To verify the variable is really allocated to heap, we might also run gc debug command, like this
 
```
smiletrl@Rulins-MacBook-Pro example3 % go build .                   
smiletrl@Rulins-MacBook-Pro example3 % GODEBUG="gctrace=1" ./example3
```
 
If we see results like below, it means new memory from the heap is being garbage collected. If we see nothing like below, it means no memory has been allocated to heap. For more details, visit [Golang GC](https://github.com/smiletrl/golang_gc/tree/master/cmd/example1).
 
```
gc 1 @0.004s 0%: 0.012+0.19+0.019 ms clock, 0.19+0.079/0.13/0.10+0.31 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
gc 2 @0.007s 1%: 0.009+0.19+0.021 ms clock, 0.15+0.060/0.052/0.15+0.34 ms cpu, 4->4->0 MB, 5 MB goal, 16 P
...
```
 
Play with this command using different slice length, and `append`. See if we can get the consistent result like command `go test -bench=. -benchmem` does.
 
Try to change `main()` code from invoking slice to array, like below. Run the gc debug command again to verify the array variable is really not being allocated to heap.
 
```
func main() {
   start := time.Now()
   caseArray()
   fmt.Printf("druation is: %+vs\n", time.Now().Sub(start).Seconds())
}
 
```
