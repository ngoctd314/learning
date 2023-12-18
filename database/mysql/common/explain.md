# Explain

## Filtered field in Explain mysql

In the `EXPLAIN` statement in MySQL, the "filtered" column provides information about the percentage of rows that the query optimizer expects to be retrieved after applying the `WHERE` clause filter. It indicates how selective the `WHERE` clause is in terms of reducing the number of rows.

`filtered` column value for each table involved in the query. The "filtered" value is a percentage ranging from 0 to 100, where:

- A value of 100 indicates that the `WHERE` clause is highly selective, and it is expected to filter out most of the rows.
- A value of 0 indicates that the `WHERE` clause is not selective, and it is not expected to filter out any rows.

The `filtered` column provides insights into how well the `WHERE` clause can reduce the number of rows early in the execution process. A higher filtered percentage generally suggests that the query is more likely to be efficient, as it can quickly eliminate irrelevant rows based on the `WHERE` condition.

Keep in mind that the "filtered" column is an estimate made by the query optimizer, and the actual number of filtered rows during execution might differ. It's helpful indicator, but the final performance depends on various factors, including the distribution of data, indexes, and the overall query structure.

## Extra

### Using index condition

Extra: Using index condition indicates that the query is utilizing an index to evaluate part of the `WHERE` clause condition. This typically happens when the `WHERE` clause involves a mix of indexed and non-indexed columns, and the optimizer decides to use the index for the indexed part of the condition.

