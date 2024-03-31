# Consistently faster and smaller compressed bitmaps with Roaring

## Summary

Compressed bitmap indexes are used in databases and search engines. Many bitmap compression techniques have been proposed, almost all relying primarily on RLE. However, on unsorted data, we can get superior performance with a hybrid compression technique that use both uncompressed bitmaps and packed arrays inside a two-level tree. An instance of this technique, Roaring has recently been proposed Due to its good performance, it has been adopted by several production platforms (Spark).

## Introduction

Sets are fundamental abstraction in software. They can be implemented in various ways, as hash sets, as trees and so forth. In databases and search engines, sets are often an integral part of indexes. For example, we may need to maintain a set of all documents or rows that satisfy some property. Besides adding or removing elements from the set, we need fast functions to compute the intersection, the union, the difference between sets, and so on.

To implement a set of integers, a particularly appealing strategy is the bitmap (also called bitset or bit vector). Commodity processors use words of W = 32 or W = 64 bits. By combining many such words, we can support large values of n. Intersections, unions and differences can then be implemented as bitwise AND, OR and AND NOT operations.

Unfornately, conventional bitmaps are only applicable when the cardinality of the set (|S|) is relatively large compared to the universe size (n), e.g., |S| > n/64. They are also suboptimal when the set is made of consecutive values (e.g. S = {1,2,3,4, ..., 99, 100}).

One popular approach has been to compress bitmaps with run-length encoding (RLE). Effectively, instead of using [n/W] words containing W bits for all bitmaps, we look for runs of consecutive words containing only ones or only zeros, and we replace them with markers that indicates which value is being repeated, and how many repetitions there are.

Roaring allows fast random access. To check for the presense of a 32-bit integer x, we seek the container corresponding the 16 most significant bits of x, using a binary search. If a bitmap container is found, we check the corresponding bit (at index x mode 2^16); if an array container is found, we use a binary search. Similarly, we can compute the intersection between two Roaring bitmaps without having to access all of the data. Indeed, suppose that we have a Roaring bitmap B1 containing only a few values, which all fall in the range [0, 2^16). This implies it uses an array container. Suppose we have another Roaring bitmap B2 over the same range but containing many values; it can only use a bitmap container. In that case, computing the intersection between B1 and B2 can be done in time O(|B1|), since it suffices to iterate over the set values of B1 and check whether the corresponding bit is set in the bitmap container of B2. Moreover, to compute the in-place union of B1 and B2, where the result is stored in the large bitmap B2, it suffices to iterate through the values of B1 and set the corresponding bits in B2 to 1, in time O(|B1|).

bimap-bitmap, array-array, array-bitmap, run-bitmap, run-array, run-run.

## Application Context

Sets can be used for many purposes. We are interested in applications that use sets of integers as part of an index. For example, one might index an attribute in a database or a word in a set of documents: for each attributed value, we have a set of numerical record identifiers. 

Indexes are most useful when there are many record identifiers. We expect the integer values in the set to span a wide range of values; i.e., at least hundreds of thousands. We are interested in cases where bitmaps are likely applicable: on average, there should be more than a few dozen integer value per set.

## Roaring BITMAP

Roaring bitmaps are used to represent sets of 32-bit unsigned integers. At a high level, a Roaring bitmap implementation is a key-value data structure where each key-value pair represents the set S of all 32-bit integers that share the same most significant 16 bits. The key is made of the shared 16 bits, whereas the value is a container storing the remaining 16 least significant bits from each member of S. No container is ever uses much more 8kB of memory. Thus several such small containers fit in the L1 CPU cache of most processors. 

Thus, when first creating a Roaring bitmap, it is usually made of array and bitmap containers. Runs are not compressed. Upon request, the storage of the Roaring bitmap can be optimized using the runOptimize function. This triggers a scan through the array and bitmap containers that convert them, it helpful, to run containers.

**Counting the number of runs** A critical step in deciding whether an array or bitmap container should be converted to a run container is to count the number of runs of consecutive numbers it contains. For array containers, we count this number by iterating through the 16-bit integers and comparing them two by two in a straightforward manner. Because array containers have at most 4096 integers, this computation is expected to be fast. 

**Algorithm 1**

## Logical Operations

**Union and intersection** 

There are many necessary logical operations, but we present primarily the union and intersection. There are most often used, and the most likely operations to cause performance bottlenecks.

An important algorithm for our purposes is the galloping intersection (also called exponential intersection) to compute the intersection between two sorted arrays of sizes c1, c2. It has complexity O(min(c1, c2) log max(c1, c2)). In this approach, we pick the next available integer i from the smaller array and seek an integer at least as big in the larger array, looking first at the next available value, the looking twice as far, and so on, until we find an integer that is not smaller than i. We then use a binary search in the larger array to find the exact location of the first integer not lower than i. We call this process a galloping search, and repeat it with each value from the smaller array.

**Algorithm 2** Algorithm to convert the set bits in a bitmap into a list of runs. We assume two-complement's 64-bit arithmetic. We use the bitwise AND and OR operations.

A galloping search makes repeated random access in a container, and it could therefore cause expensive cache misses. However, in our case, the potential problem is mitigated by the fact that all our containers fit in CPU cache.

Unions between Roaring data structures are handled in the conventional manner: we iterate through the keys in sorted order; if a key is in both input Roaring bitmaps, we merge the two containers, add the result to the output and advance in the two bitmaps. Otherwise, we clone the container corresponding to the smaller key, add it to the output and advance in this bitmap. When one bitmap runs out of keys, we append all the remaining content of the other bitmap to the output.

Though we do not use this technique, instead of cloning the container is stored and used, and a copy is only made if an attempt is made to modify the container further. 
