# Function Call Optimization

MySQL functions are tagged internally as deterministic or nondeterministic. A function is nondeterministic if, given fixed values for its arguments, it can return different results for different invocations. Examples of nondeterministic functions: RAND(), UUID().

If a function is tagged nondeterministic, a reference to it in a WHERE clause is evaluated for every row (when selecting from one table) or combination of rows (when selecting from a multiple-table join).

Nondeterministic functions may affect query performance. For example, some optimizations may not be available, or more locking might be required. The following discussion uses RAND() but applies to other nondeterministic functions as well.

```sql
CREATE TABLE t (id INT NOT NULL PRIMARY KEY, col_a VARCHAR(100));
```

Consider these two queries:

```sql
SELECT * FROM t WHERE id = POW(1, 2);
SELECT * FROM t WHERE id = FLOOR(1 + RAND() * 49);
```

Both queries appear to use a primary key lookup because of the equality comparison against the primary key, but that is true only for the first of them:

- The first query always produces a maximum of on row because POW() with constant arguments is a constant value and is used for index lookup.
- The second query constans an expression that uses the nondeterministic function RAND(), which is not constant in the query but in fact has a new value for every row of table t. Consequently, the query reads every row of the table, evaludates the predicate for each row, and outputs all rows for which the primary key matches the random value. This might be zero, one, or multiple rows, depending on the id column values and the values in the RAND() sequence. 

```sql
EXPLAIN SELECT * FROM t WHERE id = POW(1,2);
+----+-------------+-------+------------+-------+---------------+---------+---------+-------+------+----------+--------+
| id | select_type | table | partitions | type  | possible_keys | key     | key_len | ref   | rows | filtered | Extra  |
+----+-------------+-------+------------+-------+---------------+---------+---------+-------+------+----------+--------+
| 1  | SIMPLE      | t     | <null>     | const | PRIMARY       | PRIMARY | 4       | const | 1    | 100.0    | <null> |
+----+-------------+-------+------------+-------+---------------+---------+---------+-------+------+----------+--------+

EXPLAIN SELECT * FROM t WHERE id = FLOOR(1 + RAND() * 4);

+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
| id | select_type | table | partitions | type | possible_keys | key    | key_len | ref    | rows | filtered | Extra       |
+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
| 1  | SIMPLE      | t     | <null>     | ALL  | <null>        | <null> | <null>  | <null> | 4    | 25.0     | Using where |
+----+-------------+-------+------------+------+---------------+--------+---------+--------+------+----------+-------------+
```

The effects of nondeterministic are not limited to SELECT statements. This UPDATE statement uses a nondeterministic function to select rows to be modified:

```sql
UPDATE t SET col_a = some_expr WHERE id = FLOOR(1 + RAND() * 49);
```

The behavior just described has implications for performance and replication:

- Because a nondeterministic function does not produce a constant value, the optimizer cannot use strategies that might otherwise be applicable, such as index lookups. The result may be a table scan.
- InnoDB might escalate to a range-key rock rather than taking a single row lock for one matching row.
- Updates that do not execute deterministically are unsafe for replication.

The difficulties stem from the fact that the RAND() function is evaluated once for every row of the table. To avoid multiple function evaluations, use one of these techniques:

- Move the expression containing the nondeterministic function to separate statement, saving the value in a variable. In the original statement, replace the expression with a reference to the variable, which the optimizer can treat as a constant value:

```sql
SET @keyval = FLOOR(1 + RAND() * 49);
UPDATE t SET col_a = some_expr WHERE id = @keyval
```

As mentioned previously, a nondeterministic expression in the WHERE clause might prevent optimizations and result in a table scan. However, it may be possible to partially optimize the WHERE clause if other expressions are deterministic.

```sql
SELECT * FROM t WHERE partial_key=5 AND some_column=RAND();
```

If the optimizer can use partial_key to reduce the set of rows selected, RAND() is executed fewer times, which diminishes the effect of nondeterministic on optimization.
