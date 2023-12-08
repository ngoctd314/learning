# Errors in PHP 7

As with normal exceptions, these Error exceptions will bubble up until they reach the first matching catch block. If there are no matching blocks, then any default exception handler installed with set_exception_handler() will be called, and if there is no default exception handler, then the exception will be converted to a fatal error and will be handled like a traditional error.
