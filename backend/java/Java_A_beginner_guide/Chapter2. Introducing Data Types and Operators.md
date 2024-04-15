# Chapter 2. Introducing Data Types and Operators

## Java's Primitive Types

Java contains two general categories of built-in data types: object-oriented and non-object-oriented.

|Type|Width in Bits|Rage|
|-|-|-|
|byte|8|-128 to 127|
|short|16|-32,768 to 32,767|
|int|32|-2,147,483,648 to 2,147,483,647|
|long|64|-9... to 9...|

As the table shows, all of integer types are signed positive and negative values. Java does not support unsigned (positive-only) integers. Many other computer languages support both signed and unsigned integers. However, Java's designers felt that unsigned integers were unnecessary.

**Q** You say that there are four integer types: in, short, long and byte. However, I have heard that char can also be categorized as an integer type in Java. Can you explain?

**A** The formal specification for Java defines a type category called intergral types, which includes byte, short, int, long and char. They are called integral types because they all hold whole-number, binary values. However, the purpose of the first four is to represent numeric integer quantities. The purpose of char is to represent characters. Therefore, the principal uses of char and principal uses of the other integral types are fundamentally different. Because of the differences, the char type is treated separately in this book.

## Floating-Point Types

As explained in Chapter 1, the floating-point types can represent numbers that have fractional components. There are two kinds of floating-point types, float and double, which represent single- and double-precision numbers, respectively. Type float is 32 bits wide and type double is 64 bits wide.

## Characters

In Java, characters are not 8-bit quantities like they are in many other computer languages. Instead, Java uses Unicode. Unicode defines a character set that can represent all of the characters found in all human languages. In Java, char is an unsigned 16-bit type having a range of 0 to 65,535. The standard 8-bit ASCII character set is a subset of Unicode and ranges from 0 to 127. Thus, the ASCII characters are still valid Java characters.

Since char is an unsigned 16-bit type, it is possible to perform varios arithmetric manipulations on a char variable. For example, consider the following program:

```java
class Example {
    public static void main(String[] args){
      char ch;
      ch = 'X';
      System.out.println("ch contains " + ch);
      ch++;
      System.out.println("ch is now " + ch);
    }
}
```

**Q** Why does Java use Unicode?
**A** Java was designed for worldwide use. Thus, it need to use a character set that can represent all the world's languages. Unicode is the standard character set designed expressly for this purpose. Of course, the use of Unicode is inefficient for languages such as English, German, Spanish, or French, whose characters can be contained within 8 bits. But such is the price that must be paid for global portability.

## Literals

In Java, literals refer to fixed values that are represented in their human-readable form. For example, the number 100 is a literal. Literals are also commonly called constants.

By default, integer literals are of type int. If you want to specify a long literal, append an l or an L. For example, 12 is an int, but 12L is a long.

By default, floating-point literals are of type double. To specify a float literal, append an F or f to the constant. For example, 10.19F is of type float.

Although integer literals create an int value by default, they can still be assigned to variables of type char, byte, or short as long as the value being assigned can be represented by the target type. An integer literal can always be assigned to a long variable.

### Hexadecimal, Octal, and Binary Literals

As you may know, in programming it is sometimes easier to use a number system based on 8 or 16 instead of 10. The number based on 8 is called octal, and it uses the digits 0 through 7.

As a point of interest, Java also allows hexadecimal floating-point literals, but they are seldom used.

It is possible to specify an integer literal by use of binary. To do so, precede the binary number with a 0b or 0B. For example, this specifies the value 12 in binary: 0b1100.

**Q** is a string consisting of a single character the same as a character literal? For example, is "k" the same as a 'k'?

**A** No. You must not confuse strings with characters. A character literal represents a single letter of type char. A string containing only one letter is still a string. Although strings consist of characters, they are not the same type.

### A closer look at variables

### The Scope and Lifetime of Variables

So far, all of the variables that we have been using were declared at the start of the main() method. A block defines a scope. Thus, each time you start a new block, you are creating a new scope. A scope determines what objects are visible to other parts of your program. It also determines the lifetime of those objects.

In general, every declaration Java has a scope. As a result, Java defines a powerful, finely grained concept of scope. Two of the most common scopes in Java are those defined by a class and those defined by a method. A discussion of class scope (and variables declared within it) is deferred until later in this book, when classes are described. For now, we will examine only the scopes defined by or within a method.

The scope defined by a method begins with its opening curly brace. However, if that method has parameters, they too are included within the method's scope.

As a general rule, variables declared inside a scope are not visible (that is, accessible) to code that is defined outside that scope. Thus, when you declare a variable within a scope, you are localizing that variable and protecting it from unauthorized access and/or modification. A variable declared within a block is called a local variable.

Scopes can be nested. For example, each time you create a block of code, you are creating a new, nested scope. When this occurs, the outer scope encloses the inner scope.

Here is another important point to remember: variables are created when their scope is enter, and destroyed when their scope is left.

There is one quirk to Java's scope rules that may surprise you: although blocks can be nested, no variable declared within an inner scope can have the same name as a variable declared by an enclosing scope. For example, the following program, which tries to declare two separate variables with the same name, will not compile. 

```java
public static void main(String[] args) {
  int count;
  for (count =0; count < 10; count++) {
    System.out.println("This is count: " + count);
    int count; // can't declare count again because it's already declared
  }
}
```

## Operators

In Java, all objects can be compared for equality or inequality using == and !=. However, the comparison operators, <, >, <=, or >=, can be applied only to those types that support an ordering relationship. Therefore, all of the relational operators can be applied to all numeric types and to type char. However, values of type boolean can only be compared for equality or inequality, since the true and false values are not ordered. For example, true > false has no meaning in Java.

### Short-Circuit Logical Operators

Java supplies special short-circuit versions of its AND and OR logical operators that can be used to produce more efficient code. To understand why, consider the following. In an AND operation, if the first operand is false, the outcome is false no matter what value the second operand has. In an OR operation, if the first opeand is true, the outcome of the operation is true, no matter what the value of the second operand.

The short-circuit AND operator is &&, and the short-circuit OR operator is ||.

One last point: The formal specification for Java refers to the short-circuit operators as the conditional-or and the conditional-and operators, but the term "short-circuit" is commonly used.

### The assignment Operator

```java
int x, y, z;
x = y = z = 100;
```

This fragment sets the variables x, y and z to 100 using a single statement. This works because the = is an operator that yields the value of the right-hand expression. Thus, the value of z = 100 is 100, which is then assigned to y, which in turn is assigned to x. Using a "chain of assignment" is an easy way to set a group of a variables to a common value.

**Q** Since the short-circuit operators are, in some cases, more efficient than their normal counterparts, why does Java still offer the normal AND and OR operators?

**A** In some cases you will want both operands of an AND or OR operation to be evaluated because of the side effects produced.

```java
```

#### Type Conversion in Assignments

When one type of data is assigned to another type of variable, an automatic type conversion will take place if:

- The two types are compatible.
- The destination type is larger than the source type.

#### Casting Incompatible Types

A cast is an instruction to the compiler to convert one type into another.

For example, if you want to convert the type of the expression x/y to int, you can write

```java
double x, y;
(int) (x/y);
```

Here, even though x and y are of type double, the cast converts the outcome of the expression to int. The parentheses surrounding x/y are necessary. Otherwise, the cast to int would apply only to the x and not to the outcome of the division. The cast is necessary here because there is no automatic conversion from double to int.

When a cast involves a narrowing conversion, information might be lost. For example, when casting a long into a short, information will be lost if the long's value is greater than the range of a short because its high-order bits are removed. When a floating-point value is greater than the range of a short because its high-order bits are removed. When a floating-point value is cast to an integer type, the fractional component will also be lost due to truncation. For example, if the value 1.23 is assgined to an integer, the resulting value will simply be 1.

In the program, the cast of (x/y) to int results in the truncation of the fractional component, and information is lost. Next, no loss of information occurs when b is assigned the value 100 because a byte can hold the value 100. However, when the attempt is made to assign b the value 257, information loss occurs because 257 exceeds a byte's maximum value. Finally, no information is lost, but a cast is needed when assigning a byte value to a char.

### Expressions

#### Type conversion in Expressions

Within an expression, it is possible to mix two or more different types of data as long as they are compatible with each other. For example, you can mix short and long within an expression because they are both numeric types. When different types of data are mixed within an expression, they are all converted to the same type. This is accomplished through the use of Java's type promotion rules.

First, all char, byte, and short values are promoted to int. Then, if one operand is a long, the whole expression is promoted to long. If one operand is a float operand, the entire expression is promoted to float. If any of the operands is double, the result is double.

Type promotion can, however, lead to somewhat unexpected results. For example, when an arithmetic operation involves two byte values, the following sequence occurs: First, the byte operands are promoted to int. Then the operation takes place, yielding an int result.

Thus, the outcome of an operation involving two byte values will be an int. This is not what you might intuitively expect. 

```java
public static void main(String[] args) {
  byte b;
  int i;
  b = 10;
  i = b * b; // no cast needed because result is already elevated to int.
  b = (byte) (b * b); // cast is needed here to assign an int to a byte
  System.out.println(i + " " + b);
}
```

## Selftest
