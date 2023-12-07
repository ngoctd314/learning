#

**PHP tags**

When PHP parses a file, it looks for opening and closing tags, which are <?php and ?> which tell PHP to start and stop interpreting the code between them. If  a file contains only PHP code, it is preferable to omit the PHP closing tag at the end of the file. This prevents accidental whitespace or new lines being added after the PHP closing tag, which may cause unwanted effects because PHP will start output buffering when there is no intention from the programmer to send any output at that point in the script.

**Escaping from HTML**

Everything outside of a pair of opening and closing tags is ignored by the PHP parser which allows PHP files to have mixed content.


```php
<p>This is going to be ignored by PHP and displayed by the browser.</p>
<?php echo 'While this is going to be parsed.'; ?>
<p>This will also be ignored by PHP and displayed by the browser.</p>
```

This works as expected, because when the PHP interpreter hits the ?> closing tags, it simply starts outputing whatever it finds (except for the immediately following newline)


**Instruction separation**

PHP requires instructions to be terminated with a semicolon at the end statement. The closing tag of a block of PHP code automatically implies a semicolon.

**Data type**

Every single expression in PHP has one of the following built-in types depending on its value:

- null
- bool
- int
- float 
- string
- array
- object
- callable
- resource

PHP is a dynamically typed language, which means that by default these is no need to specify the type of a variable, at this will be determined at runtime. However, it is possible to statically type some aspect of the language via the use of type declarations.

**Type Juggling**

PHP does not require explicit type definition in variable declaration. In this case, the type of a variable is determined by the value it stores. That is to say, if a string is assgined to variable $var, the $var is of type string. If afterwards an int value is assgined to $var, it will be of type int.

PHP may attempt to convert the type of a value to another automatically in certain contexts. The different contexts which exist are:

- Numeric
- String
- Logical
- Integral and string
- Comparative
- Function

**Note** When a value needs to be interpreted as a different type, the value itself does not change types.

To force a variable to be evaluated as certain type, see the section on Type casting.  

**Type System**

PHP uses a nominal type system with a strong behavioral subtyping relation. The subtyping relation is checked at compile time whereas the verification of types is dynamically checked at run time.

PHP's type system supports various atomic types that can be composed together to create more complex types.

**Atomic types**

Some atomic types are built-in types which are tightly integrated with the language and cannot be reproduced with user defined types.

The list of base types is:

- Built-in types: null type, scalar types: [bool, int, float, string], array type, object type, resource type, never type, void type
- Value types: false, true
- User-defined types (generally referred to as class-types): interfaced, classes, enumerations
- Callable type

**Composite types**

It is possible to combine multiple atomic types into composite types. PHP allows types to be combined in the following ways:

- Intersection of class-types (interfaces and class names).
- Union of types.

**Intersection types**

An intersection types accepts values which satisfies multiple class-type declarations, rather than a single one. Individual types which form the intersection type are joined by the & symbol. Therefore, an intersection type comprised of the types T, U and V will be written as T&U&V.

**Union types**

A union type accepts values of multiple different types, rather than a single one. Individual types which form the union type are joined by the | symbol. Therefore, a union type comprised of the types T, U and V will be written as T|U|V. If one of the types is an intersection.

**Type aliases**

PHP supports two type aliases: mixed and iterable which correspond to the union type of object | resource | array | string | float | int | bool | null and Traverable | array respectively.

**NULL**

The null type is PHP's unit type; i.e. it has only one value: null.

Undefined, and unset() variables will resolve to the value null. 

**Syntax**

There is only one value of type null, and that is the case-insensitive constant null.

```php
<?php
$var = NULL;
$var_1 = NulL;
?>
```

**Converting to boolean**

To explicitly convert a value to bool, use the (bool) cast. Generally this is not necessary because when a value is used in a logical context it will be automatically interpreted as a value of type bool.

- the boolean false itself.
- the integer 0 (zero)
- the float 0.0 and -0.0 (zero)
- the empty string "", and the string "0"
- the array with zero elements
- the unit type NULL (including unset variables)
- Internall objects that overload their casting behavior to bool.

**Integers**

An int is a number of the set Z= {..., -2, -1, 0, 1, 2, ...}

Ints can be specified in decimal (base 10), hexadecimal (base 16), octal (base 8) or binary (base 2) notation. The negation operator can be used to denote a negative int.

**Integer overflow**

If PHP encounters a number beyond the bounds of the int type, it will be interpreted as a float instead. Also, an operation which results in a number beyond the bounds of the int type will return a float instead.

```php
<?php
$large_number = 2147483647;
?>
```

**Arrays**

An array in PHP is actually an ordered map. A map is a type that associates values to keys. This type is optimized for several different uses; it can be treated as an array, list (vector), hash table (an implementation of a map), dictionary, collection, stack, queue, and probably more. As array values can be other arrays, trees and multidimensional arrays are also possible.

Explanation of those data structures is beyond the scope of this manual, but at least one example is provided for each of them.

```php
<?php

$arr = array(
    1 => "abc",
    2 => "xyz",
);
foreach ($arr as $k => $v) {
    echo $k , ": ", $v , " ";
}

// Using the short array syntax
$arr = [
    "foo" => "bar",
    "bar" => "foo",
];
foreach ($arr as $k => $v) {
    echo $k . ":" . $v . " ";
}
```

The key can either be an int or a string. The value can be any type.

Additionally the following key casts will occur:

- Strings containing valid decimal ints, unless the number is preceded by a + sign, will be cast to the int type. E.g. the key "8" will actually be stored under 8. On the other hand "08" will not be cast, as it isn't a valid decimal integer.
- Floats are also cast to ints, which means that the fractional part will be truncated. E.g. the key 8.7 will actually be stored under 8.
- Bools are cast to ints, i.e the key true will actually be stored under 1 and the key false under 0.
- Null will be cast to the empty string, i.e key null will actually be stored under "".
- Arrays will be cast to the empty string, i.e the key null actually be stored under "".

```php
<?php

$array = array(
    "a",
    "b",
    1000 => "c",
    "d",
);

var_dump($array);
```

As you can see the last value "d" was assigned the key 7. This is because the largest integer key before that was 6.

**Accessing array elements with square bracket syntax**

Array elements can be accessed using the array[key] syntax. 

```php
<?php
// Create a simple array.
$array = array(1 => "one", 10 => "ten");

// Re-index:
$array = array_values($array);
$array[] = 7;
print_r($array);
```

**Array destructuring**

Arrays can be destructuring using the [] or list() language constructs. These constructs can be used to destructure an array into distinct variables.

```php
<?php
$source_array = ['foo','bar','baz']
[$foo, $bar, $baz] = $source_array;
?>
```

Array destructuring can be used for easy swapping of two variables.

```php
<?php

$a = array(1 => 'one', 2 => 'two', 3 => 'three');
unset($a[2]);
print_r($a);

/*
will produce an array that would have been defined as
$a = array(1 => 'one', 3 => 'three');
and NOT
$a = array(1 => 'one', 3 => 'three');
*/
$b = array_values($a);
print_r($b);
```

**Note:** This does not mean to always quote the key. Do not quote keys which are constants or variables, as this will prevent PHP from interpreting them.

```php
<?php
error_reporting(E_ALL);
ini_set('display_errors', true);
ini_set('html_errors', false);
// Simple array
$array = array(1, 2);
```

**Objects**

To create a new object, use the new statement to instantiate a class:

```php
<?php
class foo 
{
    function do_foo()
    {
        echo "Doing foo.";
    }
}

$bar = new foo;
$bar->do_foo();
```

**Converting to object**

If an object is converted to an object, it is not modified. If a value of any other type is converted to an object, a new instance of the stdClass built-in class is created. If the value was **null**, the new instance will be empty. An array converts to an object with properties named by keys and corresponding values. Note that in this case before PHP 7.2.0 numeric keys have been inaccessible unless iterated.

PHP includes complete object model. Some of its features: visibility, abstract and final classes and methods, additional magic methods, interfaces, and cloning.

PHP treats objects in the same way as references or handles, meaning that each variable contains an object reference rather than a copy of the entire object.

**The Basics**

Basic class definitions begin with the keyword class, followed by a class name, followed by a pair of curly braces which enclose the definitions of the properties and methods to the class.

The class name can be any valid label, provided it is not a PHP reserved word. A valid class name starts with a letter or underscore, followed by and number of letters, numbers or underscores.

**Readonly classes**

**new**

To create an instance of a class, the new keyword must be used. An object will always be created unless the object has a constructor defined that throws an exception on error. Classes should be defined before instantiation (and in some case this is a requirement).

If a string containing the name of a class is used with new, a new instance of that class will be created. If the class is in a namespace, its fully qualified name must be used when doing this.

```txt
Note:
If there no arguments to be passed to the class's constructor, parentheses after the class name may be omitted.
```

**Creating an instance**

```php
<?php

$instance = new SimpleClass();

// This also be done with a variable
$className = 'SimpleClass';
$instance = new $className(); // new SimpleClass()
```

**Creating an instance using an arbitrary expression**

In the given example we show multiple examples of valid arbitrary expressions that produce a class name. This shows a call to a function, string concatenation, and the ::class constant.

```php
<?php

class ClassA extends \stdClass {}
```

**Access member of newly created object**

```php
<?php
echo (new DateTime()) -> format('Y');
```

**Properties and methods**

Class properties and methods live in separate "namespaces", so it is possible to have a properties and a method with the same name. Referring to both a properties and a method has the same notation, and whether a property will be accessed or a method will be called, solely depends on the context, i.e. whether the usage is a variable access or a function call.   

```php
<?php

class Foo
{
    public $bar = 'property';
    public function bar()
    {
        return 'method';
    }
}

$obj = new Foo();
echo $obj->bar . "\n" . $obj->bar();
```

**Calling an anonymous function stored in a property**

```php
<?php
<?php

class Foo
{
    public $bar;

    public function __construct()
    {
        $this->bar = function () {
            return 42;
        };
    }
}

$obj = new Foo();
echo ($obj->bar)();
```

**extends**

A class inherit the constants, methods, and properties of another class by using the keyword extends in the class declaration. It is not possible to extend multiple classes; a class can only inherit from one base class.

The inherited constants, methods, and properties can be override by redeclaring them with the same name defined in the parent class. However, if the parent class has defined a method or constant as final, they may not be overriden.

```php
<?php

class SimpleClass
{
    public function displayVar()
    {
        echo "Var\n";
    }
}

class ExtendClass extends SimpleClass
{
    public function displayVar()
    {
        echo "Extending class\n";
        parent::displayVar();
    }
}

$extended = new ExtendClass();
$extended->displayVar();
```

**Signature compatibility rules**

When overriding a method, its signature must be compatible with the parent method. Otherwise, a fatal error is emmited. 

```php
<?php

class Base
{
    public function foo(int $a)
    {
        echo "Valid $a\n";
    }
}

class Extend1 extends Base
{
    public function foo(int $a = 5)
    {
        parent::foo($a);
    }
}

class Extend2 extends Base
{
    public function foo(int $a, $b = 5)
    {
        parent::foo($a + $b);
    }
}

$obj1 = new Extend1();
$obj1->foo();

$obj2 = new Extend2();
$obj2->foo(10, 20);
```

**Fatal error when a child method removes a parameter**

```php
<?php

class A
{
    public function test($foo, $bar)
    {
    }
}

class B extends A
{
    public function test($a, $b)
    {
    }
}

$obj = new B();
// Error when using named arguments and parameters were renamed in a child class
$obj->test(foo: 'foo', bar: 'bar');
```

**::class**

The class keyword is also used for class name resolution. To obtain the fully qualified name of a class ClassName use ClassName::class.

**Nullsafe methods and properties**

**Properties**

Class member variables are called properties. They may be referred to using other terms such as fields, but for the purposes of this reference properties will be used. 

```php
<?php
class SimpleClass
{
    public $var1 = 'hello' . 'world';
    public $var2 = <<<EOD
    hello world
    EOD;
    // invalid property declarations;
    public $var4 = self::myStaticMethod();
    public $var5 = $myVar;
    // valid property declarations
    public $var6 = myConstant;
    public $var7 = [true, false];

    public $var8 = <<<'EOD'
    hello world
EOD;
    static $var9;
    readonly int $var10;
}
```

**Type declarations**

```php
<?php
class User
{
    public int $id;
    public ?string $name;
    public function __construct(int $id, ?string $name)
    {
        $this->id = $id;
        $this->name = $name;
    }
}

$user = new User(1234, null);

var_dump($user->id);
var_dump($user->name);
```

**Accessing properties**

```php
<?php

class Shape
{
    public ?int $numberOfSides;
    public string $name;

    public function setNumberOfSides(?int $numberOfSides): void
    {
        $this->numberOfSides = $numberOfSides;
    }

    public function setName(string $name): void
    {
        $this->name = $name;
    }

    public function getNumberOfSides(): ?int
    {
        return $this->numberOfSides;
    }

    public function getName(): string
    {
        return $this->name;
    }
}

$triangle = new Shape();
$triangle->setName("triangle");
$triangle->setNumberOfSides(3);

var_dump($triangle->getName());
var_dump($triangle->getNumberOfSides());

$circle = new Shape();
$circle->setName("circle");
$circle->setNumberOfSides(null);
var_dump($circle->getName());
var_dump($circle->getNumberOfSides());
```

```php
<?php

class Test
{
    public readonly string $prop;

    public function __construct(string $prop)
    {
        // Legal initialization.
        $this->prop = $prop;
    }
}

$test = new Test("foobar");
var_dump($test->prop); // string(6) "foobar"

// $test->prop = "foobar";
```

**constant**

```php
<?php

class MyClass
{
    public const CONSTANT = 'constant value';
    public function showConstant()
    {
        echo self::CONSTANT . "\n";
    }
}

echo MyClass::CONSTANT . "\n";

$class = new MyClass();
$class->showConstant();

echo $class::CONSTANT . "\n";
```

**Autoloading Classes**

Many developers writing object-oriented applications create one PHP source file per class definition. One of the biggest anonyances is an having to write a long list of needed includes at the beginning of each script (one for each class).

**Constructors and Destructors**

```php
<?php
__construct(mixed...$value=""): void
```

PHP allows developers to declare constructor methods for classes. Classes which have a constructor method call this method on each newly created object, so it is suitable for any initialization that the object may need before it is used.

Parent constructors are not called implicitly if the child class defines a constructor. In order to run a parent constructor, a call to parent::__construct() within the child constructor is required. If the child does not define a constructor then it may be inherited from the parent class just like a normal class method (if it was not declared as private).

**Constructors in inheritance**

**Using static creation methods**

**Destructor Example**

**Visibility**

Property Visibility: public, protected or private.

Method Visibility: class methods may be defined as public, private, or protected.

```php
<?php

class MyClass
{
    public function __construct()
    {
    }
    public function MyPublic()
    {
        echo "MyPublic\n";
    }
    protected function MyProtected()
    {
        echo "MyProtected\n";
    }
    private function MyPrivate()
    {
        echo "MyPrivate\n";
    }
    public function Foo()
    {
        $this->MyPublic();
        $this->MyProtected();
        $this->MyPrivate();
    }
}

class MyClass2 extends MyClass
{
    // This is public
    public function Foo2()
    {
        $this->MyPublic();
        $this->MyProtected();
    }
}

$obj = new MyClass();
$obj->Foo();


$obj1 = new MyClass2();
$obj1->Foo2();
```

**Constant Visibility**

**Object Inheritance**

**Scope Resolution Operator(::)**

The Scope Resolution Operator (also called Paamayim Nekudotayim) or in simpler terms, the double colon, is a token that allows access to static, constant, and overriden properties or methods of a class.

When referencing these items from outside the class definition, use the name of the class.

```php
<?php

class MyClass
{
    public const CONST_VALUE = 'A constant value';
}

$classname = 'MyClass';

echo $classname::CONST_VALUE;
echo MyClass::CONST_VALUE;
```
