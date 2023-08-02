# Session Hijacking

When a website successfully authenticates a user, the browser and the server open a session. A session is an HTTP conversation in which the browser sends a series of HTTP requests corresponding to user actions, and the web server recognizes them as coming from the same authenticated user without requiring the user to log back in for each request.

If a hacker can access or forge session information that the browser sends, they can access any user's account on your site. Thankfully, modern web servers contain secure session-management code, which makes it practically impossible for an attacker to manipulate or forge a session. However, even if there are no vulnerabilities in a server's session-management capbilities, a hacker can still steal someone else's valid session while it's in progress; this is called session hijacking.

Session hijacking vulnerabilities are generally a bigger risk than the authentication vulnerabilities discussed in the previous chapter, because again, they allow an attacker to access any of your user's accounts.

Three ways hackers hijack sessions: cookie theft, session fixation, and taking advantage of weak session IDs.