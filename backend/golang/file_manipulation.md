# File manipulation

- io package provides an interface for basic I/O primitives and wraps them into shared public interfaces that abstracts the functionality.
- bufio provides an interface for buffered I/O operation with the file. Buffer is actually a temporary space in memory where data is stored and I/O operations are performed from this temporary space. It also means that if we not using bufio we are basically having unbuffered I/O operations. Typically all I/O operations are unbuffered unless specified. The key advantage of having a buffer is that it minimizes system calls as well as disk I/O and is particularly suitable for block transfer of data. This is not suitable for single character-oriented I/O operations.
