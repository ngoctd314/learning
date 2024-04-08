# Java Fundamentals

- Know the history and philosophy of Java
- Understand Java's contribution to the Internet
- Understand the importance of bytecode
- Know the Java buzzwords
- Understand the foundational principles of object-oriented programming
- Create, compile, and run a simple Java program
- Use variables
- Use the if and for control statements
- Create blocks of code
- Understand how statements are positioned, intended, and terminated
- Know the Java keywords
- Understand how statements are positioned, intented, and terminated
- Know the Java keywords
- Understand the rules for Java identifiers

**Java's Magic: The Bytecode**

The key that allowed Java to address both the security and the portability problems just described is that the output of a Java compiler is not executable code. Rather, it is bytecode. Bytecode is a highly optimized set of instructions designed to be executed by what is called the Java Virtual Machine (JVM), which is part of the Java Runtime Environment (JRE). In essence, the original JVM was designed as an interpreter for bytecode. This may come as a bit of a surprise because many modern languages are designed to be compliled into CPU-specific, executable code due to performance concerns. However, the fact that a Java program is executed by the JVM helps solve the major problems associated with web-based programs.

Translating a Java program into bytecode makes it much easier to run a program in a wide variety of environments because only the JRE (which includes the JVM) needs to be implemented for each platform. Once a JRE exists for a given system, any Java program can run on it. Remember, although the details of the JRE will differ from platform to platform, all JREs understand the same Java bytecode. If a Java program were compiled to native code, the different versions of the same program would have to exist for each type of CPU connected to the Internet.

The fact that a Java program is executed by the JVM also helps to make it secure. Because the JVM is in control, it manages program execution. Thus, it is possible for the JVM to create a restricted execution environment, called the sandbox, that contains the program, preventing unrestricted access to the machine. Safety is also enhanced by certain restrictions that exist in the Java language.

When a program is interpreted, it generally runs slower than the same program would run if compiled to executable code. However, with Java, the differential between the two is not so great. Because bytecode has been highly optimized, the use of bytecode enables the JVM to execute programs much faster than you might expect.

Although Java was designed as an interpreted language, there is nothing about Java that prevents on-the-fly compilation of bytecode into native code in order to boost performance. For this reason, the HotSpot JVM was introduced not long after Java's initial release. HotSpot includes a just-in-time (JIT) compiler for bytecode. When a JIT compiler is part of the JVM, selected portions of bytecode are compiled into executable code in real time on a piece-by-piece demand basic. That is, a JIT compiler compiles codes as it is needed during execution. Furthermore, not all sequences of bytecode are compiled - only those that will benefit from compilation. The remaining code is simply interpreted. However, the just-in-time approach still yields a significant performance boost. Even when dynamic compilation is applied to bytecode, the portability and safety features still apply because the JVM is still in charge of the execution environment.

One other point: Beginning with JDK 9, some Java env will also support an ahead-of-time compiler that can be used to compile bytecode into native code prior to execution by the JVM, rather than on-the-fly. Ahead-of-time compilation is a specialized feature and it does not replace Java's traditional, approach just described.

**Q: I have heard about a special type of Java program called a servlet. What is it?**

**A:** A Java servlet is a small program that executes on a server. Servlets dynamically extend the functionality of web server. It is helpful to understand that as useful as client-side applications can be, they are just one half of the client/server equation. Not long after the initial release of Java, it became obvious that Java would also be useful on the server side. The result was the servlet. Thus, with the advent of the servlet, Java spanned both sides of the client/server connection. Although the topic of servlets, and server-side programming in general, is beyond the scope of this beginner's guide.

**Moving Beyond Applets**

Beginning with JDK 9, the phase-out of applets was begun, with support for applets being deprecated. In the language of Java, deprecated means that a feature is still available but flagged as obsolete. Thus, a deprecated feature should not be used for new code. The phase-out become complete with the release of JDK 11 because support for applets was removed.

## The Java Buzzwords

The key considerations were summed up by the Java design team in the following list of buzzwords:

|Keyword|Description|
|-|-|
|Simple|Java has a concise, cohesive set of features that makes it easy to learn and use.|
|Secure|Java provides a secure means of creating Internet applications.|
|Portable|Java programs can execute in any environment for which there is a Java can runtime system.|
|Object-oriented|Java embodies the modern object-oriented programming philosophy.|
|Robust|Java encourages error-free programming by being strictly typed and performing run-time checks.|
|Multithreaded|Java provides integrated support for multithreaded programming.|
|Architecture-neutral|Java is not tied to a specific machine or operating system architecture|
|Interpreted|Java supports cross-platform code through the use of Java bytecode.|
|High performance|The Java bytecode is highly optimized for speed of execution.|
|Distributed|Java was designed with the distributed environment of the Internet in mind.|

## OBJECT-ORIENTED PROGRAMMING

At the center of Java is object-oriented programming (OOP). The object-oriented methodology is inseparable from Java, and all Java program are.

OOP is a powerful way to approach the job of programming. Programming methodologies have changed dramatically since the invention of the computer, primarily to accommodate the increasing complexity of programs. For example, when computers were first invented, programming was done by toggling in the binary machine instructions using the computers front panel. As long as programs were just a few hundred instructions using the computer's front panel. 

Object-oriented programming took the best ideas of structured programming and combined them with several new concepts. The result was a different way of organizing a program. In the most general sense, a program can be organized in one of two ways: around its code (what is happening) or around its data (what is being affected). Using only structured programming techniques, programs are typically organized around it code.

Object-oriented programs work the other around. They are organized around data, with the key principle being "data controlling access to code". In an object-oriented language, you define the data and the routines that are permitted to act on that data. Thus, a data type defines precisely what sort of operations can be applied to that data.

To support the principles of object-oriented programming, all OOP languages, including Java, have three traits in common: encapsulation, polymorphism, and inheritance.

### Encapsulation

Encapsulation is a programming mechanism that binds together code and the data it manipulates, that keeps both safe from outside interference and misuse.

With an object, code, data, or both may be private to that object or public. Private code or data is known to and accessible by only another part of the object. That is, private code or data cannot be accessed by a piece of the program that exists outside the object. When code or data is public, other parts of your program can access it even though it is defined within an object. Typically, the public parts of an object are used to provide a controlled interface to the private elements of the object.

Java's basic unit of encapsulation is the class. A class defines the form of an object. It specifies both the data and the code that will operate on that data. Java uses a class specification to construct objects. Objects are instances of a class. Thus, a class essentially a set of plans that specify how to build an object.

### Polymorphism

Polymorphism (from Greek, meaning "many forms") is the quality that allows one interface to access a general class of actions.

The same principle can also apply to programming. For example, consider a stack. You might have a program that requires three different types of stacks. One stack is used for integer values, one for floating-point values, and one for characters. In this case, the algorithm that implements each stack is the same, even though the data being stored differs. In a non-object-oriented language, you would be required to create three different sets of stack routines, with each set using different names.

More genrally, the concept of polymorphism is often expressed by the pharse "one interface, multiple methods". This means that it is possible to design a generic inteface to group of related activities.

### Inheritance

Inheritance is the process by which one object can accquire the properties of another object. This is important because it supports the concept of hierarchical classification.

Without the use of hierarchies, each object would have to explicitly define all of its characteristics. Using inheritance, an object need only define those qualities that make it unique within its class. It can inherit its general attributes from its parent.

## The Java Development Kit

Before you can compile and run those programs, you must have a Java Development Kit (JDK). At the time of this writing, the current release of the JDK is JDK 11. This is the version for Java SE 11. (SE stands for Standard Edition). It is also the version describe in this book. Because JDK 11 contains features that are not supported by earlier versions of Java, it is recommended that you use JDK 11 (or later) to compile and run the programs in this book. 

The JDK supplies two primary programs. The first is javac, which is the Java compiler. The second is java, which is the standard Java interpreter and is also referred to as the application launcher.

**You state that object-oriented programming is an effective way to manage large programs. However, it seems that it might add substantial overhead to relatively small ones. Since you say that all Java programs are, to some extent, object-oriented, does this impose a penalty for smaller programs**

## A frist simple program

```java
class Example {
    public static void main(String[] args) {
        System.out.println("Java drives the Web.");
    }
}
```

### Entering the Program

For most computer languages, the name of the file that holds the source code to a program is arbitrary. However, this is not the case with Java. The first thing that you must learn about Java is that the name you give to a source file is very important. For this example, the name of the source file should be Example.java.

In Java, a source file is officially called a compilation unit. It is a text file that contains (among other things) one or more class definitions. (For now, we will be using source files that contain only one class). The Java compiler requires that a source file use the .java filename extension. As you can see by looking at the program, the name of the class defined by the program is also **Example**. This is not a coincidence. In Java, all code must reside inside a class. By convention, the name of the main class should match the name of the file that holds the program. You should also make sure that the capitialization of the filename matches the class name. The reason for this is that Java is case sensitive. At this point, the convention makes it easier to maintain and organize your programs.

### Compiling the Program

To compile the Example program, execute the compiler, javac, specifying the name of the source file on the command line:

```bash
javac Example.java
```

The javac compiler creates a file called Example.class that contains the bytecode version of the program. Remember, bytecode is not excutable code. Bytecode must be executed by a JVM. Thus, the output of javac is not code that can be directly executed.

To actually run the program, you must use the Java interpreter, java. To do so, pass the class name Example as a command-line argument, as shown here: 

```bash
java Example
```

When Java source code is compiled, each individual class is put into its own output file named after the class and using the .class extension. This is why it is a good idea to give your Java source files the same name as the class the contain - the name of the source file will match the name of the .class file. When you execute the Java interpreter as just shown, you are actually specifying the name of the class that you want the interpreter to execute. It will automatically search for a file by that name that has the .class extension. If it finds the file, it will execute the code contained in the specified class.

Before moving on, it is important to mention that beginning with JDK 11, Java provides a way to run some types of simple programs directly from a source file, without explicitly invoking javac.

### The First Sample Program Line by Line
