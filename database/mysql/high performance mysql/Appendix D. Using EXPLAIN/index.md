# Using EXPLAIN

This appendix shows you how to invoke EXPLAIN to get information about the query execution plan, and how to interpret the output. The EXPLAIN command is the main way to find out how the query optimizer decides to execute queries. This feature has limitations and doesn't always tell the truth, but its output is the best information available, and it's worth studying so you can learn how your queries are executed. Learning to interpret EXPLAIN will also help you learn how MySQL's optimizer works.

## Invoking EXPLAIN
 
To use EXPLAIN, simply add the word EXPLAIN just before the SELECT keyword in your query. MySQL will set a flag on the query. When it executes the query, the flag causes it to return information about each step in the execution plan, instead of executing it.

```sql
EXPLAIN SELECT 1\G
***************************[ 1. row ]***************************
id            | 1
select_type   | SIMPLE
table         | <null>
partitions    | <null>
type          | <null>
possible_keys | <null>
key           | <null>
key_len       | <null>
ref           | <null>
rows          | <null>
filtered      | <null>
Extra         | No tables used
```

There's one row in the output per table in the query. If the query joins two tables, there will be two rows of output. An aliased table counts as a separate table, so if you join a table to itself, there will be two rows in the output. The meaning of "table" is fairly broad here: it can mean a subquery, a UNION result, and so on. You'll see later why this is so.

There are two important variations on EXPLAIN:

- EXPLAIN EXTENDED appears to behave just like a normal EXPLAIN, but it tells the server to "reverse compile" the execution plan into a SELECT statement. You can see this generated statement by running SHOW WARNINGS immediately afterward.
- EXPLAIN PARTITIONS shows the partitions the query will access, if applicable. It is available only in MySQL 5.1 and newer.

It's a common mistake to think that MySQL doest execute a query when you add EXPLAIN to it. In fact, if the query contains a subquery in the FROM clause, MySQL actually executes the subquery, places its results into a temporary table, and then finishes optimizing the outer query. It has to process all such subqueries before it can optimize the outer query fully, which is must do for EXPLAIN. This means EXPLAIN can actually cause a great deal of work for the server if the statement contains expensive subqueries or views that use the TEMPTABLE algorithm.

Bear in mind that EXPLAIN is an approximation, nothing more. Sometimes it's a good approximation, but at other times, it can be very far from the truth. Here are some of its limitations:

- EXPLAIN doesn't tell you anything about how triggers, stored functions, or UDFs will affect your query.
- It doesn't work for stored procedures, although you can extract the queries manually and EXPLAIN the individually.
- It doesn't tell you about ad hoc optimizations MySQL does during query execution.
- Some of the statistics it shows are estimates and can be very inaccurate.
- It doesn't show you everything there is to know about query's execution plan.
- It doesn't distinguish between some things with the same name. For example, it uses "filesort" for in-memory sorts and for temporary files, and it displays "Using temporary" for temporary tables on disk and in memory.
- It can be misleading

## Rewriting Non-SELECT Queries

MySQL explains only SELECT queries, not stored routine calls or INSERT, UPDATE, DELETE, or any other statements.

## The columns in EXPLAIN

EXPLAIN's output always has the same columns (except for EXPLAIN EXTENDED, which adds a filtered column in MySQL 5.1, and EXPLAIN PARTITIONS, which adds a partitions column). The variability is in the number and contents of the rows. However, to keep our examples clear, we don't always show all columns in this appendix.

## The id Column

This column always contains a number, which identifies the SELECT to which the row belongs. If there are no subqueries or unions in the statement, there is only one SELECT, so every row will show a 1 in this column. Otherwise, the inner SELECT statements generally will be numbered sequentially, according to their positions in the orginal statement.

MySQL divides SELECT queries into simple and complex types, and the complex types can be grouped into three broad classes: simple subqueries, so-called derived tables (subqueries in the FROM clause), and UNIONs. Here's a simple subquery:

```sql
mysql> EXPLAIN SELECT (SELECT 1 FROM sakila.actor LIMIT 1) FROM sakila.film;


+----+-------------+-------+
| id | select_type | table |
+----+-------------+-------+
| 1  | PRIMARY     | tbl   |
| 2  | SUBQUERY    | tbl   |
+----+-------------+-------+
```

## The select_type Column

## The table Column

This column shows which table the row is accessing. You can read this column from top to bottom to see the join order MySQL's join optimizer chose for the query.

## Derived tables and unions

## The type Column
