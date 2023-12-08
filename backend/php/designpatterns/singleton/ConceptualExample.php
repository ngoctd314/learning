<?php

/**
 * The Single class defines the 'GetInstance' method that serves as an alternative to constructor and lets clients access the same instance of this
 * class over and over.
 */
class Singleton
{
    /* *
     * The Singleton's instance is stored in a static field. This field is an array
     * because we'll allow our Singleton to have subclasses. Each item in this array
     * will be an instance of a specific Singleton's subclass.
     */
    private static $instances = [];

    /**
     * The Singleton's constructor shoulw always be private to prevent direct
     * construction calls with the `new` operator.
     */
    protected function __construct()
    {
    }
    /**
     * Singletons should not be cloneable.
     */
    protected function __clone()
    {
    }

    /**
     * Singletons should not be restorable from strings
     */
    public function __wakeup()
    {
        throw new \Exception("Cannot unserialize a singleton.");
    }

    public static function getInstance(): Singleton
    {
        $cls = static::class;
        if (!isset(self::$instances[$cls])) {
            self::$instances[$cls] = new static();
        }

        return self::$instances[$cls];
    }
}

function clientCode()
{
    $s1 = Singleton::getInstance();
    $s2 = Singleton::getInstance();

    if ($s1 === $s2) {
        echo "Singleton works, both variables contain the same instance.";
    } else {
        echo "Singleton failed";
    }
}

clientCode();
