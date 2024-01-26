# Improving Query Response Time

Improving query response time is a journey called query optimization. Query optimization takes time and effort, and there is a destination: faster query response time. To make the journey efficient - not a waste of time and effort - there are two parts: direct query optimization and indirect query optimization.

## Direct Query Optimization

Changes to queries and indexes. These changes solve a lot of performance problems, which is why the journey begins with direct query optimization.

- Range optimization
- Index Merge optimization
- Hash Join optimization
- Index Condition Pushdown optimization
- Multi-Range Read optimization
- Constant-Folding optimization
- IS NULL optimization
- ORDER BY optimization
- GROUP BY optimization
- DISTINCT optimization
- LIMIT Query optimization

## Indirect Query Optimization

Indirect query optimization is changes to data and access patterns. Instead of changing a query, you change what the query accesses and how: its data and access patterns, respectively. These changes indirectly optimize the query because query, data, and access patterns are inextricable with respect to performance. Changes to one influence the others. It's easy to prove.

Suppose you have a slow query. Data size and access patterns don't matter for this proof, so imagine whatever you like. I can reduce query response time to near-zero. (Let's say near-zero is 1 microsecond. For a computer that's a long time, but for a human it's imperceptible.) The indirect "optimization" is: TRUNCATE TABLE. With no data, MySQL can execute any query in near-zero time. That's cheating, but it nonetheless proves the point: reducing data size improves query response time.

Let's revisit the car analogy. Indirect query optimization is analogous to changing major design elements of the car. For example, weight is a factor in fuel efficiency: decreasing weight increases fuel efficiency. (Data is analogous to weight, which is why TRUNCATE TABLE dramatically increases performance - but don't use this "optimization"). Reducing weight is not a straightforward (direct) change because engineers can't magically make parts weigh less. 

A greater level of effort is why indirect query optimization is part two of the journey. If direct query optimization solves the problem, the stop - be efficient.
