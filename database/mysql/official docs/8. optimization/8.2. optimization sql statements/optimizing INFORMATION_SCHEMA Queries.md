# Optimizing INFORMATION_SCHEMA Queries

Applications that monitor databases may make frequent use of INFORMATION_SCHEMA tables. Certain types of queries for INFORMATION_SCHEMA tables can be optimized to execute more quickly. The goal is to minimize file operations (for example, scanning a directory or opening a table file) to collect the information that makes up these dynamic tables.

**1) Try to use constant lookup values for database and table names in the WHERE clause**

You can take advantage of this principle as follows:

- To lookup databases or tables, use expressions that evaluate to a constant, such as literal values, functions that return a constant, or scalar subqueries.
- Avoid queries that use a nonconstant database name lookup value (or no looup value) because they require a scan of the data directory to find matching database directory names.
- Within a database, avoid queries that use a nonconstant table name lookup value
