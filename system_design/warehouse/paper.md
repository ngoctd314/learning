# Roaring Bitmaps: implementation of an Optimized Software Library

Compressed bitmap indexes are used in systems such as Git or Oracle to accelerate queries. They represent sets and often support operations such as unions, intersections, differences, and symmetric differences. Several important systems such as Elasticsearch, Apache Spark, ... rely on a specific type of compressed bitmap index called Roaring. It benefits from several algorithms designed for the single-instruction-multiple-data (SIMD) instructions available on commodity processors. In particular, we present vectorized algorithms to compute the intersection, union, difference and symmetric between arrays.

## Introduction

Contemporary computing hardware offers performance opportunities through improved parallelism, by having more cores and better single-instruction-multiple-data (SIMD) instructions. Meanwhile, software indexes often determine the performance of big-data applications. Efficient indexes not only improve latency and throughput, but they also reduce energy usage.

Indexes are often made of sets of numerical identifiers (stored as integers). For instance, inverted indexes map query terms to document identifiers in search engines, and conventional database indexes map column values to record identifiers. We often need efficient computation of the intersection (A ∩ B), the difference (A \ B), or the symmetric difference of these sets.

The bitmap (or bitset) is a time-honored strategy to represent sets of integers concisely. Given a universe of n possible integers, we use a vector of n bits to represent any on set. On a 64-bit processor, [n/64] inexpensive bitwise operations suffice to compute set operations between two bitmaps:

- the intersection corresponds to the bitwise AND
- the union corresponds to the bitwise OR
- the difference corresponds to the bitwise ANDNOT
- the symmetric difference corresponds to the bitwise XOR

Unfornately, when the range of possible values (n) is too wide, bitmaps can be too large to be practical. For example, it might be to impractical to represent the set {1,2^31} using a bitset. For this reason, we often use compressed bitmaps.

Though there are many ways to compress bitmap, several systems rely on an approach called Roaring including Elasticsearch [2], Metamarkets's Druid [3], ... and Apache Spark [9, 10]. In turn, these systems are in widespread use.

Roaring partitions the space [0, 2^32) into chunks consisting of ranges of 2^16 integers ([0, 2^16), [2^16, 2^17), ...). For a value in the set, its least significant sixteen bits are stored in a container corresponding to its chunk (as determined by its most significant sixteen bits), using one of three possible container types:

- bitset containers made of 2^16 bits or 8kB;
- array containers made of up to 4096 sorted 16-bit integers;
- run containers made of a series of sorted <s,l> pairs indicating that all  integers in the range [s, s+l] are present.

At a high level, we can view a Roaring bitmap as a list of 16-bit numbers (corresponding to the most-significant 2B of the values present in the set), each of which is coupled with a reference to a container holding another set of 16-bit numbers corresponding to the least-significant 2B of the elements sharing the same prefix. See Fig.1.

We dynamically pick the container type to minimize memory usage. For example, when intersecting two bitset containers, we determine whether the result is an array or a bitset container on-the-fly. As we add or remove values, a container's type might change. No bitset container may store fewer than 4097 distinct values; no array container store more than 4096 distinct values. If a run container has more than 4096 values, then it must have no more than 2047 runs, otherwise the number of runs must be less than half the number of distinct values.

Roaring offers logarithmic-time random access: to check for the presense of a 32-bit integer, we seek the container corresponding to the sixteen most-significant bits using a binary search. If this prefix is not in the list, we know that the integer is not present. If a bitmap container is found, we check the corresponding bit; if an array or run container is found, we use a binary search.

As an experiment, we built an optimized Roaring bitmap library in C. Based in this work, we make two main contributions:

- We present several non-trivial algorithmic optimizations. See Table 1. In practical, we show that a collection of algorithms exploiting single-instruction-multiple-data(SIMD) instructions can enhance the performance of a data structure like Roaring in some cases, above and beyond what state-of-the-art optimizing compilers can achieve. To our knowledge, it is the first work to report on the benefits of advanced SIMD-based algorithms for compressed bitmaps.  

Though the approach we use to compute array intersections using SIMD instruction in 4.2 is not new, our work on the computation of the union 4.3, difference 4.4 and symmetric difference 4.4 of arrays using SIMD instructions might be novel of general interest.

- We benchmark our C library against a wide range of alternatives in C and C++. Our results provide guidance as to the strengths and weakness of our implementation.

## Integer-set data structures

The simplest way to represent a set of integers is as a sorted array, leading to easy implementation. Querying for the presense of a given value can be done in logarithmic time using a binary search. We can compute the intersection, union, difference, and symmetric different between two sorted arrays in linear time: O(n1 + n2) where n1 and n2 are the cardinalities of the two arrays. The intersection and difference can also computed in time O(n1logn2), which is advantageous when one array is small.

## Compressed Bitsets

A bitset (or bitmap) has both the benefits of the hash set (constant-time random access)  and of a sorted array (good locality), but suffers from impractical memory usage when the universe size is too large compared to the cardinality to the sets. So we use compressed bitmaps. Though there are alternatives the most popular bitmap compression techniques are based on the word-aligned RLE compression model inherited from Oracle. Techniques such as BBC, WAH or EWAH use special marker words to compress long sequences of identical fill words. When accessing these formats, it may be necessary to read every compressed word to determine whether it indicates a sequence of fill words, or a dirty word. A major limitation of formats like BBC, WAH, Concise in an extensive comparison. A major limitation of formats like BBC, WAH, Concise or EWAH is that random access is slow. That is to check whether a value is present in a set can take linear time O(n), where n is the compressed size of the bitmap.

Use Roaring for bitmap compression whenever possible. Do not use other bitmap compression methods such as BBC, WAH, EWAH...

## Bitmagic

## Faster Array-BitSet Operations With Bit-Manipulation instructions

Like most commodity processors, Intel and AMD processors benefit from bit-manipulation instructions. Optimizing compilers often use them, but not always in an optimal manner.

### Converting Bitsets To Arrays

Two useful bit-manipulation instructions are blsi, which sets are all the least significant 1-bit to zero (i.e x & -x in C), and tzcnt which counts the number of trailing zeros. Using the corresponding Intel intrinsics (_blsi_u64 and _mm_tzcnti_64) we can extract the locations of all 1-bits in a 64-bit word (w) to an array (out) efficiently.

### Array-BitSet Aggregates
