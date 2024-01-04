# Index Condition Pushdown Optimization

Index Condition Pushdown (ICP) is an optimization for the case where MySQL retrieves rows from a table using an index. Without ICP, the storage engine traverses the index to locate rows in the base table and returns them to the MySQL server which evaluates the `WHERE` condition for the rows. With ICP enabled, and if parts of the `WHERE` condition for the rows. With ICP enabled, and if parts of the `WHERE` condition down to the storage engine. The storage engine then evaluates the pushed index condition by using the index entry and only if this is satisfied is the the row read from the table. ICP can reduce the number of times the storage engine must access the base table and the number of times MySQL server must access the storage engine.

Applicability of the Index Condition Pushdown optimization is subject to these conditions:


