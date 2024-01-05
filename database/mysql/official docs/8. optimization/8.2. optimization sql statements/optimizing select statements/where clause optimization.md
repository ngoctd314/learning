# WHERE Clause Optimization

This section discusses optimizations that can be made for processing WHERE clauses. The examples use SELECT statements, but the same optimizations apply for WHERE clauses in DELETE and UPDATE statements.

- Removal of unnecessary parentheses:

```sql
((a AND b) AND c OR (((a AND b) AND (c AND d))))
(a AND b AND c) OR (a AND b AND c AND d)
```

- Constant folding

```sql
(a < b AND b=c) AND a = 5
b > 5 AND b=c AND a=5
```

- Constant condition removal 

```sql
(b >= 5 AND b = 5) OR (b=6 AND 5=5) OR (b=7 AND 5=6)
b=5 OR b=6
```

- COUNT(*) on a single table without a `WHERE` is retrieved directly from the table information for MyISAM and MEMORY tables. This is also done for any NOT NULL expression when used with only one table.

- HAVING is merged with WHERE if you do not use GROUP BY or aggregate functions (COUNT(), MIN(), and so on).

- For each table in a join, a simpler WHERE is constructed to get a fast WHERE evaluation for the table and also to skip rows as soon as possible

- The best join combination for joining the tables is found by trying all possibilities. If all columns in ORDER BY and GROUP BY clauses come from the same table, that table is preferred first when joining.

TODO: claim
- If there is an ORDER BY clause and a different GROUP BY clause, or if the ORDER BY or GROUP BY contains columns from tables other than the first table in the join queue, a temporary table is created.

- If you use the SQL_SMALL_RESULT modifier, MySQL uses an in-memory temporary table.

- Each table index is queried, and the best is used unless the optimizer belives that it is more efficient to use a table scan. At one time, a scan was used based on whether the best index spanned more than 30% or the table, but a fixed percentage no longer determines the choice between using an index or a scan.

