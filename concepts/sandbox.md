# Sandbox

**1. What is sandbox**

What is a sandbox? A sandbox is an isolated testing environment that enables users to run programs or open files.

In the world of cybersecurity, a sandbox environment is an isolated virtual machine in which potentially unsafe code can execute without affecting network resources or local applications.

Cybersecurity researchers use sanboxes to run suspicious code from unknown attachments and URLs and observe its behavior. Because the sandbox is an emulated environment with no access to the network, data or other applications, security teams can safely "detonate" the code to determine how it works and whether it is malicious.

Outside of cybersecurity, developers also use sandbox testing environments to run code before widespread deployment.

**2. What is the Purpose of a Sandbox**

Is a standard business production environment, a sandbox might be misunderstood or considered a needless expensive. But sanboxes are critical for several scenarios in development, cybersecurity and research. Making sure the sandbox is truly isolated and secure is more important in cybersecurity research than in software development because malware actively and aggressively scans the network for exploitable vulnerabilities.

**In development**

In development, a sandbox usually involves a development server and staging server. The development server is separated from the production environment but may still require basic network access. Developers use this server to upload code and test it as the codebase changes.

The staging server is designed to be an extract replica of production. This server is where quality assurance (QA) tests code before deploying to production. Because the staging environment is the same as the production environment, code that runs without issues in staging should run without issues in production.

**In cybersecurity research**

The sandbox must not be have any access to critical infrastructure.
