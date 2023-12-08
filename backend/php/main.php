<?php

function foo()
{
    function bar()
    {
        echo "I don't exist until foo() is called\n";
    }
}

foo();
bar();
