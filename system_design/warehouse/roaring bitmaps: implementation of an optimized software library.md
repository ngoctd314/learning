# Roaring Bitmaps: implementation of an Optimized Software Library

Compressed bitmap indexes are used in systems such as Git or Oracle to accelerate queries. They represent sets and often support operations such as unions, intersections, differences. We present an optimized software library written in C implementing Roaring bitmaps: CRoaring. It benefits from several algorithms designed for the SIMD instruction available on commodity processors. In particular, we present vectorized algorithms to compute the intersection, union, difference and symmetric difference between arrays.

## Introduction

Comtemporary computing hardware offers performance opportunities through improved parallelism, by having more cores and better SIMD instructions.
