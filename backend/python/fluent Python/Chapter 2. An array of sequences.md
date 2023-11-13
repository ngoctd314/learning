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

## List Comprehensions and Generator Expressions

A quick way to build a sequence is using a list comprehension (if the target is a list) or a generator expression (for other kinds of sequences). If you are not using these syntatic forms on a daily basis, I bet you are missing opportunities to write code that is more readable and often faster at the same time.

### List Comprehensions and Readability

A listcomp goals is always to build a new list. For Python, as for English, there are no hard-and-fast rules for clear writing.

**Local Scope Within Comprehensions and Generator Expressions**

In Python 3, list comprehensions, generator expressions, and their siblings set and dict comprehensions, have a local scope to hold the variables assigned in the for clause.

Variables remain acessiable after those comprehensions or expressions return -- unlike local variables in a function

```py 
codes1 = [ x:= symbol for symbol in symbols]
print(x)
```

### Listcomps Versus map and filter

Listcomps do everything the map and filter functions do, without the contortions of the functionally challenged Python lambda.

In Python the filter() function is used to filter elements from an iterable (such as a list) based on a specified condition. The filter() function takes a function and an iterable as arguments and returns an iterator that includes only the elements for which the function returns True.

The filter function can be applied to various situations where you want to selectively include or exclude elements from an iterable based on a condition.

In Python, the map function is used to apply a specified function to all items in an input iterable, and returns an iterator of the results.

```py
map(function, iterable)
```

```py
numbers = [1, 2, 3, 4, 5]


def square(x):
    return x**2


squared_numbers = map(square, numbers)

squared_numbers_list = list(squared_numbers)
```

The map() function can be used to apply any function to each element of an iterable providing a convenient way to transform data.

### Generator Expressions

To initialize tuples, arrays, and other types of sequences, you could also start from a listcomp, but a genexp (generator expression) saves memory because it yields items one by one using the iterator protocol instead of building a whole list just to feed another constructor.

### Tuples Are Not Just Immutable Lists
