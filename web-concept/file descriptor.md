# File descriptor

A file descriptor (FD) is a unique identifier or index used by an operating system to access a file or input/output resource. It's an abstract representation of an open file, socket, or other I/O resource that a process can use for reading from or writing to the associated resource.

In Unix-like operating systems, including Linux, file descriptors are an integral part of the I/O system. Each open file or I/O resource in a process is associated with a file descriptor.

Use ulimit -n command to view the number of file descriptors configured for your Linux system.

```bash
ulimit -n
```

On Linux, processes can use many reasons during their lifetime. The kernel keeps track of the current user's limits for most resources.