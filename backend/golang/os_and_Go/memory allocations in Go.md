# Memory allocations in Go

To understand how memory allocations in Go works, we need to understand the types of memories in programming context, which are Stack and Heap. If you are familiar with typical memory representation of C, you must already be aware of these two terms.

## Stack vs Heap

**Stack:** The stack is a memory set aside as a scratch space for the execution of thread. When a function is called, a block is reserved on top of the stack for local variables and some bookkeeping data. This block of memory is referred to as a stack frame. Initial stack memory allocation is done by OS when the program is compiled. When a function is

https://dev.to/karankumarshreds/memory-allocations-in-go-1bpa
