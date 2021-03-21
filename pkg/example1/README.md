## Example Description
 
This is an example of pointer's content escaping to heap.
 
case 1-5 demonstrates pointers' referenced variables escaping from their own function scope, and then got allocated to heap.
 
It basically shows a pointer's value could move to heap easily if this pointer lives out of its caller function scope. With inline optimization, the compiler might figure out one pointer's value still lives within the caller stack. In this repository, inline optimization is disabled to force engineers to write efficiently without relying on inline optimization.
 
 
In this example, we see the variables moved to heap explicitly with the compile tool: `go build -gcflags="-m -l"`. In the other examples, we will see variables also moved to heap without the compiler explicitly saying so.
 
## Conclusion
 
With pointer usage, if the referenced content/variable is going to be used out of [current stack frame](https://en.wikipedia.org/wiki/Call_stack), then the referenced content/variable must be allocated at heap. Otherwise, when the current stack frame has popped out, the referenced content/variable will be lost. Then the pointer (returned/defined from current stack frame, and used out of current stack frame) can't reach its value any more, which results in the [dangling pointer scenario](https://en.wikipedia.org/wiki/Dangling_pointer).
 
Here we come to a good rule of thumb: <b>Do not use pointer unless it's really imperative to use pointer semantic</b>
