# Chapter 1. Reliable, Scalable, and Maintainable Applications

Many applications today are data-intensive, as opposed to compute-intensive. Raw CPU power is rarely a limting factor for these applications - bigger problems are usually the amount of data, the complexity of data, and the speed at which it is changing.

A data-intensive application is typically built from standard building blocks thaht provide commonly needed functionality. For example, many applications need to:

- Store data to that they, or another application, can find it again later (databases)
- Remember the result of an expensive operation, to speed up reads (caches)
- Allow users to search data by keyword or filter it in various ways (search indexes)
- Send a message to another process, to be handled asynchronously (stream processing)
- Periodically crunch a large amount of accumulated data (batch processing)

## Thinking About Data Systems

We typically think of databases, queues, caches, etc as being very different categories of tools. Although a database anda a message queue have some superficial similarity - both store data for some time - they have very different access patterns, which means different performance characteristics, and thus very different implementations.

If you have an application-managed caching layer (using Memcached or similar), or a full-text search server (such as Elasticsearch or Solr) separate from you main database, it is normally the application code's responsibility to keep those caches and indexes in sync with the main database.

In this book, we focus on three concerns that are important in most software systems:

**Reliability**

The system should continue to work correctly (performing the correct functino at the desired level of performance) even in the face of adversity (hardware or software faults, and even human error).

**Scalability**

As the system grows (in data volume, traffic volume, or complexity) there should be reasonable ways of dealing with that growth.

**Maintainability**

Over time, many different people will work on the system (engineering and operatings, both maintaining current behavior and adapting the system to new use cases), and they should all be able to work on it productively.

## Reliability

For software, typical expectations include:

- The application performs the function that the user expected.
- It can tolerate the user making mistakes or using the software in unexpected ways. 
- Its performance is good enough for the required use case, under the expected load and data volume.
- The systems prevents any unauthorized access and abuse.

If all those things together mean "working correctly" then we can understand reliability as meaning, roughly, "continuing to work correctly, even when things go wrong".

The things can go wrong are called faults, and systems that anticipate faults and can cope with them are called fault-tolerant or resilient. The former term is slightly misleading: it suggests that we could make a system tolerant of every possible kind of fault, which in reality is not feasible. If the entire planet Earth (and all servers on it) were swallowed by a black hold, tolerance of that fault would require web hosting in space. So it only makes sense to talk about tolering certain types of faults.

Note that a fault is not the same as failure. A fault is usually defined as one component of system deviating from its spec, whereas a failure is when the system as a whole stops providing the required service to the user. It is impossible to reduce the probability of a fault to zero; therefore it is usually best to design fault-tolerance mechanisms that prevent faults from causing failures. In this book we cover several techniques for building reliable systems from unreliable parts.

### Hardware Faults

When we think of causes of system failure, hardware faults quickly come to mind. Hard disks crash, RAM becomes faulty, the power grid has a blackout, someone unplugs the wrong network cable. Anyone who has worked with large datacenters cal tell you that these things happen all the time when you have a lot of machines.

Hard disks are reported as having a mean time to failure (MTTF) of about 10 to 50 years. Thus, on a storage cluster with 10000 disks, we should expect on average one disk to die per day.

Our first response is usually to add redundancy to the individual hardware components in order to reduce the failure rate of the system. Disks may be set up in a RAID configuration, servers may have dual power supplies and hot-swappable CPUs, and datacenters may have batteries and diesel generators for backup power. When one component dies, the redundant component can take its place while the broken component is replaced. This approach cannot complete prevent hardware problems from causing failures, but it is well understood and can often keep a machine running uninterrupted for years.

Until recently, redundancy of hardware components was sufficient for most applications, since it makes total failure of a single machine fairly rare. As long as you can restore a backup onto a new machine fairly quickly, the downtime in the case of failure is not catastrophic in most applications. Thus, multi-machine redundancy was only required by a small number of applications for which high availability was absolutely essential.

However, as data volumes and application's computing demands have increased, more applications have begun using larger numbers of machines, which proportionally increases the rate of hardware faults. Moreover, in some cloud platforms such as AWS it is fairly common for vm instances to become unavailable without warning, as the platforms are designed to prioritize flexibility and elasticity over single machine reliability. 

### Software Errors

We usually think of hardware faults as being random and independent from each other: one machine's disk failing does not imply that another machine's disk is going to fail. There may be weak correlations (for example due to a common cause, such as the temperature in the server rack), but otherwise it is unlikely that a large number of hardware components will fail at the same time.

Another class of fault is systematic error within the system. Such faults are harder to anticipate, and because they are correlated across nodes, they tend to cause many more system failures than uncorrelated hardware faults.

- A software bug that causes every instance of an application server to crash when given a particular bad input. For example, consider the leap second on June 30, 2012, that cause many applications to hang simultaneously due to a bug in the Linux kernel.
- A runaway process that uses up some shared resource - CPU time, memory, disk, space, or network bandwidth.
- A service that the system depends on that slows down, becomes unresponsive, or starts returning corrupted responses.
- Cascading failures, where a small fault in one component triggers a fault in another component, which in turn triggers further faults.

The bugs that cause these kinds of software faults often lie dormant for a long time until they are triggered by an unusual set of circumtances. In those circumtances, it is revealed that the software is making some kind of assumption about its environment.

### Human Errors

Humans are known to be unreliable.

How do we make our systems reliable, in spite of unreliable humans? The best systems combine several approaches:

- Design systems in a way that minimizes opportunities for error. APIs, and admin interfaces make it easy to do "the right thing" and discourage "the wrong thing". However, if the interfaces are too restrictive people will work around them, negating their benefit, so this is tricky balance to get right.
- Decouple the places where people make the most mistakes from the places where they can cause failures. In particular, provide fully featured non-production sandbox environments where people can explore and experiment safely, using real data without effecting real users.

...

### How Important Is Reliability?

Reliability is not just for luclear power stations and air traffic control software more mundane applications

## Scalability

Even if a system is working reliably today, that doesn't mean it will necessarily work reliably in the future. One common reason for degradation is increased load: perhaps the system has grown from 10,000 concurrent users to 100,000 concurrent users, or from 1M to 10M. Perhaps it is processing much larger volumes of data than it did before.

Scalability in the term we use to describe a system's ability to cope with increased load. Note, however, that it is not a one-dimensional label that we can attach to a system: it is meaningless to say "X is scalable" or "Y doesn't scale". Rather, discussing scalability means considering questions like "If the system grows in a particular way, what are our options for coping with the growth?" and "How can we add computing resources to handle the addition load?"

### Describing Load

First, we need to succinctly describe the current load on the system; only then can we discuss growth questions (what happens if our load doubles?). Load can be described with a few numbers which we call load parameter.
