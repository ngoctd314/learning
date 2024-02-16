# Comparison of B-Tree and Hash Indexes

Understanding the B-tree and hash data structures can help predict hwo different queries perform on different storage engines that use these data structs in their indexes.

## B-Tree Index Characteristics

A B-tree index can be used for column comparisons in expressions that use the =, >, >=, <, <=, or BETWEEN operations. The index also can be used for LIKE comparisons if the argument to LIKE is a constant string that does not start with a wildcard character.

```sql
SELECT * FROM tbl_name WHERE key_col LIKE 'Patrick%';
SELEcT * FROM tbl_name WHERE key_col LIKE 'Pat%_ck%';
```
The following SELECT statements do not use indexes

```sql
SELECT * FROM tbl_name WHERE key_col LIKE '%Patrick%';
SELECT * FROM tbl_name WHERE key_col LIKE other_col;
```
In the first statement, the LIKE value begins with a wildcard character. In the second statement, the LIKE value is not a constant.

If you use LIKE '%string%' and string is longer than three characters, MySQL uses the Turbo Boyer-Moore algorithm to initialize the pattern for the string and then uses this pattern to perform the search more quickly.

A search using col_name IS NULL employs indexes if col_name is indexes.

To be able to use an index, a prefix of the index must be used in every AND group.

## Hash Index Characteristics

Hash indexes have somewhat different characteristics from those just discussed:

- They are used only for equality comparisons that use the = or <=> operators (but are very fast). They are not used for comparison operators such as < that find a range of values. Systems that rely on this type of single-value lookup are known as "key-value stores"; to use MySQL for such applications, use hash indexes wherever possible. 
- The optimizer cannot use a hash index to speed up ORDER BY operations. (This type of index cannot be used to search for the next entry in order).
- MySQL cannot determine approximately how many rows there are between two values.
- Only whole keys can be used to search for a row.(With a B-tree, any leftmost prefix of the key can be used to find rows.)
