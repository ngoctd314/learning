# More Data Types and Operators

## Array

## Multidimensional Arrays

### Two-Dimentional Arrays

### Irregular Arrays

When you allocate memory for a multidimensional array, you need to specify only the memory for the first (leftmost) dimension. You can allocate the remaining dimensions separately. For example, the following code allocates memory for the first dimension of table when it is declared.

It allocates the second dimension manually.

```java
int[][] table = new int[3][];

table[0] = new int[4];
table[1] = new int[4];
table[2] = new int[4];
```

### Alternative Array Declaration Syntax

There is a second form that can be used to declare an array:

### Assigning Array References

As with other objects, when you assign one array reference variable to another, you are simply changing what object that variable refers to. You are not causing a copy of the array to be made, nor are you causing the contents of one array to be copied to the other. For example, consider this program:

### Using the length Member

Recall that in Java, arrays are implemented as objects. One benefit of this approach is that each array has associated with in a length instance variable that contains the number of elements that the array can hold.

```java
public static void main(String[] args) throws java.io.IOException {
    int[] list = new int[10];
    System.out.println(list.length);
}
```

### The For-Each Style for Loop

```java
public static void main(String[] args) throws java.io.IOException {
    int len = 10;
    int[] list = new int[len];
    for (int i = 0; i < len; i++) {
        list[i] = i * i;
    }
    for (int v : list) {
        System.out.print(v + " ");
    }
}
```