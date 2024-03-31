1. A key ingredient in the performance of Roaring bitmaps are the new bit-count processor instructions (such as popcnt) that became available on desktop processors more recently (2008). These new instructions allow us to quickly compute the densite of new chunks, and to efficiently extract the location of the set bits from a bitmap.

2. Convert bitmap, array -> run

We can illstrate the core operation of the algorithm using a single 32-bit word containing 6 runs of consecutive ones: 

Ci                  = 000111101111001011111011111000001
Ci << 1             = 001111011110010111110111110000010
(Ci << 1) ANDNOT Ci = 001000010000010100000100000000010

We can verify that bitCount((Ci << 1) ANDNOT Ci) = 6, that is, we have efficiently computed the numbers of runs.

3. Container fit in CPU cache

4. Union Bitmap vs Bitmap

A union between two bitmap containers is straightforward: we execute the bitwise OR between all pairs of corresponding words. There are 1024 words in each container, so 1024 bitwise OR operations are need. At the same time, we compute the cardinality of the result using the bitCount function on the generated words.

5. Bitmap vs Array

Unions are also efficient: we create a copy of the bitmap and iterate over the array, setting the corresponding bits.

6. Array vs Array

For unions, if the sum of the cardinalities of the array containers is 4096 or less, we merge the two sorted arrays into a new array container that has its capacity set to the sum of the cardinalities of the input arrays. Otherwise, we generate an initially empty bitmap container. Though we cannot know whether the result will be a bitmap container (whether the cardinality is larger than 4096), as a heuristic, we suppose that it will be so. Iterating through the values of both arrays, we set the corresponding bits in the bitmap to 1. Using the bitCount function, we compute cardinality, and then convert the bitmap into an array container if the cardinality is at most 4096.

Another feature of Roaring is that some of these logical operations can be executed in place. In-place computations avoid unneccessary memory allocations and improve data locality.

- The union of a bitmap container with any other container can be written out in the input bitmap container.
- Though array containers do not support in-place operations, we find it efficient to support in-place unions in a run container with repect to either another run container or an array container. In these cases, it is common that the result of the union is either smaller or not much larger than combined sizes of the inputs.

A common operation in applications is the aggregation of a long list of bitmaps. When the problem is to compute the intersection of many bitmaps, we can expect a naive algorithm to work well with Roaring: we can compute the intersection of the first two bitmaps, then intersection the result with the third bitmap, and so forth. With each new intersection, the result might become smaller, and Roaring can often efficiently compute the intersection between bitmap having small cardinality and bitmaps having larger cardinalities, as already stated. Computing the union of many bitmaps requires more care. As already remarked in Chambi et at, it is wasteful to update the cardinality each and every time when computing the union between several bitmap containers. Though the bitCount function is fast, it can still use a significant fraction of the running time: Chambi et al report that it reduces the speed by about 30%. Instead, we proceed with what we call a "lazy union". We compute the union as usual, except that some unions between containers are handled differently:

We consider two strategies to compute the union of many bitmaps. One approach is a naive two-by-two union: we first compute the union of the first two bitmaps, then the union of the result and then the third bitmap and so forth, doing the computation in-place if possible. The benefit of this approach is that we always keep just one intermediate result in memory. In some instances, however, we can get better results with other algorithms. For example, we can use heap: put all original bitmaps in a min-heap, and repeatedly poll the two smallest bitmaps, compute their union, and put them back in the heap, as long as the heap contain containes more than one bitmap. This approach may create many more intermediate bitmaps, but it can also be faster in some instances.


Priority Query

SIMD

popCount

1 bit position
