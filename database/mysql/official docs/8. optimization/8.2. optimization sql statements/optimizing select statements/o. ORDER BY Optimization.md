# ORDER BY Optimization

This section describes when MySQL can use an index to satisfy an ORDER BY clause, the filesort operation used when an index cannot be used, and execution plan information available from the optimizer about ORDER BY.

## Use of Indexes to Satisfy ORDER BY

In some cases, MySQL may use an index to satisfy an ORDER BY clause and avoid the extra sorting involved in performing a filesort operation.

The index may also be used even if the ORDER BY does not match the index exactly, as long as all unused portions of the index and all extra ORDER BY columns are constants in the WHERE clause. If the index does not contain all columns accessed by the query, the index is used only if index access is cheaper than other access methods.

Assuming that there is an index on(key_part1, key_part2), the following queries may use the index to resolve the ORDER BY part. Whether the optimizer actually does so depends on whether reading the index is more efficient than a table scan if columns not in the index must also be read.

- In this query, the index on (key_part1, key_part2) enables the optimizer to avoid sorting:

```sql
SELECT * FROM t1
    ORDER BY key_part1, key_part2
```

However, the query uses `SELECT` *, which may select more columns than key_part1 and key_part2. In that case, scanning an entire index and looking up table rows to find columns in the index may be more expensive than scanning the table and sorting the results. If so, the optimizer is not likely to use the index. If SELECT * selects only the index columns, the index is used and sorting avoided.

If t1 is an InnoDB table, the table primary key is implicitly part of the index, and the index can be used to resolve the ORDER BY for this query:

```sql
SELECT pk, key_part1, key_part2 FROM t1
    ORDER BY key_part1, key_part2;
```

- In this query, key_part1 is constant, so all rows accessed through the index are in key_part2 order, and an index on (key_part1, key_part2) avoids sorting if the **WHERE** clause is selective enough to make an index range scan cheaper than a table scan:

```sql
SELECT * FROM t1 
    WHERE key_part1 = constant    
    ORDER BY key_part2;
```

- In the next two queries, whether the index is used to similar to the same queries without DESC shown previously:

```sql
SELECT * FROM t1
    ORDER BY key_part1 DESC, key_part2 DESC;

SELECT * FROM t1 
    WHERE key_part1 = constant
    ORDER BY key_part2 DESC;
```

- In the next two queries, key_part1 is compared to a constant. The index is used if the **WHERE** clause is selective enough to make an index range scan cheaper than a table scan:

```sql
SELECT * FROM t1 
    WHERE key_part1 > constant
    ORDER BY key_part1 ASC;

SELECT * FROM t1
    WHERE key_part1 < constant
    ORDER BY key_part1 DESC;
```

- In the next query, the ORDER BY does not name key_part1, but all rows selected have a constant key_part1 value, so the index can still be used:

```sql
SELECT * FROM t1
    WHERE key_part1 = constant1 AND key_part2 > constant2
    ORDER BY key_part2
```

In some cases, MySQL cannot use indexes to resolve the ORDER BY, although it may still use indexes to find the rows that match the WHERE clause.

- The index used to fetch the rows differs from the one used in the ORDER BY:

```sql
SELECT * FROM t1 WHERE key2=constant ORDER BY key1;
```

- The query joins many tables, and the columns in the ORDER BY are not all from the first constant table that is used to retrieve rows.

- The query has different ORDER BY and GROUP BY expressions.

- There is an index on only a prefix of a column named in the ORDER BY clause. In this case, the index cannot be used to fully resolve the sort order. For example, if only the first 10 bytes of CHAR(20) column are indexed, the index cannot distinguish values past the 10th byte and a filesort is needed.

- The index does not store rows in order. For example, this is true for a HASH index in a MEMORY table.

Availability of an index for sorting may be affected by the use of column aliases. Suppose that the column t1.a is indexed. In this statement, the name of the column in the select list is a. It refers to t1.a, as does the reference to a in the ORDER BY, so the index on t1.a can be used:

```sql
SELECT a FROM t1 ORDER BY a;
```

In this statement, the name of the column in the select list is also a, but it is the alias name. It refers to ABS(a), as does the reference to a in the ORDER BY, so the index on t1.a cannot be used:

```sql
SELECT ABS(a) AS a FROM t1 ORDER BY a;
```

By default, MySQL sorts GROUP BY col1, col2, ... queries as if you also included ORDER BY col1, col2, ... in the query. If you include an explicit ORDER BY clause that contains the same column list, MySQL optimizers it away without any speed penalty, although the sorting still occurs.

If a query includes GROUP BY but you want to avoid the overhead of sorting the result, you can suppress sorting by specifying ORDER BY NULL.

```sql
SELECT a, COUNT(*) FROM bar GROUP BY a ORDER BY NULL;
```

```sql
CREATE TABLE tbl (id int AUTO_INCREMENT PRIMARY KEY, a int, b int, c int, key idx(a, b));
INSERT INTO tbl (a, b, c) VALUES (1, 1 , 1), (2, 2, 2), (3, 3, 3);

EXPLAIN SELECT * FROM tbl ORDER BY a, b;
```

However, the query uses SELECT *, which may select more columns than key_part1 and key_part2. In that case, scanning an entire index and looking up table rows to find columns not in the index may be more expensive than scanning the table and sorting the results. If so, the optimizer is not likely to use the index. If SELECT * selects only the index columns, the index is used and sorting avoided.

In InnoDB table, the table primary key is implicitly part of the index, and the index can be used to resolve the ORDER BY for this query:

```sql
EXPLAIN SELECT id, a, b FROM tbl ORDER BY a, b;
```

In this query, key_part1 is constant, so all rows accessed through the index are in key_part2 order, and an index on (key_part1, key_part2) avoids sorting if the `WHERE` clause is selective enough to make an index range scan cheaper than a table scan:

```sql
SELECT * FROM t1
    WHERE key_part1 = constant
    ORDER BY key_part2;
```

In the next two queries, whether the index is used is similar to the same queries without DESC shown previously:

```sql
SELECT * FROM t1
    ORDER BY key_part1 DESC, key_part2 DESC;
```

In the next two queries, key_part1 is compared to a constant. The index is used if the `WHERE` clause is selective enough to make an index range scan cheaper than a table scan:

```sql
SELECT * FROM t1
    WHERE key_part1 > constant
    ORDER BY key_part1 ASC;

SELECT * FROM t1
    WHERE key_part1 < constant
    ORDER BY key_part1 DESC;
```

In the next query, the ORDER BY does not name key_part1, but all rows selected have a constant key_part1 value, so the index can still be used:

```sql
SELECT * FROM t1
    WHERE key_part1 = constant1 AND key_part2 > constant2
    ORDER BY key_part2;
```

In some cases, MySQL cannot use indexes to resolve the ORDER BY, although it may still use indexes to find the rows that match the WHERE clause. Examples:

- The query uses ORDER BY on different indexes:

```sql
SELECT * FROM t1 ORDER BY key1, key2
```

- The query uses ORDER BY on nonconsecutive parts of an index:

```sql
SELECT * FROM t1 WHERE key2=constant ORDER BY key_part1, key_part3
```

- The query mixes ASC and DESC

```sql
SELECT * FROM t1 ORDER BY key_part1 DESC, key_part2 ASC;
```

- The index used to fetch the rows differs from the one used in the ORDER BY:

```sql
SELECT * FROM t1 WHERE key2=constant ORDER BY key1;
```

- The query uses ORDER BY with an expression that includes terms other than the index column name:

```sql
SELECT * FROM t1 ORDER BY ABS(key);
SELECT * FROM t1 ORDER BY -key;
```

- The query joins many tables, and the columns in the ORDER BY are not all from the first nonconstant table that is used to retrieve rows. (This is the first table in the EXPLAIN output that does not have a const join type).

- The query has different ORDER BY and GROUP BY expressions.

- There is an index on only a prefix of a column named in the ORDER BY clause. In this case, the index cannot be used to fully resolve the sort order. For example, if only the first 10 bytes of a CHAR(20) column are indexed, the index cannot distinguish values past the 10th byte and a filesort is needed.

- The index does not store rows in order. For example, this is true for a HASH index in a MEMORY table. 

By default, MySQL sorts GROUP BY col1, col2, ... queries as if you also included ORDER BY col1, col2, ... in the query. If you include an explicit ORDER BY clause that contains the same column list, MySQL optimizes it away without any speed penalty, although the sorting still occurs.

If a query includes GROUP BY but you want to avoid the overhead of sorting the result, you can suppress sorting specifying ORDER BY NULL. For example:

```sql
INSERT INTO foo;
SELECT a, COUNT(*) FROM bar GROUP BY a ORDER BY NULL;
```

The optimizer may still choose to use sorting to implement grouping operations. ORDER BY NULL suppresses sorting of the result, not prior sorting done by grouping operations to determine the result.

## Use of filesort to Satisfy ORDER BY

If an index cannot be used to satisfy an ORDER BY clause, MySQL performs a filesort operation that reads table rows and sorts them. A filesort constitutes an extra sorting phase in query execution.

To obtain memory for filesort operations, the optimizer allocates a fixed amount of sort_buffer_size bytes up front. Individual sessions can change the session value of this variable as desired to avoid excessive memory use, or to allocate more memory as necessary.

A filesort operation uses temporary disk files as necessary if the result set is too large to fit in memory. Some types of queries are particularly suited to completely in-memory filesort operations. For example, the optimizer can use filesort to efficiently handle in memory, without temporary files, the ORDER BY operation for queries (and subqueries) of the following form:

```sql
SELECT ... FROM single_table ... ORDER BY non_index_column [DESC] LIMIT [M, ]N;
```

Such queries are common in web applications that display only a few rows from a larger result set.

```sql
SELECT col1, ... FROM t1 ... ORDER BY name LIMIT 10;
SELECT col1, ... FROM t1 ... ORDER BY RAND() LIMIT 15;
```

## Influencing ORDER BY Optimization

## ORDER BY Execution Plan Information Available

With EXPLAIN, you can check whether MySQL can use indexs to resolve an ORDER BY clause.

- If the Extra column of EXPLAIN output does not contain Using filesort, the index is used and a filesort is not performed.
- If the Extra column of EXPLAIN output contains Using filesort, the index is not used and a filesort is performed.

In addition, if a filesort is performed, optimizer trace output includes a filesort_summary block.

```json
"filesort_summary": {
    "rows": 100,
    "examined_rows": 100,
    "number_of_tmp_files": 0,
    "sort_buffer_size": 25192,
    "sort_mode": "<sort_key, packaged_additional_fields>"
}
```
