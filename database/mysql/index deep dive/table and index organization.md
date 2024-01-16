# Table and Index Organization

- The physical organization of indexes and tables
- The structure and use of the index and table pages, index and table rows, buffer pools, and disk cache.
- The characteristics of disk I/Os, random and sequential
- The significance and synchronous and asynchronous I/Os
- The similarities and differences between database management systems
- Pages and table clustering, index rows, index-only tables and page adjacency
- The very confusing but important issue of the term cluster
- Alternatives to B-tree indexes
- Bitmap indexes and hashing

## Index and Table pages

Index and table rows are grouped together in pages; these are often 4K in size, this being a rather convenient size to use for most purposes, but other page sizes may be used.

Buffer pools and I/O activity are based on pages; for example, an entire page will be read from disk into a buffer pool. This means that several rows, not just one, are read into the buffer pool with a single I/O. We will also see that several pages may be read into the pool by just one I/O.

## Index rows

An index row is a useful concept when evaluating access paths. For a unique index, such as the primary key index, it is equivalent to an index entry in the leaf page; the column values are copied from the table to the index, and a pointer to the table row added. Usually, the table page number forms a part of this pointer. For a nonunique index, the index rows for a particular index value should be visualized as individual index entries. What is actually stored in a nonunique index is, in most cases, the CITY value followed by several pointers.

## Index structure

## Table rows

Each index row points to a corresponding row in the table; the pointer usually identifies the page in which the row resides together with some means of identifying its position within the page. Each table row contains some control information to define the row and to enable the DBMS to handle insertions and deletions, together with the columns themselves.

The sequence in which the rows are positioned in the table, as a result of a table load or row inserts, may be defined so as to be the same as that of one of its indexes. In this case, as the index 

Obviously, only one of the indexes can be defined to determine the sequence of the table rows in this way.

## Buffer pools and disk I/O

One of the primary objectives of relational database management systems is to ensure that data from tables and indexes is readily available when required. To enable this objective to be achieved as far as possible buffer pools, held in memory, are used to minimize disk activity. Each DBMS may have several pools according to the type, table or index, the the page size. The buffer pool manager will attempt to ensure that frequently used data remains in the pool to avoid the necessity of additional reads from disk. How effective this is will be extremely important with respect to the performance of SQL statements, and so will be equally important for the purposes of this book.

## Reads from the DBMS Buffer Pool

If an index or table page is found in the buffer pool, the only cost involved is that of the processing of the index or table rows. This is highly dependent on whether the row is rejected or accepted by the DBMS, the former incurring very little processing, the latter incurring much more as we will see in due course.

## Ramdom I/O from Disk Drives

Again, we must remember that a page will contain several rows; we may be interested in all of these rows, just a few of them, or even only a single row - the cost will be the same, roughly 10ms. If the disk drives are heavily used, this figure might be considerably increased as a result of having to wait for the disk to become available.

## Reads from the Disk Server Cache

Fortunately, disk servers in use today provide their own memory (or cache) in order to reduce this huge cost in terms of elapsed time. 

## Sequential Reads from Disk Drives

So far, we have only considered reading a single index or table page into the buffer pool. There will be many occasions when we actually want to read several pages into the pool and process the rows in sequence. The DBMS will be aware that several index or table pages should be read sequentially and will identify those that are not already in buffer pool. It will then issue multiple-page I/O requests, where the number of pages in each request will be determined by the DBMS; only those pages not already in the buffer pool will be read because those that are already in the pool may contain updated data has not yet been written back to disk.

There are two very important advantages to reading pages sequentially:

- Reading several pages together means that the time per page will be reduced; with current disk severs, the value may be as low as 0.1 ms for 4K pages (40 MB/s)
- Because the DBMS knows in advance which pages will be required, the reads can be performed before the pages are actually requested; this is called prefetch.

## Assisted Random Reads

We have seen how heavy the cost of random reads can be, and how buffer pools and disk caches can help to minimize this cost.

## Automatic Skip Sequential

By definition, an access pattern will be skip-sequential if a set of noncontiguous rows are scanned in one direction. The I/O time per row will thus be automaically shorter than with random access;

## List Prefetch

## DBMS SPECIFICS

### Pages

The size of the table pages sets an upper limit to the length of table rows. Normally, a table row must fit in one table page; an index row must fit in one leaf page. If the average lenght of the rows in a table is more than one third of the page size, space utilization suffers. Only one row with 2100 bytes fits in a 4K page, for instance. The problem of unusable space is more pronounced with indexes. As new index row must be placed in a leaf page 

**Page Adjacency**

Are the logically adjacent pages (such as leaf page 1 and leaf page 2) physically adjacent on disk? Sequential read would be very fast if they are.

In some older DBMSs, such as SQL/DS and the early versions of SQL Server, the pages of an index or table could be spread all over a large file. The only difference in the performance of random and sequential read was then due to the fact that a number of logically adjacent rows resided in the same page.
