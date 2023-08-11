# Recursion

## Guide line

- What's the simplest possible input?
- Play around with examples and visualize
- Relate hard cases to simpler cases
- Generalize the pattern
- Write code by combining recursive pattern with the base case

## Recursion

Recursion is a tool not often used by imperative language developers because it is thought to be slow and to waste space. But as you'll see, there are several techniques that can be used to minimize or eliminate these problems.

For new computer science students, the concept of recursive programming is often difficult. Recursive thinking is difficult because is almost seems like circular reasoning. It's also not an intuitive process; when we give instructions to other people, we rarely direct them recursively.

For those of you who are new to computer programming, here's a simple definition of recursion: Recursion occurs when a function calls itself directly or indirectly.

## Basic steps of recursive programs

Every recursive program follows the same basic sequence of steps:

1. Initialize the algorithm. Recursive programs often need a seed value to start with. This is accomplished by using a parameter passed to the function or by providing a gateway function that is non-recursive but that sets up the seed values for the recursive calculation.

2. Check to see whether the current value(s) being processed match the base case. If so, process and return the value.

3. Redefine the answer in terms of a smaller or simpler sub-problem or sub-problems.

4. Run the algorithm on the sub-problem.

5. Combine the results in the formulation of the answer.

6. Return the results.

## Using an inductive definition

Sometimes when writing recursive programs, finding the simpler sub-problem can be tricky. Dealing with inductively-defined data sets, however, makes finding the sub-problem considerably easier. An inductive-defined data set is a data structure defined in terms of itself -- this is called an inductive definition.

For example, linked lists are defined in terms of themselves. A linked list consists of a node structure that contains two members: the data it holding and a pointer to another node structure (or NULL, to terminate the list). Because the node structure contains a pointer to a node structure within it, it is said to defined inductively. 

With inductive data, it is fairly easy to write recursive procedures. Notice how like our recursive programs, the definition of a linked list also contains a base case -- in this case, the NULL pointer. Since a NULL pointer terminates a list, we can also use the NULL pointer condition as a base case for many of our recursive functions on linked lists.

## Linked list example

Let's look at a few examples of recursive functions on linked lists. Suppose we have a list of numbers, and we want to sum them. Let's go though each step of the recursive sequence and identify how it applies to our summation function:

1. Initialize the algorithm. This algorithm's seed value is the first node to process and is passed as a parameter to the function.

2. Check for the base case. The program needs to check and see if the current node is the NULL list. If so, we return zero because the sum of all members of an empty list is zero.

3. Redefine the answer in terms of a simpler sub-program. We can define the answer as the sum of the rest of the list plus the contents of the current node. To determine the sum of the rest of the list, we call this function again with the next node.

4. Combine the results. After the recursive call completes, we add the value of the current node to the results of the recursive call.

```go
func sum_list(l *Node) int {
    if l == NULL {
        return 0
    }
    return l.Val + sum_list(l.Next)
}
```
You make be thinking that you know how write this program to perform faster or better without recursion. We will get to speed and space issues of recursion later on.

