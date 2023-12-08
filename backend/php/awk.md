```php
<?php

class Test
{
    public static function getNew()
    {
        return new Test();
    }
}

class Child extends Test
{
}

echo get_class(Test::getNew()), "\n"; // Test
echo get_class(Child::getNew()); // Test
```

```php
<?php

class Test
{
    public static function getNew()
    {
        return new static();
    }
}

class Child extends Test
{
}

echo get_class(Test::getNew()), "\n"; // Test
echo get_class(Child::getNew()); // Child
```

```php
<?php
class Test
{
    public static function getNew()
    {
        return new static();
    }
}

class Child extends Test
{
}

$obj3 = Test::getNew();
var_dump($obj3 instanceof Test);

$obj4 = Child::getNew();
var_dump($obj4 instanceof Child);
```
