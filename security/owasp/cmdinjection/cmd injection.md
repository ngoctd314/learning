# Command Injection

Another type of injection attack is command injection, which attackers can use to exploit a website that makes insecure command line calls to the underlying operating system. If your web application makes cmd line calls, make sure to construct your command strings securely. Otherwise, attackers can craft HTTP requests that execute arbitrary os commands, and seize control of your application.

For many programming languages, constructing command strings to invoke operating systems is actually pretty unusual.

## Anatomy of a Command Injection Attack

If your website makes use of command line calls, make sure an attacker can't trick the web server into injection extra commands into the execution.

A command injection targeting Golang applications intends to run system operations without the knowledge of the app's developers. The motive can be reveal hidden system or application information. With extreme motives, hackers can even take over host machine itself from running commands to change their privilege level and passwords.

Command injection -> Change permissions, take ownership, run system cmd, steal source code.

Typically, once a hacker has set their crosshair on your Golang application, they send a couple commands to the server. These commands, if your application is vulnerable, reveal more information that functions as a launchpad for more harmful commands.

**Some exploitable command**

1. Command: whoami

2. Command: uname -a

This is a query of the os on which your're hosting yoru Golang application. Knowning this piece of information narrows the hacker's commands down to your OS.

3. Command: ifconfig

This is perhaps the worst information for a hacker to accquire from your server - your entire network configuration.

4. Command: netstat -an

This command reveals all connections currently feeding and fetching packets to the server.

## Command Injection Example
