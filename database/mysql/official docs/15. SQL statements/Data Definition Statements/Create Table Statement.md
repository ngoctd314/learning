# Create Table Statement

MySQL supports foreign keys, which permit cross-referencing related data across tables, and foreign key constraints, which help keep the related data consistent.

A foreign key relationship involves a parent table that holds the initial colume values, and a child table with column values that reference the parent column values that reference the parent colume values. A foreign key constraints is defined on the child table.

The essential syntax for a defining a foreign key constraint in a CREATE TABLE or ALTER TABLE statement includes the following:

## Identifiers

Foreign key constraint naming is governed by the following rules:

- The `CONSTRAINT symbol` value is used, if defined.
- If the `CONSTRAINT symbol` clause is not defined, or a symbol is not included following the CONSTRAINT keywork a constraints name name is generated automatically.

Prior to MySQL 8.0.16, if the CONSTRAINT symbol clause was not defined, or a symbol was not included following the CONSTRAINT keyword, both InnoDB and NDB storage engines would use the FOREIGN_KEY index_name if defined. In MySQL 8.0.16 and higher, the FOREIGN_KEY index_name is ignored.

- The CONSTRAINT symbol value, if defined, must be unique in the database. A duplicate symbol results in an error similar to: ERROR 1005 (HY000): Can't create table 'test.fk1' (errno: 121).

- NDB Cluster stores foreign names using the same lettercase with which they are created. Pior to version 8.0.20, when processing *SELECT* and other SQL statements, NDB compared the names of foreign keys in such statements with the names as stored in a case-sensitive fashion.  

## Conditions and Restrictions

Foreign key constraints are subject to the following conditions and restrictions:

- Parent and child tables must use the same storage engine, and they cannot be defined as temporary tables.
- Creating a foreign key constraint requires the *REFERENCES* privilege on the parent table.
- Corresponding colums in the foreign key and the referenced key must have similar data types. The size and sign of fixed precision types such as INTEGER and DECIMAL must be the same. The length of string types need not be the same. For nonbinary (character) string columns, the character set and collation must be the same.
- MySQL supports foreign key references between on column and another within a table. (A column cannot have a foreign key reference to itself). In these cases, a "child table record" refers to a dependent record within the same table.
- MySQL requires indexes on foreign keys and references keys so that foreign key checks can be fast and not required table scan. In the referencing table, there must be an index where the foreign key columns are listed as the first columns in the same order. Such as index is created on the referencing table automatically if it does not exist. This index might be silently dropped later if you create another index that can be used to enforce the foreign key constraint. index_name, if given, is used as described previously.
- **InnoDB** permits a foreign key to reference any index column or group of columns. However, in the referenced table, there must be an index where the referenced colums are the first colums in the same order.
- Index prefixes on foreign key columns are not supported. Consequently, BLOB and TEXT columns cannot be included in a foreign key because indexes  those columns must always include a prefix length.
- A table in a foreign key relationship cannot be altered to use another storage engine. To change the storage engine, you must drop any foreign key constraints first.
- A foreign key constraint cannot reference a virtual generated column.

## Referential Actions

When an **UPDATE** or **DELETE** operation affects a key value in the parent table that has matching rows in the child table, the result depends on the referential action specified by `ON UPDATE` and `ON DELETE` subclauses of the `FOREIGN KEY` clause. Referential actions include:

- `CASCADE`: Delete or update the row from the parent table and automatically delete or update the matching rows in the child table. Both ON DELETE CASCADE and ON UPDATE CASCADE are supported. Between two tables, do not define several `ON UPDATE CASCADE` clauses that act on the same column in the parent table or in the child table.

If a FOREIGN KEY clause is defined on both tables in a foreign key relationship, making both tables a parent and child, an `ON UPDATE CASCADE` or `ON DELETE CASCADE` subclause defined for one FOREIGN KEY clause.

- `SET NULL`: Delete or update the row from the parent table and set the foreign key column or columns in the child table to `NULL`. Both `ON UPDATE CASCADE` or `ON DELETE CASCADE` clauses are supported.

If you specify a SET NULL action, make sure that you have not declared the columns in the child table as NOT NULL.  

- `RESTRICT`: Rejects the delete or update operation for the parent table. Specifying RESTRICT (or NO ACTION) is the same as omitting the `ON DELETE` or `ON UPDATE` clause.

- `NO ACTION`: A keyword from standard SQL. For InnoDB, this is equivalent to `RESTRICT`; the delete or update operation for the parent table is immediately rejected if there is a related foreign key value in the referenced table.

- `SET DEFAULT`: This action is recognized by the MySQL parser, but both InnoDB and NDB reject table definitions containing `ON DELETE SET DEFAULT` or `ON UPDATE SET DEFAULT` clauses.

For storage engines that support foreign keys, MySQL rejects any INSERT or UPDATE operation that attempts to create a foreign key value in a child table if there is no matching candidate key value in the parent table.

For an `ON DELETE` or `ON UPDATE` that is not specified, the default action is always `NO ACTION`

As the default, an `ON DELETE NO ACION` or `ON UPDATE NO ACTION` clause that is specified explicitly does not appear in `SHOW CREATE TABLE` output or in tables dumped with mysqldump. `RESTRICT`, which is an equivalent non-default keyword, appears in `SHOW CREATE TABLE` output and in tables dumped with mysqldump. `RESTRICT`, which is an equivalent non-default keyword, appears in `SHOW CREATE TABLE` output and in tables dumped with mysqldump.

## Foreign Key Constraints Examples

```sql
CREATE TABLE parent (
    id INT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE child (
    id INT,
    parent_id INT,
    INDEX par_ind (parent_id),
    FOREIGN KEY (parent_id) REFERENCES parent(id) ON DELETE CASCADE
) ENGINE=INNODB;
```

This is a more complex example in which a product_order table has foreign keys for two other tables. One foreign key references a two-colum index in the product table. The other references a single-column index in the customer table:

```sql
CREATE TABLE product (
    category INT NOT NULL, id INT NOT NULL,
    price DECIMAL,
    PRIMARY KEY (category, id)
) ENGINE=INNODB;


CREATE TABLE customer (
    id INT NOT NULL,
    PRIMARY KEY (id),
) ENGINE=INNODB;

CREATE TABLE product_order (
    no INT NOT NULL AUTO_INCREMENT,
    product_category INT NOT NULL,
    product_id INT NOT NULL,
    customer_id INT NOT NULL,

    PRIMARY KEY (no),
    INDEX (product_category, product_id),
    INDEX (customer_id),

    FOREIGN KEY (product_category, product_id)
        REFERENCES product (category, id)
        ON UPDATE CASCADE ON DELETE RESTRICT,

    FOREIGN KEY (customer_id)
        REFERENCES customer (id)
) ENGINE=INNODB;
```

## Adding Foreign Key constraints

## Dropping Foreign Key Constraints

## Foreign Key Checks

## Locking

## Foreign Key Definitions and Metadata

## Foreign Key Definitions and Metadata

## Foreign Key Errors
