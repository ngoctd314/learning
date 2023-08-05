# Command Execution

Many web applications call operating system processes via the command line. If your application calls out to the OS, you need to be sure command strings are securely constructed.

Many websites make use of command line calls to read files, send emails, and perform other native operations. If your site transforms untrusted input into shell commands, you need to take care to sanitize the input.

If you don't, an attacker will able to craft HTTP requests that execute whatever command they want.

**Being vulnerable to remote code execution is a severe risk. Better learn to protect ourselves eh?**

If an attacker can execute arbitrary code on your servers, your system are almost certainly going to be compromised. You need to take great care when designing how your web server interacts with the underlying os.

## RISKS

Command injection is a major security lapse, and the last step along the road to complete system takeover. After gaining access, an attacker will attempt to escalate their privileges on the server, install malicious scripts, or make your server part of a botnet to be used at a later date.

Command injection vulnerabilities often occur in order, legacy code, such as CGI scripts.

## PROTECTION

If your application calls out to the os, you need to be sure command strings are securely constructed, or else you risk having malicious instructions injected by an attacker.

**1. Try To Avoid Command Line Calls Altogether**

**2. Escape Inputs Correctly**

**3. Restrict the Permitted Commands**

**4. Perform Thorough Code Reviews**

**5. Run with Restricted Permissions**