# Index merge optimization

The Index Merge access method retrieves rows with multiple range scans and merges their results into one. This access method merges index scans from a single table only, not scans across multiple tables. The merge can produce unions, intersections, or unions-of-intersections of its underlying scans.

## Index Merge Intersection Access Algorithm

This access algorithm is applicable when a `WHERE` clause is converted to several range conditions on different keys combined with `AND`, and each condition is one of the following:

- An N-part expression of this form, where the index has exactly N parts (that is, all index parts are covered):

```sql
key_part1 = const1 AND key_part2 = const2 ... AND key_partN = constN
```

- And range condition over the primary key of an InnoDB table.

## Index Merge Union Access Algorithm

The criteria for this algorithm are similar to those for the Index Merge intersection algorithm. The algorithm is applicable when the table's `WHERE` clause is converted to several range conditions on different keys combined with OR, and each condition is one of the following:

- An N-part expression of this form, where the index has exactly N parts (that is, all index parts are covered):

```sql
key_part1 = const1 OR key_part2 = const2 ... OR key_partN = constN
```

- Any range condition over a primary key of an InnoDB table.

## Index Merge Sort-Union Access Algorithm

This access algorithm is applicable when the `WHERE` clause is converted to several range conditions combined by OR, but the Index Merge union algorithm is not applicable.

```sql
SELECT * FROM tbl_name
    WHERE key_col1 < 10 OR key_col2 < 20;

SELECT * FROM tbl_name
    WHERE (key_col1 > 10 OR key_col2 = 20) AND nonkey_col = 30;
```
