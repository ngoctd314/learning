# Optimizing Subqueries with the EXISTS Strategy

Certain optimizations are applicable to comparisons that use the IN (or=ANY) operator to test subquery results. This section discuss these optimizations, particularly with regard to the challenges that NULL values present. The last part of the discussion suggests how you can help the optimizer:

Consider the following subquery comparison:

```sql
outer_expr IN (SELECT inner_expr FROM ... WHERE subquery_where)
```

MySQL evaluates queries "from outside to inside". That is, it first obtains the value of the outer expression outer_expr, and then runs the subquery and captures the rows that it produces.

A very useful optimization is to "inform" the subquery that the only rows of interest are those where the inner expression inner_expr is equal to outer_expr. This is done by pushing down an appropriate equality into the subquery's WHERE clause to make it more restrictive.
