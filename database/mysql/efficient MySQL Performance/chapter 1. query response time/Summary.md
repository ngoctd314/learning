# Summary

- Performance is query response time: how long it takes MySQL to execute a query.
- Query response time is the North Star of MySQL performance because it is meaningful and actionable.
- Query metrics originate from the slow query log or the Performance Schema.
- The Performance Schema is the best source of query metrics.
- Query metrics are grouped and aggregated by digest: normalized SQL statements.
- Improving query response time (query optimization) is a two-part journey: direct query optimization, then indirect query optimization.

Direct query optimization is changes to queries and indexes.

Indirect query optimization is changes to data and access patterns.

- At the very least, review the query profile and optimize slow queries when performance affects customers, before and after code changes, and once a month.

- To make MySQL go faster, you must decrease response time (free time to do more work) or increase load (push MySQL to work harder).
