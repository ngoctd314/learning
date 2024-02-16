# Multiple-Column Indexes

MySQL can create composite indexes (that is, indexes on multiple columns). An index may consist of up to 16 columns. For certain data types, you can index a prefix of the column.

MySQL can use multiple-column indexes for queries that test all the columns in the index, or queries that test just the first column, the first two columns, the first three columns, and so on. If you specify the columns in the right order in the index definition, a single composite index can speed up several kinds of queries on the same table.

**Note**

As an alternative to a composite index, you can introduce a column that is "hashed" based on information from other columns. If this column is short, reasonably unique, and indexed, it might be faster than a "wide" index on many columns. In MySQL, it is very easy to use this extra column:

```sql
SELECT * FROM tbl_name
    WHERE hash_col=MD5(CONCAT(val1, val2))
    AND col1=val1 AND col2=val2;
```

Suppose that a table has following specifications:

```sql
CREATE TABLE tests (
    id INT NOT NULL,
    last_name CHAR(30) NOT NULL,
    first_name CHAR(30) NOT NULL,
    PRIMARY KEY (id),
    INDEX name (last_name, first_name)
);
```

The name index is an index over the last_name and first_name columns. The index can be used for lookups in queries that specify values in a known range for combinations of last_name and first_name values. It can also be used for queries that specifies just a last_name value because that column is a leftmost prefix of the index. 

Therefore, the name index is used for lookups in the following queries: 

```sql
SELECT * FROM tests WHERE last_name = 'Jones';

SELECT * FROM tests WHERE last_name = 'Jones' AND first_name = 'John';

SELECT * FROM tests WHERE last_name = 'Jones' AND (first_name = 'John' OR first_name = 'Joh');

SELECT * FROM tests WHERE last_name = 'Jones' AND first_name >= 'M' AND first_name < 'N';
```

However, the name index is not used for lookups in the following queries

```sql
SELECT * FROM tests WHERE first_name = 'John';
SELECT * FROM tests WHERE last_name = 'John' OR first_name = 'John';
```
Suppose that u issue the following SELECT statement: 

```sql
SELECT * FROM tbl_table
    WHERE col1=val1 AND col2=val2;
```

If a multiple-column index exists on col1 and col2, the appropriate rows can be fetched directly. If separate single-column indexes exist on col1 and col2, the optimizer attempts to  use the Index Merge optimization, or attempts to find the most restrictive index by deciding which index excludes more rows and using that index to fetch the rows.

If the table has a multiple-column index, and leftmost prefix of the index can be used by the optimizer to look up rows. For example, if you have a three-column index on (col1, col2, col3), you have  indexes seach capabilities on (col1), (col1, col2) and (col1, col2, col3).

MySQL cannot use the index to perform lookups if the columns do not form a leftmost prefix of the index. Suppose that you have the SELECT statements shown here:

```sql
SELECT * FROM tbl_name WHERE col1=val1; 
SELECT * FROM tbl_name WHERE col1=val1 AND col2=val2;

SELECT * FROM tbl_name WHERE col2=val2;
SELECT * FROM tbl_name WHERE col2=val2 AND col3=val3;
```

If an index exists on (col1, col2, col3), only the first two queries use the index. The third and fourth queries do involve indexes columns, but do not use an index to perform lookups because (col2) and (col2, col3) are not leftmost prefixes of (col1, col2, col3).
