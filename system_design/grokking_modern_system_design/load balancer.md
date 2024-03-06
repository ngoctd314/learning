# Load balancer

## Introduction to Load Balancers

### What is load balancing?

Millions of requests cound arrive per second in a typical data center. To serve these requests, thousands (or a hundred thousand) servers work together to share the load of incoming requests.

Here, it's important that we consider how the incoming requests will be divided among all the available servers.

A **load balancer (LB)** is the answer to the question. The job of the load balancer is to fairly divide all client's requests among the pool of available servers. Load balancers perform this job to avoid overloading or crashing servers.

The load balancing layer is the first point of contact within a data center after the firewall. A load balancer may not be required if a service entertains a few hundred or even a few thousand requests per second. However, for increasing client requests, load balancers provide the following capabilities:

- Scalability: By adding servers, the capacity of the application/service can be increased seamlessly. Load balancers make such upscaling or downscaling transparent to the end users.
- Availability: Even if some servers go down or suffer a fault, the system will remains available. One of the jobs of the load balancers is to hide faults and failures of servers.
- Performance: Load balancers can forward requests to servers with a lesser load so the user can get a quicker response time.

### Placing load balancers

Generally, LBs sit between clients and servers. Requests go through to servers and back to clients via the load balancing layer. However, that isn't the only point where load balancers are used.

### Services offered by load balancers

LBs not only enable services to be scalable, available, and highly performant, they offer some key services like the following:

- Health checking: LBs use the heartbeat protocol to monitor the health and, therefore, reliability of end-servers. Another advantage of health checking is the improved user experience.
- TLS termination: LBs reduce the burden on end-servers by handling TLS termination with the client.
- Predictive analytics: LBs can predict traffic patterns through analytics performed over traffic passing through them or using statistics of traffic obtained over time.
- Service discovery: An advantage of LBs is that the client's requests are fowarded to appropriate hosting servers inquiring about the service registry.
- Security: LBs may also improve security by mitigating attacks like DoS at different layers of the OSI model.

**Question**

What if load balancers fail? Are they not a single point of failure (SPOF)?

Load balancers are usually deployed in pairs as a means of disaster recovery. If one load balancer fails, and there's nothing to failover to, the overall service will go down. Generally, to maintain high availability, enterprises use clusters of load balancers that use heartbeat communication to check the health of load balancers at all times. On failure of primary LB, the backup can take over. But, if the entire cluster fails, manual rerouting can also be performed in case of emergencies.

## Global and Local Load Balancing

### Global server load balancing

GSLB ensures that globally arriving traffic load is intelligently forwarded to a data center. For example, power or network failure in a data center requires that all the traffic be rerouted to another data center. GSLB takes forwarding decisions based on the user's geographic locations.

### Load balancing in DNS

We understand that DNS can respond with multiple IP addresses for a DNS query. DNS uses a simple technique of reordering the list of IP addresses in response to each DNS query. Therefore, different users get a reordered IP address list. It results in users visiting a different server to entertain their requests.

### The need for local load balancers

DNS plays a vital role in balancing the load, but it suffers from the following limitations:

- The small size of the DNS packet (512 bytes) isn't enought to include all possible IP addresses of the servers.
- There's limited control over the client's behavior. Clients may select arbitrarily from the received set of IP addresses. Some of the received IP addresses may belong to busy data centers.

### What is local load balancing?

Local load balancers reside within a data center. They behave like a reverse proxy and make their best effort to divide incoming requests among pool of available servers.

## Advanced Details of Load Balancers

### Algorithms of load balancers

Load balancers distribute client requests according to an algorithm. Some well-known algorithms are given below:

**- Round-robin scheduling:** In this algorithm, each request is forwarded to a server in the pool in a repeating sequential manner.
**- Weighted round-robin:** If some servers have a higher capability of serving client's requests, then it's preferred to use a weighted round-robin algorithm. In a weighted round-robin algorithm, each node is assigned a weight. LBs forward clients requests according to the weight of the node. The higher the weight, the higher the number of assignments.
**- Least connections:** In certain cases, even if all the servers have the same capacity to serve clients, uneven load on certain servers is still possibility. For example, some clients have many a request that requires longer to serve. Or some client may subsequent requests on the same connection. In that case, we can use algorithms like least connections where newer arriving requests are assigned to servers with fewer existing connections. 
**- Least response time:** In performance-sensitive services, algorithms such at least response time are required. This algorithm ensures that the server with the least response time is requested to serve the clients.
**- IP hash:** Some applications provide a different level of service to users based on their IP addresses.
**- URL hash:**

### Static versus dynamic algorithms

Algorithms can be static or dynamic depending on the machine's state.

**Static algorithms** don't consider the changing state of the servers. Therefore, task assignment is carried out based on existing knowledge about the server's configuration. Naturally, these algorithms aren't complex, and they get implemented in a single router or commodity machine where all the requests arrive.

**Dynamic algorithms** are algorithms that consider the current or recent state of the servers. Dynamic algorithms maintain state by communicating with the server, which adds a communication overhead.

### Stateful versus stateless LBs

While static and dynamic algorithms are required to consider the health of the hosting servers, a state is maintained to hold session information of different clients with hosting servers.

**Stateful load balancing**

As the name indicates, state load balancing involves maintaining a state of the sessions established between clients and hosting servers. Stateful LBs increase complexity and limit scalability because session information of all the clients is maintained across all the load balancers. That is, load balancers share their state information with each other to make forwarding decisions.

**Stateless load balancing**

Stateless load balancing maintains no state and is, therefore, faster and lightweight. Stateless LBs use consistent hashing to make forwarding decisions.

### Types of load balancers

Depending on the requirements, load balancing can be performed at the network/transport and application layer of the open systems interconnection (OSI) layers.

- Layer 4 load balancers
- Layer 7 load balancers
