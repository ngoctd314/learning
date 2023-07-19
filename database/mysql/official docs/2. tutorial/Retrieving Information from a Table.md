# Retrieving Information from a Table

## Working with NULL values

The NULL value can be surprising until you get used to it. Conceptually, NULL means "a missing unknown value" and it is treated somewhat differently from other values.

To test the NULL, use the IS NULL and IS NOT NULL operators:

```sql
SELECT 1 IS NULL, 1 IS NOT NULL;
```

You cannot use arithmetic comparison operators such as =, <=, or <> to test for NULL.

```sql
SELECT 1 = NULL, 1 <> NULL, 1 < NULL, 1 > NULL;

```
```txt
+----------+----------+----------+-----------+-----------+-----------+
| 1 = null | 1 > null | 1 < null | 1 >= null | 1 <= null | 1 <> null |
+----------+----------+----------+-----------+-----------+-----------+
| <null>   | <null>   | <null>   | <null>    | <null>    | <null>    |
+----------+----------+----------+-----------+-----------+-----------+
```

Because the result of any arithmetic comparison with NULL is also NULL, you cannot obtain any meaningful results from comparisons.

In MySQL, 0 or NULL mean false and anything else means true. The default truth value from a boolean operation is 1.

Two NULL values are regarded as equal in GROUP BY.

When doing an ORDER BY, NULL values are presented first if you do ORDER BY ... ASC and last if you do ORDER BY ... DESC.

A common error when working with NULL is to assume that it is not possible to insert a zero or an empty string into a column defined as NOT NULL, but this is not the case. There are in fact values, whereas NULL means "not having a value". You can test this easily enough by using IS [NOT] NULL as shown.

## Pattern matching

MySQL provides standard SQL pattern matching as well as a form of pattern matching based on extended regular expression similar to those used by Unix utilities such as vi, grep, and sed.

SQL pattern matching enables you to use _ to match any single character and % to match an arbitrary number of characters. Do not use = or <> when you use SQL patterns. Use LIKE or NOT LIKE comparison operators instead.

```sql
SELECT * FROM tbl WHERE name like '%pattern%';
SELECT * FROM tbl WHERE name like '_____';
SELECT * FROM tbl WHERE REGEXP_LIKE(name, 'pattern')
```

## Counting Rows

`COUNT(*)` counts the number of rows


- If the ONLY_FULL_GROUP_BY SQL mode is enabled, an error occurs:

```text
mysql> SET sql_mode = 'ONLY_FULL_GROUP_BY';
mysql> SELECT owner, COUNT(*) FROM pet;
ERROR 1140, "In aggregated query without GROUP BY, expression #1 expression #1 of SELECT list contains non aggregated column 'db.test_nulls.name'; this is incompatible with sql_mode=only_full_group_by"
```

- If ONLY_FULL_GROUP_BY is not enabled, the query is processed by treating all rows as a single group, but the value selected for each named column is nondeterministic. The server is free to select the value from any row:

```txt
mysql> SET sql_mode = '';
mysql> SELECT owner, COUNT(*) FROM pet;

+-------+----------+
| name  | count(*) |
+-------+----------+
| random| 7        |
+-------+----------+
```
