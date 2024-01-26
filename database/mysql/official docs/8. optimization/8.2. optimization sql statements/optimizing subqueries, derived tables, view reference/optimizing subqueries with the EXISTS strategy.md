# Optimizing Subqueries with the EXISTS Strategy

Certain optimizations are applicable to comparisons that use the IN (or=ANY) operator to test subquery results. This section discusses these optimizations, particularly with regard to the challenges that NULL values present. The last part of the discussion suggests how you can help the optimizer:

Consider the following subquery comparison:

```sql
outer_expr IN (SELECT inner_expr FROM ... WHERE subquery_where)
```

MySQL evaluates queries "from outside to inside". That is, it first obtains the value of the outer expression outer_expr, and then runs the subquery and captures the rows that it produces.

A very useful optimization is to "inform" the subquery that the only rows of interest are those where the inner expression inner_expr is equal to outer_expr. This is done by pushing down an appropriate equality into the subquery's WHERE clause to make it more restrictive. The converted comparison looks like this:

```sql
EXISTS (SELECT 1 FROM ... WHERE subquery_where AND outer_expr=inner_expr)
```

After the conversion, MySQL can use the pushed-down equality to limit the number of rows it must examine to evaluate the subquery.

```sql
(oe_1, ..., oe_N) IN
    (SELECT ie_1, ..., ie_N FROM ... WHERE subquery_where)
```

Becomes

```sql
EXISTS (SELECT 1 FROM ... WHERE subquery_where
                          AND oe_1 = ie_1
                          AND ...
                          AND oe_N = ie_N)
```

For simplicity, the following discussion assumes a single pair of outer and inner expression values.

The conversion just described has its limitations. It is valid only if we ignore possible NULL values. That is, the "pushdown" strategy works as long as of these conditions are true:

- outer_expr and inner_expr cannot be NULL.
- You need not distinguish NULL from FALSE subquery results. If the subquery is a part of an OR and AND expression in the WHERE clause, MySQL assumes that you do not care.

