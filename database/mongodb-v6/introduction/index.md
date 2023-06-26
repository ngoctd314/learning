# Introduction to MongoDB

The advantages of using documents are:

- Documents correspond to native data types in many programming language
- Embedded documents and arrays reduce need for expendsive joins
- Dynamic schema supports fluent polymorphism

## Key Features

**High Performance**

MongoDB provides high performance data persistence. In particular,

- Support for embedded data models reduces I/O activity on database system
- Indexes support faster queries and can include keys from embedded documents and arrays.

**High Availability**

MongoDB's replication facility, called replica set, provides

- Automatic failover
- Data redundancy

A replica set is a group of MongoDB servers that maintain the same data set, providing redundancy and increasing data availability.

**Horizontal Scalability**

MongoDB provides horizontal scalability as part of its core functionality

- Sharding distributes data across a cluster of machines.
- Starting in 3.4, MongoDB supports creating zones of data based on the shard key. In a balanced cluster, MongoDB directs reads and writes covered by a zone only to those shards inside  the zone.

**Support for Multiple Storage Engines**

- WiredTiger Storage Engine
- In-Memory Storage Engine

