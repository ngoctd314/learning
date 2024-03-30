# roaring

## When should you use a bitmap?

Sets are a fundamental abstraction in software. They can be implemented in various ways, as hash sets, as trees, and so forth. In databases and search engines, sets are often an integral part of indexes. For example, we may need to maintain a set of all documents or rows (represented by numerial identifier) that satisfy some property. Besides adding or removing elements from the set, we need fast functions to compute the intersection, the union, the difference between sets, and so on.

To implement a set of integers, a particular appealing strategy is the bitmap (also called bitset or bit vector). Using n bits, we can represent any set made of the integers from the range [0, n): the ith bit is set to one if integer i is present in the set.

However, a bitset, even a compressed one is not always applicable. For example, if you have 1000 random-looking integers, then a simple array might be the best representation. We refer to this case as the "sparse" scenario.

## When should you use compressed bitmaps?

An uncompressed BitSet can use a lot of memory. For example, if you take a BitSet and set the bit at position 1000000 to true and you have just over 100kB. That is over 100kB to store the position of one bit. This is wasteful even if you do not care about memory: suppose that you need to compute the intersection between this BitSet and another one that has a bit at position 1000001 to true, then you need to go through all these zeros, whether you like it or not. That can become very wasteful.

This being said, there are definitely cases where attempting to use compressed bitmaps is wasteful. For example, if you have a small universe size. If your bitmaps represent sets of integers from [0, n) where n is small (n = 64 or n = 128). If you can use uncompressed BitSet and it does not blow up your memory usage, the compressed bitmaps are probably not useful to you. In fact, if you do not need compression, then a BitSet offers remarkable speed.

## How does Roaring compares with the alternatives?

Unfortunately, when the range of possible values (n) is too wide, bitmaps can be too large to be practical. For example, it might be impractical to represent the set {1,2^31} using a bitset. For this reason, we often use compressed bitmaps.

Though there are many ways to compress bitmaps, several systems rely on an approach called Roaring include Elasticsearch... In turn, these systems are in widespread use, Roaring is used in machine learning, data visualization, in natural language processing.

Roaring partitions the space [0, 2^23) into chunks consisting of ranges of 2^16 integers ([0, 2^16], [2^16, 2^17], ...). For a value in the set, its least significant sixteen bits are stored in a container corresponding to its chunk (as determined by its most significant sixteen bits), using one of three possible container types:

- bitset containers made of 2^16 bits or 8kB;
- array containers made of up to 4096 sorted 16-bit integers.

At a high level, we can view a Roaring bitmap as a list of 16-bit numbers (corresponding to the most-significant 2B of the values present in the set), each of which is coupled with a reference to a container holding another set of 16-bit numbers 

## Integer-set data structures


## Introduction

We build a new Roaring hybrid that combines uncompressed bitmaps, packed arrays and RLE compressed segments. The result is a new Roaring format that compresses better.

Besides adding or removing elements from the set, we need fast functions to compute the intersection, union, the difference between sets, and so on.

Conventional bitmaps are only applicable when the cardinality of the set (|S|) is relatively large compared to the universe size (n), e.g |S|> n/64. They are also suboptimal when the set is made of consecutive values (e.g, S = {1,2,3,4,...99,100})

When processing such RLE-compressed formats, one may need to read every compressed word to determine whether a value is in the set. Moreover, computing the intersection or union between two bitmaps B1 and B2 has complexity O(|B1| + |B2|) where |B1| and |B2| are the compressed sizes of the bitmaps. This complexity is worse than that of a hash set, where we can compute an intersection with an expected-time complexity of O(min(|S1|, |S2|), where |S1| and |S2| are the cardinalities of the sets. Indeed, it suffices to iterate over the smallest sets, and for each value, check whether it is in the larger set. Similarly, we can compute an in-place union, where the result is store in the largest hash set, by inserting all of the values from the small set in the large set, in expected time O(min(|S1|, |S2|). It is comparatively more difficult to compute in-place unions with RLE-compressed bitmap may require a complete scan of the entire bitmap in O(|B|) time. Such as scan can be hundreds of times slower than checking for the presence of a value in an uncompressed bitmap or hash map.

Roaring partitions the space [0, n) into chunks or 2^16 integers [0, 2^16), [2^16, 2*2^16), ... ~ 1526 chunks cho 1e8.

Each set value is stored in a container corresponding to its chunk. Roaring stores dense and sparse chunks differently. Dense chunks (containing more than 4096 integers) (0) are stored using conventional bitmap containers (made of 2^16 bits or 8kB) whereas sparse chunks use smaller containers made of packed sorted arrays 16-bit integers. All integers in a chunk share the same 16 most-significant bits. The containers are sorted in an array along with the most-significant bits.

**Though we refer to a Roaring bitmap as a bitmap, it is a hybrid data structure, combining uncompressed bitmaps with sorted arrays.**

Roaring allows fast random access. To check for the presense of a 32-bit integer x, we seek the container corresponding to the 16  most significant bits to x, using a binary search. If a bitmap container is found we check the corresponding bit (at index x mod 2^16); if an array container is found, we use a binary search. Similarly, we can compute the intersection between two Roaring bitmaps without having to access all of the data. Indeed, suppose that we have a Roaring bitmap B1 containing only a few values, which all fall in the range [0, 2^16). This implies is uses an array container. Suppose we have another Roaring bitmap B2 over the same range range but containing may values; it can only use a bitmap container. In that case, computing the intersection between B1 and B2 can be done in time O(|B1|), since it suffices to iterate over the set values of B1 and check whether the corresponding bit is set in the bitmap container B2. Moreover, to compute the in-place union of B1 and B2, whether the result is stored in the bitmap (B2), it suffices to iterate through the value of B1 and set the corresponding bits in B2 to 1, in time O(|B1|)

Operations between containers must produce new containers that are either arrays or bitmap containers. Because converting between container types may be expensive, we found it useful to predict the container type as part of the computation. For example, if we must compute the union between two array containers such that the sum of their cardinalities exceeds 4096, we preemptively create a bitmap container and store the result of the union. Only if the resulting cardinality falls below 4096 do we convert the result back to an array container.

However, the original Roaring has a limitation in some scenarios because it does not compress long runs of values. Indeed, given a bitset made of a few long runs.

Roaring bitmap implementation is a key-value data structure where each key-value pair represents the set S of all 32-bit integers that share the most significant 16 bits. The key is made of the shared 16 bits, whereas the value is a container storing the remaining 16 least significant bits for each member of S. No container ever uses much more than 8kB of memory.

The key-value store is implemented as two arrays: an array of packaged 16-bit values representing the keys and an array of containers. The arrays expand dynamically in a standard manner when there are insertions

In a system such as Druid, the bitmaps are created, stored on disk and then memory-mapped as need. When we serialize the bitmaps, we interleave with the 16-bit keys, the cardinalities of the corresponding containers: cardinalities are stored as 16-bit values. If need, we also use an uncompressed bitmap containing at least one bit per container to indicate whether the corresponding container is a run container.

- A bitmap container is an object made of 1024 64-bit words (using 8kB) representing an uncompressed bitmap, able to store all sets of 16-bit integers. The container can be serialized as an array of 64-bit words. We also maintain a counter to record how many bits are set to 1.


```go
type bitmapContainer struct {
	cardinality int
	bitmap      []uint64
}
```

In some cases, the range of value might not cover the full range [0, 2^16) and a smaller bitmap might sufficient - thus improving compression. However, the bitmap containers would then need to grow and shrink dynamically. For simplicity, we use fixed-size bitmap containers.

Modern processors has bit-count instructions - such as popcnt for x64 processors and cnt for the 64-bit ARM architecture  - that can do this count using sometimes as little as a single clock cycle.

- An array container is an object containing a counter keeping track of the number of integers, followed by a packed array of sorted 16-bit unsigned integers. It can be serialized as an array of 16-bit values.

```go
type arrayContainer struct {
	content []uint16
}
```

- Our new addition, the run container, is made of a packed array of pairs of 16-bit integers. The first value of each pair represents a starting value, whereas the second value is the length of a run. For example, we would store the values 11, 12, 13, 14, 15 as the pair 11, 4 where 4 means that beyond 11 itself, there are 4 contiguous values that follow. This is packed array, we need to maintain the number of runs stored in the packed array. Like the array container, the run container is stored in a dynamic array. During serialization, we write out the number of runs, followed by the corresponding packed array.

When starting from an empty Roaring bitmap, if a value is added, and array container is created. When inserting a new value in an array container, it the cardinality exceeds 4096, then the container is transformed into a bitmap container. On the other hand, if a value is removed from a bitmap container so that its size falls to 4096 integers, then it is transformed into an array container. When ever a container becomes empty, it is removed from the top-level key-value structure along with the corresponding key.

```txt
> 4096 -> bitmapContainer
< 4096 -> arrayContainer
```
Thus, when first creating a Roaring bitmap, it is usually made of array and bitmap containers.  Runs are not compressed. Upon request, the storage of the Roaring bitmap can be optimized using the runOptimize function. This triggers a scan through the array and bitmap containers that converts them, it helpful, to run containers. In a given application, this might be done prior to storing the bitmaps as immutable objects to be queried. Run containers may also arise from calling a functions to add a range of values.

To decide the best container type, we are motived to minimize storage. In serialized form, a run container uses 2 + 4r bytes give r runs, a bitmap container always uses 8192 bytes and an array container uses 2c + 2 bytes, where c is the cardinality.

- All array containers are such that they use no more space then they would as a bitmap container: they contain no more than 4096 values.
- Bitmap containers use less space than they would as array containers: they contain more than 4096 values.
- A run container is only allowed to exist if it is smaller than either the array container or the bitmap container the could equivalently store the same values. If the run container has cardinality greater than 4096 values, then it must contain no more than [(8192-2)/4] = 2047 runs. If the run container has cardinality no more than 4096, then the number of runs must be less than half the cardinality. 


## Logical operators

### Union and intersection

There are many necessary logical operations, but we present primarily the union and intersection. They are most often used, and the most likely operations to cause performance bottlenecks. An important algorithm for out purposes is the galloping (also called exponential intersection) the intersection between two sorted arrays of size c1, c2. It has complexity O(min(c1, c2)log(max c1, c2)). In this approach, we pick the next available integer i from the smaller array and seek an integer at least as big in the larger array, looking first at the next available value, then looking twice as far, and so on, until we find an integer that is not smaller than i.

A galloping search makes repeated random accesses in a container, and it could therefore cause expensive cache misses. However, in our case, the potential problem is mitigated by the fact that all our containers fit in CPU cache.

Intersections between two input Roaring bitmaps start by visiting the keys from both bitmaps, starting from the beginning. If a key is found in both input bitmaps, the corresponding containers are intersected and the result (if non-empty) is added to the output.

Unions between Roaring data structures are handled in the conventional manner: we iterate through the keys in sorted order; if a key is in both input Roaring bitmaps, we merge the two containers, and the result to the output and advance in the two bitmaps. When one bitmap runs out of keys, we append all the remaining content of the other bitmap to the output.

Though we do not use this technique, instead of cloning the containers during unions, we could use a copy-on-write approach whereas a reference to container is stored and used, and a copy is only made if an attempt is made to modify the container further. This approach can be implemented by adding a bit vector containing one bit per container. Initially, this bit is set to 0

We first briefly review the logical operations between bitmap and array containers, referring the reader to Chambi et al. [7] for algorithmic details.

Bitmap vs Bitmap: To compute the intersection between two bitmaps, we first compute the
cardinality of the result using the bitCount function over the bitwise AND of the
corresponding pairs of words. If the intersection exceeds 4096, we materialize a bitmap
container by recomputing the bitwise AND between the words and storing them in a new
bitmap container. Otherwise, we generate a new array container by, once again, recomputing
the bitwise ANDs, and iterating over their 1-bits. We find it important to first determine the
right container type as, otherwise, we would sometimes generate the wrong container and
then have to convert it—an expensive process. The performance of the intersection operation
between two bitmaps depends crucially on the performance of the bitCount function.
A union between two bitmap containers is straightforward: we execute the bitwise OR
between all pairs of corresponding words. There are 1024 words in each container, so
1024 bitwise OR operations are needed. At the same time, we compute the cardinality of
the result using the bitCount function on the generated words

A union between two bitmap containers is straightforward: we execute the bitwise OR
between all pairs of corresponding words. There are 1024 words in each container, so
1024 bitwise OR operations are needed. At the same time, we compute the cardinality of
the result using the bitCount function on the generated words.

Bitmap vs Array: The intersection between an array and a bitmap container can be computed
quickly: we iterate over the values in the array container, checking the presence of each 16-bit
integer in the bitmap container and generating a new array container that has as much capacity
as the input array container. The running time of this operation depends on the cardinality of
the array container. Unions are also efficient: we create a copy of the bitmap and iterate over
the array, setting the corresponding bits.

serialize size vs real size

+ remove header => optimize disk

Bitmap vs Bitmap: To compute the intersection between two bitmaps, we first compute the cardinality of the result using the bitCount function over the bitwise AND of the corresponding pairs of words. If the intersection exceeds 4096, we materialize a bitmap container by recomputing the bitwise AND between the words and storing them in a new bitmap container. Otherwise, we generate a new array container by, once again, recomputing the bitwise ANDs, and iterating over their 1-bits. We find it important to first determine the right container type as, otherwise, we would sometimes generate the wrong container and then have to convert it-an expensive process. The performance of the intersection operation between two bitmaps depends crucially on the performance of the bitCount function. 

A union between two bitmap containers is straightforward: we execute the bitwise OR between all pairs of corresponding words. There are 1024 words in each container, so 1024 bitwise OR operations are need. At the same time, we compute the cardinality of the result using the bitCount function on the generated words.

Compressed bitmap indexes are used in systems such as Git or Oracle to accelerate queries. They represent
sets and often support operations such as unions, intersections, differences, and symmetric differences. Several
important systems such as Elasticsearch, Apache Spark, Netflix’s Atlas, LinkedIn’s Pinot, Metamarkets’
Druid, Pilosa, Apache Hive, Apache Tez, Microsoft Visual Studio Team Services and Apache Kylin rely
on a specific type of compressed bitmap index called Roaring. We present an optimized software library
written in C implementing Roaring bitmaps: CRoaring. It benefits from several algorithms designed for
the single-instruction-multiple-data (SIMD) instructions available on commodity processors. In particular,
we present vectorized algorithms to compute the intersection, union, difference and symmetric difference
between arrays. We benchmark the library against a wide range of competitive alternatives, identifying
weaknesses and strengths in our software. Our work is available under a liberal open-source license.

## Summary

Compressed bitmap indexes are used in systems such as Git or Oracle to accelerate queries. They represent sets and often support operations such as unions, intersections, differences, and symetric differences. Several important systems such as Elasticsearch, Apache Spark... rely on a specific type of compressed bitmap index called Roaring. It benifits from several algorithms designed for the single-instruction-multiple-data (SIMD) instructions available on commodity processors. In particular, we present vectorized algorithms to compute the intersection, union, difference and symmetric difference between arrays.

Keywords: bitmap indexes, vectorization; SIMD intructions; database indexes; Jaccard index.

## Introduction

Comtemporary computing hardware offers performance opportunities through improved parallelism, by having more cores and better single-instruction-multiple-data (SIMD) instructions. Meanwhile, software indexes often determine the performance of big-data applications. Efficient indexes not only improve latency and throughput, but they also reduce energy usage.

Indexes are often made of sets of numerical identifiers (stored as integers). For instance, inverted indexes map query terms to document identifiers in search engines, and convertional database indexes map column values to record identifiers. We often need efficient computation of the intersection (A ∩ B), the union (A ∪ B), the difference (A \ B), or the symetric difference ((A \ B)  ∪ (B \ A)) of these sets.

The bitmap (or bitset) is a time-honored strategy to represent sets of integers concisely. Given a universe of n possible integers, we use a vector of n bits to represent any one set. On a 64-bit processor, [n/64] inexpensive bitwise operations suffice to compute set operations between two bitmaps:

- the intersection corresponds to the bitwise AND;
- the union corresponds to the bitwise OR;
- the difference corresponds to the bitwise ANDNOT;
- the symmetric difference corresponds to the bitwise XOR;

Unfortunately, when the range of possible values (n) is too wide, bitmaps can be too large to be practical. For example, it might be impractical to represent the set {1, 2^31} using a bitset. For this reason, we often use compressed bitmaps.

Though there are many ways to compress bitmaps, several systems rely on an approach called Roaring including Elasticsearch. In turn, the systems are in widespread use: Elasticsearch provides the search capabilities of Wikipedia. Additionally, Roaring is use in machine learning, data visualization, in natural language processing, in geographical information systems.

Roaring partitions the space [0, 2^32) into chunks consisting of ranges of 2^16 integers ([0x2^16, 1x2^16), [1x2^16, 2x2^17), ). For a value in the set, its least significant sixteen bits are stored in a container corresponding to its chunk, using one of three possible container types:

- bitset containers made of 2^16 bits or 8kB
- array containers made up to 4096 sorted 16-bit integers
- run containers made of a series of sorted <s, l> pairs indicating that all integers in the range [s, s+l] are present.

At a high level, we can view a Roaring bitmap as a list of 16-bit numbers (corresponding to the most-significant 2B) of the values present in the set, each of which is coupled with a reference to a container holding another set of 16-bit numbers corresponding to the least significant 2B of the elements sharing the same prefix.

We dynamically pick the container type to minimize memory usage. For example, when intersecting two bitset containers, we determine whether the result is an array of a bitset container on-the-fly. As we add or remove values, a container's type might change. No bitset container may store fewer than 4097 distinct values; no array container may store more than 4096 distinct values. If a run container has more than 4096 distinct values then it must have no more than 2047 runs, otherwise the number of runs must be less than half of the number of distinct values.

Roaring offers logarithmic-time random access: to check for the presence of a 32-bit integer, we seek the container corresponding to the sixteen most-significant bits using a binary search. If this prefix is not in the list, we know that the integer is not present. If a bitmap container is found, we check the corresponding bit; if an array or run container is found, we use a binary search.

Two main contributions:

- We present several non-trivial algorithm optimizations. In table 1, we show that a collection of algorithms exploiting single-instruction-multiple-data (SIMD) instructions can enhance the performance of a data structure like Roarings in some cases, above and beyond that state-of-the-art optimizing compilers can achieve. To our knowledge, it is the first work to report on the benefits of advanced SIMD-based algorithms for compressed bitmaps.

Though the approach we use to compute array intersections using SIMD instructions in is not new, our work on the computation of the union, difference and symmetric difference of arrays using SIMD instructions might be novel and of general interest.

- We benchmark our C library against a wide range of alternatives in C and C++. Our results provide guidance as to the strengths and weaknesses of our implementation.

inputs: [2^16+1,  2^16+2, 2^16+3, ..2^16 + 99] => run container

inputs: [1,2,4,5,6,7,8,...4096] => bitmap container

## Integer set data structures

The simplest way to represent a set of integers is as a sorted array, leading to an easy implementation. Querying for the presense of a given value can be done in logarithmic time using a binary search. Efficient set operations are already supported in standard libraries. We can compute the intersection, union, difference, and symmetric between two sorted arrays in linear time: O(n1 + n2) where n1 and n2 are the cardinalities of the two arrays. The intersection and difference can also be computed in time O(n1logn2), which is advantageous when one array is small.

|containers|optimization|section|prior work|
|-|-|-|-|
|bitset -> array|converting bitsets to arrays|$3.1|-|
|bitset + array|setting, flipping or resetting the bits of a bitset at indexes specified by an array, with an without cardinality tracking|$3.2|-|
|bitset|computing the cardinality using a vectorized Harley-Seal algorithm|$4.1.1|[24]|
|bitset + bitset|computing AND/OR/XOR/ANDNOT between two bitsets with cardinality using a vectorized Harley-Seal algorithm|$4.1.2|[24]|
|array+array|computing the intersection between two arrays using a vectorized algorithm|$4.2|[22,23]|
|array+array|computing the union between two arrays using vectorized algorithm|$4.3|-|
|array+array|computing the difference between two arrays using a vectorized algorithm|$4.4|-|
|array+array|computing the symmetric difference between two arrays using a vectorized algorithm|$4.5|-|

**Compressed bitset**

A bitset (or bitmap) has both the benefits of the hash set (constant-time random access) and of a sorted array (good locality), but suffers from impractical memory usage when the universe size is too large compared to the cardinality of the sets. So we compressed bitmaps. Though there are alternatives [25], the most popular bitmap compression techniques are based on the word-aligned RLE compression model inherited from Oracle (BBC [26]): WAH [27], Consice [28], EWAH [29], COMPAX [31], VLC [32], VAl-WAH [33], among consecutive bits,. Using W0=8, he uncompressed bitmap 0000000001010000 fillWord(W0)dirtyWord(01010000). Techniques such as BBC, WAH or EWAH use special formats, it may be necessary to read every compressed word to determine whether it indicates a sequence of fill words, or a dirty word. EWAH was found to have superior performance to WAH and Concise [33] in an extensive comparison. A major limitation of formats like BBC, WAH, Concise or EWAH is that random access is slow. A major limitation of formats like BBC, WAH, Consice or EWAH is that random access is slow. That is, to check whether a value is present in a set can take linear time O(n), where n is the compressed size of the bitmap. When intersecting a small set with a more voluminous one, there formats may have suboptimal performance.

Use Roaring for bitmap compresion whenever possible.

**BitMagic**

Mixing container types in one data structure is not unique to Roaring.

The BitMagic library is probably the most closely related data structure [41] with a publicly available implementation. Like Roaring, it is a two-level data structure similar to RIDBit. There are, however, a few differences between Roaring and BitMagic:

- In BitMagic, special pointer values are used to indicate a full (containing 2^16 values) or empty container; Roaring does not need to mark empty containers (they are omitted) and it can use a run container to represent a full container efficiently.
- Roaring relies on efficient heuristics to generate memory-efficient container. For example, when computing the union between two array containers, we guess whether the output is likely to be more efficiently represented as a bitset container, os opposed to an array. BitMagic does not attempt to optimize the resulting containers: e.g., the intersection of two bitset containers is a bitset container, even when another container type could be more economical.

Roaring keeps a cardinality counter updates for all its bitset containers. BitMagic keeps track of the set cardinality, but not at the level of the bitset containers.

- Roaring uses a key-container array, with one entry per non-empty container; BitMagic's top-level array has [n/2^16] entries to represent a set of values in [0, n). 

Overall, BitMagic is simpler than Roaring, but we expect that is can sometimes use more memory. BitMagic includes the following SIMD-based optimizations on x64 processors with support for the SSE2 and SSE4 instruction sets:

- It uses manually optimized SIMD instructions to compute and AND, OR, XOR and ANDNOT operations between two bitset containers.
- It uses a mix of SIMD instructions and the dedicated population-count instruction (popcnt) for its optimized functions that compute only the cardinality of the result from AND, OR, XOR and ANDNOT
- It uses SIMD instructions to negate a bitset quickly.

### FASTER ARRAY-BITSET OPERATIONS WITH BIT-MANIPULATION INSTRUCTION

Like most commodity processors, Intel and AMD processors benefit from bit-manipulation instructions [43]. Optimizing compilers often use them, but not always in an optimal manner.

**3.1. Converting Bitsets To Arrays**

Two useful bit-manipulation instructions are blsi, which sets all the least significant 1-bit to zero (i.e..., x & -x in C)

```go
func (bc *bitmapContainer) fillArray(container []uint16) {
	//TODO: rewrite in assembly
	pos := 0
	base := 0
	for k := 0; k < len(bc.bitmap); k++ {
		bitset := bc.bitmap[k]
		for bitset != 0 {
			t := bitset & -bitset
			container[pos] = uint16((base + int(popcount(t-1))))
			pos = pos + 1
			bitset ^= t
		}
		base += 64
	}
}

func popcount(x uint64) int {
	// Implementation: Parallel summing of adjacent bits.
	// See "Hacker's Delight", Chap. 5: Counting Bits.
	// The following pattern shows the general approach:
	//
	//   x = x>>1&(m0&m) + x&(m0&m)
	//   x = x>>2&(m1&m) + x&(m1&m)
	//   x = x>>4&(m2&m) + x&(m2&m)
	//   x = x>>8&(m3&m) + x&(m3&m)
	//   x = x>>16&(m4&m) + x&(m4&m)
	//   x = x>>32&(m5&m) + x&(m5&m)
	//   return int(x)
	//
	// Masking (& operations) can be left away when there's no
	// danger that a field's sum will carry over into the next
	// field: Since the result cannot be > 64, 8 bits is enough
	// and we can ignore the masks for the shifts by 8 and up.
	// Per "Hacker's Delight", the first line can be simplified
	// more, but it saves at best one instruction, so we leave
	// it alone for clarity.
	const m = 1<<64 - 1
	x = x>>1&(m0&m) + x&(m0&m)
	x = x>>2&(m1&m) + x&(m1&m)
	x = (x>>4 + x) & (m2 & m)
	x += x >> 8
	x += x >> 16
	x += x >> 32
	return int(x) & (1<<7 - 1)
}
```

Such code is useful when we need to convert a bitset container to an array container. We can ensure that only a handful of instructions are needed per bit set in the bitset container.

### 3.2 Array-Bitset Aggregates

Roaring has bitset containers, which we implemented as arrays of 64-bit words. Some of the most common operations on bitsets involve getting, setting, flipping or clearing the value of a single bit in one of the words. Sometimes, it is also necessary to determine whether the value of the bit was changed.

```go
type bitmapContainer struct {
	cardinality int
	bitmap      []uint64
}
```

We are interested in the following scenario: given a bitset and an array of 16-bit values, we wish to set (or clear or flip) all bits at the indexes corresponding to the 16-bit values, we wish to set (or clear of flip) all bits at the indexes corresponding to the 16-bit values. For example, if the array is made of the values {1,3,7,96,130}, then we might want to set the bits at indexes 1,3,7,96,130 to 1. This operation is equivalently to computing the union between a bitset and an array container.  

- If we do not care about the cardinality, and merely want to set the bit, we can use the simple C expression (w[pos] >> 6 |= UNIT67_C(1) << (pos & 63))

## Vectorized Processing

Modern commodity processors use parallelism to accelerate processing. SIMD instructions offer a particular form of processor parallelism [45] that proves advantageous for processing large volumes of data [46]. Whereas regular intructions operate on a single machine word (e.g., 64 bits), SIMD instructions operate on larger registers (e.g., 256 bits) that can be used to present "vectors" containing several distinct values. For example, a single SIMD instruction can add sixteen 16-bit (16x16) integers in one 32B vector register to the corresponding 16-bit integer in another vector register using a single cycle.

SIMD instructions are ideally suited for operations between bitset containers. When computing the intersection, union, difference or symmetric difference between two words (AND, OR, AND NOT, or XOR) and, optionally, save the result to memory. All of these operations have corresponding SIMD instructions. So, instead of working 64-bit words, we work over larger words (256 bits), dividing the number of instruction (/4) and giving significantly faster processing.

Historically, SIMD instructions have gotten wider and increasingly powerful. The Pentium 4 was limited to 128-bit instructions.

## Vectorized Population Count Over  Bitsets

We can trivially vectorize operations between bitsets. Indeed, it sufficies to compute bitwise operations over vectors instead of machine words. By aggressively unrolling the resulting loop, we can produce highly efficient code. Optimizing compilers can often automatically vectorize such code. It is more difficult, however, to also compute the cardinality of the result efficiently. Ideally, we would like to vectorize simultaneously.

## Vectorized Harley-Seal Population Count Commodity processors have dedicated instructions to count the 1-bits in a word (the "population-count"): popcnt for x64 processors and cnt for the 64-bit RAM architecture.

Processors have dedicated  instructions to count the 1-bits in a word (the "population-count"): popcnt for x64 processors and cnt for the 64-bit ARM architecture. On recent Intel processors, popcnt has a throughput of one instruction per cycle for both 32-bit and 64-bit integers.

## Vectorized Intersections Between Arrays

Because array containers represent integers values as sorted arrays of 16-bit integers, we can put to good use an algorithm based on a vectorized string comparison function 
