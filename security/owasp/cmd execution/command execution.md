# Command Execution

Many web applications call operating system processes via the command line. If your application calls out to the OS, you need to be sure command strings are securely constructed.

Many websites make use of command line calls to read files, send emails, and perform other native operations. If your site transforms untrusted input into shell commands, you need to take care to sanitize the input.