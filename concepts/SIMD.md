# SIMD

## What is SIMD

Special processor instructions that perform some operation on more than 1 variable at a time.

Scalar

x[0] = x[0] + y[0]
x[1] = x[1] + y[1]
x[2] = x[2] + y[2]

SIMD

x = vector_add(x, y)

## What can SIMD do?

- Load/store
- Operations on Registers

Arithmetrics on 1 registers

Arithmetrics on 2 registers

Logical operations on 2 registers

Data type conversion

Variable permutations

Computations

## How to Use SIMD?

We don't do anything at all and leave everything to our compiler.

- Auto-vectorization
- Assembly commands
- Intrinsic functions
- Dedicated libraries

C and C++ compilers are so specialized that they are able to detect certain operations and perform so-called auto vectorization so they automatically translate our to operate on bigger chunks of data than just one. However that's not always possible the compiler however smart it can be it cannot predict what do we want to do with certain operations so in general we need to instruct the compiler on our own what do we want to do we as programmers have number of ways how we can interact with the processor to make it execute these special instructions on the very low level we can use this instructions directly so processors can be programmed using Assembly language and these SIMD instruction are special Assembly commands that we can use either in Assembly programs or in asm blocks inside c and c plus plus code. 

## Why is SIMD useful in DSP?

Now that we know that SIMD are special processor instructions that allow use to operate on more than one variable at once we may think why is it so useful in dsp so the first reason that comes to my mind is that a lot of digital signal processing algorithms are already designed with vectors in mind they use the vector notation and then somehow easily translated to SIMD implementation another reason is that in digital signal processing we often want to do the same operation over and over again but on different data that is why it is very beneficial that we can use SIMD to perform the same operation on more than one variable at once. Another reason is that in digital signal processing especially in audio processing we have what is called the block processing. So we receive samples in blocks and then we also output a block of samples. Using SIMD we can divide these blocks into smaller blocks and perform the dsp algorithm that we want on these smaller chunks instead of one by one, but these blocks are also often of length makes them easy in for vectorization.

**Summary**

- SIMD = operation on more than one variable at once.
- Different processors = different SIMD instructions.
- SIMD can make your DSP code significantly faster at the cost of

+ code complexity
+ portability
+ expert knowledge
