# LIMIT Query Optimization

If you need only a specified number of rows from a result set, use a LIMIT clause in the query, rather than fetching the whole result set and throwing away the extra data.

MySQL sometimes optimizes a query that has a LIMIT row_count clause and no HAVING clause:

- If you select only a few rows with LIMIT, MySQL uses indexes in some cases when normally it would prefer to do a full table scan.
- If you combine LIMIT row_count with ORDER BY, MySQL stops sorting as soon as it has found the first row_count rows of the sorted result, rather than sorting the entire result. If ordering is done by using an index, this is very fast. If a filesort must be done, all rows that match the query without the LIMIT clause are selected, and most or all of them are sorted, before the first row_count are found.
- If you combine LIMIT row_count with DISTINCT, MySQL stops as soon as it finds row_count unique rows.
- In some cases, a GROUP BY can be resolved by reading the index in order (or doing a sort on the index), then calculating summaries until the index value changes. In this case, LIMIT row_count does not calculate any unnecessary GROUP BY values.
- If the server uses temporary tables to resolve a query, it uses the LIMIT **row_count** clause to caculate how much space is required.
- If an index is not used for ORDER BY but a LIMIT clause is also present, the optimizer may be able to avoid using a merge file and sort the rows in memory using an in-memory filesort operation. 

If multiple rows have identical values in the **ORDER BY** columns, the server is free to return those rows in any order, and may do so differently depending on the overall execution plan. In other words, the sort order of those rows is nondeterministic with respect to the nonordered columns In other words, the sort order of those rows is nondeterministic with respect to the nonordered columns.

One factor that affects execution plans is LIMIT, so an ORDER BY query with and without LIMIT may return rows in different orders. Consider this query, which is sorted by the category column but nondeterministic with respect to the id and rating columns:

```sql
SELECT * FROM ratings ORDER BY category;
+----+----------+--------+
| id | category | rating |
+----+----------+--------+
|  1 |        1 |    4.5 |
|  5 |        1 |    3.2 |
|  3 |        2 |    3.7 |
|  4 |        2 |    3.5 |
|  6 |        2 |    3.5 |
|  2 |        3 |    5.0 |
|  7 |        3 |    2.7 |
+----+----------+--------+
```

Including LIMIT may affect order of rows within each category value.

```sql
SELECT * FROM ratings ORDER BY category LIMIT 5;
+----+----------+--------+
| id | category | rating |
+----+----------+--------+
|  1 |        1 |    4.5 |
|  5 |        1 |    3.2 |
|  3 |        2 |    3.7 |
|  4 |        2 |    3.5 |
|  6 |        2 |    3.5 |
+----+----------+--------+
```

In each case, the rows are sorted by the ORDER BY column, which is all that is required by the SQL standard.

If it is important to ensure the same row order with and without LIMIT, include additional columns in the ORDER BY clause to make the order deterministic. For example, if id values are unique, you can make rows for a given category value appear in id order by sorting like this:
