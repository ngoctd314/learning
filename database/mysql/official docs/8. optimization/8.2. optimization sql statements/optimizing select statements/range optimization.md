# Range optimization

The range access method uses a single index to retrieve a subset of table rows that are contained within one or several index value intervals. It can be used for a single-part or multiple-part index.

## Range Access Method for Single-Part Indexes

For a single-part index, index value intervals can be conveniently represented by corresponding conditions in the `WHERE` clause, denoted as range conditions rather than "intervals".

The definition of a range condition for a single-part index is as follows:

- For both BTREE and HASH indexes, comparison of a key part with a constant value is a range condition when using the =, <=>, IN(), IS NULL, or IS NOT NULL operators.
- Additionally, for BTREE indexes, comparison of a key  part with a constant value is a range condition when using the >, <, >=, <=, BETWEEN, !=, or <> operators, or LIKE comparisons if the argument to LIKE is a constant string that does not start with a wildcard character.

"Constant value" in the preceding descriptions means one of the following:

- A constant from the query string
- A column of a const or system table from the same join
- The result of an uncorrelated subquery
- And expression composed entirely from subexpressions of the preceding types

## Range Access Method for Multiple-Part Indexes

Range conditions on a multiple-part index are an extension of range conditions for a single-part index. A range condition on a multiple-part index restricts index rows to lie within one or several key tuple intervals.
