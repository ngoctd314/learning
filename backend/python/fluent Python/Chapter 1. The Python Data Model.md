# The Python Data Model

One of the best qualities of Python is its consistence. After working with Python for a while, you are able to start marking informed, correct guesses about features that are new to you.

The Python interpreter invokes special methods to perform basic object operations, often triggered by special syntax. The special method names are always written with leading and trailing double underscores. For example, the syntax obj[key] is supported by the __getitem__ special method.

We implement special methods when we want our objects to support and interact with fundamental language constructs such as:

+ Collections
+ Attribute access
+ Iteration (including async iteration using async for)
+ Operator overloading
+ Function and method invocation
+ String representation and formatting
+ Asynchronous programming using await
+ Object creation and destruction
+ Managed contexts using the with and async with statements

**Magic and Dunder**

Dunder is a shortcut for "double underscore before and after."

## A Pythonic Card Deck

```py
class FrenchDeck:
    ranks = [str(n) for n in range(2, 11)] + list("JQKA")
    suits = "spades diamonds clubs hearts".split()

    def __init__(self):
        self._cards = [Card(rank, suit) for suit in self.suits for rank in self.ranks]

    def __len__(self):
        return len(self._cards)

    def __getitem__(self, position):
        return self._cards[position]


deck = FrenchDeck()
print(len(deck))
print(deck[-1])
```

We've just seen two advantages of using special methods to leverage the Python Data Model:

- Users of your classes don't have to memorize arbitrary method names for standard operations (How to get the number of items? Is it .size(), .length(), or what?)
- It's easier to benefit from the rich Python standard library and avoid reinventing the wheel, like the random.choice function.

Because our __getitem__ delegates to the [] operator of self._cards, our deck automatically supports slicing.

Just by implementing the __getitem__ special method, our deck is also iterable:


