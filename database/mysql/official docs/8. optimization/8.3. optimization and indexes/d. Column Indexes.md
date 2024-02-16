# Column Indexes

The most common type of index involves a single column, storing copies of the values from that column in a data structure, allowing fast lookups for the rows with the corresponding column values. The B-tree data structure lets the index quickly find a specific value, a set of values, or a range of values, corresponding to operations such as =, >, <=, BETWEEN, IN, and so on, in a WHERE clause.

## Index Prefixes

With col_name(N) syntax in an index specification for a string column, you can create an index that uses only the first N characters of the column. Indexing only a prefix of column values in this way can make the index file much smaller. When you index a BLOB or TEXT column, you must specify a prefix length for the index.

```sql
CREATE TABLE test (blob_col BLOB, INDEX(blob_col(10)));
```

Prefixes can be up to 1000 bytes long (767 bytes for InnoDB tables, unless you have innodb_large_prefix set). 

**Note**

Prefix limits are measured in bytes, whereas the prefix length in CREATE TABLE, ALTER TABLE, and CREATE INDEX statements is interpreted as number of characters for nonbinary string types (CHAR, VARCHAR, TEXT) and number of bytes for binary string types (BINARY, VARBINARY, BLOB). Take this into account when specifying a prefix length for a nonbinary string column that uses a multibyte character set.

If a search term exceeds the index prefix length, the index is used to exclude non-matching rows, and the remaining rows are examined for possible matches.

## FULLTEXT Indexes

FULLTEXT indexes are used for full-text searches. Only the InnoDB and MyISAM storage engines support FULL TEXT indexes and only for CHAR, VARCHAR, and TEXT columns. Indexing always takes place over the entire column and column prefix indexing is not supported.

## Spatial Indexes

You can create indexes on spatial data types. MyISAM and InnoDB support R-tree indexes on spatial types. Other storage engines use B-trees for indexing spatial types.

## Indexes in the MEMORY Storage Engine

The MEMORY storage engine uses HASH indexes by default, but also supports BTREE indexes.
