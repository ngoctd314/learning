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

|Problem|Description|
|-|-|
|Lost updates|Occur when two transactions select the same row and then update the row based on the values originally selected. Since each transaction is unaware of the other, the later update overwrites the earilier update.|
|Dirty reads|Occur when a transaction selects data that hasn't ben commited by another transaction. For example, transaction A changes a row. Transaction B then selects the changed row before transaction. A commits the change. If transaction A then rolls back the change, transaction B has selected data that doesn't exist in the database|
|Nonrepeatable reads|Occur when two SELECT statements that try to get the same data get different values because another transaction has updated the data in the time between the two statements. For example, transaction A selects a row, transaction B then updates the row. When transaction A selects the same row again, the data is different.|
|Phantom reads|Occur when you perform an update or delete on a set of rows at the same time that another transaction is performing an insert or delete that affects one or more rows in that same set of rows.|

**Description**

- In a large system with many users, you should expect for these kinds of problems to occur.

#### 14.2.3. How to set the transaction isolation level

The simplest way to prevent concurrency problems is to change the default locking behavior. To do that, you use the SET TRANSACTION ISOLATION LEVEL statement to set the transaction isolation level.

If you use the SERIALIZABLE option all four concurrency problems will be prevented.

When you set the isolation level to SERIALIZABLE, each transaction is completely isolated from every other transaction an concurrency is severely restricted. The server does this by locking each resource, preventing other transactinos from accessing it. Since each transaction must wait for the previous transactino to commit, the transactions are executed serially, one after another.

Since the SERIALIZABLE level eliminates all concurrency problems, you may think this is always the best option. However, this option requires more overhead to manage all of the locks, so the access time for each transaction is increased. For some systems, this may cause significant performance problems. As a result, you typically want to use the SERIALIZABLE isolation level only for situations in which phantom reads aren't acceptable.


## 15. How to create stored procedures and functions

## 16. How to create triggers and events

Triggers can be executed before or after and INSERT, UPDATE, or DELETE (write operator) statment is executed on a table. As a result, they provide a powerful way to enforce data consistency, log changes to the database, and implement business rules. Events can be executed at a scheduled time. As a result, they provide a convenient way to automatically perform any task that needs to be run regularly such as scheduled maintenance of tables.

### 16.1. How to work with triggers

A trigger is named block of code that is executed, or fired, automatically when a particular type of SQL statement is executed. When using MySQL, a trigger is fired when an INSERT, UPDATE, or DELETE statement is executed on a table.


#### 16.1.1. How to create a Before trigger

```sql
CREATE TRIGGER trigger_name
    {BEFORE|AFTER} {INSERT|UPDATE|DELETE} ON trigger_name
    FOR EACH ROW 
    sql_block
```

```sql
DELIMITER //
CREATE TRIGGER vendors_before_update
    BEFORE UPDATE ON vendors
    FOR EACH ROW
BEGIN
    SET NEW.vendor_state = UPPER(NEW.vendor_state);
END//
```

**Description**
- A trigger is named block of code that executes, or fires, in response to an INSERT, UPDATE, or DELETE statement.
- You can fire a trigger before or after an INSERT, UPDATE, or DELETE statement is executed on a table.
- You must specify a FOR EACH ROW clause. This creates a row-level trigger that fires once for each row that's modified
- You can use the OLD and NEW keywords to get and set the values for the columns that are stored in the old row and the new row

#### 16.1.1.2. How to use a trigger to enforce data consistency

Triggers are commonly used to enforce data consistency.

```sql
DELIMITER //
CREATE TABLE invoices_before_uppdate
    BEFORE UPDATE ON invoices
    FOR EACH ROW
BEGIN
    DECLARE sum_line_item_amount DECIMAL(9,2);
    SELECT SUM(line_item_amount)
    INTO sum_line_item_amount
    FROM invoice_line_items
    WHERE invoice_id = NEW.invoice_id;

    IF sum_line_item_amount != NEW.invoice_total THEN
        SIGNAL SQLSTATE 'HY000'
        SET MESSAGE_TEXT = 'Line item total must match invoice total.';
    END IF;
END //
```

### 16.2. How to work with events

An event, or scheduled event, is a named block of code that executes, or fires, according to the event scheduler. By default, the event scheduler is off.

**Description**

- An event, or scheduled event, is a named block of code that executes, or fires according to the event scheduler.