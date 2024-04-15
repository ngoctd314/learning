# Program statements

## The traditional switch statement

```java
switch (expression) {
    case constant1:
        break;
    case constant2:
        break;
    default:
        statement sequence
}
```

For versions of Java prior to JDK 7, the expression controlling the switch must resolve to type byte, short, int, char, or an enumeration. However, today, expression can also be of type String. This means that modern versions of Java can use a string to control a switch.

Technically, the break statement is optional, although most applications of the switch will use it. When encountered within the statement and resume at the next statement outside the switch. However, if a break statement does not end the statement sequence associated with a case, then all statements at and following the matching case will be executed until a break (or the end of the switch) is encountered. Thus, a case without a break will "fall through" to the next case.

**Ask the Expert**

**Q** Under what conditions should I use an if-else-if ladder rather than a switch when coding a multiway branch?

**A** In general, use an if-else-if ladder when the conditions controlling the selection process do not rely upon a single value.

```java
if (x < 10) // ...
else if (y != 0) // ...
else if (!done) // ...
```

This sequence cannot be recorded into a switch because all three conditions involve different variables - and differing types. 

**Missing Pieces**

```java
class Empty {
    public static void main(String[] args) {
        int i;
        for (i = 0; i < 10; ) { // The iteration expression  is missing
            System.out.println("Pass #" + i);
            i++; // increment loop control var
        }
    }
}
```

```java
class Empty2 {
    public static void main(String[] args) {
        int i;
        i = 0; // move initialization out of loop
        for (; i < 10;) {
            System.out.println("Pass #" + i);
            i++; // increment loop control var
        }
    }
}
```

### Loops with No Body

In Java, the body associated with a for loop (or any other loop) can be empty. This because a null statement is syntactically valie. Body-less loops are often useful.

```java
class Empty3 {
    public static void main(String[] args) {
        int i;
        int sum = 0;
        // sum the numbers through 5
        for (i = 1; i <=5; sum += i++); // No body in this loop!
        System.out.println("Sum is " + sum);
    }
}
```

### Declaring Loop Control Variables Inside the for loop

Often the variable that controls a for loop is needed only for the purposes of the loop and is not used elsewhere. When this is the case, it is possible to declare the variable inside th initialization portion of the for.

```java
class ForVar {
    public static void main(String[] args) {
        int sum = 0; 
        int fact = 1;

        // compute the factorial of the numbers through 5
        for (int i = 1; i <= 5; i++) { // The variable i is declared inside the for statement
            sum += i; // i is known throughout the loop 
            fact *= i;
        }
        // but i is not known here
        System.out.println("Sum is " + sum);
        System.out.println("Factorial is " + fact);
    }
}
```

When you declare a variable inside a for loop, there is one important point to remember: the scope of that variable ends when the for statement does. (That is, the scope of the variable is limited to the for loop). Outside the for loop, the variable will cease to exist. Thus, in the preceding example, i is not accessible outside the for loop.

**Q** Given the flexibility inherent in all of Java's loops, what criteria should I use when selecting a loop? That is, how do I choose right loop for a specific job?
**A** Use a for loop when performing a known number of iterations based on the value of a loop control variable. Use the do-while when you need a loop that will always perform at least one iteration. The while is best used when the loop will repeate until some condition becomes false.

### Use break to Exist a Loop

Here are two other points to remember about break. First, more than one break statement may appear in a loop. However, be careful. Too many break statements have the tendency to destructure your code. Second, the break that terminates a switch statement affects only that switch statement and not any enclosing loops.

### Use break as a Form of goto

In addition to its uses with the switch statement and loops, the break statement can be employed by itself to provide a "civilized" form of the goto statement. 