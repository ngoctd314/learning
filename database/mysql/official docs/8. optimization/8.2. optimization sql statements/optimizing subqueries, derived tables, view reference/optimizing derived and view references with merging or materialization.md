# optimizing derived tables and view references with merging or materialization

The optimizer can handle derived table references using two strategies

- Merge the derived table into the outer query block.
- Materialize the derived table to an internal temporary table.

```sql
SELECT * FROM (SELECT * FROM t1) AS derived_t1;
```

With merging of the derived, that query is executed similar to 

```sql
SELECT * FROM t1;
```

Example 2:

```sql
SELECT *
    FROM t1 JOIN (SELECT t2.f1 FROM t2) AS derived_t2 ON t1.f2 = derived_t2.f1
    WHERE t1.f1 > 0;
```

With merging of the derived table derived_t2, that query is executed similar to:

```sql
SELECT t1.*, t2.f1
FROM t1 JOIN t2 ON t1.f2 = t2.f1
WHERE t1.f1 > 0 
```
