# Extended EXPLAIN Output Format

The EXPLAIN statement produces extra("extended") information that is not part of EXPLAIN output but can be viewed by issuing a SHOW WARNINGS statement following EXPLAIN. 

The Message value in SHOW WARNINGS output displays how the optimizer qualifies table and column names in the SELECT statement, what the SELECT looks like after the application of rewriting and optimization rules, and possibly other notes about the optimization process.

The extended information displayable with a SHOW WARNINGS statement following EXPLAIN is produced only or SELECT statements. SHOW WARNINGS displays empty result for other explainable statements (DELETE, INSERT, REPLACE, and UPDATE).

Here is an example extended EXPLAIN output:


