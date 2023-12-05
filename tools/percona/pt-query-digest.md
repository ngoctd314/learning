# pt-query-digest

**pt-query-digest** - Analyze MySQL queries from logs, processlist, and tcpdump.

```bash
pt-query-digest [OPTIONS] [FILES] [DSN]
```

**pt-query-digest** analyzes MySQL queries from slow, general, and binary log files. It can also analyze queries from SHOW PROCESSLIST and MySQL protocol data from tcpdump. By default, queries are grouped by fingerprint and reported in descending order of query time. If no FILES are given, the tool reads STDIN. The optional DSN is used for certain options like --since and --until.

Report the slowest queries from slow.log:

```bash
pt-query-digest slow.log
```

Report the slowest queries from the processlist on host1:

```bash
pt-query-digest --processlist h=host1
```

Capture MySQL protocol data with tcpdump, then report the slowest queries:

```bash
tcpdump -s 65535 -x -nn -q -tttt -i any -c 1000 port 3306 > mysql.tcp.txt

pt-query-digest --type tcpdump mysql.tcp.txt
```

Save query data from slow.log to host2 for later review and trend analysis:

```bash
pt-query-digest --review h=host2 --no-report slow.log
```

## RISKS

Percona is mature, proven in the real world, and well tested, but all database tools can pose a risk to the system and the database server.

## ATTRIBUTES

pt-query-digest works on events, which are collection of key-value pairs attributes. You'll recognize most of the attributes right away: Query_time, Lock_time, and so on. You can just look at a slow log and see them. However, there are some that don't exist in the slow log, and show logs may actually include different kinds of attributes.

```txt
Attribute    pct   total     min     max     avg     95%  stddev  median
============ === ======= ======= ======= ======= ======= ======= =======
Count         10       2
Exec time      2   347us   128us   219us   173us   219us    64us   173us
Lock time      8     2us     1us     1us     1us     1us       0     1us
Rows sent      2      12       6       6       6       6       0       6
Rows examine   0      12       6       6       6       6       0       6
Query size     9     112      56      56      56      56       0      56
```

See "ATTRIBUTES REFERENCE" near the end of this documentation for a list of common and --type specific attributes. A familiarity with these attributes is necessary this working with --filter, --ignore-attributes, and other attribute-related options.

With creative use of --filter, you can create new attributes derived from existing attributes. A familiarity with these attributes is necessary for working with --filter, --ignore-attributes

## OUTPUT

The default --output is a query analysis report. The --[no]report option controls whether or not this report is printed. Sometimes you may want to parse all the queries but suppress the report, for example when using --review or --history.

There is one paragraph for each class of query analyzed. A "class" of queries all have the same value for the --group-by attribute which is fingerprint by default. A fingerprint is an abstracted version of the query text with literals removed, whitespace collapsed, and so forth.

The output described here is controlled by --report-format. That option allows you to specify what to print and in what order. The default output in the default order is described here.

The report, by default, begins with a paragraph about the entire analysis run The information is very similar to what you'll see for each class of queries in the log, but it doesn't have some information that would be too expensive to keep globally for the analysis.

```txt
80ms user time, 10ms system time, 37.46M rss, 46.97M vsz
Current date: Tue Dec  5 10:08:34 2023
Hostname: ngoctd10
Files: slow_query.log
Overall: 20 total, 10 unique, 0.26 QPS, 0.00x concurrency ______________
Time range: 2023-12-05T02:34:37 to 2023-12-05T02:35:55
Attribute          total     min     max     avg     95%  stddev  median
============     ======= ======= ======= ======= ======= ======= =======
Exec time           12ms    60us     3ms   616us     2ms   714us   366us
Lock time           23us       0     4us     1us     2us     1us     1us
Rows sent            600       0     381      30   49.17   78.57    0.99
Rows examine       2.04k       0   1.64k  104.25  151.03  356.66    1.96
Query size         1.11k      11     199   56.65  192.76   53.86   36.69
```

- user time: The amount of time that the CPU spends executing code in user space on behalf of a specific process. User time includes the time spent executing the actual application code, also known as user-mode code. This is the time spent running the specific tasks and operations requested by the application or user.

Indicates how much CPU time is spent executing the actual tasks requested by the application or user: It's directly related to the workload imposed by the application.

- system time: The amount of time that CPU spends executing code in kernel space on behalf of a specific process. System time includes the time spent executing kernel-mode code. The kernel is the core of the os, and sytem time accounts for the time spent in kernel-level operations, such as handling system calls, managing I/O operations, and other tasks that require elevated privileges.

Reflects the CPU time spent on system-level operations, often involving interactions with the os kernel. It includes activities like file I/O, memory management, and other kernel-level tasks.

In the output from pt-query-digest the "user-time" and "system time" values indicate the CPU time consumed by the process during the execution of the analyzed queries. These values can help in understanding the resource utilization pattern and identifying whether the workload is more focused on user-level tasks or involves significant system-level operations.

User Time: 80ms, System Time: 10ms.

In this context, it suggests that a significant portion of the CPU time is spent executing user-level code (application or query-related tasks), while a smaller portition is spent in system-level operations (kernel-related tasks). The total of these (user time + system time) contributes to the overall CPU time consumed by the process.

- Overall: k1 total, k2 unique, k3 QPS, k4 concurrency

k1 total: There were k1 queries in total during the specified time range.

k2 query: There were k2 unique queries in k1

k3 QPS: There were k3 query per second

k4 concurrency: concurrency level

|Attribute|Meaning|
|-|-|
|Exec time|Total execution time of queries|
|Lock time|Total time spent waiting for locks|
|Rows sent|Number of rows sent to the client|
|Rows examine|Number of rows examined by the server|
|Query size|Size of the queries in bytes|

- Stddev: Standard deviation

Standard deviation is measure of the amount of variation or despersion in a set of values. It quantifies how much individual values differ from the mean (average) of the set.

A low standard deviation indicates that the values tend to be close to the mean, while a high standard deviation suggests that the values are more spread out.

- Median: 

The median is the middle value of a dataset when it is ordered. In other words, it is the value below which half of a dataset when it is ordered. In other words, it is the value below which half of the data falls and above which half of the data falls.

The median is less sensitive to extreme values (outliers) compared to the mean. It provides a measure of central tendecy that is not influenced by extremely high or low values.

- Rank: The query's rank within the entire set of queries analyzed
- Query ID: The query's fingerprint
- Response time: The total response time, and percentage of overall total
- Calls: The number of times this query was executed
- R/Call: The mean response time per execution
- V/M: The variance-to-mean ratio of response time
- Item: The distilled query

**QUERY REVIEW**

A query --review is the process of storing all the query fingerprints analyzed.

- You can add md to classes of queries, such as marking them for follow-up, adding notes to queries, or marking them with an issue ID for your issue tracking system.
- You can refer to the stored values on subsequent runs so you'll know whether you've seen a qeury 

--limit: limit analysis

**FINGERPRINTS**

A query fingerprint is the abstracted form of a query, which makes it possible to group similar queries together. Abstracting a query removes literal values, normalizes whitespace, and so on.

```sql
SELECT name, password FROM user WHERE id = '123';
select name, password  from user where id = 5;
```

Both of those queries will fingerprint to

```sql
select name, password from user where id = ?;
```

Once the query's fingerprint is known, we can then talk about a query as though it represents all similar queries.

--order-by

--limit

type: Array; default: 95%:20

Limit output to the given percentage or count.

If the argument is an integer; report only the top N worst queries. If the argument is an integer followed by the %sign, report that percentage of the worst queries. If the percentage is followed by a colon and another integer, report the top percentage or the number specified by that integer, whichever comes first.

--log

type: string

Print all output to this file when daemonized.
