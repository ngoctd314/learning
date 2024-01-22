# Avoiding Full Table Scans

The output from EXPLAIN shows ALL in the type column when MySQL uses a full table scan to resolve a query. This usually happens under the following conditions:

- The table is so small that it is faster to perform a table scan than to bother with a key lookup.
- There are no usable restrictions in the ON or WHERE clause for indexed columns.
- You are comparing indexed columns with constant values and MySQL has calculated (based on the index tree) that the constants cover too large a part of the table and that a table scan would be faster.
- You are using a key with low cardinality (many rows match the key value) through another column. In this case, MySQL assumes that by using the key it is likely to perform many key lookups and that a table scan would be faster.

For small tables, a table scan often is appropriate and the performance impact is negligible. For large tables, try the following techniques to avoid having the optimizer incorrectly choose a table scan:

- Use ANALYZE TABLE tbl_name to update the key distributions for the scanned table.
- Use FORCE TABLE for the scanned table to tell MySQL that table scans are very expensive compared to using the given index:

```sql
SELECT * FROM t1, t2 FORCE INDEX(index_for_column)
    WHERE t1.col_name=t2.col_name;
```
