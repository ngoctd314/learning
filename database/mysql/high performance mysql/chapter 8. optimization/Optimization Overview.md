# Optimizating Overview

Database performance depends on several factors at the database level, such as tables, queries, and configuration settings. These software constructs result in CPU and I/O operations at the hardware level, which you must minimize and make as efficient as possible.

Typical users aim to get the best database performance out of their existing software and hardware configurations. Advanced users look for opportinities to improve the MySQL software itself.

## Optimizing at the Database Level

- Are the tables structured properly? In particular, do the columns have the right data types, and does each table have the appropriate columns for the type of work?
- Are the right indexes in place to make queries efficient
- Are you using the appropriate storage engine for each table, and taking advantage of the strengths and features of each storage engine you use? In particular, the choice of a transaction storage engine such as InnoDB or nontransaction one such as MyISAM can be very important for performance and scalability. 
- Does each table use an appropriate row format? This choice also depends on the storage engine used for the table. In particular, compressed tables use less disk space and so require less disk I/O to read and write the data. Compression is available for all kinds of workloads with InnoDB tables, and for read-only MyISAM tables.
- Does the application use an appropriate locking strategy? For example, by allowing shared access when possible so that database operations can run concurrently, and requesting exclusive access when appropriate so that critical operations get top priority.
- Are all memory areas used for caching sized correctly? That is, large enough to hold frequently accessed data, but not so large that they overload physical and cause paging. 

## Optimizing at the Hardware Level

Any database application eventually hits hardware limits as the database becomes more and more busy. A DBA must evaluate whether it is possible to tune the application or reconfigure the server to avoid these bottenecks, or whether more hardware resources are required. System bottlenecks typically arise from these sources:

- Disk seeks. It takes time for the disk to find a piece of data. With modern disks, the mean time for this is usually lower than 10ms, so we can in theory do about
- Disk reading and writing. When the disk is the correct position, we need to read or write the data. With modern disks, one disk delivers at least 10-20MB/s throughput. This is easier to optimize than seeks because you can read in parallel from multiple disks.
- CPU cycles. When the data is in main memory, we must process it to get our result. Having large tables compared to the amount of memory is the most common limiting factor. But with small tables, speed is usually not the problem.
- Memory bandwidth. When the CPU needs more data that can fit in the CPU cache, main memory bandwidth becomes a botteneck. This is an uncommon bottleneck for most systems, but one to be aware of.

## Balancing Portability and Performance

To use performance-oriented SQL extensions in a portable MySQL program, you can wrap MySQL-specific keywords in a statement with /*! */ comment delimiters. Other SQL servers ignore the commented keywords.
