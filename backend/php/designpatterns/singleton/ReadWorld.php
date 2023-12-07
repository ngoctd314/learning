<?php

class Singleton
{
    private static $instances = [];

    protected function __construct()
    {
    }
    protected function __clone()
    {
    }
    public function __wakeup()
    {
        throw new \Exception("Cannot unserialize singleton");
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

class Logger extends Singleton
{
    private $fileHandle;
    protected function __construct()
    {
        $this->fileHandle = fopen("php://stdout", 'w');
    }
}
