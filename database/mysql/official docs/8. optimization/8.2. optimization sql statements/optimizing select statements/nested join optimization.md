# Nested Join Optimization

The syntax for expressing joins permits nested joins. The following discussion refer to the join syntax.

The syntax of **table_factor** is extended in comparision with the SQL standard. The latter accepts only table_reference, not a list of them inside a pair of parentheses. This is a conservative extension if we consider each comma in a list of table_reference items as equivalent to an inner join.

```sql
SELECT * FROM t1 LEFT JOIN (t2, t3, t4)
    ON (t2.a=t1.a AND t3.b=t1.b AND t4.c=t1.c);
```

Is equivalent to:

```sql
SELECT * FROM t1 LEFT JOIN (t2 CROSS JOIN t3 CROSS JOIN t4)
    ON (t2.a=t1.a AND t3.b=t1.b AND t4.c=t1.c);
```

In MySQL, CROSS JOIN is syntactically equivalent to INNER JOIN; they can replace each other. In standard SQL, they are not equivalent. INNER JOIN is used with an ON clause; CROSS JOIN is used otherwise.

In general, parentheses can be ignored in join expressions containing only inner join operations. Consider this join expression:

```sql
t1 LEFT JOIN (t2 LEFT JOIN t3 ON t2.b=t3.b OR t2.b IS NULL)
    ON t1.a=t2.a;
```
