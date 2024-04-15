# Chapter 4. Introduction Classes, Objects, and Methods

Before you can go much further in your study of Java, you need to learn about the class. The class is the essence of Java. It is foundation upon which the entire Java language is built because the class defines the nature of an object. As such, the class forms the basis for object-oriented programming in Java. Within a class are defined data and code that acts upon that data. The code is contained in methods. Because classes, objects, and methods are fundamental to Java, they are introduced in this chapter. Having a basic understanding of these features will allow you to write more sophisticated programs and better understand certain key Java elements. 

## Class Fundamentals

Java uses a class specification to construct objects. Objects are instances of a class. Thus, a class is essentially a set of plans that specify how to build an object. It is important to be clear on one issue: a class is a logical abstraction. It is not until an object of that class has been created that a physical representation of that class exists in memory.

One other point: Recall that the methods and variables that constitute a class are called members of the class. The data members are also referred to as instance variables.

## The general Form of a Class

```java
class Example {
  public static void main(String[] args) throws java.io.IOException {
    // After this statement executes, p refers to an instance of Vehicle. Thus, it
    // will have physical reality.
    // Each time you create an instance of a class, you are creating an object that
    // contains its own copies
    // of the instances variables.
    // Each object has its own copies of the instance variables defined by its
    // class. Thus, the contents of the variables in one object can differ from the
    // contents of the variables in another. There is not connection between the two
    // objects except for the fact that they are both objects of the same type.
    People p = new People();
    System.out.println(p);
  }
}

// A class definition creates a new data type. In this case, the new data type
// is called People
class People {
  String name;
}
```

### How Objects Are Created

```java
Vehicle minivan = new Vehicle();
```

This declaration performs two functions. First it declares a variable called minvan of the class type Vehicle. This variable doesn't define an object. Instead, it is simply a variable that can refer to an object. Second, the declaration creates an instance of the object and assigns to minivan a reference to that object. This is done by using the new operator.

The new operator dynamically allocates (that is, allocates at run time) memory for an object and returns a reference to it. This reference is, essentially, the address in memory of the object allocated by new. This reference is then sorted in a variable. Thus, in Java, all class objects must be dynamically allocated.

The two steps combined in the preceding statement can be rewritten like this to show each step invididually:

```java
Vehicle minivan;
minivan = new Vehicle();
```

The first like declares minivan as a reference to an object of type Vehicle. Thus, minivan is a variable that can be refer to an object, but it is not an object itself. At this point, minivan does not refer to an object. The next line creates a new Vehicle object and assigns a reference to it to minivan. Now, minivan is linked with an object.

### Reference Variables and Assignment

In an assignment operation, object reference variables act differently that do variables of a primitive type, such as int. When you assign one primitive-type variable to another, the situation is straightforward. The variable on the left receives a copy of the value of the variable on the right. When you assign one object reference variable to another, the situation is a bit more complicated because you are changing the object that the reference variable refers to. The effect of this difference can cause some counterintuitive results. For example, consider the following fragment: 

```java
Vehicle car1 = new Vehicle();
Vehicle car2 = car1;
```

```java
class Example {
  public static void main(String[] args) throws java.io.IOException {
    People p1 = new People();
    People p2 = p1;
    p2.name = "author";
    System.out.println(p1.name);
  }
}

class People {
  String name;
}
```

Although car1 and car2 both refer to the same object, they are not linked in any other way. For example, a subsequent assignment to car2 simply changes the object to which car2 refers. 

```java
Vehicle car1 = new Vehicle();
Vehicle car2 = car1;
Vehicle car3 = new Vehicle();

car2 = car3; // now car2 and car3 refer to the same object.
```

After this sequence executes, car2 refers to the same object as car3. The object refered to by car1 is unchanged.

### Methods

As explained, instance variables and methods are constituents of classes. Methods are subroutines that manipulate the data defined by the class and, in many cases, provide access to that data. In most cases, other parts of your program will interact with a class through its methods. 

If the method does not return a value, its return type must be void. The name of the method is specified by name.

A void method can return in one of two ways - its closing curly brace is reached, or a return statement is executed.

### Returning a Value

### Using Parameters

### Constructors

A constructor initializes an object when it is created. It has the same name as its class and is syntactically similar to a method. However, constructors have no explicit return type. Typically, you will use a constructor to give initial values to the instance variables defined by the class, or to perform any other startup procedures required to create a fully formed object.

All classes have constructors, whether you define one or not, because Java automatically provides a default constructor. In this case, non-initialized member variables have their default values, which are zero, null, and false for numeric types, references types, and booleans respectively. Once you define your own constructor, the default constructor is no longer used.

```java
class Example {
  public static void main(String[] args) throws java.io.IOException {
    People p = new People();
    p.name();
    p.age();
  }
}

class People {
  String name;
  int age;

  void name() {
    System.out.println("Your name is: " + name);
    return;
  }

  void age() {
    System.out.println("Your age is: " + age);
  }
}
```

This constructor assigns the instance variable x of MyClass the value 10. This constructor is called by new when an object is created. For example, in the line

```java
MyClass t1 = new MyClass();
```

### Parameterized Constructors

### The new Operator Revisited

In the context of an assignment, the new operator has this general form:

```java
classvar = new className(args);
```

If a class does not define its own constructor, new will use the default constructor supplied by Java. Thus, new can be used to create an object of any class type. The new operator returns a reference to the newly created object, which (in this case) is assgined to class-var.

Since memory is finite, it is possible that new will not be able to allocate memory for an object because insufficient memory exists. If this happens, a run time exception will occur. 

### Garbage Collection

As you have seen, objects are dynamically allocated from a pool of free memory by using the new operator. As explained, memory is not infinite; and the free memory can be exhausted. Thus, it is possible for new to fail because there is insufficient free memory to create the desired object. For this reason, a key component of any dynamic allocation scheme is the recovery of free memory from unused objects, making that memory available for subsequent reallocation.  

Java's garbage collection system reclaims objects automatically - occuring transparently, behind the scenes, without any programmer intervention. It works like this: When no references to an object exist, that object is assumed to be no longer needed, and the memory occupied by the object is released. This recycled memory can then be used for a subsequent allocation.

Garbage collection occurs only sporadically during the execution of your program. It will not occur simply because one or more objects exist that are no longer used. For efficiency, the gc will usually run only when two conditions are met: there are objects to recycle, and there is a reason to recycle them. Remember, garbage collection takes time, so the Java run-time system does it only when it is appropriate. Thus, you can't know precisely when gc will take place.

### The this Keyword

