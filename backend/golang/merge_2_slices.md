slice: x, y

cần merge => z

1) Order is important: 

+ Copy style

If x >> y => Performant (overhead of zero value is small)

If x << y => Not Performant => Use Append style

+ Append style

Ram alloc over head

2) Order is not important:

+ Copy style 

Try to pass the longer slice as the first copy, so that the zero overhead will be saved. 

+ Append style

Try to pass the shorter slice as the first argument, so that some memory might be saved.
