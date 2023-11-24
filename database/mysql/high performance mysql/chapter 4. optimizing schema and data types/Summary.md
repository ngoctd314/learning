# Summary

To write sargable  queries:
- Avoid using functions or calculations on indexes columns in the WHERE clause.
- Use direct comparisons when possible, instead of wrapping the column in a function.
- If we need to use a function on a column, consider creating a computed column or a function-based index, if the database system supports it.

Good schema design is pretty universal, but of course MySQL has special impl details to consider. In a nutshell, it's good idea to keep things as small and simple as you can.

- Try to avoid extremes in your design, such as a schema that will force enormously complex queries, or tables with oddles and oodles of columns.
- Use small, simple, appropriate data types, and void NULL unless it's actually the right way to model your data's reality.
- Try to use the same data types to store similar or related values, especially if they'll be used in a join condition.
- Watch out for variable-length strings, which might cause pessimistic full-length memory allocation for temporary tables and sorting.
- Try to use integers for identifiers if you can.
- Avoid the legacy MySQL-isms such as specifying precisions for floating-point numbers or display widths for integers.
- Be careful with ENUM and SET. They're handy, but they can be abused, and they're tricky sometimes. BIT is best avoided.

Normalization is good, but denomarlization (duplication of data, in most cases) is sometimes actually necessary and benefical. We'll see more examples of that in the next chapter. And precomputing, caching, or generating summary tables can also be a big win.

Finally, ALTER TABLE can be painful because in most cases, it locks and rebuilds the whole table. We showed a number of workarounds for specific cases; for the general case, you'll have to use other techniques, such as performing the ALTER on a replica and then promoting it to master.
