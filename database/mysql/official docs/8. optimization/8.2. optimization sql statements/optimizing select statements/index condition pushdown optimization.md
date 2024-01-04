# Index Condition Pushdown Optimization

Index Condition Pushdown (ICP) is an optimization for the case where MySQL retrieves rows from a table using an index. Without ICP, the storage engine traverses the index to locate rows in the base table and returns them to the MySQL server which evaluates the `WHERE` condition for the rows. With ICP enabled, and if parts of the `WHERE` condition for the rows. With ICP enabled, and if parts of the `WHERE` condition down to the storage engine. The storage engine then evaluates the pushed index condition by using the index entry and only if this is satisfied is the the row read from the table. ICP can reduce the number of times the storage engine must access the base table and the number of times MySQL server must access the storage engine.

Applicability of the Index Condition Pushdown optimization is subject to these conditions:

```sql
CREATE TABLE tbl (id int AUTO_INCREMENT PRIMARY KEY, a int, b int, key idx(a));
INSERT INTO tbl (a, b) VALUES (1, 1), (2, 2), (3, 1), (4, 1), (1, 3), (2, 2), (3, 4);

EXPLAIN SELECT * FROM tbl WHERE a = 1 AND b = 3;
+----+-------------+-------+------------+------+---------------+-----+---------+-------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key | key_len | ref   | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+-----+---------+-------+------+----------+-------------+
| 1  | SIMPLE      | tbl   | <null>     | ref  | idx           | idx | 5       | const | 2    | 14.29    | Using where |
+----+-------------+-------+------------+------+---------------+-----+---------+-------+------+----------+-------------+
```
