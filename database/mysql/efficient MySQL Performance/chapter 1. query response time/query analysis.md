# Query Analysis

## Rows examined

Rows examined is the number of rows that MySQL accessed to find matching rows. It indicates the selectivity of the query and the indexes.

To understand rows examined, let's look at two examples. First, let's use the following table, t1, and three rows:

```sql
CREATE TABLE t1 (
    id int NOT NULL,
    c char(1) NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB;
+----+---+
| id | c |
+----+---+
| 1  | a |
| 2  | b |
| 3  | c |
+----+---+
```

Column id is the primary key, and column c is not indexed.

The query `SELECT c FROM t1 WHERE c = 'b'` matches one row but examines three rows because there is no unique index on column c. Therefore, MySQL has no idea how many rows match the WHERE clause. We can see that only one row matches, but MySQL doesn't have eyes, it has indexes.

For the second example, let's use the following table, t2, and seven rows:

```sql
CREATE TABLE t2 (
    id int NOT NULL,
    c char(1) NOT NULL,
    d varchar(8) DEFAULT NULL,
    PRIMARY KEY (id),
    KEY c (c)
);
+----+------+--------+
| id | c    | d      |
+----+------+--------+
| 1  | a    | apple  |
| 2  | a    | ant    |
| 3  | a    | acorn  |
| 4  | a    | apron  |
| 5  | b    | banana |
| 6  | b    | bike   |
| 7  | c    | car    |
+----+------+--------+
```

How many rows will query `SELECT d FROM t2 WHERE c = 'a' AND d = 'acorn'` examine? The answer is: four. MySQL uses the nonunique index on column c to look up rows matching the condition c = 'a', and that matches four rows. And to match the other condition, d = 'acorn', MySQL examines each of those four rows. As a result, the query examines four rows but matches (and returns) only one row.

Rows examined only tells half the story. The other half is rows sent. 

## Rows sent

Rows sent is the number of rows returned to the client - the result set size. Rows sent is most meaningful in relation to rows examined

- Rows sent = Rows examined

The ideal case is when rows sent and rows examined are equal and the value is relatively small, especially as a percentage of total rows, and query response time is acceptable.

- Rows sent < Rows examined

Fewer rows sent than examined is a reliable sign of poor query or index selectivity. If the difference is extreme, it likely explains slow response time.

- Rows sent > Rows examined

It’s possible, but rare, to send more rows than were examined

Rows sent is rarely a problem by itself. Modern networks are fast and the MySQL protocol is efficient. This is usually small, but it can be large if the query returns BLOB or JSON columns.

## Rows affected

Row affected is the number of rows inserts, updated, or deleted.
