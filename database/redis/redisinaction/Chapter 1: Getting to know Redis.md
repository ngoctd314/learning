# Chapter 1: Getting to know Redis

Redis is an in-memory remote database that offers high performance, replication, and a unique data model to produce a platform for solving problems.

## 1.1 What is Redis?

When I say that Redis is a database, I'm only telling a partial truth. Redis is a very fast non-relational database that stores a mapping of keys to five different types of values. Redis supports in-memory persistent storage on disk, replication to scale read performance, and client-side sharing to scale write performance.

### 1.1.1 Redis compared to other databases and software

### 1.1.2 Other features

When using an in-memory database like Redis, one of the first questions that's asked is "What happens when my server gets turned off?" Redis has two different forms of persistence available for writing in-memory data to disk in a compact format. The first method is a point-in-time dump either when certain conditions are met (a number of writes in a given period) or when one of the two dump-to-disk commands is called. The other method uses an append-only file that writes every command that alters data in Redis to disk as it happens. Depending on how careful you want to be with your data, append-only writing can be configured to never sync, sync once per second, or sync at the completion of every operation.

Even though Redis is able to perform well, due to its in-memory design there are situations where you may need Redis to process more read queries than a single Redis server can handle. To support higher rates of read performance (along with handling failover if the server that Redis is running on crashes), Redis supports master/slave replication where slaves connect to the master and receive an initial copy of the full database. As writes are performed on the master, they're sent to all connected slaves for updating the slave datasets in real time. 

### 1.1.3 Why Redis?


