# Index Condition Pushdown Optimization

Index Condition Pushdown (ICP) is an optimization for the case when MySQL retrieves rows from a table using an index. Without ICP, the storage engine traverses the index to locate rows in the base table and returns them to the MySQL server which evaluates the WHERE condition for the rows. With ICP enabled, and if parts of the `WHERE` condition can be evaluated by using only columns from the index, the MySQL server pushes this part of the `WHERE` condition down to the storage engine. The storage engine then evaluates the pushed index condition by using the index entry and only if this satisfied is the row read from the table. ICP can reduce the number of times the storage engine must access the base table and the number of times the storage engine must access the base table and the number of times the MySQL server must access the storage engine.

Applicability of the Index Condition Pushdown optimization is subject to these conditions:

- ICP is used for the range, ref, eq_ref, and ref_or_null access methods when there is a need to access full table rows.
- ICP can be used for InnoDB and MyISAM tables, including partitioned InnoDB and MyISAM tables.
- For InnoDB tables, ICP is used only secondary indexes. The goal of ICP is to reduce the number of full-rows reads and thereby reduce I/O operations. For InnoDB clustered indexes, the complete record is already read into the InnoDB buffer. Using ICP in this case does not reduce I/O.
- ICP is not supported with secondary indexes created on virtual generated columns. InnoDB supports secondary indexes on virtual generated columns.
- Conditions that refer to subqueries cannot be pushed down.
- Conditions that refer to stored functions cannot be pushed down. Storage engines cannot invoke stored functions.
- Triggered conditions cannot be pushed down.

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

To understand how this optimization works, first consider how an index can proceeds when Index Condition Pushdown is not used:

1. Get the next row, first by reading the index tuple, and then by using the index tuple to locate and read the full table row.
2. Test the part of the `WHERE` condition that applies to this table. Accept or reject the row based on the test result.

Using Index Condition Pushdown, the scan proceeds like this instead:

1. Get the next row's index tuple (but not the full table row).
2. Test the part of the `WHERE` condition that applies to this table and can be checked using only index columns. If the condition is not satisfied, process the index tuple to locate and read the full table row.
3. If the condition is satisfied, use the index tuple to locate and read the full table row.
4. Test the remaining part of the WHERE condition that applies to this table. Accept or reject the row based on the test result.

EXPLAIN output shows Using index condition in the Extra column when Index Condition Pushdown is used. It does not show Using index because that does not apply when full table rows must be read.

```sql
CREATE TABLE `people` (
  `id` int NOT NULL AUTO_INCREMENT,
  `zipcode` varchar(10) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx` (`zipcode`,`lastname`)
)
```

```go
type Data struct {
	ID       int    `db:"id"`
	Zipcode  string `db:"zipcode"`
	Lastname string `db:"lastname"`
	Address  string `db:"address"`
}

func main() {
	var in []Data
	for i := 1; i <= 1000 ; i++ {
		in = append(in, Data{
			Zipcode:  strconv.Itoa(i * 10000),
			Lastname: fmt.Sprintf("%detrunia%d", i, i),
			Address:  fmt.Sprintf("%dMain Street%d", i, i),
		})
	}

	_, err := mysqlConn.NamedExec("INSERT INTO people (zipcode, lastname, address) VALUES (:zipcode, :lastname, :address) ", in)
	if err != nil {
		log.Println(err)
	}
}
```

```sql
SELECT * FROM people
  WHERE zipcode='95054'
  AND lastname LIKE '%etrunia%'
  AND address LIKE '%Main Street%';
```

MySQL can use index to scan through people with zipcode = '95054'. The second part (lastname LIKE '%etrunia%') cannot be used to limit the number of rows that must be scanned, so without Index Condition Pushdown, this query must retrieve full table rows of all people who have zipcode = '95054'
