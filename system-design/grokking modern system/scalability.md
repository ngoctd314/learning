# Scalability

## What is scalability?

Scalability is the ability of a system to handle an increasing amount of workload without compromising performance. A search engine, for example, must accommodate increasing numbers of users, as well as the amount of data it indexes.

The workload can be of different types, including the following:
- Request workload: This is the number of requests served by the system.
- Data/storage workload: This is the amount of data stored by the system. 

## Dimensions

## Different approaches of scalability

### Vertical scalability - scaling up

Vertical scaling, also known as "scaling up", refers to scaling by providing additional capabilities (for example, additional CPUs or RAM) to an existing device. Vertical scaling allows us to expand our present hardware or software capability, but we can only grow it to the limitations of our server. The dollar cost of vertical scaling is usually high because we might need exotic components to scale up.

### Horizontal scalability - scaling out

Horizontal scaling, also known as "scaling out" refers to increasing the number of machines in the network. We use commodity nodes for this purpose because of their attractive dollar-cost benefits. The catch here is that we need to build a system such that many nodes could collectively work as if we had a single, huge server.
