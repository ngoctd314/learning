# Chapter 2. An array of sequences

## Overview of Built-In Sequences

The standard library offers a rich selection of sequence types implemented in C:

**Container sequences**

Can hold items of different types, including nested containers. Some examples: list, tuple, and collections.deque.

**Flat sequences**

Hold items of one simple type. Some examples: str, bytes, and array.array.

A container sequence holds references to the objects it contains, which may be of any type, while a flat sequence stores the value of its contents in its own memory space.

Flat sequences are more compact, but they are limited to holding primitive machine values like bytes, integers and floats.

Every Python object in memory has a header with metadata. The simplest Python object, a float, has a value field and two metadata fields:

- ob_refcnt: the object's reference count.
- ob_type: a pointer to the object's type.
- ob_fval: a C double holding the value of the float.

On a 64-bit Python build, each of those fields takes 8 bytes. That's why an array of floats is much more compact than a tuple of floats: the array is a single object holding the raw values of the floats, while the tuple consists of several objects - the tuple itself and each float object contained in it.

## List 
