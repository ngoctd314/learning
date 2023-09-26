## 4.8. More on Defining Functions

### 4.8.1. Default Argument Values

The most useful form is to specify a default value for one or more argument.

```py
def ask_ok(prompt, retries=4, reminder='Please try again!')
```

This function can be called in several ways:

- giving only the mandatory argument: ask_ok('Do you really want to quit?')
- giving one of the optional arguments: ask_ok('OK to overwrite the file?', 2)
- or even giving all arguments: ask_ok('OK to overwrite the file?', 2, 'Come on, only yes or no!')

```py
i = 5

# The default values are evaluated at the point of function definition in the defining scope
def f(arg=i):
    print(arg)

i = 6
f()
```

**Important warning:** The default value is evaluated only once. This makes a difference when the default is a mutable object such as a list, dictionary, or instances of most classes.

```py
def f(a, L=[]):
    L.append(a)
    return L


print(f(1))
print(f(2))
print(f(3))
# [1]
# [1, 2]
# [1, 2, 3]
```

If you don't want the default to be shared between subsequent calls, you can write the function like this instead:

```py
def f(a, L=None):
    if L is None:
        L = []
    L.append(a)
    return L
```

### 4.8.2. Keyword Arguments

Functions can also be called using keyword arguments of the form kwarg=value.

```py
def parrot(voltage, state="a stiff", action="voom", type="Norwegian Blue"):
```

Accepts one required argument (voltage) and three optional arguments (state, action, and type)

When a final formal parameter of the form \*\*name is present, it receives a dictionary (Mapping Types - dict) containing all keyword arguments except for those corresponding to a formal parameter. This may be combined with a formal parameter of the form \*name which receives a tuple containing the positional arguments beyond the formal parameter list. (\*name must occur before \*\*name). For example, if we define a function like this:

```py
def cheeseshop(kind, *arguments, **keywords):
    print("-- Do you have any", kind, "?")
    print("-- I'm sorry, we're all out of", kind)
    print("-" * 40)
    for arg in arguments:
        print(arg)
    print("-" * 40)
    for kw in keywords:
        print(kw, ":", keywords[kw])


cheeseshop(
    "Limburger",
    "It's very runny, sir.",
    "It's really very",
    shopkeeper="Michael Palin",
    client="John Cleese",
    sketch="Cheese Shop Sketch",
)
```

### 4.8.3. Special parameters

By default, arguments may be passed to a Python function either by position or explicitly by keyword.

```py
def f(pos1, pos2, / , pos_or_kwd, *, kwd1, kwd2):
#   Positional-only
#                     Positional or keyword
#                                     Key word only
```

**4.8.3.1. Positional-or-Keyword Arguments**

If / and \* are not present in the function definition, arguments may be passed to a function by position or by keyword.

**4.8.3.2. Positional-Only Parameters**

Looking at this in a bit more detail, it is possible to mark certain parameters as positional-only. If positional-only, the parameters order matters, and the parameters cannot be passed by keyword. Positional-only parameters are placed before a / (forward-slash). The / is used to logically separate the positional-only parameters from the rest of the parameters. If there is no / in the function definition, there are no positional-only parameters

**4.8.3.3. Keyword-only Arguments**

To mark parameters as keyword-only, indicating the parameters must be passed by keyword argument, place an \* in the arguments list just before the first keyword-only parameter.

**4.8.3.4. Function Examples**

Consider the following example function definitions paying close attention to the markers / and \*:

```py
def standard_arg(arg):
    print(arg)

def pos_only_arg(arg, /):
    print(arg)

def kwd_only_arg(*, arg):
    print(arg)

def combined_example(pos_only, /, standard, *, kwd_only):
    print(pos_only, standard, kwd_only)
```

**4.8.3.5. Recap**

The use case will determine which parameters to use in the function definition:

```py
def f(pos1, pos2, /, pos_or_kwd, *, kwd1, kwd2):
```

As guidance:

- Use positional-only if you want the name of the parameters to not be available to the user. This is useful when parameter names have no real meaning, if you want to enforce the order of the arguments when the function is called or if you need to some positional parameters and arbitrary keywords.
- Use keyword-only when names have meaning and the function definition is more understandable by being explicit with names or you want to prevent users relying on the position of the argument being passed.
- For an API, use positional-only to prevent breaking API changes if the parameter's name is modified in the future.

### Arbitrary Argument Lists

```py
ls = list(range(3, 6))
print(ls)

args = [3, 10]
ls = list(range(*args))
print(ls)
```

### Lambda Expressions

Small anonymous functions can be created with the lambda keyword. This function returns the sum of its two arguments: lambda a, b: a + b.

### Documentation Strings

```py
def my_function():
    """Do nothing, but document it
    No, really, it doesn't do anything
    """


print(my_function.__doc__)
```

### Function Annotations

Function annotations are completely optional metadata information about the types used by user-defined functions.

Annotations are stored in the **annotations** attribute of the function as a dictionary and have no effect on any other part of the function.

```py
def f(ham: str, eggs: str = "eggs") -> str:
    print("Annotations:", f.__annotations__)
    print("Arguments:", ham, eggs)
    return ham + " and " + eggs


f("spam")
```

### Intermezzo: Coding style

Now that you are about to write longer, more complex pieces of Python, it is a good time to talk about coding style. Most languages can be written (or more concise, formatted) in different styles; some are more readable than others. Making it easy for others to read your code is always a good idea, and adopting a nice coding style helps tremendously for that.

- Use 4-space indentation, and no tabs
- Wrap lines so that they don't exceed 79 characters
- Use blank lines to separate functions and classes, and larger blocks of code inside functions
- When possible, put comments on a line of their own
- Use docstrings
- ...
