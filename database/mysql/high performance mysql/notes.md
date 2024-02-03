# MySQL High Performance

[Chapter 1. MySQL Architecture and History](#chapter-1.-mysql-architecture-and-history)<br>
[Chapter 2. Benchmarking](#chapter-2.-benchmarking)<br>
[Chapter 3. Profiling Server Performance](#chapter-3.-profiling-server-performance)<br>
[Chapter 5. Indexing for High Performance](#chapter-5.-indexing-for-high-performance)<br>

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

You can use the SHOW TABLE STATUS command (or in MySQL 5.0 and newer versions, query the INFORMATION_SCHEMA tables) to display information about tables.

```sql
SELECT * FROM INFORMATION_SCHEMA.TABLES WHERE TABLE_SCHEMA = 'learn' AND TABLE_NAME = 'tbl'\G
```

#### The InnoDB Engine

InnoDB was designed for processing many short-lived transactions that usually complete rather than being rolled back. Its performance and automatic crash recovery make it popular for nontransactional storage needs, too.

**InnoDB's history**

**InnoDB overview**

InnoDB stores its data in a series of one or more data files that are collectively known as a tablespace. A tablespace is essentially a black box that InnoDB manages all by itself. In MySQL 4.1 and newer versions, InnoDB can store each table's data and indexes in separate files. InnoDB can also use raw disk partitions for building its tablespace, but modern filesystems make this unnecessary.

##### InnoDB's history

##### InnoDB overview

## Chapter 2. Benchmarking

## Chapter 3. Profiling Server Performance

## Chapter 4. Optimizing Schema and Data Types

Good logical and physical design is the cornerstone of high performance, and you must design your schema for the specific queries you will run. This often involves trade-offs. For example, a denormalized schema can speed up some types of queries but slow down others. Adding counter and summary tables is a great way to optimize queries, but they can be expensive to maintain.

Book for database design: Clare Churcher's book Beginning Database Design (Apress).

### Choosing Optimal Data Types

*Smaller is usually better*

In general, try to use the smallest data type that can correctly store and represent your data. Smaller data types are usually faster, because they use less space on the disk, in memory, and in the CPU cache. They also generally require fewer CPU cycles to process.

Make sure you don't underestimate the range of values you need to store, though, because increasing the data type range on multiple places in your schema can be a painful and time-consuming operation. If you're in doubt as to which is the best data type to use, choose the smallest one that you don't think you'll exceed.

*Simple is good*

Fewer CPU cycles are typically required to process operations on simpler data types. For example, integers are cheaper to compare than characters, because character sets and collations (sorting rules) make character comparisons complicated. Here are two examples: you should store dates and times in MySQL's built-in types instead of strings, and you should use integers for IP addresses.

*Avoid NULL if possible*

A lot of tables include nullable columns even when the application does not need to store NULL (the absence of a value), merely because it's the default. It's usually best to specify columns as NOT NULL unless you intend to store NULL in them.

It's harder for MySQL to optimize queries that refer to nullable columns, because they make indexes, index statistics, and value comparisons more complicated. A nullable column uses more storage space and requires special processing inside MySQL. When a nullable column is indexed, it requires an extra byte per entry and can even cause a fixed-size index.

#### Whole Numbers

TINYINT, SMALLINT, MEDIUMINT, INT, or BIGINT.

#### Real Numbers

#### String Types

Each string column can have its own character set and set of sorting rules for that character set, or collation.

#### VARCHAR and CHAR types

The two major string types are VARCHAR and CHAR, which store character values. Unfornately, it's hard to explain exactly how these values are stored on disk and in memory, because the implementations are storage engine-dependent. We assume you are using InnoDB and/or MyISAM. If not, you should read the documentation for your storage engine.

Be aware that a storage engine may store a CHAR or VARCHAR value differently in memory from how it stores that value on disk, and that the server may translate the value into yet another storage format when it retrieves it from the storage engine.

**VARCHAR**

VARCHAR uses 1 or 2 bytes to record the value's length.

VARCHAR helps performance because it saves space. However, because the rows are variable-length, they can grow when you update them, which can cause extra work. If a row grows and no longer fits in original location, the behavior is storage engine-dependent. For example, InnoDB may need to split the page to fit the row into it.

In version 5.0 and newer, MySQL presers trailing spaces when you store and retrieve values.

**CHAR**

CHAR is fixed-length: MySQL always allocates enough space for the specified number of characters. When storing a CHAR value, MySQL removes any trailing spaces. Values are padded with spaces as needed for comparisons.

**Generosity Can Be Unwise**

Storing the value 'hello' requires the same amount of space in a VARCHAR(5) and a VARCHAR(200) column. Is there any advantage to using the shorter column? 

As it turns out, there is a big advantage. The larger column can use much more memory, because MySQL often allocates fixed-size chunks of memory to hold values internally. This is especially bad for storing or operations that use in-memory temporary tables. The same thing happens with filesorts that use on-disk temporary tables.

The best strategy is to allocate only as much space as you really need.

#### BLOB and TEXT types

BLOG and TEXT are string data types designed to store large amounts of data as either binary or character strings, respectively.

The only difference between the BLOG and TEXT families is that BLOB types store binary data with no collation or character set, but TEXT types have a character set and collation.

MySQL can't index the full length of these data types and can't use the indexes for sorting. (You'll find more on these topics in the next chapter).

#### Using ENUM instead of a string type

#### Date and Time Types

#### Bit-Packed Data Types

#### Choosing Identifiers

*Integer types*

Integers are usually the best choice for identifiers, because they're fast and they work with AUTO_INCREMENT

*ENUM and SET*

The ENUM and SET types are generally a poor choice for identifiers. ENUM and SET columns are appropriate for holding information such as an order's status, a product's type, or a person gender.

*String types*

Avoid string types for identifiers if possible, because they take up a lot of space and are generally slower than integer types.

You should also be very careful with completely "random" strings, such as those produced by MD5(), SHA1(), or UUID(). Each new value you generate with them will be distributed in arbitrary ways over a large space, which can slow INSERT and some types of SELECT queries.

- They slow INSERT queries because the inserted value has to go in a random location in indexes. This causes page splits, random disk accesses, and clustered index fragmentation for clustered storage engines.
- They slow SELECT queries because logically adjacent rows will be widely dispersed on disk and in memory.
- Random values cause caches to perform poorly for all types of queries because they defeat locality of reference, which is how caching works. If the entire dataset is equally "hot", there is no advantage to having any particular part of the data cached in memory, and if the working set does not fit in memory, the cache will have a lot of flushes and misses. 

#### Special Types of Data

### Schema Design Gotchas in MySQL

Although there are universally bad and good design principles, there are also issues that arise from how MySQL is implemented, and that means you can make MySQL-specific mistakes, too. This section discusses problems that we've observed in schema designs with MySQL.

*Too many columns*

MySQL's storage engine API works by copying rows between the server and the storage engine in row buffer format; the server the decodes the buffer into columns. But it can be costly to turn the row buffer into the row data structure with the decoded columns. 

*Too many joins*

MySQL has a limitation of 61 tables per join.

*The all-powerful ENUM*

Beware of overusing ENUM. Here's an example we saw:

```sql
CREATE TABLE tbl (
    country enum('', '0', '1', '2', ...,'31')
);
```

The schema was sprinkled liberally with this pattern. This would probably be a questionable design decision.

*The ENUM in disguise*

*NULL not invented here*

### Normalization and Denormalized

#### Pros and Cons of a Normalization Schema

- Normalized updates are usually faster than denormalized updates.
- When the data is well normalized, there's little or no duplicated data, so there's less data to change.
- Normalized tables are usually smaller, so they fit better in memory and perform better.

#### Pros and Cons of a Denormalized Schema

A denormalized schema works well because everything is in the same table, which avoids joins.

#### A Mixture of Normalized and Denormalized

### Cache and Summary

#### Materialized Views

#### Counter tables

### Speeding Up ALTER TABLE

### Summary

Good schema design is pretty universal, but of course MySQL has special implementation details to consider. In a nulshell, it's a good idea to keep things as small and simple as you can.

## Chapter 5. Indexing for High Performance

### Indexing Basics

#### Types of Indexes

##### B-Tree indexes
