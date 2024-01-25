# Internal Temporary Table Use in MySQL

In some cases, the server creates internal temporary tables while processing statements. Users have no direct control over when this occurs.

The server creates temporary tables under conditions such as these:

- Evaluation of UNION statements, with some exceptions described later.
- Evaluation of some views such those that use the TEMPTABLE algorithm, UNION, or aggregation.
- Evaluation of derived tables. 
- Tables created for subquery or semijoin materialization.
- Evaluation of statements that contain an ORDER BY clause and a different GROUP BY clause, or for which the ORDER BY or GROUP BY contains columns from tables other than the first table in the join queue.
- Evaluation of DISTINCT combined with ORDER BY may require a temporary table.
- For queries that use the SQL_SMALL_RESULT modifier, MySQL uses an in-memory temporary table, unless the query also contains elements (described later) that require on-disk storage.
- Evaluation of multiple-table UPDATE statements.
- Evaluation of GROUP_CONCAT() or COUNT(DISTINCT) expressions.

Some query conditions prevent the use of an in-memory temporary table, in which case the server uses an on-disk table instead:


