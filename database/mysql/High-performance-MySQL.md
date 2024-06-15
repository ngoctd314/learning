[[_TOC]]

# Chapter 1. MySQL Architecture and History

MySQL can power embedded applications, data warehouses, content indexing and delivery software, highly available redundant systems, online transaction processing (OLTP), and much more.

## MySQL's Logical Architecture

A good mental picture of how MySQL's components work together will help you understand the server.

The topmost layer contains the services that aren't unique to MySQL. They're services most network-based client/server tools or servers need: connection handling, authentication, security, and so forth.

The second layer is where things get interesting. Much of MySQL's brains are here, including the code for query parsing, analysis, optimization, caching, and  all the built-in functions (dates, times, math, and encryption). Any functionality provided across storage engines lives at this level: stored producers, triggers, and views.

The third layer contains the storage engines. They are responsible for storing and retrieving all data stored "in" MySQL. The server communicates with them through the storage engine API. This interface hides differences between storage engines and makes them largely transparent at the query layer. The API contains a couple of dozen low-level functions that perform operations such as "begin a transaction" or "fetch the row that has this primary key". The storage engines don't parse SQL or communicate with each other; the simple response to request from server. One exception is InnoDB, which does parse foreign key definitions, because the MySQL server doesn't yet implement them itself.

![alt](./assets/high-performance-mysql/mysql-arch.png)

### Connection Management and Security

Each client connection gets its own thread within the server process. The connection's queries execute within that single thread, which in turn resides on one core or CPU.

=> If a query is too large, you can divide it into multiple queries, and then merge the results in the application. This approach creates opportunities to handle large queries more effectively and enables parallel execution.

The server cache threads, so they don't need to be created and destroyed for each new connection.

When clients (applications) connect to the MySQL server, the server needs to authenticate them. Authentication is based on username, originating host, and password.

### Optimization and Execution

MySQL parses queries to create an internal structure (the parse tree), and then applies a variety of optimizations. These can include rewriting the query, determining the order in which it will read tables, choosing which indexes to use, and so on. You can pass hints to the optimizer through special keywords in the query, affecting its decision making process. You can also ask the server to explain various aspects of optimization. This lets you know what decisions the server is making and gives you a reference point for reworking queries, schemas and settings to make everything run as efficiently as possible.

The optimizer does not really care what storage engine a particular table uses, but the storage engine does effect how the server optimizes the query. The optimizer asks the storage engine about some of its capabilities and the cost of certain operations, and for statistics on the table data. 

Before even parsing the query, though, the server consults the query cache, which can store only SELECT statements, along with their result sets. If anyone issues a query that's identical to one already in the cache, the server doesn't need to parse, optimize, or execute the query at all - it can simply pass back the stored result set.

## Concurrency Control

Anytime more than one query needs to change data at the same time, the problem of concurrency control arises. MySQL has to do this at two levels: the server level and the storage engine level. Concurrency level is a big topic which a large body of theoretical literature is devoted, so we will just give you a simplified overview of how MySQL deals with concurrent readers and writers, so you have the context you need for the rest of this chapter.

### Read/Write Locks

Reading from the mailbox isn't as troublesome. There's nothing wrong with multiple clients reading the same mailbox simultaneously; because they aren't making changes, nothing is likely to go wrong. But what happens if someone tries to delete message number 25 while programs are reading the mailbox? It depends, but a reader could come away with corrupted or inconsistent view of the mailbox. So, to be safe, even reading from a mailbox require special care.

The solution to this classic problem of concurrency control is rather simple. Systems that deal with concurrency read/write access typically implement a locking system that consists of two lock types. These locks are usually known as shared locks and exclusive locks, or read locks and write locks.

Without worrying about the actual locking technology, we can describe the concepts as follows. Read locks on a resource are shared, or mutually non-blocking: many clients can read from a resource at the same time and not interfere with each other. Write locks, on the other hand, are exclusive - i.e., they block both read locks and other write locks - because the only safe policy is to have a single client writing to the resource at a given time and to prevent all reads when a client is writing.

In the database world, locking happens all the time: MySQL has to prevent one client reading a piece of data while another is changing it. It performs this lock management internally in a way that is transparent much of the time.

### Lock Granularity

One way to improve the concurrency of a shared resource is to be more selective about what you lock. Rather than locking the entire resource, lock only the part that contains the data you need to change. Better yet, lock only the exact piece of data you plan to change. Minimizing the amount of data that you lock at any one time lets changes to a given resource occur simultaneously, as long as they don't conflict with each other.

The problem is locks consume resources. Every lock operation - getting a lock, checking to see whether a lock is free, releasing a lock, and so on - has overhead. If the system spends too much time managing locks instead of storing and retrieving data, performance can suffer.

A locking strategy is a compromise between lock overhead and data safety, and that compromise effects performance. Most commercial database servers don't give you much choice: you get what is known as row-level locking in your tables, with a variety of often complex ways to give good performance with many locks.

### Table locks

The most basic locking strategy available in MySQL, and the one with the lowest overhead is table locks. It locks the entire table. When a client wishes to write to a table (insert, delete, update, etc), it acquires a write lock. This keeps all other read and write operations at bay. When nobody is writing, readers can obtain read locks, which don't conflict with other read locks.

Table locks have variations for good performance in specific situations. For example, READ LOCAL table locks allow some types of concurrent write operations. Write locks also have a higher priority than read locks, so a request for a write lock will advance to the front of the lock queue even if readers are already in the queue (write locks can advance past read locks in the queue, but read lock cannot advance past write locks).

Although storage engines can manage their own locks, MySQL itself also uses a variety of locks that are effectively table-level for various purposes. For instance, the server uses a table-level lock for statements such as ALTER TABLE, regardless of the storage engine. 

```go
type Table struct {
    m sync.RWMutex
    rows [int]Row
}
```

### Row locks

The locking style that offers greatest concurrency (and carries the greatest overhead) is the use of row locks.

Row locks are implemented in the storage engine, not the server. The server is completely unaware of locks implemented in the storage engines, and as you'll see later in this chapter and throughout the book, the storage engines all implement locking in their own ways.

## Transactions

A transaction is a group of SQL queries that are treated atomically, as a single unit of work. If the database engine can apply the entire group of queries to a database, it does so, but if any of them can't be done because of a crash or other reason, none of them is applied. It's all or nothing. 

**Atomicity**

A transaction must function as a single indivisible unit of work so that the entire transaction is either applied or rolled back. It's all or nothing.

**Consistency**

The database should always move from one consistent state to the next. If the transaction is never committed, none of the transaction's changes are ever reflected in the database. 

**Isolation**

The results of a transaction are usually invisible to other transactions until the transaction is complete. When we discuss isolation levels, you'll understand why we said usually invisible.

**Durability**

Once committed, a transaction's changes are permanent. This means the changes must be recorded such that data won't be lost in a system crash. Nothing is 100% durable. 

Just as with increased lock granularity, the downside of this extra security is that the database server has to do more work. A database server with ACID transactions also generally requires more CPU power, memory, and disk space than one without them. If you don't really need transaction, you might be able to get higher performance with a non-transactional storage engine for some kinds of queries. You might be able to use **LOCK TABLES** to give the level of protection you need without transactions.

### Isolation Levels

Isolation is more complex than it looks. The SQL standard defines four isolation levels with specific rules for which changes are and aren't visible inside and outside a transaction. Lower isolation levels typically allow higher concurrency and have lower overhead.

```sql
CREATE TABLE `persons` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

#### READ UNCOMMITTED

In the **READ UNCOMMITTED** isolation level, transactions can view the results of un-committed transactions. At this level, many problems can occur unless you really, really know that you are doing and have a good reason for doing it. This level is rarely used in practice, because its performance isn't much better than the other levels, which have many advantages. Reading uncommitted is also known as a dirty read.

```go
func main() {
	waitForUpdate := make(chan struct{}, 1)
	waitForSelect := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		rs, _ := tx.Exec("UPDATE persons SET name=? WHERE id =?", "test", 1)
		slog.Info("UPDATE persons", "results", rs)
		waitForUpdate <- struct{}{}
		<-waitForSelect
		tx.Commit()
		slog.Info("COMMIT UPDATE")
	}()
	go func() {
		defer wg.Done()

		<-waitForUpdate
        // read data is uncommitted
		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadUncommitted,
		})
		rows, _ := tx.Query("SELECT name FROM persons WHERE id=?", 1)
		for rows.Next() {
			var name string
			_ = rows.Scan(&name)
			slog.Info("SELECT name", "value", name)
		}
		tx.Commit()
		slog.Info("COMMIT SELECT")
		waitForSelect <- struct{}{}
	}()
	wg.Wait()
}
```

```txt
2024/06/05 06:58:09 INFO UPDATE persons results="{Locker:0xc00011e090 resi:0xc000188000}"
2024/06/05 06:58:09 INFO SELECT name value=test
2024/06/05 06:58:09 INFO COMMIT SELECT
2024/06/05 06:58:09 INFO COMMIT UPDATE
```

#### READ COMMITTED

The default isolation level for most database systems (but not MySQL) is READ COMMITTED. It satisfies the simple definition of isolation used earlier: a transaction will see only those changes made by transactions that were already committed when it began, and its changes won't be visible to other until it has committed. This level still allows what's known as a non-repeatable read. This means you can run the same statement twice and see different data.

```go
func main() {
	waitForUpdate := make(chan struct{}, 1)
	waitForSelect := make(chan struct{}, 1)
	waitForUpdateCommit := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		rs, _ := tx.Exec("UPDATE persons SET name=? WHERE id =?", "test", 1)
		slog.Info("UPDATE persons", "results", rs)
		waitForUpdate <- struct{}{}
		<-waitForSelect
		tx.Commit()
		waitForUpdateCommit <- struct{}{}
		slog.Info("COMMIT UPDATE")
	}()
	go func() {
		defer wg.Done()

		<-waitForUpdate
		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
		{
			rows, _ := tx.Query("SELECT name FROM persons WHERE id=?", 1)
			for rows.Next() {
				var name string
				_ = rows.Scan(&name)
				slog.Info("result: SELECT name", "value", name)
			}
		}
		waitForSelect <- struct{}{}
		<-waitForUpdateCommit
		{
			rows, _ := tx.Query("SELECT name FROM persons WHERE id=?", 1)
			for rows.Next() {
				var name string
				_ = rows.Scan(&name)
				slog.Info("result: SELECT name", "value", name)
			}
		}
		tx.Commit()
		slog.Info("COMMIT SELECT")
	}()
	wg.Wait()
}
```

```txt
2024/06/05 07:13:53 INFO UPDATE persons results="{Locker:0xc0000ae090 resi:0xc000104000}"
2024/06/05 07:13:53 INFO result: SELECT name value=""
2024/06/05 07:13:53 INFO COMMIT UPDATE
2024/06/05 07:13:53 INFO result: SELECT name value=test
2024/06/05 07:13:53 INFO COMMIT SELECT
```

#### REPEATABLE READ

REPEATABLE READ solves the problems that READ UNCOMMITTED allows. It guarantees that any rows a transaction reads will "look the same" in subsequent reads within the same transaction, but in theory it still allows another tricky problem: phantom reads. Simply put, a phantom read can happen when you select some range of rows, another transaction inserts a new row into the range, and then you select the same range again; you will then see the "phantom" row. InnoDB slove the phantom read problem with mvcc. 

REPEATABLE READ is MySQL's default transaction isolation level.

#### SERIALIZABLE

The highest level of isolation, SERIALIZABLE, solves the phantom read problem by forcing transactions to be ordered so that they can't possibly conflict. In a nutshell, SERIALIZABLE, places a lock on every row it reads. At this level, a lot of timeouts and lock contention can occur. We've rarely seen people use this isolation level, but your application's needs might force you to accept the decreased concurrency in favor of the data stability that results.

|Isolation level|Dirty reads possible|Non-repeatable reads possible|Phantom reads possible|Locking reads|Locking writes|
|-|-|-|-|-|-|
|READ UNCOMMITTED|Yes|Yes|Yes|No|Yes|
|READ COMMITTED|No|Yes|Yes|No|Yes|
|REPEATABLE READ|NO|No|Yes|No|Yes|
|SERIALIZABLE|No|No|No|Yes|Yes|

```go
func main() {
	db.Exec("DROP TABLE IF EXISTS test")
	db.Exec("CREATE TABLE test (id int auto_increment primary key, a int)")
	db.Exec("INSERT INTO test (id, a) VALUES (1, 0)")

	txA, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	row := txA.QueryRow("SELECT a from test WHERE id = 1")
	var a int
	row.Scan(&a)
	fmt.Printf("txA: a: %d\n", a)
	txA.Exec("UPDATE test SET a = 1 WHERE id = 1 ")

	txB, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadUncommitted})
	row = txB.QueryRow("SELECT a from test WHERE id = 1")
	row.Scan(&a)
	fmt.Printf("txB: a: %d\n", a)
	txB.Commit()

	txA.Commit()
}
```

```txt
txA: a: 0
txB: a: 1
```

```go
func main() {
	db.Exec("DROP TABLE IF EXISTS test")
	db.Exec("CREATE TABLE test (id int auto_increment primary key, a int)")
	db.Exec("INSERT INTO test (id, a) VALUES (1, 0)")

	txA, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	row := txA.QueryRow("SELECT a from test WHERE id = 1")
	var a int
	row.Scan(&a)
	fmt.Printf("txA: a: %d\n", a)
	txA.Exec("UPDATE test SET a = 1 WHERE id = 1 ")

	txB, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	row = txB.QueryRow("SELECT a from test WHERE id = 1")
	row.Scan(&a)
	fmt.Printf("txB: a: %d\n", a)
	txB.Commit()

	txA.Commit()
}
```

```txt
txA: a: 0
txB: a: 0
```

```go
func main() {
	db.Exec("DROP TABLE IF EXISTS test")
	db.Exec("CREATE TABLE test (id int auto_increment primary key, a int)")
	db.Exec("INSERT INTO test (id, a) VALUES (1, 0)")

	txA, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelReadUncommitted})

	row := txA.QueryRow("SELECT a from test WHERE id = 1")
	var a int
	row.Scan(&a)
	fmt.Printf("txA: a: %d\n", a)
	txA.Exec("UPDATE test SET a = 1 WHERE id = 1 ")

	txB, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
	row = txB.QueryRow("SELECT a from test WHERE id = 1")
	row.Scan(&a)
	fmt.Printf("txB: a: %d\n", a)
	txB.Commit()

	txA.Commit()
}
```

```txt
txA: a: 0
Deadlock
```

### Deadlocks

A deadlock is when two or more transactions are mutually holding are requesting locks on the same resources, creating a cycle of dependencies. Deadlocks occur when transactions try to lock resources in a different order. They can happen whenever multiple transactions lock the same resources. 

To combat this problem, database systems implement various forms of deadlock detection and timeouts. The more sophisticated systems, such as the InnoDB storage engine, will notice circular dependencies and return an error instantly. This can be a good thing - otherwise, deadlocks would manifest themselves as very slow queries. Others will give up after the query exceeds a lock wait timeout, which is always good. The way InnoDB currently handles deadlocks is to roll back the transaction that has fewest exclusive row locks.

Deadlocks cannot be broken without rolling back one of the transactions, either partially or wholly. They are a fact of life in transactional systems, and your applications should be designed to handle them.

### Transaction Logging

### Transaction In MySQL

## Multiversion Concurrency Control

## MySQL's Storage Engines

## A MySQL Timeline

## MySQL's Development Model

## Summary

# Chapter 2. Benchmarking MySQL
