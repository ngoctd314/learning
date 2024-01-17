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

But the point of this example is the FrenchDeck class. It's short, but it packs a punch. First, like any standard Python collection, a deck responds to the len() function by returning the number of cards in it.

```py
deck = FrenchDeck()
len(deck) # 52
```

Reading specific cards from the deck - say, the first or the last - is easy, thanks to the __getitem__ method:

```py
deck[0]
# Card(rank='2', suit='spades')
```

We've just seen two advantages of using special methods to leverage the Python Data Model:

- Users of your classes don't have to memorize arbitrary method names for standard operations (How to get the number of items? Is it .size(), .length(), or what?)
- It's easier to benefit from the rich Python standard library and avoid reinventing the wheel, like the random.choice function.

Because our __getitem__ delegates to the [] operator of self._cards, our deck automatically supports slicing.

Just by implementing the __getitem__ special method, our deck is also iterable:

```py
deck = FrenchDeck()

for card in deck:
    print(card)
```

Iteration is often implicit. If a collection has no __contains__ method, the in operator does a sequential scan.

```py
deck = FrenchDeck()

print(Card("Q", "hearts") in deck)
print(Card("Q", "bearts") in deck)
```

Although FrenchDeck implicitly inherits from the object class, moss of its functionality is not inherited, but comes from leveraging the data model and composition. By implementing the special methods __len__ and __getitem__, our FrenchDeck behaves like a standard Python sequence, allowing it to benefit from core language features using random.choice, reversed, and sorted.

## How Special Methods Are Used

They are meant to be called by the Python interpreter, and not by you. You don't write my_object.__len__(). You write len(my_object) and, if my_object is an instance of a user-defined class, then Python calls the __len__ method you implemented.


