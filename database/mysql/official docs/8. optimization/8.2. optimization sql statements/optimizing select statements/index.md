Queries, in the form of SELECT statements, perform all the lookup operations in the database. Tuning these statements is a top priority, whether to achieve sub-second response times for dynamic web pages, or to chop hours off the time to generate huge overnight reports.

The main considerations for optimizing queries are:

- To make a slow SELECT ... WHERE query faster, the first thing to checks is whether you can add an index. Set up indexes on columns used in the WHERE clause, to speed up evaluation, filtering, and the final retrieval of results. To avoid wasted disk space, construct a small set of indexes that speed up many related queries used in your application.

Indexes are especially important for queries that reference different tables, using features such as joins and foreign keys.

- Isolate and tune any part of the query, such as function call, that takes excessive time. Depending on how the query is structured, a function could be called once for every row in the result set, or even once for every row in the table.

- Minimize the number of full table scans in your queries, particularly for big tables.

- Keep table statistics up to date by using the ANALYZE TABLE statement periodically, so the optimizer has the information needed to construct an efficient execution plan.

- Learn the tuning techniques, indexing techniques, and configuration parameters that are specific to the storage engine for each table.

- You can optimize single-query transactions for InnoDB tables.

- Avoid transforming the query in ways that make it hard to understand.

- If a performance issue is not easily solved by one of the basic guidelines, investigate the internal details of the specific query by reading the EXPLAIN plan and adjusting your indexes, WHERE clauses, join clauses, and so on.

- Adjust the size and properties of the memory areas that MySQL uses for caching. With efficient use the InnoDB buffer pool, MyISAM key cache, and the MySQL query cache, repeated queries run faster because the results are retrieved from memory the second and subsequent times.

- Deal with locking issues, where the speed of your query might be affected by other sessions accessing the tables at the same time.
