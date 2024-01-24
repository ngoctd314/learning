# Optimizing data size

Design your tables to minimize their space on the disk. This can result in huge improvements by reducing the amount of data written to and read from disk. Smaller tables normally require less main memory while their contents are being actively processed during query execution. Any space reduction for table data also results in smaller indexes that can be processed faster.

## Table Columns

- Use the most efficient (smallest) data types possible. MySQL has many specialized types that save disk space and memory. For example, use the smaller integer types if possible to get smaller tables. MEDIUMINT is often a better choice than INT because a MEDIUMINT column uses 25% less space.
- Declare columns to be NOT NULL if possible. It makes SQL operations faster, by enabling better use of indexes and eliminating overhead for testing whether each value is NULL. You also save some storage space, one bit per column. If you really need NULL values in your tables, use them. Just avoid the default setting that allows NULL values in every column.

## Row Format

- InnoDB tables are created using the DYNAMIC row format by default.

## Indexes

- The primary index of a table should be as short as possible. This makes identification of each row easy and efficient. For InnoDB tables, use primary key columns are duplicated in each secondary index entry, so a short primary key save considerable space if you have many secondary indexes.

- Create only the indexes that you need to improve query performance. Indexes are good for retrieval, but slow down insert and update operations. If you access a table mostly by searching on a combination of columns, create a single composite index on them rather than a separate index for each column. The first part of the index should be the column most used. If you always use many columns when selecting from the table, the first column in the index should be the one with the most duplicates, to obtain better compression of the index. 

- If it is very likely that a long string column has a unique prefix on the first number of characters, it is better to index only this prefix, using MySQL's support for creating an index on the leftmost part of the column. Shorter indexes are faster, not only because they require less disk space, but because they also give you more hits in the index cache,and thus fewer disk seeks.

## Joins

## Normalization

- Normally, try to keep all data nonredundant (observing what is referred to in database theory as third normal form).
- If speed is more important than disk space and the maintenance costs of keeping multiple copies of data. 
