## Example Description

This is an example of pointer's content escaping to heap.

case 1-5 demonstrates pointers' referenced variables escaping from their own function scope, and then got allocated to heap.

It basically shows a pointer's value could move to heap easily if this pointer lives out of its caller function scope. With inline optimization, the compiler might figure out one pointer's value still lives within the caller stack. In this repository, inline optimization is disabled to force engineers to write efficiently without relying on inline optimization.

Here we come to a good rule of thumb: <b>not use pointer unless it's really imperative to use pointer semantic</b>

In this example, we see the variables moved to heap explicitly with the compile tool: `go build -gcflags="-m -l"`. In the other examples, we will see variables also moved to heap without the compiler explicitly saying so.
