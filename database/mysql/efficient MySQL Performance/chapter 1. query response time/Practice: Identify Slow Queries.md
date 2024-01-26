# Practice: Identify Slow Queries

The goal of this practice is to identify slow queries using pt-query-digest: a command-line tool that generates a query profile and query reports from a slow query log.

```sql
SET GLOBAL long_query_time=0;
SET GLOBAL slow_query_log=ON;
SELECT @@GLOBAL.slow_query_log_file;
```

Zero in the first statement, SET GLOBAL long_query_time=0; causes MySQL to log every query. Be careful: on a busy server, this can increase disk I/O and use gigabytes of disk space. If needed, use a slightly larger value like 0.0001 (100 microseconds) or 0.001 (1 millisecond).
