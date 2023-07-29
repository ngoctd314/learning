# More SQL skills as you need them

## 6. How to code summary queries

How to use a special type of function called an aggregate function. Aggregate functions allow you to do jobs like calculate averages, summarize totals, or find the highest value for a given column, and you'll use them in summary queries.


### 6.1. How to work with aggregate functions

aggregate functions, which operation on a series of values and return a single summary value. Because aggregate function typically operate on the values in columns, they are sometimes referred to as column functions. A query that contains one or more aggregate functions is typically referred to as a summary query.

|Function syntax|Result|
|-|-|
|AVG([ALL|DISTINCT] expression)|The average of the non-null values in the expression|
|SUM([ALL|DISTINCT] expression)|The total of the non-null values in the expression|
|MIN, MAX([ALL|DISTINCT] expression)|The lowest non-null value in the expression|
|COUNT([ALL|DISTINCT] expression)|The number of non-null values in the expression|
|COUNT(*)|The number of rows selected by the query|

**Description**

- Aggregate functions, also called column functions, perform a calculation on the values in a set of selected rows.
- A summary query is a SELECT statement that includes one or more aggregate functions
- The expression you specify the the AVG and SUM functions must result in a numeric value. The expression for the MIN, MAX, and COUNT functions can result in a numeric, date, or string value.
- All aggregate functions except for COUNT(*) ignore null values

### 6.2. Queries that use aggregate functions

**Description**

- To count all of the selected rows, you typically use the COUNT(*) function. Alternatively, you can use the COUNT function with the name of any column that can't contain null values.
  

### 6.3. How to group and summarize data

#### 6.3.1. How to code the GROUP BY and HAVING clauses

The GROUP BY clause determines how the selected rows are grouped, and the HAVING clause determines which groups are included in the final results. These clauses are coded after the WHERE clause but before the ORDER BY clause. Because the WHERE clause is applied before the rows are grouped, and the ORDER BY clause is applied after the rows are grouped.

**Description**

- The GROUP BY clause groups the rows of a result set based on one or more columns or expressions.
- If you include aggregate function in the SELECT clause, the aggregate is calculated for each group specified by the GROUP BY clause.
- The HAVING clause specifies a search condition for a group or an aggregate. MySQL applies this condition after it groups the rows that satisfy the search condition in the WHERE clause.
- When a SELECT statement includes a GROUP BY clause, the SELECT clause can include the columns used for grouping, aggregate functions, and expressions that result in a constant value.

#### 6.3.2. Queries that use the GROUP BY and HAVING clauses

**Description**

- With MySQL 8.0.12 and earlier, the GROUP BY clause sorted the columns in ascending sequence by default. Then, to change the sort sequence, you could code the DESC keyword after the column name in the GROUP BY clause. In addition, to get your results faster, you could code an ORDER BY NULL clause to prevent MySQL from sorting the rows in the GROUP BY clause.
- With MySQL 8.0.13 and later, the columns in a GROUP BY clause

#### 6.3.3. How the HAVING clause compares to the WHERE clause

You can limit the groups included in a result set by coding a search condition in the HAVING clause. In addition, you can apply a search condition to each row before it's included in a group. To do that, you code the search condition in the WHERE clause just as you would for any SELECT statement.

```sql
SELECT vendor_name,
    COUNT(*) AS invoice_qty,
    ROUND(AVG(invoice_total), 2) AS invoice_avg
FROM vendors JOIN invoices
    ON vendors.vendor_id = invoices.vendor_id
GROUP BY vendor_name
HAVING AVG(invoice_total) > 500
ORDER BY invoice_qty DESC;
```
This example the invoices in the Invoices table by vendor name and calculates a count and average invoice amount for each group. Then, the HAVING clause limits the groups in the result set to those that have an average invoice total greater than $500.

```sql
SELECT vendor_name,
    COUNT(*) AS invoice_qty,
    ROUND (AVG(invoice_total), 2) AS invoice_avg
FROM vendors JOIN invoices
    ON vendors.vendor_id = invoices.vendor_id
WHERE invoice_total > 500
GROUP BY vendor_name
ORDER BY invoice_qty DESC
```
This example the WHERE clause limits the invoices included in the groups to those that have an invoice total greeter that $500. In other words, the search condition in this example is applied to every row.

HAVING clause can include aggregate functions as shown in the first example, but the WHERE clause can't. That's because the search condition in a WHERE clause is applied before the rows are grouped. Second, although the WHERE clause can refer to any column in the base tables, the HAVING clause can only refer to columns included in the SELECT clause. HAVING doesn't filter the base tables.

**Description**

- When you include a WHERE clause in a SELECT statement that uses grouping and aggregates, MySQL applies the search condition before it groups the rows and calculates the aggregates.
- When you include a HAVING clause in a SELECT statement that uses grouping and aggregates, MySQL applices the search condition after it groups the rows and calculates the aggregates.
- A WHERE clause can refer to any column in the base tables.
- A HAVING clause can only refer to a column included  in the SELECT clause.
- A WHERE clause can't contain aggregate functions
- A HAVING clasue can contain aggregate functions

#### 6.3.4. How to code aggregate window functions

##### 6.3.4.1. How the aggregate window functions work


## 7. How to code subqueries

Subqueries allow you to build queries that would be difficult or impossible to build otherwise.
