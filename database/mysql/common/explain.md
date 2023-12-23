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

### Using join buffer (Block Nested Loop)

When you see "Using join buffer (Block Nested Loop)" in the execution plan (`EXPLAIN` output) of a MySQL query, it indicates that the query execution involves a block nested loop join algorithm.

The Block Nested Loop Join is a type of nested loop join algorithm used by the MySQL query optimizer when performing a join operation between two tables. Here's a brief explanation:

**1. Nested Loop Join**

- In a nested loop join, MySQL iterates through each row of the outer table and, for each row, searches for matching rows in the inner table.
- The "nested" part comes from the fact that there is an inner loop for each row of the outer table.

**2. Block Nested Loop Join**

- In a Block Nested Loop Join, rather than processing a single row at a time, MySQL processes a block of rows (a set of rows) from the outer table in the each iteration.
- This can be more efficient than processing rows one by one, especially when dealing with large dataset.

The use of a join buffer in the Block Nested Loop Join is related to how the inner table (the smaller table in terms of the number of rows) is accessed and read. The join buffer helps to manage the rows from the inner table efficiently during the join operation.

While the Block Nested Loop Join can be effective in certain scenarios, it's essential to consider factors such as the size of the tables, available indexes, and the overall query complexity. Other join algorithms, like Hash Join or Merge Join, might be more suitable in different situations.

A block-nested loop (BNL) is an algorithm used to join two relations in a relational database.

This algorithm is variation of the simple nested loop joins and two relations R and S (the "outer" and "inner" join operations, respectively). Suppose |R| < |S|. In a traditional nested loop join, S will be scanned once for every tuple of R. If there are many qualifying R tuples, and particularly if there is no applicable index for the join key on S, this operation will be very expensive.

The block nested loop join algorithm improves on the simple nested loop join by only scanning S once for every group of R tuples. Here groups are disjoint set of tuples in R and the union of all gropus has the same tuples as R. For example, one variant of the block nested loop join reads an entire page of R tuples into memory and loads them into a hash table. It then scans S, and probes the hash table to find S tuples that match any of the tuples in the current page of R. This reduces the number of scans of S that are necessary.

```txt
algorithm block_nested_loop_join is
    for each page pr in R do
        for each page ps in S do
            for each tuple r in pr do
                for each tuple s in ps do
                    if r and s satisfy the join condition then
                        yield tuple <r,s>
```

## EXPLAIN EXTENDED mysql

The `EXPLAIN EXTENDED` statement in MySQL is used to obtain extra information about how MySQL executes a query. It provides additional details beyond the standard `EXPLAIN` output, including the access plan, how the optimizer resolves aliases, and more.

Here's an example of how you can use `EXPLAIN EXTENDED`

```sql
EXPLAIN EXTENDED
SELECT * FROM your_table WHERE your_condition;
```

After executing this query, you can then use the following command to view the extended information:

```sql
SHOW WARNINGS;
```

The extended information will be displayed as part of the warning messages.


