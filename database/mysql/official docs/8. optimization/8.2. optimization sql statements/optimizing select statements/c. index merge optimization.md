# Index merge optimization

The Index Merge access method retrieves rows with multiple range scans and merges their results into one. This access method merges index scans from a single table only, not scans across multiple tables. The merge can produce unions, intersections, or unions-of-intersections of its underlying scans.

Example queries for which Index Merge may be used:

```sql
SELECT * FROM tbl_name WHERE key1 = 10 OR key2 = 20;

SELECT * FROM tbl_name 
    WHERE (key1 = 10 OR key2 = 20) AND non_key = 30;

SELECT * FROM t1, t2
    WHERE (t1.key1 IN (1,2) OR t1.key2 LIKE 'value%') 
    AND t2.key1 = t1.some_col
```

**Note**

The Index Merge optimization algorithm has the following known limitations:

- If your query has a complex `WHERE` clause with deep AND/OR nesting and MySQL does not choose the optimal plan, try distributing terms using the following identify transformations:

```sql
(X AND Y) OR Z => (X OR Z) AND (Y OR Z)
(X OR Y) AND Z => (X AND  Z) OR (Y AND Z)
```

- Index Merge is not applicable  to full-text indexes. 

The Index Merge access method has several algorithms, which are displayed in the Extra a field of EXPLAIN output:

- Using intersect(...)
- Using union(...)
- Using sort_union(...)

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

Examples:

```sql
SELECT * FROM innodb_table
    WHERE primary_key < 10 AND key_col1 = 20;

SELECT * FROM tbl_name
    WHERE key1_part1 = 1 AND key1_part2 = 2 AND key2 = 2;
```

The Index Merge intersection algorithm performs simultaneous scans on all used indexes and produces the intersection of row sequences that it receives from the merged index scans.

If all columns used in the query are covered by the used indexes, full table rows are not retrived (EXPLAIN output using  Using index in Extra field in this case). Here is an example of such a query:

```sql
SELECT COUNT(*) FROM t1 WHERE key1 = 1 AND key2 = 1;
```

If the used indexes do not cover all columns  used in the query, full rows are retrieved only when the range conditions for all used keys are satisfied.

If one of the merged conditions is a condition over the primary key of an InnoDB table, it is not used for row retrieval, but is used to filter out rows retrieved using other conditions.

The Index Merge Intersection access algorithm is used in MySQL to efficiently retrieve rows that satisfy the conditions specified in a query involving multiple indexes. This algorithm is employed when the query involves multiple conditions, and each condition can be satisfied by a different index.

Here's a simplified explaination of how MySQL implements the Index Merge Intersection access algorithm:

**1. Index Scan:** For each condition in the WHERE clause that can be satisfied by an index, MySQL performs an index scan using that index. This means that is identifies and retrieves rows that match the condition based on the specified index.

**2. Intersection:** Once the index scans are complete for all conditions, MySQL computes the intersection of the row sets obtained from each index. The intersection includes only the rows that satisfy all the conditions specified in the WHERE clause.

**3. Result set:** The final result set consists of the rows that are present in the intersection set. These rows satify all the conditions specified in the query.

**4. Optimization:** MySQL's query optimizer decideds whether to use the Index Merge Intersection access algorithm based on factors such as index statistics, cardinality, and the cost of accessing each index.

## Index Merge Sort-Union Access Algorithm

The criteria for this algorithm are similar to those for the Index Merge intersection algorithm. The algorithm is applicable when the table's WHERE clause is converted to several range conditions on different keys combined with OR, and each condition is one of the following:

- An N-part expression of this form, where the index has exactly N parts (that is, all index parts are covered):  

```sql
key_part1 = const OR key_part2 = const2 ... OR key_partN = constN
```

- Any range condition over a primary key of an InnoDB table.

- A condition for with the Index Merge intersection algorithm is applicable.

## Index Merge Sort-Union Access Algorithm

This access algorithm is applicable when the `WHERE` clause is converted to several range conditions combination by OR, but the Index Merge union algorithm is not applicable.

Examples:

```sql
SELECT * FROM tbl_name
    WHERE key_col1 < 10 OR key_col2 < 20;

SELECT * FROM tbl_name
    WHERE (key_col1 > 10 OR key_col2 = 20) AND nonkey_col = 30;
```

The difference between the sort-union algorithm and the union algorithm is that the sort-union algorithm must first fetch row IDs for all rows and sort them before returning any rows.
