# Cross-site script attacks

If your web server is secure, a hacker's next best injectiontarget is the web browser. If an attacker can find a way to inject malicious JavaScript into a user's browser while the user views your website, that user is in for a bad time. We call this type of code injection a cross-site script (XSS) attack.

JavaScript can read or modify any part of a web page, so there's a lot of attacker can do with cross-site scripting vulnerabilities. If JS can read HTTP session information, they can hijack a user's session entirely, allowing them to log in as that user remotely.

## Stored Cross-Site Scripting Attacks

