Reference: https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

# How the Go runtime implements maps efficiently (without generics)

## What is a map function?

To understand how a map works, let's first talk about the idea of the map function. A map function maps one value to another. Given one value, called a key, it will return a second, the value.

## Go's map is a hashmap

The specific map implementation i'm going to talk about is the hashmap, because this is the implementation that the Go runtime uses. A hashmap is a classic data structure offering O(1) lookups on average and O(n) in the worst case. That is, when things are working well, the time to execute the map function is a near constant. 

The size of constant is part of the hashmap design and the point at which the map moves from O(1) to O(n) access time is determined by its hash function.

## The hash function

What is a hash function? A hash function takes a key of an unknown length and returns a value with a fixed length.

hash(key) -> integer

this hash value is almost always an integer for reasons that we'll see in a moment.

Hash and map functions are similar. They both take a key and return a value. However in the case of the former, it returns a value derived from the key, not the value associated with the key.

## Important properties of a hash function

It's important to talk about the properties of a good hash function as the quality of the hash function determines how likely the map function is to run near O(1).

When used with a hashmap, hash functions have two important properties. The first is stability. The hash function must be stable. Given the same key, your hash function must return the same answer. If doesn't you will not be able to find things you put into the map.

The second property is good distribution. Given two near identical keys, the result should be wildly different. This is important for two reasons. This is important for two reasons. Firstly, as we'll see, values in a hashmap should be distributed evenly across buckets, otherwise the access time is not O(1). Secondly as the user can control some of the aspects of the input to the hash function, they may be able to control some of the aspects of the input to the hash function, they may be able to control the output of the hash function, leading to poor distribution which has been a DDos vector for some languages. This property is also known as collision resistence.

## The hashmap data structure

The second part of hashmap is the way data is stored.

The classical hashmap is an array of buckets each of which contains a pointer to an arary of key/value entries. In this case our hashmap has eight buckets (as this is the value that the Go implementation uses) and each bucket can hold up eight entries each (again drawn from the Go implementation). Using powers of two allows the use of cheap bit masks and shifts rather than expensive division.

As entries are added to a map, assuming a good hash function distribution, then the buckets will fill at roughly the same rate. Once the number of entries across each bucket passes some percentage of their total size, known as the load factor, then the map will grow by doubling the number of buckets and redistributing the entries across them.

With this data structure in mind, if we had a map of project names to Github stars, how would we about inserting a value into the map?
