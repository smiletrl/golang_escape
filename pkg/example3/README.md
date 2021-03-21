## Example Description
 
This is an example of variable escaping to heap without escape analysis explicitly saying so.

Run escape analysis command:

```
smiletrl@Rulins-MacBook-Pro example3 % go build -gcflags="-m -l" case.go
# command-line-arguments
./case.go:11:17: make([]employer, 10) does not escape
```

Above command doesn't indicate any variable escaping to heap.

Now run the benchmark memory test:

```
pkg: github.com/smiletrl/golang_escape/pkg/example3
cpu: Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz
BenchmarkSlice1-16    	276773486	         4.747 ns/op	       0 B/op	       0 allocs/op
BenchmarkSlice2-16    	17964181	        60.60 ns/op	      48 B/op	       1 allocs/op
PASS
ok  	github.com/smiletrl/golang_escape/pkg/example3	2.978s
```

It indicates one heap allocation for `getemployerSlice2()`. And the allocated memory is 48Byte.

We could look at the assumbly language for this code, and explore why.

```
smiletrl@Rulins-MacBook-Pro example3 % go tool compile -S case.go
"".getemployerSlice1 STEXT nosplit size=14 args=0x0 locals=0x0 funcid=0x0
	0x0000 00000 (case.go:10)	TEXT	"".getemployerSlice1(SB), NOSPLIT|ABIInternal, $0-0
	0x0000 00000 (case.go:10)	FUNCDATA	$0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (case.go:10)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
	0x0000 00000 (case.go:10)	XORL	AX, AX
	0x0002 00002 (case.go:12)	JMP	7
	0x0004 00004 (case.go:12)	INCQ	AX
	0x0007 00007 (case.go:12)	CMPQ	AX, $10
	0x000b 00011 (case.go:12)	JLT	4
	0x000d 00013 (case.go:12)	RET
	0x0000 31 c0 eb 03 48 ff c0 48 83 f8 0a 7c f7 c3        1...H..H...|..
"".getemployerSlice2 STEXT size=241 args=0x0 locals=0x70 funcid=0x0
	0x0000 00000 (case.go:23)	TEXT	"".getemployerSlice2(SB), ABIInternal, $112-0
	0x0000 00000 (case.go:23)	MOVQ	(TLS), CX
	0x0009 00009 (case.go:23)	CMPQ	SP, 16(CX)
	0x000d 00013 (case.go:23)	PCDATA	$0, $-2
	0x000d 00013 (case.go:23)	JLS	231
	0x0013 00019 (case.go:23)	PCDATA	$0, $-1
	0x0013 00019 (case.go:23)	SUBQ	$112, SP
	0x0017 00023 (case.go:23)	MOVQ	BP, 104(SP)
	0x001c 00028 (case.go:23)	LEAQ	104(SP), BP
	0x0021 00033 (case.go:23)	FUNCDATA	$0, gclocals·69c1753bd5f81501d95132d08af04464(SB)
	0x0021 00033 (case.go:23)	FUNCDATA	$1, gclocals·edcc1c706859af29673c95fedcaaa670(SB)
	0x0021 00033 (case.go:25)	MOVQ	$0, "".e+64(SP)
	0x002a 00042 (case.go:25)	XORPS	X0, X0
	0x002d 00045 (case.go:25)	MOVUPS	X0, "".e+72(SP)
	0x0032 00050 (case.go:25)	MOVUPS	X0, "".e+88(SP)
	0x0037 00055 (case.go:26)	LEAQ	go.string."adam"(SB), AX
	0x003e 00062 (case.go:26)	MOVQ	AX, "".e+64(SP)
	0x0043 00067 (case.go:26)	MOVQ	$4, "".e+72(SP)
	0x004c 00076 (case.go:27)	MOVQ	$23, "".e+80(SP)
	0x0055 00085 (case.go:28)	LEAQ	go.string."ceo"(SB), AX
	0x005c 00092 (case.go:28)	MOVQ	AX, "".e+88(SP)
	0x0061 00097 (case.go:28)	MOVQ	$3, "".e+96(SP)
	0x006a 00106 (case.go:30)	LEAQ	type."".employer(SB), AX
	0x0071 00113 (case.go:30)	MOVQ	AX, (SP)
	0x0075 00117 (case.go:30)	MOVUPS	X0, 8(SP)
	0x007a 00122 (case.go:30)	MOVQ	$0, 24(SP)
	0x0083 00131 (case.go:30)	MOVQ	$1, 32(SP)
	0x008c 00140 (case.go:30)	PCDATA	$1, $1
	0x008c 00140 (case.go:30)	CALL	runtime.growslice(SB)
	0x0091 00145 (case.go:30)	MOVQ	40(SP), AX
	0x0096 00150 (case.go:30)	PCDATA	$0, $-2
	0x0096 00150 (case.go:30)	CMPL	runtime.writeBarrier(SB), $0
    ...
```

The interesting line is `0x008c 00140 (case.go:30)	CALL	runtime.growslice(SB)`. It means `append()` has tried to grow the slice.

Line 24 has declared one empty slice with length & capacity to be zero. Line 30 is going to expand the slice length to hold the new slice content at line 25. So [`runtime.growslice`](https://github.com/golang/go/blob/release-branch.go1.16/src/runtime/slice.go#L125) has been called. This function then invokes [`mallocgc`](https://github.com/golang/go/blob/f39c4deee812b577ffb84b78e62ce4392d2baeb1/src/runtime/malloc.go#L905), which will make heap allocation.

You may find more append details at [this great blog](https://blog.golang.org/slices-intro) and [Effective go](https://golang.org/doc/effective_go#slices).

## Conclusion

Heap allocation might happen when escape analysis has given up `analysis`. Citing from [randall77](https://github.com/randall77)

> Escape analysis is only concerned with flow of data. emps does not escape here, it only flows to itself.
There's a related but distinct question of heap allocation vs. stack allocation. Generally to stack allocate a variable, it needs to not escape and have a size known at compile time. The allocation that happens here (in growslice) fails the second test, not the first.
