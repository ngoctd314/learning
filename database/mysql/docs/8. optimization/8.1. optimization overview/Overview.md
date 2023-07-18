# Optimization Overview

Database performance depends on several factors at the database level, such as tables, queries, and configuration settings.

## Optimization at the Database Level

The most important factor in making a database application fast is its basic design:

- Are the tables structured properly? In particular, do the columns have the right data types, and does each table have the type of work? For example, applications that perform frequent updates often have many tables with few columns, while applications that analyze large amounts of data often have few tables with many columns.
- Are the right indexes in place to make queries efficient?
- Are you usin the appropriate storage engine for each table, and taking advantage of the strengths and features of each storage engine you use? in particular, the choice of a transactional storage engine such as InnoDB or a nontransactional one such as MyISAm can be very important for performance and scalability.

## Optimization at the Hardware Level

Any database application eventually hits hardware limits as the database becomes more and more busy. A DBA must evaluate whether it is possible to tune the application or reconfigure the server to avoid these bottlenecks, or wheather more hardware resources are required. System bottlenecks typically arise from these sources:
- Disk seeks: It takes time for the disk to find a piece of data. With modern disks, the mean time for this is usually lower than 10ms, so we can in theory do about 100 seeks a second. This time improves slowly with new disks and is very hard to optimize for a single table. The way to optimize seed time is to distributed the data onto more than one disk. 
- Disk reading and writing. When the disk is at the correct position, we need to read or write the data. With modern disks, one disk delivers at least 10-20 MB/s throughput. This is easier to optimize than seeks because you can reed in parallel from multiple disks.
- Cpu cycles. When the data is in main memory, we must process it to get our result. Having large tables compared to the amount of memory is the most common limiting factor. But with small tables, speed is usually not the problem.
- Memory bandwidth. When the CPU needs more data than can fit in the CPU cache, main memory bandwidth becomes a bottleneck. This is an uncommon bottleneck for most systems, but one to be aware of.

## Balancing Portability and Performance