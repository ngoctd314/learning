# The Go Memory Model

https://go.dev/ref/mem

## Introduction

The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.

### Advice

Programs that modify data being simultaneously accessed by multiple goroutines must serialize such access.

To serialize access, protect the data with channel operations or other synchronization primitives such as those in the sync and sync/atomic packages.

**If you must read the rest of this document to understand the behavior of your program, you are being too clever.**

**Don't be clever.**

### Informal Overview

Go approaches its memory model in much the same way as the rest of the language, aiming to keep the semantics simple, understandable, and useful. This section gives a general overview of the approach and should suffice for most programmers. The memory model is specified more formally in the next section.

A data race is defined as a write to a memory location happening concurrently with another read or write to that same location, unless all the accesses involved are atomic data accesses as provided by the sync/atomic package. As noted already, programmers are strongly encouraged to use appropriate synchronization to avoid data races. In the absense of data races, Go programs behave as if all the goroutines were multiplexed onto a single processor. This property is sometimes referred to as DRF-SC: data-race-free programs execute in a sequentailly consistent manner.

While programmers should write Go programs without data races, there are limitations to what a Go implementation can do in response to a data race. An implementation may always react to a data race by reporting the race and terminating the program. Otherwise, each read of a single-word-sized or sub-word-sized memory location must observe a value actually written to that location (perhaps by a concurrent executing goroutine) and not yet overwritten. These implementation constraints make Go more like Java or Javascript, in that most races have a limited number of outcomes.

## Memory Model

The following formal definition of Go's memory model closely follows the approach presented by Hans-J.Boehm. The definition of data-race-free programs and the guaranteed of sequential consistency for race-free programs are equivalent to the ones in that work.

A memory operation is modeled by four details:

- its kind, indicating whether it is an ordinary data read, an ordinary data write, or a synchronization operation such as an atomic data access, a mutex operation, or a channel operation.
- its location in the program.
- the memory location or variable being accessed, and.
- the values read or written by the operation.

Some memory operations are read-like, including read, atomic read, mutex read, and channel receive. Other memory operations are write-like, including write, atomic write, mutex unlock, channel send, and channel close. Some, such as atomic compare-and-swap, are both read-like and write-like.

A goroutine execution is modeled as a set of memory operations executed by a single goroutine.

**Requirement 1:** The memory operations in each goroutine must correspond to a correct sequential execution of that goroutine, given the values read from and written to memory. That execution must be consistent with sequenced before relation.

## Implementation Restrictions for Programs Containing Data Races

## Synchronization

### Initialization

### Goroutine creation

### Goroutine destruction

### Channel communication

### Locks

### Once

### Atomic Values

### Finalizers

### Additional Mechanisms

## Incorrect synchronization

## Incorrect compilation

## Conclusion
