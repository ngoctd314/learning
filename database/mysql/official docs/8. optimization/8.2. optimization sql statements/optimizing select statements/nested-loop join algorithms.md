# Nested-Loop Join Algorithms

MySQL executes joins between tables using a nested-loop algorithm or variations on it.

## Nested-Loop Join Algorithm

A simple nested-loop join (NLJ) algorithm reads rows from the first table in a loop one at a time, passing each row to a nested loop that processes the next table in the join.

```txt
Table   Join Type
t1      range
t2      ref
t3      ALL
```

If a simple NLJ algorithm is used, the join is processed like this:

```go
for each row in t1 matching range {
    for each row in t2 matching reference key {
        for each row in t3 {
            if row satisfies join conditions, sends to client
        }
    }
}
```

Because the NLJ algorithm passes rows one at a time from outer loops to inner loops, it typically reads tables processed in the inner loops many times.

## Block Nested-Loop Join Algorithm

A Block Nested-Loop (BNL) algorithm uses buffering of rows read in outer loops to reduce the number of times that tables in inner loop must be read. For example, if 10 rows are read into a buffer and the buffer is passed to the next inner loop, each row read in the inner loop can be compared against all 10 rows in the buffer. This reduces by an order or magnitude the number of times the inner table must be read.

MySQL join buffering has these characteristics:

- Join buffering can be used when the join is of type ALL or index (in other words, when no possible keys can be used, and a full scan is done, of either the data or index rows respectively), or range. Use of buffering is also applicable to outer joins.
- A join buffer is never allocated for the first nonconstant table, even if would be of type ALL or index.
- Only columns of interest to a join are stored in its join buffer, not whole rows.
