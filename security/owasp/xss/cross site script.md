# Cross-site script attacks

If your web server is secure, a hacker's next best injectiontarget is the web browser. If an attacker can find a way to inject malicious JavaScript into a user's browser while the user views your website, that user is in for a bad time. We call this type of code injection a cross-site script (XSS) attack.

JavaScript can read or modify any part of a web page, so there's a lot of attacker can do with cross-site scripting vulnerabilities. If JS can read HTTP session information, they can hijack a user's session entirely, allowing them to log in as that user remotely.

Typically, an XSS attack attempts to fetch content from previously unknown sources and bring it into the browser. Then, whatever payload is embedded in the cross-site script runs commands or renders content in place of the developer's original intent.

IFrame are a common XSS attack target on Golang applications.

## Stored Cross-Site Scripting Attacks

Websites routinely generate and render HTML using information stored in a database. Retail websites will store product information in a database, and social media sites will store user conversations. Websites will take content from the database according to the URL user has navigated to, and interpolate it into the page to produce the finished HTML.

Any page content coming from the database is a potential attack vector for hackers. Attackers will attempt to inject Javascript code into the database so that the web server will write out the JavaScript when it render HTML. We call this type of attack a stored cross-site scripting attack: the Javascript is written to the database, but executed in the browser when an unsuspecting victim views a particular page on the site.

Malicious JavaScript can be planted in a database by using the SQL injection method, but attackers will more common insert malicious code through legitimate avenues. For instance, if a website allows users to post comments, the site will store the comment text in a database and display it back to other users who view the same comment thread. In this scenario, an easy way for a hacker to perform a cross-site scripting attack is to write a comment containing a `<script>` tag to the database. If the website fails to construct HTML securely, the `<script>` tag will get written out whenever the page is rendered to other users, and the JS will be executed in the victim's browser.

### Mitigation 1: Escape HTML Characters

### Mitigation 2: Implement a Content Security Policy

## Reflect Cross-Site Scripting Attacks

### Mitigation: Escape Dynamic Content from HTTP Requests

## DOM-Based Cross-Site Scripting Attacks
