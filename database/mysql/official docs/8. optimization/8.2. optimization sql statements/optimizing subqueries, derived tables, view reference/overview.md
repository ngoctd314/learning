# Optimizing Subqueries, Derived Tables, View References and Common Table Expressions

The MySQL query optimizer has different strategies available to evaluate subqueries:

**For a subquery used with an IN, = ANY, or EXISTS predicate, the optimizer has these choices:**

- Semijoin
- Materialization
- Exists strategy

**For a subquery used with a NOT IN, <> ALL or NOT EXISTS predicate, the optimizer has these choices**

- Materialization
- EXISTS strategy

For a derived table, the optimizer has these choices (which also apply to view references and common table expressions):

- Merge the derived table into the outer query block
- Materialize the derived table to an internal temporary table

**Semijoin**

A semijoin is an operation in relational databases where one table is filtered based on the existence of matching rows in another table. It's a type of join operation that doesn't necessarily return the actual matching rows from the second table but rather filters the rows from the first table based on whether a match exists in the second table.

- Matching Rows: Initially, the database performs an operation similar to an inner join between the two tables involved in the semijoin. It identifies the rows in the first table that have matching values in the second table based on the specified join condition.
- Filtering: Instead of returning the combined rows from both tables like in a regular join, a semijoin only returns the rows from the first table that have matches in the second table. The actual rows from the second table are not included in the result; only the rows from the first table are retained.

Semijoins are useful in scenarios where you want to filter one table based on the existence of related rows in another table without needing the actual data from the second table in the final result set. They are commonly used in SQL queries to optimize performance and reduce the amount of data that needs to be processed. 

In SQL, semijoins can expressed using various techniques, including EXISTS subqueries, IN subqueries, or using the INTERSECT operator in some databases.

```sql
SELECT *
FROM tbl1
WHERE EXISTS (
    SELECT 1 
    FROM tbl2 WHERE tbl1.col = tbl2.col
);
```

**Materialize**

Materialize, in the context of databases and query optimization, refers to the process of physically storing intermediate query results in temporary tables or memory structures during query execution. This technique is used to optimize query performance by reducing the need for repeated calculations or expensive operations. 

Here's how materialize typically works:

**1. Query Evaluation:** When a complex query is executed, it may involve multiple operations such as joins, aggregations, and sorting. 

**2. Intermediate Results:**
