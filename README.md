# Motivation

When read data from io.Reader, we make a `[]byte` first, then read data, then use data, then drop the `[]byte`. This will generate huge GC, 
and influence application performance.

In order to reduce GC, we can build a memory pool. Get a `[]byte` from pool, after used, then back the `[]byte` to pool.

# Goal

## Memory Effective

Don't wast too much memory. For example, want a `[]byte` with capacity 1025, get a `[]byte` with capacity 2048, it's a memory wast.

## Easy Understand

Easy understand over performance. In order to improve 2ns, make the code hard understand, is not a good idea. In fact, 2ns is ignorable 
as a real transaction will take 20ms+.

# Design

## Bucket

Capacity of `[]byte` been divided into 5 bucket. Each bucket is responsible for capacity between `(min, max]`.

## Alignment

Each bucket has a alignment, it define how to align `[]byte`. We can't return exactly requested capacity, we will
return the `[]byte` which capacity bigger or equal than requested capacity. For example we request 9 byte, the alignment
is 8 byte, then we will return 16 byte. 

## Class

Each bucket has N class. Each class is a linked list with same size node. For example, node of class 1 has one alignment,
node of class 2 has one 2*alignment, node of class n has one n*alignment.

# Example

## Request a buffer

Buffer with size as 0:

`
buf, err := gmem.Malloc(1024)
`

Buffer with size as given:

`
buf, err := gmem.MallocWithSize(1024, 1024)
`

Free a buffer:

`
buf := make([]byte,0,256)
gmem.Free(buf)
`

# Benchmark

`
BenchmarkMallocWithSize1KB-8            22390994                52.1 ns/op            32 B/op          1 allocs/op
BenchmarkMallocWithSize128KB-8          22735800                53.5 ns/op            32 B/op          1 allocs/op
BenchmarkMallocWithSize1MB-8            21597294                51.7 ns/op            35 B/op          1 allocs/op
`

# Reference

- https://github.com/google/tcmalloc/blob/master/tcmalloc/common.h


