# Put Back-of-the-envelope Numbers in Perspective

## Why do we use back-of-the-envelope calculations?

Some examples of a back-of-the-envelope calculation could be:

- The number of concurrent TCP connections a server can support.
- The number of requests per second (RPS) a web, database, or cache server can handle.
- The storage requirements of a service.

Since we need good estimations in many design problems:

- The types of data center servers
- The realistic access latencies of different components.
- The estimation of RPC that a server can handle.
- Examples of bandwidth, servers, and storage estimation.

## Types of data center servers

**Data centers** don't have a single type of server. Enterprise solutions use commodity hardware to save cost and develop scalable solutions.

## Web servers

For scalability, the web servers are decoupled from the application servers. **Web servers** are the first point of contract after load balancers.

## Application servers

Application servers run the core application software and business logic. The difference between web servers and application servers is somewhat fuzzy. Application servers primarily provide dynamic content, whereas web servers mostly serve static content to the client, which is mostly a web browser. They can require extensive computational and storage resources. Storage resources can be volatile and non-volatile.

## Storage servers

With the explosive growth of Internet users, the amount of data stored by giant services has multiplied. Additionally, various types of data are now being stored in different storage units.

## Standard numbers to remember

## Requests estimation
