# pt-query-digest

**pt-query-digest** - Analyze MySQL queries from logs, processlist, and tcpdump.

```bash
pt-query-digest [OPTIONS] [FILES] [DSN]
```

**pt-query-digest** analyzes MySQL queries from slow, general, and binary log files. It can also analyze queries from SHOW PROCESSLIST and MySQL protocol data from tcpdump. By default, queries are grouped by fingerprint and reported in descending order of query time. If no FILES are given, the tool reads STDIN.
