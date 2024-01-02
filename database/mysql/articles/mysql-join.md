# MySQL join

## Introduction to MySQL join clauses

A relational database consists of multiple related tables linking together using common columns, which are known as fk columns. Because of this, the data in each table is incomplete from business perspective.

For example, in the sample database, we have the orders and orderdetails tables that are linked using the orderNumber column.

To get complete order information, you need to query data from both orders and orderdetails tables.

That's why joins come into the play.

A join is a method of linking data between one (self-join) or more tables based on the values of the common column between the tables.

MySQL supports the following types of joins:

1. Inner join
2. Left join
3. Right join
4. Cross join

## MySQL INNER JOIN clause

The following shows the basic syntax of the inner join clause that joins two tables table_1 and table_2:

```sql
SELECT column_list
FROM table_1
INNER JOIN table_2 ON join_condition;
```

The inner join clause joins two tables based on a condition which is known as a join predicate.

The inner join clause compares each row from the first table with every row from the second table. If values from both rows satisfy the join condition, the inner join clause creates a new row whose column contains all columns of the two rows from both tables and includes this new row in the result set. In other words, the inner join clause includes only matching rows from both tables.

If the join condition use the equality operator ( = ) and the column names in both tables used for matching are the same, and you can use `USING` clause instead:

```sql
SELECT column_list
FROM tbl_1
INNER JOIN tbl_2 USING (column_name);
```

The following statement uses an inner join clause to find members who are also the committee members:

```sql
SELECT
    m.member_id,
    m.name AS member,
    c.commitee_id,
    c.name AS committee
FROM
    members m
INNER JOIN committees c ON c.name = m.name;
```

```sql
+-----------+--------+--------------+-----------+
| member_id | member | committee_id | committee |
+-----------+--------+--------------+-----------+
|         1 | John   |            1 | John      |
|         3 | Mary   |            2 | Mary      |
|         5 | Amelia |            3 | Amelia    |
+-----------+--------+--------------+-----------+
3 rows in set (0.00 sec)
```

In this example, the inner join clause uses the values in the `name` columns in both tables `members` and `committees` to match. The following Venn diagram illustrates the inner join:

Because both tables use the same column to match, you can use the `USING` clause as shown in the query:

```sql
SELECT
    m.member_id,
    m.name AS member,
    c.committee_id,
    c.name AS committee
FROM
    members m
INNER JOIN committees c using(name);

+-----------+--------+--------------+-----------+
| member_id | member | committee_id | committee |
+-----------+--------+--------------+-----------+
| 1         | John   | 1            | John      |
| 3         | Mary   | 2            | Mary      |
| 5         | Amelia | 3            | Amelia    |
| 1         | John   | 5            | John      |
+-----------+--------+--------------+-----------+
```

## MySQL LEFT JOIN clause

Similar to an inner join, a left join also requires a join predicate. When joining two tables using a left join, the concepts of left and right tables are introduced.

The left join selects data starting from the left table. For each row in the left table, the left join compares with every row in the right table.

If the values in the two rows satisfy the join condition, the left join clause creates a new row whose columns contain all columns of the rows in both tables and includes this row in the result set.

If the values in the two rows are not matched, the left join clause still creates a new row whose columns contain columns of the row in the left table and `NULL` for columns of the row in the right table.

In other words, the left join selects all data from the left table whether there are matching rows exist in the right table or not.

In case these are no matching rows from the right table found, the left join uses `NULLs` for columns of the row from the right table in the result set.

To find members who are not the committee members, you add a `WHERE` clause and `IS NULL` operator as follows:

```sql
SELECT
    m.member_id,
    m.name AS member,
    c.committee_id,
    c.name AS committee
FROM members m
LEFT JOIN committees c USING(name)
WHERE c.committee_id IS NULL
```

## MySQL CROSS JOIN clause

Unlike the inner join, left join, and right join, the cross join clause does not have a join condition.

The cross join makes a Cartesian product of rows from the joined tables. The cross join combines each row from the first table with very row from the right table to make the result set.

Suppose the first table has n rows and second has m rows. The cross-join that joins the tables will return nxm rows.

```sql
SELECT select_list
FROM table_1
CROSS JOIN table_2;
```
