# Use of Index Extensions

InnoDB automatically extends each secondary index by appending the primary key columns to it. Consider this table definition:

```sql
CREATE TABLE t1 (
    i1 INT NOT NULL DEFAULT O,
    i2 INT NOT NULL DEFAULT O,
    d DATE DEFAULT NULL,
    PRIMARY KEY (i1, i2),
    INDEX k_d (d)
) ENGINE = InnoDB;
```

This table defines the primary key on columns (i1, i2). It also defines a secondary index k_d on column (d), but internally InnoDB extends this index and treats it as columns (d, i1, i2).

The optimizer takes into account the primary key columns of the extended secondary index when determining how and whether to use that index. This can result in more efficient query execution plans and better performance.

The optimizer can use extended secondary indexes for ref, merge, and index_merge index access, for Loose Index Scan access, for join and sorting optimization, and for MIN(), MAX() optimization.
