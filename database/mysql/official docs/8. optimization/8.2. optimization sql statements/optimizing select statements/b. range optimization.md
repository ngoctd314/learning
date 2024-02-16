# Range optimization

The range access method uses a single index to retrieve a subset of table rows that are contained within one or several index value intervals. It can be used for a single-part or multiple-part index.

## Range Access Method for Single-Part Indexes

For a single-part index, index value intervals can be conveniently represented by corresponding conditions in the `WHERE` clause, denoted as range conditions rather than "intervals".

The definition of a range condition for a single-part index is as follows:

- For both BTREE and HASH indexes, comparison of a key part with a constant value is a range condition when using the =, <=>, IN(), IS NULL, or IS NOT NULL operators.
- Additionally, for BTREE indexes, comparison of a key  part with a constant value is a range condition when using the >, <, >=, <=, BETWEEN, !=, or <> operators, or LIKE comparisons if the argument to LIKE is a constant string that does not start with a wildcard character.
- For all index types, multiple range conditions combined with OR or AND form a range condition.

"Constant value" in the preceding descriptions means one of the following:

- A constant from the query string
- A column of a const or system table from the same join
- The result of an uncorrelated subquery
- And expression composed entirely from subexpressions of the preceding types

Some nonconstant values may be converted to constants during the optimizer constant propagation phase.

MySQL tries to extract range conditions from the WHERE clause for each of the possible indexes. During the extraction process, conditions that cannot be used for constructing the range condition are dropped, conditions that produce overlapping ranges are combined, and conditions that produce empty ranges are removed.

Consider the following statement, where key1 is an indexed column and nonkey is not indexed:

In general (and as demonstrated by the preceding example), the condition used for a range scan is less restrictive than the `WHERE` clause. MySQL performs an additional check to filter out rows that satisfy the range condition but not the full WHERE clause.

The range condition extraction algorithm can handle nested AND/OR constructs of arbitrary depth, and its output does not depend on the order in which conditions appear in `WHERE` clause.

## Range Access Method for Multiple-Part Indexes

Range conditions on a multiple-part index are an extension of range conditions for a single-part index. A range condition on a multiple-part index restricts index rows to lie within one or several key tuple intervals. Key tuple intervals are defined over a set of key tuples, using ordering from the index.

## Equality Range Optimization of Many-Valued Comparisons

Consider these expressions, where col_name is an indexed column:

```sql
col_name IN (val1, ..., valN)
col_name = val1 OR ... OR col_name = valN
```

Each expression is true if col_name is equal to any of several values. These comparisons are equality range comparisons (where the "range" is a single value). The optimizer estimates the cost of reading qualifying rows for equality range comparison as follows:

- If there is a unique index on col_name, the row estimate for each range is 1 because at most one row can have the given value.
- Otherwise, any index on col_name is nonunique and the optimizer can estimate the row count for each range using dives into the index or index statistic.

The eq_range_index_dive_limit system variable enables you to configure the number of values at which the optimizer switches from one row estimation strategy to the other. To permit use of index dives for comparisons of up to N equality ranges, set eq_range_index_dive_limit to N+1. To disable statistics and always use index dives regardless of N, set eq_range_index_dive_limit to 0.

To update table index statistic for best estimates, use ANALYZE TABLE.

## Range Optimization of Row Constructor Expressions

The optimizer is able to apply the range scan access method to queries of this form:

## Limiting Memory Use for Range Optimization

To control the memory available to the range optimizer, use the range_optimizer_max_mem_size system variable.
