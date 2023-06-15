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