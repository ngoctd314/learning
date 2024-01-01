# Explain

```txt
{EXPLAIN | DESCRIBE | DESC}
    tbl_name [col_name | wild]

{EXPLAIN | DESCRIBE | DESC}
    [explain_type]
    {explainable_stmt | FOR connection connection_id}

{EXPLAIN | DESCRIBE | DESC} ANALYZE [FORMAT = TREE] select_statement

explain_type: {
    FORMAT = format_name
}

format_name: {
    TRADITIONAL | JSON | TREE
}

explainable_stmt: {
    SELECT statement
  | TABLE statement
  | DELETE statement
  | INSERT statement
  | REPLACE statement
  | UPDATE statement
}
```

The DESCRIBE and EXPLAIN statments are synonyms. In practice, the DESCRIBE keyword is more often used to obtain information about table structure, whereas EXPLAIN is used to obtain a query execution plan (that is, an explanation of how MySQL would execute a query).

## Obtaining Table Structure Information

DESCRIBE provides information about the columns in a table:

## Obtaining Execution Plan Information

The EXPLAIN statement provides information about how MySQL executes statements:

When EXPLAIN is used with an explainable statement, MySQL displays information from the optimizer about the statement execution plan. That is, MySQL explains how it would process the statement, including information about how tables are joined and in which order.

The FORMAT option can be used to select the output format. TRADITIONAL presents the output in tabular format. This is the default if no FORMAT option is present. JSON format displays the information in JSON format.

EXPLAIN requires the same privileges required to execute the explained statement.

With the help of EXPLAIN, you can see where you should add indexes to tables so that the statement executes faster by using indexes to find rows. You can also use EXPLAIN to check whether the optimizer joins the tables in an optimal order. To give a hint to the optimizer to use a join order corresponding to the order in which the tables are named in a SELECT statement, begin the statement with SELECT STRAIGHT_JOIN rather than just SELECT.

If you have a problem with indexes not being used when you belive that they should be, run ANALYZE TABLE to update table statistics, such are cardinality of keys, that can affect the choices the optimizer makes. 

## Obtaining information with EXPLAIN ANALYZE

## EXPLAIN Output Format

The EXPLAIN statement provides information about how MySQL executes statements.

EXPLAIN returns a row of information for each table used in the SELECT statement. It lists the tables in the output in the order that MySQL would read them while processing the statement. This means that MySQL reads a row from the first table, then finds a matching row in the second table, and then in the third table, and so on. When all tables are processed, MySQL outputs the selected columns and backtracks through the table list until a table is found for which there are more matching rows. The next row is read from this table and the process continues with the next table.

## Explain Output Columns

This section describes the output columns produced by EXPLAIN. Later sections provide additional information about type and Extra columns.

|Column|JSON name|Meaning|
|-|-|-|
|id|select_id|The SELECT identifier|
|select_type|None|The SELECT type|
|table|table_name|The table for the output row|
|partitions|partitions|The matching partitions|
|type|access_type|The join type|
|possible_keys|possible_keys|The possible indexes to choose|
|key|key|The index actually chosen|
|key_len|key_length|The columns compared to the index|
|ref|ref|The columns compared to the index|
|rows|rows|Estimate of rows to be examined|
|filtered|filtered|Percentage of rows filtered by table condition|
|Extra|None|Additional information|

- id (JSON name: select_id)

The SELECT identifier. This is the sequential number of the SELECT within the query. The value can be NULL if the row refers to the union result of other rows.

- select_type (JSON name: none)

The type of SELECT, which can be any of those shown in the following table. A JSON-formatted EXPLAIN exposes the SELECT type as a property of a query_block.

- table (JSON name: table_name)

- type (JSON name: access_type)

The join type.

- possible_keys (JSON name: possible_keys)

The possible_keys column indicates the indexes from which MySQL 

- key(JSON name: key)

The key column indicates the key(index) that MySQL actually decide to use. If MySQL decides to use one of the possible_keys value. This can happen if none of the possible_keys indexes are suitable for looking up rows, but all the columns selected by the query are columns of some other index. That is, the named index covers the selected columns, so although it is not used to determine which rows to retrive 

For InnoDB, a secondary index might cover the selected columns even if the query also selects the primary key because InnoDB stores the primary key value with each secondary index. If ke is NULL, MySQL found no index to use for executing the query more efficiently.

To force MySQL to use or ignore an index listed in the possible_keys column, use FORCE INDEX, USE INDEX or IGNORE INDEX in your query.

- key_len(JSON name: key_length)

The key_len column indicates the length of the key that MySQL decided to use. The value of key_len enables you to determine how many parts of a multiple-part key MySQL actually uses. If the key column says NULL, the key_len column also says NULL.

Due to the key storage format, the key length is one greater for a column that can be NULL than for a NOT NULL column.

- ref(JSON name:ref)

The ref column shows which columns or constants are compared to the index named in the key column to select rows from the table. If the value is func, the value used is the result of some function.

- rows (JSON name: rows)

The rows column indicates the number of rows MySQL believes it must be examine to execute the query. For InnoDB tables, this number is an estimate, and may not always be exact.

- filtered(JSON name: filtered)

The filtered column indicates an estimated percentage of table rows that are filtered by the table condition. The maximum value is 100, which means no filtering of rows occurred. Values decreasing from 100 indicate increasing amounts of filtering. rows shows the estimated number of rows exmained and rows x filtered shows the number of rows that are joined with the following table.

- Extra(JSON name:none)

This column contains additional information about how MySQL resolves the query.

## EXPLAIN Join Types

|join type|mean|
|-|-|
|system|The table has only one row (= system table). This is a special case of the const join type.|
|const|The table has at most one matching row, which is read at the start of the query. Because there is only one row, values from the column in this row can be regard as constants by the rest of the optimizer. const tables are very fast because they are read only once. const is used when you compare all parts of PRIMARY KEY or UNIQUE index to constant values.|
|eq_ref|One row is read from this table for each combination of rows from the previous tables. Other than the system and const types, this is the best possible join type. Is is used when all parts of an index are used by the join and the index is a PRIMARY KEY of UNIQUE NOT NULL index. eq_ref can be used for indexed columns that are compared using the = operator. The comparison value can be a constant or an expression that uses columns from tables that are read before this table.|
|ref|All rows with matching index values are read from this table for each combination of rows from the previous tables. ref is used if the join uses only a leftmost prefix of the key of if the key is not a PRIMARY KEY or UNIQUE index (in other words, if the join cannot select a single row based on the key value). If the key that is used matches only a few rows, this is a good join type.|
|fulltext|The join is performed using a FULLTEXT index.|
|ref_or_null|This join is like ref, but with the additional that MySQL does an extra search for rows that contain NULL values. This join type optimization is used most often is resolving subqueries.|index_merge|This join type indicates that the Index Merge optimization is used. In this case, the key column in the output row contains a list of indexes used, and key_len contains a list of longest key parts for indexes used.|
|unique_subquery|This type replaces eq_ref for some IN subqueries of the following: value IN (SELECT primary_key FROM single_table WHERE some_expr).|
|index_subquery|-|
|range|Only rows that are in a given range are retrieved, using an index to select the rows. The key column in the output row indicates which index is used. The key_len contains the longest key part that was used. The ref column is NULL for this type. range can be used when a key is compared to a constant using any of the =, <>, >=, <=, IS NULL, <=>, BETWEEN, LIKE, or IN().|
|index|The index join type is the same as ALL, except that the index tree is scanned. This occurs two ways: If the index is a covering index for the queries and can be used to satify all data required from the table, only the index tree is scanned. In this case, the Extra column says Using Index. An index-only scan usually is faster than ALL because the size of the index usually is smaller than the table data. A full table scan is performed using reads from the index to look up data rows in index order. Use index does not appear in the Extra column. MySQL can use this join type when the query uses only columns that are part of a single index.|
|ALL|A full table scan is done for each combination of rows from the previous tables. This is normally not good if the table is the first table not marked const, and usually very bad in all other case. Normally, you can avoid ALL by adding indexes that enable row retrieval from the table based on constant values or column values from earilier tables.|

## EXPLAIN Extra Information

The Extra column of EXPLAIN output contains additional information about how MySQL resolves the query. The following list explains the values that can appear in this column. Each item also indicates for JSON-formatted output which property displays the Extra value. For some of these, there is a specific property.

If you want to make your queries as fast as possible, look out for Extra column values of Using filesort and Using temporary, or, in JSON-formatted EXPLAIN output, for using_filesort and using_temporary_table properties equal to true.

- Backward index scan (JSON: backward_index_scan)

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

## Type

In the context of MySQL's `EXPLAIN` statement, the `type` column in the output represents the access method that MySQL uses to retrieve rows from tables. The values in the `type` column indicate the type of index or scan that MySQL uses for a particularly table access. Here are the common values for the `type` column:

*1. system:* This is a const join type. It's used when the table has only one row, and the optimizer knows this at compile time. 

*2. const:* This is also a constant type. It's used when the table has at most one matching row, and the optimizer knows this at compile time. This is a very efficient access method because it reads only one row.

Both "system" and "const" join types are used when the optimizer can ascertain that the accessed table has only one row.

- "System" is used when the optimizer knows this information during the query optimization phase.
- "Const" is used when the condition of having at most one matching row is known at compile time.

In practical terms, you may encounter "const" frequently, and it is often associated with queries that involve accessing a table by its primary key or a unique index, where the optimizer can determine that only one row will match the conditions. 

```sql
CREATE TABLE example_table (
    id INT PRIMARY KEY,
    name VARCHAR(255)
);

-- Insert a single row into the table
INSERT INTO example_table (id, name) VALUES (1, 'John Doe');

-- Query using the "system" join type
EXPLAIN SELECT * FROM example_table WHERE id = 1;
```

*3. eq_ref:* This join type is used for indexed lookups on unique or primary key indexes. It occurs when all parts of an index are used by the join, and only one matching row is expected.

```sql
create table tbl1 (id int auto_increment primary key);
create table tbl2 (id int auto_increment primary key);

explain select * from tbl1 join tbl2 on tbl2.id = tbl1.id;
+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------------+------+----------+--------+
| id | select_type | table | partitions | type   | possible_keys | key     | key_len | ref                   | rows | filtered | Extra  |
+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------------+------+----------+--------+
| 1  | SIMPLE      | tbl1  | <null>     | ALL    | PRIMARY       | <null>  | <null>  | <null>                | 1    | 100.0    | <null> |
| 1  | SIMPLE      | tbl2  | <null>     | eq_ref | PRIMARY       | PRIMARY | 4       | learn_explain.tbl1.id | 1    | 100.0    | <null> |
+----+-------------+-------+------------+--------+---------------+---------+---------+-----------------------+------+----------+--------+
```

*4. ref:* This join type is used for indexed lookups on non-unique indexes. It occurs when only a subset of the index is used in the join condition.

```sql
alter table tbl2 add index idx_name(name);
explain select * from tbl2 where tbl2.name = 'ngoctd';
+----+-------------+-------+------------+------+---------------+----------+---------+-------+------+----------+--------+
| id | select_type | table | partitions | type | possible_keys | key      | key_len | ref   | rows | filtered | Extra  |
+----+-------------+-------+------------+------+---------------+----------+---------+-------+------+----------+--------+
| 1  | SIMPLE      | tbl2  | <null>     | ref  | idx_name      | idx_name | 1023    | const | 1    | 100.0    | <null> |
+----+-------------+-------+------------+------+---------------+----------+---------+-------+------+----------+--------+

explain select * from tbl1 join tbl2 on tbl2.name = tbl1.name;
+----+-------------+-------+------------+------+---------------+----------+---------+-------------------------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key      | key_len | ref                     | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+----------+---------+-------------------------+------+----------+-------------+
| 1  | SIMPLE      | tbl1  | <null>     | ALL  | <null>        | <null>   | <null>  | <null>                  | 1    | 100.0    | Using where |
| 1  | SIMPLE      | tbl2  | <null>     | ref  | idx_name      | idx_name | 1023    | learn_explain.tbl1.name | 1    | 100.0    | <null>      |
+----+-------------+-------+------------+------+---------------+----------+---------+-------------------------+------+----------+-------------+
```

*5. range:* This join type is used when the optimizer can use an index to retrieve only rows that are within a certain range.

```sql
explain select * from tbl2 where name like 'a%';
+----+-------------+-------+------------+-------+---------------+----------+---------+--------+------+----------+-----------------------+
| id | select_type | table | partitions | type  | possible_keys | key      | key_len | ref    | rows | filtered | Extra                 |
+----+-------------+-------+------------+-------+---------------+----------+---------+--------+------+----------+-----------------------+
| 1  | SIMPLE      | tbl2  | <null>     | range | idx_name      | idx_name | 1023    | <null> | 1    | 100.0    | Using index condition |
+----+-------------+-------+------------+-------+---------------+----------+---------+--------+------+----------+-----------------------+
```

*6. index:* This join type is similar to `ALL`, but it's more efficient. It occurs when the optimizer accesses the entire index to fullfill the query, without reading the actual data rows.

*7. all:* This is the least efficient join type. It occurs when the optimizer performs a full table scan, reading all rows.

```sql
drop index idx_name on tbl2;

explain select * from tbl2 where name = 'a';
+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key    | key_len | ref    | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
| 1  | SIMPLE      | tbl2  | <null>     | ALL  | <null>        | <null> | <null>  | <null> | 4    | 25.0     | Using where |
+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
```

## Ref

In the `EXPLAIN FORMAT=JSON` output in MySQL, the `cost_info` section provides information about the estimated costs associated with the query execution. Here's an explaination of the fields in the `cost_info` section.

- `read_cost`: This field represents the cost of reading data during the query execution. It's an estimate of the cost associated with accessing and retriving data from storage.
- `eval_cost`: The `eval_cost` field represents the cost of evaluating expressions during the query execution. It includes the cost of evaluating conditions, expressions, and other computations.
- `prefix_cost`: The `prefix_cost` field in the sum of the `read_cost` and `eval_cost`. It represents the total cost associated with the access method and expression evaluation.
- `data_read_per_join`: This field indicates the estimated amount of data read per join. It provides information about the volumne of data that needs to be read and processed during the execution of the query.

## Extra

The Extra column of EXPLAIN output contains additional information about how MySQL resolves the query. The following list explains the values that can appear in this column. Each item also indicates for JSON-formatted output which property displays the Extra value. For some of these, there is a specific property.

If you want to make your queries as fast as possible, look out for Extra column values of Using filesort and Using temporary, or, in JSON-formatted EXPLAIN output, for using_filesort and using_temporary_table properties equal to true.

- Backward index scan (JSON: backward_index_scan)

The optimizer is able to use a descending index on an InnoDB table. Shown together with Using index.

```sql
CREATE TABLE tbl (id int unsigned AUTO_INCREMENT PRIMARY KEY, cnt int);
INSERT INTO tbl (cnt) VALUES (3), (1), (2), (9), (10);

EXPLAIN SELECT * FROM tbl ORDER BY cnt DESC;
+----+-------------+-------+------------+-------+---------------+---------+---------+--------+------+----------+----------------------------------+
| id | select_type | table | partitions | type  | possible_keys | key     | key_len | ref    | rows | filtered | Extra                            |
+----+-------------+-------+------------+-------+---------------+---------+---------+--------+------+----------+----------------------------------+
| 1  | SIMPLE      | tbl   | <null>     | index | <null>        | idx_cnt | 5       | <null> | 5    | 100.0    | Backward index scan; Using index |
+----+-------------+-------+------------+-------+---------------+---------+---------+--------+------+----------+----------------------------------+
```

- Distinct(JSON property: distinct)

MySQL is looking for distinct values, so it stops searching for more rows for the current row combination after it has found the first matching row.


- Full scan on NULL key???

This occurs for subquery optimization as a fallback strategy when the optimizer cannot use an index-lookup access method.

- Impossible HAVING(JSON property:message)

The HAVING clause is always false and cannot select any rows.

- Impossible WHERE 

- LooseScan (m...n) (JSON property: message)

- no matching row in const table (JSON property:message)

For a query with a join, there was an empty table or a table with no rows satisfying a unique index condition.

- Not exists(JSON property:message)

MySQL was able to do a LEFT JOIN optimization

- Using filesort(JSON property: using_filesort)

MySQL must do an extra pass to find out how to retrieve the rows in sorted order. The sort is done by going through all rows according to the join type and sorting the sort key and pointer to the row for all rows that match the *WHERE* clause. The keys then are sorted and the rows are retrieved in sorted order. 

- Using index (JSON property:using_index)

The column information is retrieved from the table using only information in the index tree without having to do an additional seek to read the actual row. This strategy can be used when the query uses only columns that are part of a single index.

- Using index condition (JSON property:using_index_condition)

Tables are read by accessing index tuples and testing them first to determine whether to read full table rows. In this way, index information is used to defer ("push down") reading full table rows unless it is necessary.

- Using where (JSON property:attached_condition)

A WHERE clause is used to restrict which rows to match against the next table or send to the client. Unless you specifically intend to fetch or examine all rows from the table, you may have something wrong in your query if the Extra value is not Using where and the table join type is ALL or index.

Using where has no direct counterpart in JSON-formatted output; the attached_condition property contains any WHERE condition used.

- Using temporary (JSON property:using_temporary_table)

To resolve the query, MySQL needs to create a temporary table to hold the result. This typically happens if the query contains GROUP BY and ORDER BY clauses that list column differently.

- Using sort_union(...), Using  union(...), Using intersect(...) (JSON property:message)
