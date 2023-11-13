# Generator expressions in Python are a concise way to create generators, which are a type of iterable Python. They provide a memory-efficient way to iterate over large sequences of data.

# A generator expression looks similar to a list comprehension, but it uses parentheses `()` instead of square like brackets `[]`. The main difference is that a generator expression produces values on-the-fly, one at a time, and doesn't create a list in memory.

list_comp = [x**2 for x in range(5)]

gen_exp = (x**2 for x in range(5))

print(list_comp)
print(gen_exp)

for item in gen_exp:
    print(item)

# The list comprehension creates a list [0, 1, 4, 9, 16]
# The generator expression (x**2 for x in range(5)) creates a generator object and the values are generate on-the-fly when iterated
# When iterating over the generator expression, values are produced one at a time without creating a list in memory
# Generator are useful when you need to iterate over a large sequence of data, and you don't want to store the entire sequence in memory. They are memory-efficient and can be more suitable for certain scenarios compared to creating a list.
# Keep in mind that a generator can be iterated only once, once it's exhausted, you need to create a new generator if you want to iterate over the data again.

symbols = "abcde"
gen = (ord(symbol) for symbol in symbols)
print(gen)
# print(list(gen))
print(tuple(gen))

# If the generator expression is the single argument in a function call, there is no need to duplicate the enclosing parentheses.
# The array constructor takes two arguments, so the parentheses around the generator expression are mandatory. The first argument of the array constructor defines the storage type used for the numbers in the array.

colors = ["black", "white"]
sizes = ["S", "M", "L"]

for tshirt in (f"{c} {s}" for c in colors for s in sizes):
    print(tshirt)
