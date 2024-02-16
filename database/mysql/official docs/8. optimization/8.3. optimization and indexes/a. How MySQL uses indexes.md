# How MySQL Uses Indexes

Indexes are used to find rows with specific column values quickly. Without an index, MySQL must begin with the first row and then read through the entire table to find the relevant rows. If the table has an index for the columns in question, MySQL can quickly determine the position to seek to in the middle of the data file without having to look at all the data. This is much faster than reading every row sequentially.

Most MySQL indexes (PRIMARY KEY, UNIQUE INDEX, and FULLTEXT) are stores in B-trees. Memory tables also support hash indexes; InnoDB use inverted lists for FULL TEXT indexes. 

MySQL uses indexes for these operations:

- To find the rows matching a WHERE clause quickly.
- To eliminate rows from consideration. If there is a choice between multiple indexes, MySQL normally uses the index that finds the smallest number of rows.
- If the table has a multiple-column index, any leftmost prefix of the index can be used by the optimizer to look up rows. For example, if you have a three-column index on (col1, col2, col3), you have indexed search capabilities on (col1), (col1, col2) and (col1, col2, col3)
- To retrieve rows from other tables when performing joins.
- To find the MIN() or MAX() value for a specific indexed column key_col. This is optimized by a preprocessor that checks whether you are using WHERE key_part_N = constant on all key parts that occur before in the index. In this case, MySQL does a single key lookup for each MIN() or MAX() expression and replaces it with a constant. If all expressions are replaced with constants, the query returns at once.

```sql
SELECT MIN(key_part2), MAX(key_part2)
    FROM tbl_name WHERE key_part1=10;
```

Indexes are less important for queries on small tables, or big tables where report queries process most or all of the rows, reading sequentially is faster than working through an index. Sequential reads minimize disk seeks, even if not all the rows are needed for the query.

## 8.3.2. Primary Key Optimization

The primary key for a table represents the column or set of columns that you use in your most vital queries. It has an associated index, for fast query performance. Query performance benefits from the NOT NULL optimization, because it cannot include any NULL values. With the InnoDB storage engine, the table data is physically organized to do ultra-fast lookups and sorts based on the primary key or columns.  

## 8.3.3. SPATIAL Index Optimization

## 8.3.4. Foreign Key Optimization

If a table has many columns, and you query many different combinations of columns, it might be efficient to split the less-frequently used data into separate tables with a few columns each, and relate back to the main table by duplicating the numeric ID column from the main table.

## 8.3.5. Column Indexes

The B-tree data structure lets the index quickly find a specific value, a set of values, or a range of values, corresponding to operators such as =, >, <=, BETWEEN, IN, and so on, in a WHERE clause.

### 8.3.5.1. Index Prefixes

With col_name(N) syntax in an index specification for a string column, you can create an index that use only the first N characters of the column. Indexing only a prefix of columns values in this way can make the index file much smaller. When you create index a BLOB or TEXT column, you must specify a prefix length for the index. 

```sql
CREATE TABLE test(blob_col BLOB, INDEX(blob_col(10)));
```
