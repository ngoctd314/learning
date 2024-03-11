# Building blocks

## The bottom-up approach for modern system design

System design problems usually have some similarities, though specific details are often unique. We have extracted these similarities across design problems as the basic building blocks we'll be covering. One example of a building block is a load-balancing.

**1. Domain Name System:** This building block focuses on how to design hierarchical and distributed naming systems for computers connected to the Internet via different Internet protocols.

**2. Load Balancers:** Here, we'll understand the design of a load balancer, which is used to faily distribute incoming client's requests among a pool of available servers. It also reduces load and can bypass failed servers.

**3. Databases:**

**4. Key-Value Store:** It is a non-relational database that stores data in the form of a key-value pair. Here, we'll explain the design of a key-value store along with important concepts such as a chieving scalability, durability, and configurability.

**5. Content Delivery Network:** In this chapter, we'll design a concent delivery network (CDN) that's used to keep viral content such as videos, images, audio, and webpages. If efficiently delivers content to end users while reducing latency and burden on the data centers.

**6. Sequencer:** In this building block, we'll focus on the design of a unique IDs generator with a major focus on maintaining causality. It also explains three different methods for generating unique IDs.

**7. Service Monitoring:** Monitoring systems are critical in distributed systems because they help analyze the system and alert the stackholders if a problem occurs. Monitoring if often useful to get early warning systems so that system administrators can act ahead of an impending problem becoming a huge issue.

**8. Distributed Caching:**

A distributed caching system where multiple cache servers coordinate to store frequently accessed data.

**9. Distributed Messaging Queue:**

Design of a queue consisting of multiple servers, which is used between interacting entities called producers and consumers. It helps decouple producers and consumers, results in independent scalability, and enhances reliability.

**10. Publish-Subscribe System:**

In this building block, we'll focus on the design of an asynchronous service-to-service communication method called a pub-sub system. It is popular in serverless, microservices architectures and data processing systems.

**11. Rate Limiter:**

Here, we'll design a system that throttles incoming requests for a service based on the predefined limit. It is generally used as a defensive layer for services to avoid their excessive usage-whether intended or unintended.

**12. Blob Store:**

This building block focuses on a storage solution for unstructure data - for example, multimedia files and binary executables.

**13. Distributed Search:**

A search system takes a query from a user and returns relevant content in a few seconds or less. This building block focuses on the three integral components: crawl, index, and search.

**14. Distributed logging:**

Logging is an I/O intensive operation that is time-consuming and slow. Here, we'll design a system that allows services in a distributed system to log their events efficiently. The system will be made scalable and reliable.

**15. Distributed Task Scheduling:**

We'll design a distributed task scheduler system that mediates between tasks and resources. It intelligently allocates resources to tasks to meet task-level and system-level goals. It's often used to offload background processing to be completed asynchronously.

**16. Sharded Counters:**

This building block demonstrates an efficient distributed counting system to deal with millions of concurrent read/write requests.

## Conventions

**- Function requirements:** These represent the features a user of the designed system will be able to use. For example, the system will allow a user to search for content using the search bar.
**- Non-functional requirements:** High availability, low latency, scalability, and so on.
