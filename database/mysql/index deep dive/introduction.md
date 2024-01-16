# Introduction

- To understand how SQL optimizers decide what table and index scans should be performed to process SQL statements as efficiently as possible.
- To be able to quantify the work being done during these scans to enable satisfactory index design.
- Type and background of audience for whom the book is written.
- Initial thoughts on the major reasons for inadequate indexing.
- Systematic index design.

## Another book about SQL performance!

Relational databases have been around now for over 20 years, and that's precisely

## INADEQUATE INDEXING

For many years, inadequate indexing has been the most common cause of performance disappointments. The most widespread problem appears to be that indexes do not have sufficient columns to support all the predicates of a WHERE clause. Frequently, there are not enough indexes on a table; some SELECTs may have no useful index; sometimes an index has the right columns but in the wrong order.

It is relatively easy to improve the indexing of a relational database because no program changes are required. However, a change to a production system always carries some risk. Furthermose, while a new index is being created, update programs may experience long waits because they are not able to update a table being scanned for a CREATE INDEX.

- The index design topics are short, perhaps only a few pages.
- The negative side effects of indexes are emphasized; indexes consume disk space and they make inserts, updates, and deletes slower.
- Index design guidelines are vague and sometimes questionable. Some writers recommend indexing all restrictive columns. Others claim that index design is an art that can only be mastered through trial and error.

## MYTHS AND MISCONCEPTIONS

Even recent books, such as one published as late as 2002, suggest that only the root page of a B-tree index will normally stay in memory. This was an appropriate 20 years ago, when memory was typically so small that the database buffer pool could contain only a few hundred pages, perhaps less than a megabyte. Today, the size of the database buffer pools may be hundreds of thousands of pages, one gigabyte (GB) or more; the read caches of disk servers are typically even larger - 64 GB, for instance. Although databases have grown as disk storage has become cheaper, it is nwo realistic to assume that all the nonleaf pages of a B-tree index will usually remain in memory or the read cache. Only the leaf pages will normally need to be read from a disk drive; this, of course, makes index maintenance much faster.

**Myth 1: No More Than Five Index Levels**

**Myth 2: No more than six indexes per Table**

**Myth 3: Volatile Columns Should Not Be Indexed**

Index rows are held in key sequence
