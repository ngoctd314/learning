# Shell Features

## The Shell Versus Programs

When you run a command, it might invoke a Linux program (like who), or instead it might be a built-in command, a feature of the shell itself. You can tell the difference with the type command: 

```sh
type who
# who is /usr/bin/who
type cd
# cd is a shell builtin
```