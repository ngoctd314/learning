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


