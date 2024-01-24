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

The table_open_cache and max_connections system variables affect the maximum number of files the server keeps open. If you increase one or both of these values, 
