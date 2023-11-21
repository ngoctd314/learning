Reference: https://en.wikipedia.org/wiki/Locality_of_reference

In cs, locality of reference, also known as the principle of locality, is the tendency of a processor to access the same set of memory locations repetitively over a short period of time. There are two basic types of locality - temporal and spatial locality. Temporal locality refers to the reuse of specific data and/or resources within a relatively small time duration. Spatial locality (also termed data locality) refers to the use of data elements within relatively close storage locations. Sequential locality, a special case of spatial locality occurs when data elements are arranged and accessed linearly, such as traversing the elements in a one-dimensional array.

Locality is a type of predictable behavior that occurs in cs. Systems that exhibit strong locality of reference are great candidates for performance optimization through the use of techniques such as caching, prefetching for memory and advanced branch predictors of a processor core. 

## Types of locality

There are several different types of locality of reference:

- Temporal locality: if at one point a particular memory location is referenced, then it is likely that the same location will be referenced again in the near future. There is temporal proximity between adjacent references to the same memory location. In this case it is common to make efforts to store a copy of the reference data in faster memory storage, to reduce the latency of subsequent references. Temporal locality is a special case of spatial locality, namely when the prospective location is identical to the present location. 

- Spatial locality: If a particular storage location is referenced at a particular time, then it is likely that nearby memory locations will be referenced in the near future. In this case it is common to attempt to guess the size and shape of the area around the current reference for which it is worthwhile to prepare faster access for subsequent reference.

**Memory locality(data locality)** Spatial locality explicitly relating to memory.


