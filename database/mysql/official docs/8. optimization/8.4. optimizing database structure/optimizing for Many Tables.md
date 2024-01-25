# Optimizing for Many Tables

Some techniques for keeping individual queries fast involve splitting data aross many tables. When then number of tables runs into the thousands or even millions, the overhead of dealing with all these tables becomes a new performance.

## How MySQL Opens and Closes Tables

When you execute a mysqladmin status command, you should see something like this:

```txt
Uptime: 426 Running threads: 1 Questions: 11082
Reloads: 1 Open tables: 12
```

The Open tables value of 12 can be somewhat puzzling if you have fewer than 12 tables.

MySQL is multithreaded, so there may be many clients issuing queries for a given table simultaneously. To minimize the problem with multiple client sessions having different states on the same table, the table is opened independently by each concurrent session. This uses additional memory but normally increases performance. With MyISAM tables, one extra file descriptor is required for the data file for each client that has the table open. (By constrast, the index file descriptor is shared between all sessions.)

The table_open_cache and max_connections system variables affect the maximum number of files the server keeps open. If you increase one or both of these values, you may run up against a limit imposed by your operating system on the per-process number of open file descriptors. Many operating systems permit you to increases the open-files limit, although the method varies widely from system to system. Consult your operating system documentation to determine whether it is possible to increase the limit and how to do so.

table_open_cache is related to max_connections. For example, for 200 concurrent running connections, specify a table cache size of at least 200*N, where N is the maximum number of tables per join in any of the queries which you execute. You must also reserve some extra file descriptors for temporary tables and files.

Make sure that your operating system can handle the number of open file descriptors implied by the table_open_cache setting. If the table_open_cache is set too high, MySQL may run out of file descriptors and exhibit symptoms such as refusing connections or failing to perform queries.

## Disadvantages of Creating Many Tables in the Same Database

If you have many MyISAM tables in the same database directory, open, close, and create operations are slow. If you execute SELECT statements on many different tables, there is a little overhead when the table cache is full, because for every table that has to be opened, another must be closed. You can reduce this overhead by increasing the number of entries permitted in the table cache.
