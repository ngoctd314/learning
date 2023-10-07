Reference: https://dave.cheney.net/2018/05/29/how-the-go-runtime-implements-maps-efficiently-without-generics

# How the Go runtime implements maps efficiently (without generics)

## What is a map function?

To understand how a map works, let's first talk about the idea of the map function. A map function maps one value to another. Given one value, called a key, it will return a second, the value.

## Go's map is a hashmap

The specific map implementation i'm going to talk about is the hashmap, because this is the implementation that the Go runtime uses. A hashmap is a classic data structure offering O(1) lookups on average and O(n) in the worst case. That is, when things are working well, the time to execute the map function is a near constant. 

The size of constant is part of the hashmap design and the point at which the map moves from O(1) to O(n) access time is determined by its hash function.


