# MySQL High Performance

[Chapter 1. MySQL Architecture and History](#chapter-1.-mysql-architecture-and-history)<br>
[Chapter 2. Benchmarking](#chapter-2.-benchmarking)<br>
[Chapter 3. Profiling Server Performance](#chapter-3.-profiling-server-performance)<br>

## Chapter 1. MySQL Architecture and History

### MySQL's Logical Architecture

Client -> Connection / thread handling -> Query cache -> Parser -> Optimizer -> Storage engine.

The third layer contains the storage engines. They are responsible for storing and retriving all data stored "in" MySQL. Like the various filesystems available for GNU/Linux, each storage engine has it own benefits and drawbacks.

#### Connection Management and Security

Each client connection gets its own thread within the server process. The connection's queries execute within that single thread, which in turn resides on one core or CPU. The server caches threads, so they don't need to be created and destroyed for each new connection.

#### Optimization and Execution

MySQL parse queries to create an internal structure (the parse tree), and then applies a variety of optimizations.

The optimizer does not really care what storage engine a particular table uses, but the storage engine does affect how the server optimizes the query. The optimizer asks the storage engine about some of its capabilities and the cost of certain operations, and for statistics on the table data. 

Before even parsing the query, though, the server consults the query cache, which can store only SELECT statements, along with their result sets. If anyone issues a query that's identical to one already in the cache, the server doesn't need to parse, optimize, or execute the query at all - it can simply pass back the stored result set.

### Concurrency Control

Anytime more than one query needs to change data at the same time, the problem of concurrency control arises. MySQL has to do this at two levels: the server level and the storage engine level.

#### Read/Write Locks

Systems that deal with concurrent read/write access typically implement a locking system that consits of two lock types. These locks are usually known as shared locks and exclusive lock, or read locks and write locks. Read locks on a resource are shared, or mutually nonblocking: many clients can read from a resource at the same time and not interfere with each other. Write locks, on the other hand, are exclusive - they block both read locks and other write locks - because the only safe policy is to have a single client writing to the resource at a given time and to prevent all reads when a client is writing.

In the database world, locking happens all the time: MySQL has to prevent one client from reading a piece of data while another is changing it. 

#### Lock Granularity

One way to improve the concurrency of a shared resource is to be more selective about what you lock. Rather than locking the entire resource, lock only the part that contains the data you need to change. Better yet, lock only the exact piece of data you plan to change. Minimizing the amount of data that you lock at any one time lets changes to a given resource occur simultaneously, as long as they don't conflict with each other.

The problem is locks consume resources. Every lock operation - getting a lock, checking to see a lock is free, releasing a lock, and so on - has overhead. If the system spends too much time managing locks instead of storing and retrieving data, performance can suffer.

A locking strategy is a compromise between lock overhead and data safety, and that compromise affects performance. Most commercial database server don't give you much choice: you get what is known as row-level locking in your tables, with a variety of often complex way to give good performance with many locks.

**Table locks**

The most basic locking strategy available in MySQL, and the one with the lowest overhead, is table locks. It locks the entire table. When nobody is writing, readers can obtain read locks, which don't conflict with other read locks.

Table locks have variations for good performance in specific situations. For example, READ LOCAL table locks allow some types of concurrent write operations. Write locks also have a higher priority than read locks, so a request for a write lock will advance to the front of the lock queue even if readers are already in the queue (write locks can advance past read locks in the queue, but read locks cannot advance past write locks).

Although storage engines can manage their own locks, MySQL itself also uses a variety of locks that are effectively table-level for various purposes.

**Row locks**

The locking style that offers the greatest concurrency (and carries the greatest overhead) is the use of row locks. Row-level locking, as this strategy is commonly known, is available in the InnoDB. Row locks are implemented in the storage engine, not the server. The server is completely unaware of locks implemented in the storage engines.

TODO

### Transactions

A transaction is a group of SQL queries that are treated atomically, as a single unit of work. If the database engine can apply the entire group of queries to a database, it does so, but if any of them can't be done because of a crash or other reason, none of them is applied. It's all or nothing.

Transactions aren't enough unless the system passes the ACID test. ACID stands for Atomicity, Consistency, Isolation, and Durability. These are tightly related criteria that a well-behaved transaction processing system must meet:

*Atomicity*

A transaction must function as a single indivisible unit of work so that the entire transaction is either applied or rolled back. When transactions are atomic, there is no such thing as a partially commpleted transaction: it's all or nothing.

*Consistency*

The database should always move from one consistence state to the next.

*Isolation*

The results of a transaction are usually invisible to other transactions until the transaction is complete.

*Durability*

Once commited, a transaction's changes are permanent. This means the changes must be recorded such that data won't be lost in a system crash. Durability is a slight fuzzy concept, however, because there are actually many levels. Some durability strategies provide a stronger safety guarantee than others, and nothing is ever 100% durable.

#### Isolation Levels

Each storage engine implements isolation levels slightly differently, and they don't necessarily match what you might expect if you're used to another database product.

### Multiversion Concurrency Control

### MySQL's Storage Engines

When you create a table, MySQL stores the table definition in a .frm file with the same as the table. Thus, when you create a table named MyTable, MySQL stores the table definition in MyTable.frm. Because MySQL uses the filesystem to store the table definition in MyTable.frm. Each storage engine stores the table's data and indexes differently, but the server itself handles the table definition.

## Chapter 2. Benchmarking
