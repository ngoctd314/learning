# InnoDB and MyISAM Index Statistic Collection

Storage engines collect statistic about tables for use by the optimizer. Table statistics are based on value groups, where a value group is a set of rows with the same key prefix value.

MySQL uses the average value group size in the following ways:

- To estimate how many rows must be read for each ref access.
- To estimate how many rows a partial join produces; that is, the number of rows that an operation of this form produces:

```sql
(...) JOIN tbl_name ON tbl_name.key = expr;
```

As the average value group size for an index increases, the index is less useful for those two purposes because the average number of rows per lookup increases: For the index to ge good for optimization purposes, it is best that each index value target a small number of rows in the table. When a given index value yields a large number of rows, the index is less useful and MySQL is less likely to use it.

The average value group size is related to table cardinality, which is the number of value groups. The SHOW INDEX statement displays a cardinality value based on N/S, where N is the number of rows in the table and S is the average value group size. That ratio yields an approximate number of value groups in the table.

For a join based on the <=> comparison
