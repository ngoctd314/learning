# MySQL Indexes: A Visual Introduction

## InnoDB Tables Are Indexes

InnoDB tables are B-tree indexes organized by the primary key. Rows are index records stored in leaf nodes of the index structure. Each index record has metadata (denoted by "...") used for row locking, transaction isolation, and so on.

Secondary indexes are B-tree indexes, too but leaf nodes store primary key values. When MySQL uses secondary index to find a row, it does a second lookup on the primary key to read the full row.

## Table Access Methods

Using an index to look up rows is one of three table access methods. Since tables are indexes, an index lookup is the best and most common access method. But sometimes, depending on the query, and index lookup is not possible and the only resource is an index scan or table scan.

There are three access method: index lookup, index scan, and table scan. For an index lookup, there are several access types: ref, eq_ref, range and so forth.

**Index lookup**

An index lookup finds specific rows - or ranges of rows - by leveraging the ordered structure and algorithmic access of an index. This is the fastest access method because it's precisely what indexes are designed for: fast and efficient access to large amounts of data.

**Index scan**

When an index lookup is not possible, MySQL must use brute force to find rows: read all rows and filter out non-matching ones. Before MySQL resorts to reading every row using the primary key, it tries to read rows using a secondary index. This is called an index scan.

There are two types of index scan. The first is a full index scan, meaning MySQL reads all rows in index order. Reading all rows is usually terrible for performance, but reading them in index order can avoid sorting rows when the index order matches the query ORDER BY.

Scanning the secondary index in order might be sequential reads, but primary key lookups are almost certainly random reads. Accessing rows in index order does not guarantee sequential reads; more than likely, it incurs random reads.

The second type of index scan is an index-only scan: MySQL reads column values (not full rows) from the index. This requires a covering index, which is covered later. It should be faster than a full index scan because it doesn't require primary key lookups to read full rows; it only reads column values from the secondary index, which is why it requires a covering index. 

Don't optimize for an index scan unless the only alternative is a full table scan. Otherwise, avoid index scans.

**Table scan**

A full table scan reads all rows in primary key order. When MySQL cannot do an index lookup or an index scan, a table scan is the only option. This is usually terrible for performance, but it's also easy to fix because MySQL is adept at using indexes and has many index-based optimizations.

### Leftmost Prefix Requirement

In MySQL lingo we way, "The primary key is appended to secondary indexes" even though it's not literally appended. (You can literally append it by creating index (a, b, id), but don't do that.) "Appended to" really means that secondary index leaf nodes contain primary key values. This is imporant because it increases the size of every secondary index: primary key values are duplicated in secondary indexes. Larger indexes require more memory, which means fewer indexes can fit in memory.

### EXPLAIN: Query Execution Plan

The MySQL EXPLAIN command shows a query execution plan (or, EXPLAIN plan) that describes how MySQL plans to execute the query: table join order, table access method, index usage, and other important details.

### WHERE

### GROUP BY

MySQL can use an index to optimize GROUP BY because values are implicitly grouped by index order.

|a|b|
|-|-|
|Ag|B|
|Ag|B|
|Al|B|
|Al|B|
|Al|Br|
|Ar|B|
|Ar|Br|
|Ar|Br|
|At|Bi|
|Au|Be|

### ORDER BY

### Covering Indexes

### Join Tables

MySQL uses an index to join tables, and this usage is fundamentally the same as using an index for anything else. The main difference is the source of values used in join conditions for each table. This becomes more clear when visualized, but first we need a second table to join. 

MySQL can join a table using any access method, but an index lookup using the eq_ref access type is the best and fastest because it matches only one row. The eq_ref access type has two requirements: a primary key or unique not-null secondary index and equality condition on all index columns.

Table join order is critical because MySQL joins tables in the best order possible, not the order tables are written in the query. You must use EXPLAIN to see the table join order. EXPLAIN prints tables in the order from top (first table) to bottom (last table).

Never guess or presume the table join order because small changes to a query can yield a significantly different table join order or query execution plan. 

## Indexing: How to Think Like MySQL

Indexes and indexing are different topics. The previous section introduced indexes: standard B-tree indexes on InnoDB tables for WHERE, GROUP BY, ORDER BY, covering indexes, and table joins.

### Know the Query

```sql
SHOW CREATE TABLE;
SHOW TABLE STATUS;
SHOW INDEXES;
```

Answer the following questions:

**Query**

- How many rows should the query access? (Rows examined)
- How many rows should the query return? (Rows sent)
- Which columns are selected (returned)? (SELECT ?)
- What are the GROUP BY, ORDER BY, and LIMIT clauses (if any)?
- Are there subqueries? (If yes, repeat the process for each.)

**Table access (per-table)**

- What are the table condition?
- Which index should the query use?
- What other indexes could the query use?
- What is the cardinality of each index?
- How large is the table - data size and row count?

Those questions help you mentally parse the query because that's what MySQL does: parse the query. This is especially helpful for seeing helpful for seeing complex queries in simpler terms: tables, table conditions, indexes, and SQL clauses.

This information helps you piece together a puzzle that, one complete, reveals query response time. To improve response time, you'll need to change some pieces. But before doing that, the next step is to assemble the current pieces with the help of EXPLAIN.

### Understand with EXPLAIN

The second step is to understand the current query execution plan reported by EXPLAIN. Consider each table and its conditions with respect to it indexes, starting with the index that MySQL chose: the key field in the EXPLAIN output. If the possible_keys field lists other indexes, think about how MySQL would access rows using those indexes - always with the leftmost prefix requirement in mind.

Always EXPLAIN the query. Make this a habit because direct query optimization is not possible without EXPLAIN.

The query and its response time are puzzle, but you have all the pieces: execution plan, table conditions, table structures, table sizes, index cardinalities, and query metrics. Keep connecting the pieces until the puzzle is complete - until you can see the query.

### Optimize the Query

The third step is direct query optimization: change the query, its indexes, or both. This is where all the fun happens, and there's no risk yet because these changes are made in development or staging, not production. Be certain that your development or staging env has data that is representative of production because data size and distribution affect how MySQL chooses indexes.

At first, it might seem like the query cannot be modified because it fetches the correct rows, so the query is written correctly. A query "is what it is," right? Not always; the same result can be achieved with different methods. A query as a result - literally, a result set - and a method of obtaining that result. These two are closely related but independent. Knowing that is tremendously helpful when considering how to modify a query. Start by clarifying the intended result of the query. A clear result allows you to explore new ways of writing the query that achieve the same result.

There can be multiple ways to write a query that execute differently but return the same result.

Adding or modifying a index is a trade-off between access methods and query specific optimizations. For example, do you trade an ORDER BY optimization for a range scan? 

### Deploy and Verify

The last time is to deploy the changes and verify that they improve response time. But first: know how to roll back the development - and be ready to do so - in case the changes have unintended side effects. 

Always know how to  and be ready to - roll back a development to production.

## It Was a Good Index Until

If nothing changes, a good index will stay a good index until the end of time.

### Queries Changed

More indexes use more RAM which, ironically, decreases the RAM available for each index. The second problem is a decrease in write performance because, when MySQL writes data, it must check, update, and potentially reorganize. An inordinate number of indexes can severely degrade write performance.

To find existing duplicate indexes, use pt-duplicate-key-checker: it safely finds and reports duplicate indexes.

### Excessive, Duplicate, and Unused

Cardinality is the number of unique values in an index.

Use SHOW INDEX to see index cardinality.

Selectivity is cardinality divided by the number of rows in the table. Selectivity ranges from 0 to 1, where 1 is a unique index: a value for every row. MySQL doesn't show index selectivity; you have to calculate it manually using SHOW INDEX for cardinality and SHOW TABLE STATUS for the number for rows.

### Extreme Selectivity

An index with extremely low selectivity provides little leverage because each unique value could match a large number of rows. A classic example is an index on a column with only two possible values: yes or no, true or false, coffee or tea, on so on. If the table has 100000 rows, then selectivity is practically zero: 2 / 100000 = 0.00002. It's an index, but not a good one because each value could match many rows. How many? Flip the division: 100000 rows / 2 unique values = 50000 rows per value. If MySQL were to use this index (which is unlikely), a single index lookup could match 50000 rows. That presumes values are evenly distributed, but what if 99,999 rows have value coffee and only 1 row has value tea? Then the index works great for tea but terribly for coffee.

If a query uses an index with extremely low selectivity, see if you can create a better, more selective index; or consider rewriting the query to use a more selective index; or, think about altering the schema to organize the data better with respect to access patterns.

### It's a Trap! (When MySQL Chooses Another Index)

In very rare cases, MySQL chooses the wrong index. The index itself is never inaccurate; it's only the index statistics that are inaccurate.

Index statistics are estimates about how values are distributed in the index. MySQL does random dives into the index to sample pages. (A page is a 16KB unit of logical storage. Almost everything is stored in pages.) If index values are evenly distributed, then a few random dives accurately represent the whole index.

Running ANALYZE TABLE is safe and usually very fast, but be careful on a busy server: it requires a flush lock that can block all queries accessing the table.

## Table Join Algorithms

```sql
SELECT * FROM t1 
    JOIN t2 ON t1.A = t2.B 
    JOIN t3 ON t2.B = t3.C
```

Suppose that EXPLAIN reports the join order as t1, t2, and t3. The nested-loop join algorithm works like the pseudocode:

```go
func find_rows(table, index, condition string) []rows {
    // Return array of rows in table matching conditions
    // using index for lookup or table scan if NULL
}

for find_rows(t1, some_index, "WHERE ...") {
    for find_rows(t2, index_on_B, "WHERE B= <t1.A>") {
        return find_rows(t3, NULL, "WHERE C = <t2.B>")
    }
}
```

Using NLJ algorithm, MySQL begins by using some_index to find matching rows in the outermost table: t1. For each matching row in table t1, MySQL joins table t2 by using an index on the join column, index_on_B, to lookup rows matching t1.A. For each matching row in the table t2, MySQL join table t3 using the same process.

When no more rows in t3 match the join column value from table t2, the next matching row from t2 is used. When no more rows in t2 match the join column value from table t1, the next matching row from t1 is used. When no more rows in t1 match, the query completes.

The nested-loop join algorithm is simple and effective, but there's one problem, the innermost table is accessed very frequently, and the full join makes that access very slow. In this example, table t3 is accessed for every matching row in t1 multipled by every matching row in t2. If both t1 and t2 has 10 matching rows, then t3 is accessed 100 times. The block nested-loop join algorithm addresses this problem. Join column values from matching rows in t1 and t2 are saved in a join buffer. (The join buffer size is set by system variable join_buffer_size) When the join buffer is full, MySQL scans t3 and joins each t3 row that matches join column values in the join buffer.

As of MySQL 8.0.20, the hash join algorithm replaces the block nested-loop join algorithm. Hash join creates an in-memory hash table of join tables, like table t3 in this example. MySQL use the hash table to look up rows in the join table, which is extremely fast because a hash table lookup rows in the join table, which is extremely fast because a hash table lookup is a constant time operation.

EXPLAIN indicates a hash join by printing "Using join buffer (hash join)" in the Extra field.

## Summary

- Indexes provide the most and the best leverage for MySQL performance.
- Do not scale up hardware to improve performance until exhausting other options.
- Tuning MySQL is not necessary to improve performance with a reasonable configuration.
- An InnoDB table is a B-tree index organized by the primary key.
- MySQL accesses a table by index lookup, index scan, or full table scan - index lookup is the best access method.
