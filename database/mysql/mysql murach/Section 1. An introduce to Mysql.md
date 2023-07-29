# An introduction to MySQL

## An introduction to relational databases

Relationships exist between the primary key in one table and the foreign key in another table. The foreign key is simply one or more columns in a table that refer to a primary key in another table.

One-to-many relationships are the most common, two tables can also have a one-to-one or many-to-many relationship. If a table has a one-to-one relationship with another table, the data in the two tables could be stored in a single table. Because of that, one-to-one relationships are used infrequently.

In constrast, a many-to-many relationship is usually implemented by using an intermediate table that has a one-to-many relationships which the two tables in the many-to-many relationship. In the other words, a many-to-many relationship can usually be broken down into one-to-many relationship.

If you define a foreign key for a table in MySQL, you can have the foreign key enforce referential integrity. When MySQL enforces referential integrity, it makes sure that any changes to the data in the database don't create invalid relationships between tables. This helps to maintain the integrity of the data that's stored in the database.

When you define a column in a table, you assign properties to it. You must choose data type, identify whether the column can store a null value, you can also assign a default value to each column, each table can also contain a numeric column whose value is generated automatically by the DBMS. In MySQL, a column like this is called an auto increment column.

A join lets you combine data from two or more tables into a single result set

The most common type of join in an inner join. This type of join returns rows from both tables only if their related columns match.

An outer join returns rows from one table in the join even if the other table doesn't contain a matching row.

## How to retrieve data from a single table

By default, a column in the result set is given the same name as the column in the base table. If that's not what you want, you can specify a substitute name, or column alias, for the column.

### How to code arithmetic expressions

|Operator|Name|Order of precedence|
|-|-|-|
|*|Multiplication|1|
|/|Division|1|
|Div|Integer Division|1|
|% (MOD)|Modulo (remainder)|1|
|+|Addition|2|
|-|Subtraction|2|

### How to use the CONCAT function to join strings

An expression can include any of the functions that are supported by MySQL. A function performs an operation and returns a value.


With MySQL you don't have to code a FROM clause. This makes it easy to test expressions that include arithmetic operators and functions.

The DISTINCT keyword prevents duplicate (identical) rows from being included in the result set

### How to code the WHERE clause

Earlier in this chapter, I mentioned that to improve performance, you should code your SELECT statements so they retrieve only the columns you need.


**How to use the comparison operators**

- You can use a comparison operator to compare any two expressions. Since MySQL automatically converts the data for comparison.

- If the result of a comparison is a true value, the row being tested is included is the result set. If it's a false or null value, the row isn't included.

- Character comparisons performed on MySQL databases are not case-sensitive. So, for example, 'CA' and 'ca' are considered equivalent.

- If you compare a null value using one of these comparison operators, the result is always a null value.

**How to use the AND, OR and NOT logical operators**

- You can use the AND and OR logical operators to create compound conditions that consist of two or more conditions. You use the AND operator to specify that the search must satify both of the conditions, and you use the OR operator to specify that the search must satisfy at least one of the conditions.

- You can use the NOT operator to negate a condition. Because this can make the search condition unclear, you should rephrase the condition if possible so it doesn't use NOT.

- When a MySQL evaluates a compound condition, it evaluates the operations in this sequence: NOT, AND and OR. You can use parentheses to override this order of precedence or to clarify the sequence in which the operations are evaluated.

**How to use the IN operator**

```sql
WHERE test_expression [NOT] IN ( sub query | expression )

WHERE vendor_state NOT IN ('CA', 'NV', 'OR')

WHERE vendor_id IN 
    (SELECT vendor_id
    FROM invoices
    WHERE invoice_date = '2018-07-18'
    )
```

**How to use the BETWEEN operator**

- You can use the BETWEEN phrase to test whether an expression falls within a range of values
- You can use the NOT operator to test for an expression that's not within the given range.

**How to use the Like and REGEXP operators**

To retrieve rows that match a specific string pattern, or mask, you can use the LIKE or REGEXP operators.

Both the LIKE and REGEXP operators provide powerful functionality for finding information in a database. However, searches that use these operators sometimes run slowly since they can't use a tables' indexes.

**How to use the IS NULL clause**

- A null value represents a value that's unknown, unavailable or nto applicable. It isn't the same as a zero or an empty string.

**How to code the ORDER BY clause**

- The ORDER BY clause specifies how you want the rows in the result set sorted. You can sort by one or more columns, and you can sort each column in either ASC or DESC. ASC is default.

**- Null values appear first in the sort sequence, even if you're using DESC**

- By default, in an asc sort, special characters appear first in the sort sequence, followed by numbers, then letters.

- You can sort by any column in the base table regardless of whether it's included in the SELECT clause.

**How to sort by an alias, expression or column number**

```sql
-- An ORDER BY clause that uses an alias
SELECT vendor_name
    CONCAT (vendor_city, ', ', vendor_state, ' ', vendor_zip_code) AS address
FROM vendors
ORDER BY address, vendor_name

-- An ORDER BY clause that uses an expression
SELECT vendor_name
    CONCAT (vendor_city, ', ', vendor_state, ' ', vendor_zip_code) AS address
FROM vendors
ORDER BY CONCAT(vendor_contact_last_name, vendor_contact_first_name)

-- An ORDER BY clause that uses column positions
SELECT vendor_name
    CONCAT (vendor_city, ', ', vendor_state, ' ', vendor_zip_code) AS address
FROM vendors
ORDER BY 2, 1
```

**How to code the LIMIT clause**

```sql
-- LIMIT clause that starts with the first row
SELECT vendor_id, invoice_total
FROM invoices
ORDER BY invoice_total DESC
LIMIT 5

-- LIMIT clause that starts with the thrid row
SELECT invoice_id, vendor_id, invoice_total
FROM invoices
ORDER BY invoice_id
LIMIT 2, 3
```

- You can use the LIMIT clause to limit the number of rows returned by the SELECT statement. This clause takes one or two integer arguments.
- If you code a single argument, it specifies the maximum row count, beginning with the first row. If you code both argument, the offset specifies the first row to return, where the offset of the frist row is 0
- If you want to retrieve all of the rows from a certain offset to the end of the result set, set code -1 for the row count


## How to retrieve data from two or more tables

A join lets you combine columns from two or more tables into a single result set. The join condition indicates how the two tables should be compared. in most cases, they're compared based on the relationship between the primary key of the first table and a foreign key of the second table.

### How to work with inner joins

```sql
SELECT invoice_number, vendor_name
FROM vendors INNER JOIN invoices
    ON vendors.vendor_id = invoices.vendor_id
ORDER BY invoice_number
```

**Description**

- A join combines columns from two or more tables into a result set based on the join conditions you specify. For an inner join, only those rows that satisfy the join condition are included in the result set.

- A join condition names a column in each of the two tables involved in the join and indicates how the two columns should be compared. In most cases, you use the equal operator to retrieve rows with matching columns. However, you can also use any of the other comparison operators in a join condition.

- Tables are typically joined on the relationship between the PK in one table and a FK in the other table. However, you can also join tables based on relationships not defined in the db. (ad hoc relationships)

- If the two columns in a join condition have the same name, you must qualify them with the table name.

### How to use table aliases

When you name a table to be joined in the FROM clause, you can refer to the table by an alias.

```sql
SELECT select_list
FROM table_1 a1
    [INNER] JOIN table_2 a2
        ON a1.column_name operator a2.column_name
    [[INNER] JOIN table_3 a3
        ON a3.column_name operator a3.column_name]
```

**Description**

- A table alias is an alternative table name assigned in the FROM clause. You can use an alias which is typically just a letter or two, to make a SQL statement easier to code and read.
- If you assign an alias to a table, you must use that alias refer to the table throughout your query. You can't use the original table name.
- You can use an alias for one table in a join without using an alias for another table.

### How to join to a table in another database

```sql
SELECT vendor_name, customer_last_name, customer_first_name, vendor_state, vendor_city
FROM vendors v
    JOIN om.customers c
    ON v.vendor_zip_code = c.customer_zip
ORDER BY state, city
```
**Description**

- A MySQL server can store tables in multiple databases. These databases are sometimes referred to as schemas
- When you run a SELECT statement against one database, you can join to a table in another database if you have appropriate privileges. To do that, you must prefix the table name in the other database with the name of that database.

### How to use compound join conditions

Althgough a join condition typically consists of a single comparision, you can include two or more comparisons in a join condition using the AND and OR operators.

```sql
SELECT customer_firstname, customer_last_name
FROM customers c JOIN employees e
    ON c.customer_first_name = e.first_name
    AND c.customer_last_name = e.last_name
```

**Description**

- A join condition can inclulde two or more conditions connected by AND or OR operators

### How to use a self-join

A self-join joins a table to itself. Although self-joins are rare, they are sometimes useful for retrieving data that can't be retrieved any other way. For example, a self-join that returns from the Vendors table where the vendor is in a city and state that has at least one other vendor. In other words, it does not return a vendor if that vendor is the only vendor in that city and state.

This statement includes the DISTINCT keyword. That way, a vendor appears only once in the result set. Otherwise, a vendor would appear once for every other row with a matching city and state. For example, if a vendor is in a city and state that has nine other vendors in that city and state, this query would return nine rows for that vendor.

This example also shows how you can use columns other than key columns in a join condition. Keep in mind, however, that this is an unusual situation and you're not likely to code joins like this often.

```sql
SELECT DISTINCT v1.vendor_name, v1.vendor_city, v1.vendor_state
FROM vendors v1 JOIN vendors v2
    ON v1.vendor_city = v2.vendor_city AND -- same city
        v1.vendor_state = v2.vendor_state AND -- same state
        v1.vendor_name <> v2.vendor_name  -- exclude rows that match a vendor with itself
ORDER BY v1.vendor_state, v1.vendor_city
```

Result:
|vendor_name|vendor_city|vendor_state|
|-|-|-|
|Computer Library|Phoenix|AZ|
|AT&T|Phoenix|AZ|
|Wells Fargo Bank|Phoenix|AZ|
|Aztek Label|Anaheim|CA|
|Blue Shield of Califainia|Anaheim|CA|
|Abbey Office|Fresno|CA|
|California Business|Fresno|CA|
|Postmaster|Fresno|CA|

**Description**

- A self-join is a join that joins a table with itself

### How to join more than two tables

It's common for programmers to need to join data from more than two tables. 

```sql
SELECT vendor_name, invoice_number, invoice_date, line_item_amount, account_description
FROM vendors v
    JOIN invoices i
        ON v.vendor_id = i.vendor_id
    JOIN invoice_line_items li
        ON i.invoice_id = li.invoice_id
    JOIN general_ledger_accounts gl
        ON li.account_number = gl.account_number
WHERE invoice_total - payment_total - credit_total > 0
ORDER BY vendor_name, line_item_amount DESC
```

**Description**

- You can think of a multi-table join as a series of two-table joins proceeding from left to right.

### How to work with outer joins

When you use a left outer join, the result set includes all the rows from the first, or left table. Similarly, when you use a right outer join the result set includes all the rows from the second, or right table.

```sql
SELECT select_list
FROM table_1
    {LEFT|RIGHT} JOIN table_2
        ON join_condition_1
    {LEFT|RIGHT} JOIN table_3
        ON join_condition_2
```

**Description**

- An outer join retrieves all rows that satisfy the join condition, plus unmatched rows in the left or right table
- When a row with unmatched columns is retrieved, any columns from the other table that are included i the result set are given null values.

### Other skills for working with joins

**How to join tables with the USING keyword**

```sql
SELECT department_name, last_name, project_number
FROM departments
    JOIN employees USING (department_number)
    LEFT JOIN projects USING (employee_id)
ORDER BY department_name
```

- You can use the USING keyword to simplify the syntax for joining tables 
- The join can be an inner join or an outer join
- The tables must be joined by a column that has the same name in both tables
- To include multiple columns, separate them with commas 
- The join must be an equijoin, which means that the equals operator is used to compare the two columns

**How to join tables using the NATURAL keyword**

When you code a natural join you don't specify the column that's used to join the two tables. Instead, the database automatically joins the two tables based on all columns in the two tables that have the same name.

```sql
SELECT select_list
FROM table_1
    NATURAL JOIN table_2
    [NATURAL JOIN table_3]...
```

- You can use the NATURAL keyword to create a natural join that joins two tables based on all columns in the two tables that have the same name.
- Although the code for a natural join is shorter than the code for joins that use the ON or USING clause, a natural join only works correctly for certain types of database structures.

**How to use cross joins**

A cross join joins each row from the first table with each row from the second table.

**How to work with unions**

Like a join, a union combines data from two or more tables. Instead of combining columns from base tables, however a union combines rows from two or more result sets.

- A union combines the result sets of two or more SELECT statements into one result set.
- Each result set must return the same number of columns, and the corresponding columns in each result set must have compatible data types.
- By default, a union eliminates duplicate rows. If you want to include duplicate rows, code the ALL keyword
- The column names in the final result set are taken from the first SELECT clause. Column aliases assigned by the other SELECT clauses have no effect on the final result.
- To sort the rows in the final result set, code an ORDER BY clause after the last SELECT statement. This clause must refer to the column names assigned in the first SELECT clause.

**A union that simulates a full outer join**

A full outer join returns unmatched rows from both the left and righ tables. Although MySQL doesn't provide language for coding a full outer join, you can simlate a full outer join by coding a union that combines the result sets for a left outer join and a right outer join.

```sql
SELECT name
FROM a
    LEFT JOIN b
    ON a.id = b.id
UNION
SELECT name
FROM a
    RIGHT JOIN b
    ON a.id = b.id
ORDER BY name
```
- When you use a full outer join, the result set includes all the rows from both tables
- MySQL doesn't provide language keywords for full outer joins, but you can simulate a full outer join by using the UNION keyword to combine the result sets from a left outer join and a right outer join.

## 5. How to insert, update and delete data

### 5.1. How to create a copy of a table

```sql
CREATE TABLE invoices_copy AS
SELECT id
FROM invoices
```
When you use this technique to create tables, MySQL only copies the column definitions and data. In other words, MySQL doesn't retain other parts of the column definitions such as pk, fk and indexes.

**Description**

- You can use the CREATE TABLE AS statement to create a new table based on the result set defined by a SELECT statement
- Each column name in the SELECT clause must be unique. If you use caculated values in the select list, you must name the column
- You can code the other clauses of the SELECT statement just as you would for any other SELECT statement, including grouping, aggregates, joins, and subqueries.
- When you use the CREATE TABLE AS statement to create a table, only the column definitions and data are copied. Definitions of primary keys, foreign keys, indexes, and so on are not included in the new table.

### 5.2. How to insert new rows

**Description**

- You use the INSERT statement to add one or more rows to a table
- To insert a null value into a column, you can use the NULL keyword. To insert a default value or to have MySQL generate a value for an auto increment column you can use DEFAULT keyword.
- If you include a column list, you can omit columns with default values and null values. Then, the default value or null value is assigned automatically. You can also omit an auto increment column.

### 5.3. How to use a subquery in an INSERT statement

```sql
INSERT INTO invoice_archive
SELECT *
FROM invoices
WHERE invoice_total - payment_total - credit_total = 0
```

**Description**

- A subquery is a SELECT statement that's coded within another SQL statement
- To insert rows selected from one or more tables into another table, you can code a subquery in place of the VALUES clause. Then MySQL inserts the rows returned by the subquery into the target table. For this to work, the target table must already exist.

