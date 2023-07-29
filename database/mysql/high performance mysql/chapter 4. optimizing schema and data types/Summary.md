# Summary

Good logical and physical design is cornerstone of high performance, and you must design your schema for the specific queries you will run.

To write sargable  queries:
- Avoid using functions or calculations on indexes columns in the WHERE clause.
- Use direct comparisons when possible, instead of wrapping the column in a function.