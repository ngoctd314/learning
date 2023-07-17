# Updating Arrays in State

Arrays are mutable in JS, but you should treat them as immutable when you store them in state. Just like with objects, when you want to update an array stored in state, you need to create a new one (or make a copy of an existing one), and then set state to use the new array.

**You will learn**

- How to add, remove, or change items in an array in React state
- How to update an object inside of an array
- How to make array copying less repetive with Immer

## Updating ararys without mutation

In JS, arrays are just another kind of object. Like with objects, you should treat arrays in React state as read-only. This means that you shouldn't reassign items inside an array like arr[0] = 'bird', and you also shouldn't use methods that mutate the array, such as push() and pop().

Instead, everytime  you want to update an array, you'll want to pass a new array to your state setting function. To do that, you can create a new array, you'll want to pass a new array to your state setting function. To do that, you can create a new array from the original array in your state by calling its non-mutating methods like filter() and map(). Then you can set your state to the resulting new array.

||avoid (mutates the array)|prefer (returns a new array)|
|-|-|-|
|adding|push, unshift|concat, [...arr]|
|removing|pop, shift, splice|filter, slice|
|replacing|splice, arr[i] = ...|map|
|sorting|reverse, sort|copy the array first|

Unfornately, slice and splice are named similarly but are very different 

- Slice lets you copy an array or part of it
- Splice mutates the array (to insert or delete items)
