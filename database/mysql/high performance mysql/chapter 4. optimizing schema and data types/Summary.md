# Summary

Good logical and physical design is cornerstone of high performance, and you must design your schema for the specific queries you will run.

To write sargable  queries:
- Avoid using functions or calculations on indexes columns in the WHERE clause.
- Use direct comparisons when possible, instead of wrapping the column in a function.
- If we need to use a function on a column, consider creating a computed column or a function-based index, if the database system supports it.