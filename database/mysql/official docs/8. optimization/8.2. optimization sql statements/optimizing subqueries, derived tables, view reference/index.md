# Optimizing Subqueries, Derived Tables, View References and Common Table Expressions

The MySQL query optimizer has different strategies available to evaluate subqueries:

**For a subquery used with an IN, = ANY, or EXISTS predicate, the optimizer has these choices:**

- Semijoin
- Materialization
- Exists strategy

**For a subquery used with a NOT IN, <> ALL or NOT EXISTS predicate, the optimizer has these choices**

- Materialization
- EXISTS strategy

For a derived table, the optimizer has these choices(which also apply to view references and common table expressions):

- Merge the derived table into the outer query block
- Materialize the derived table to an internal temporary table
