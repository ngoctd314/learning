# Stored Program Development

## 13. Language skills for writing stored programs

**Four types of stored programs**

|Type|Description|
|-|-|
|Stored procedure|Can be called from an application that has access to the database|
|Stored function|Can be called from a SQL statement|
|Trigger|Is executed in response to an INSERT, UPDATE, or DELETE statement on a specified table|
|Event|Is executed at a scheduled time|

### 13.1. An introduction to stored programs

#### 13.1.1. A script that creates and calls a stored procedure

```sql
-- This script creates a stored procedure named test that doesn't accept any parameters.

-- USE statement, which selects the AP database
USE ap;

DROP PROCEDURE IF EXISTS test;

-- Change statement delimiter from semicolon to double front slash
DELIMITER //

-- The CREATE PROCEDURE statement creates the procedure. To indicate that this procedure doesn't accept any 
-- parameter
CREATE PROCEDURE test()
-- The code within the CREATE PROCEDURE statement is defined by a block of code that begins with
-- the BEGIN keyword and ends with the END keyword.
BEGIN
    -- The declare statement defines a variable 
    DECLARE sum_balance_due_var DECIMAL(9,2);
    SELECT SUM(invoice_total - payment_total - credit_total)
    INTO sum_balance_due_var
    FROM invoices
    WHERE vendor_id = 95;

    IF sum_balance_due_var > 0 THEN
        SELECT CONCAT('Balance due: $', sum_balance_due_var) AS message;
    ELSE
        SELECT 'Balance paid in full' AS message;
    END IF;
END //

-- Change statement delimiter from double front slash to semicolon
DELIMITER ;

CALL test();
```

**Description**

- A stored program consists of one or more SQL statements stored in the database for later use.
- Within a stored program, you can write procedural code that controls the flow of execution. That includes if/else constructs, loops, and error-handling code.

**A summary of statements for coding stored programs**

After the SQL statements for writing procedural code, this figure presents one SQL statement that you're already familiar with that's commonly used within stored programs: the SELECT statement.

### 13.2. How to write procedural code

#### 13.2.1. How to declare and set variables

- A variable stores a value that can change as a stored program executes.
- A variable must have a name that's different from the names of any columns used in any SELECT statement within the stored program.

#### 13.2.2. How to use a cursor *

By default, SQL statements work with entire result set rather than individual rows. However, you may sometimes need to work with the data in a result set one row at a time.

## 14. How to use transactions and locking

What happens when two users try to update the same data at the same time?

### 14.1. How to work with transactions

A transaction is a gropu of SQL statements taht you combine into a single logical unit of work.

Before you being using MySQL to work with transactions, you should relize that some storage engines don't support transactions. As a result, the skills presented in this topic only apply to storage engines such as InnoDB that support transactions. 

#### 14.1.1. How to commit and rollback transactions

By default, a MySQL session uses autocommit mode, which automatically commits INSERT, UPDATE, and DELETE statements immediately after you execute them.

**Description**

- By default, MySQL runs in autocommit mode, which automatically commits changes in the database immediately after each INSERT, UPDATE or DELETE statement is executed.

#### 14.1.2. How to work with save points

A SAVEPOINT statement is used to identify a save point before each of the three INSERT statements that are included in the script. As a result, the script includes three save points.

Use the ROLLBACK TO SAVEPOINT statement to roll back all or part of a transaction.

```sql
USE ap;

START TRANSACTION;

SAVEPOINT before_invoice;

INSERT INTO invoices
VALUES (115, 34, 'ZXA-080', '2015-01-18', 14092.59, 0, 0, 3, '2015-04-18', NULL);

SAVEPOINT before_line_item1;

INSERT INTO invoice_line_items
VALUES (115, 1, 160, 4447.23, 'HW upgrade');

SAVEPOINT before_line_item2;

INSERT INTO invoice_line_items
VALUES (115, 2, 167, 9645.36, 'OS upgrade');

ROLLBACK TO SAVEPOINT before_line_item2;

ROLLBACK TO SAVEPOINT before_line_item1;

ROLLBACK TO SAVEPOINT before_invoice;

COMMIT; -- doesn't commit any changes to the database.
```

**Description**

- When you use save points, you can roll back a transaction to the beginning of the transaction or to a particular save point.
- You can use the SAVEPOINT statement to create a save point with the specified name
- You can use the ROLLBACK TO SAVEPOINT statement to roll back a transaction to the specified save point
- Save points are useful when a single transaction contains so may SQL statements that rolling back the entire transaction would be inefficient.

### 14.2. How to work with concurrency and locking

#### 14.2.1. How concurrency and locking are related

**Description**

- Concurrency is the ability of a system to support two or more transactions working with the same data at the same time
- MySQL can automatically prevent some concurrency problems by using locks. A lock stops the execution of another transaction if it conflicts with a transaction that is already running.
- Concurrency is a problem only when the data is being modified. When two or more SELECT statments read the same data, the SELECT statements dont' affect each other.

#### 14.2.2. The four concurrency problems that locks can prevent


## 15. How to create stored procedures and functions

### 