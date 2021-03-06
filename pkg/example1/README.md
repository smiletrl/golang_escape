## Example Description

 case 1-5 demonstrates pointers' referenced variables escaping from their own function scope, and then got allocated to heap.

It basically shows a pointer's value could move to heap easily if this pointer lives out of its caller function scope. With inline optimization, the compiler might figure out one pointer's value still lives within the caller stack. In this repository, inline optimization is disabled to force engineers to write efficiently without relying on inline optimization.

In this example, we will see the variables moved to heap explicitly with the compile tool: `go build -gcflags="-m -l"`. In the following examples, we will see variables also moved to heap without the compiler explicitly saying so.
