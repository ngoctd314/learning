# Concurrency Problem in MySQL

Arise when multiple transactions or queries attempt to access or modify the same data simultaneously. These concurrency issues can lead to data inconsistencies, incorrect results, or conflicts.

## 1. Dirty Reads

A dirty read occurs when one transaction reads data that has been modified by another transaction but has not be committed yet. This can lead to reading inconsistent or incorrect data. Occur when a transaction selects data that hasn't been commited by another transaction. For example, transaction A changes a row. Transaction B then selects the changed row before transaction A commits the change. If transaction A then rolls back the change. If transaction A then rolls back the change, transaction B has selected data that doesn't exist in the database.

### 1.1. Dirty read with Isolation level ReadUncommitted

```go
trigger := make(chan struct{}, 1)
readUncommited := make(chan struct{}, 1)

go func() {
    <-trigger
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelReadUncommitted,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `SELECT * FROM users`
    listUser := []User{}
    if err := tx.Select(&listUser, query); err != nil {
        log.Println("select error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }

    log.Println("list user read commited", listUser)
    readUncommited <- struct{}{}

    if err := tx.Commit(); err != nil {
        log.Println("commit error", err)
    }
}()

go func() {
    listUser := []User{{Name: fmt.Sprintf("name_%s", time.Now().Format(time.DateTime)), Age: 2023}}
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelReadUncommitted,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `INSERT INTO users VALUES (:name, :age)`
    if _, err := tx.NamedExec(query, listUser); err != nil {
        log.Println("insert error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }
    trigger <- struct{}{}

    log.Println("insert success")
    <-readUncommited
    // dirty read here
    if err := tx.Rollback(); err != nil {
        log.Println("rollback error", err)
    }
    // if err := tx.Commit(); err != nil {
    // 	log.Println("commit error", err)
    // }
}()
```

### 1.2. Dirty read with isolation level READ COMMITED

```go
triggerSelectUser := make(chan struct{}, 1)
readCommited := make(chan struct{}, 1)

go func() {
    <-triggerSelectUser
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelReadCommitted,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `SELECT * FROM users`
    listUser := []User{}
    if err := tx.Select(&listUser, query); err != nil {
        log.Println("select error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }

    log.Println("list user read commited", listUser)

    if err := tx.Commit(); err != nil {
        readCommited <- struct{}{}
        log.Println("commit error", err)
    } else {
        log.Println("commit SELECT users")
    }
}()

go func() {
    listUser := []User{{Name: fmt.Sprintf("name_%s", time.Now().Format(time.DateTime)), Age: 2023}}
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelReadCommitted,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `INSERT INTO users VALUES (:name, :age)`
    if _, err := tx.NamedExec(query, listUser); err != nil {
        log.Println("insert error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }
    triggerSelectUser <- struct{}{}

    log.Println("insert success")
    <-readCommited
    if err := tx.Rollback(); err != nil {
        log.Println("rollback error", err)
    } else {
        log.Println("rollback INSERT users")
    }
    // if err := tx.Commit(); err != nil {
    // 	log.Println("commit error", err)
    // } else {
    // 	log.Println("commit INSERT users")
    // }

}()
```

### 1.3. Dirty read with isolation level REPETABLE READ

```go
triggerSelectUser := make(chan struct{}, 1)
readCommited := make(chan struct{}, 1)

go func() {
    <-triggerSelectUser
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelRepeatableRead,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `SELECT * FROM users`
    listUser := []User{}
    if err := tx.Select(&listUser, query); err != nil {
        log.Println("select error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }

    log.Println("list user read commited", listUser)

    readCommited <- struct{}{}
    if err := tx.Commit(); err != nil {
        log.Println("commit error", err)
    } else {
        log.Println("commit SELECT users")
    }
}()

go func() {
    listUser := []User{{Name: fmt.Sprintf("name_%s", time.Now().Format(time.DateTime)), Age: 2023}}
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelDefault,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `INSERT INTO users VALUES (:name, :age)`
    if _, err := tx.NamedExec(query, listUser); err != nil {
        log.Println("insert error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }
    triggerSelectUser <- struct{}{}

    log.Println("insert success")
    <-readCommited
    if err := tx.Rollback(); err != nil {
        log.Println("rollback error", err)
    } else {
        log.Println("rollback INSERT users")
    }
    // if err := tx.Commit(); err != nil {
    // 	log.Println("commit error", err)
    // } else {
    // 	log.Println("commit INSERT users")
    // }

}()
```
### 1.4. Dirty read with isolation level SERIALIZABLE

```go
go func() {
    <-triggerSelectUser
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelSerializable,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `SELECT * FROM users`
    listUser := []User{}
    log.Println("SELECT is blocked here")
    if err := tx.Select(&listUser, query); err != nil {
        log.Println("select error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }

    log.Println("list user serializable", listUser)

    if err := tx.Commit(); err != nil {
        log.Println("commit error", err)
    } else {
        log.Println("commit SELECT users")
    }
}()

go func() {
    listUser := []User{{Name: fmt.Sprintf("name_%s", time.Now().Format(time.DateTime)), Age: 2023}}
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelDefault,
    })
    if err != nil {
        log.Fatal(err)
    }

    query := `INSERT INTO users VALUES (:name, :age)`
    if _, err := tx.NamedExec(query, listUser); err != nil {
        log.Println("insert error", err)
        if err := tx.Rollback(); err != nil {
            log.Println("rollback error", err)
        }
        return
    }
    triggerSelectUser <- struct{}{}

    log.Println("exec INSERT")
    // if err := tx.Rollback(); err != nil {
    // 	log.Println("rollback error", err)
    // } else {
    // 	log.Println("rollback INSERT users")
    // }
    if err := tx.Commit(); err != nil {
        log.Println("commit error", err)
    } else {
        log.Println("commit INSERT users")
    }

}()
```

## 2. Lost Updates

A lost update occurs when two transactions try to update the same data concurrently, and one transaction update overwrites the changes made by the other transaction. As a result, one update is lost. Occur when two transactions select the same row and then update the row based on the values originally selected. Since each transaction is unaware of the other, the later update overwrites the earlier update.

### 2.1. Lost updates with isolation level READ UNCOMMITED, READ COMMITED, REPEATABLE READ ,SERIALIZABLE

```go
wg := sync.WaitGroup{}
n := 100
wg.Add(n)
for i := 0; i < n; i++ {
    go func() {
        defer wg.Done()
        tx, err := db.BeginTxx(ctx, &sql.TxOptions{
            Isolation: sql.LevelReadCommitted,
            ReadOnly:  false,
        })
        if err != nil {
            log.Fatal(err)
        }

        user := User{}
        if err := tx.Get(&user, "SELECT * FROM users WHERE name = ?", "ngoctd"); err != nil {
            fmt.Println("select error", err)
            if err := tx.Rollback(); err != nil {
                fmt.Println("rollback error", err)
            }
            return
        }

        if _, err := tx.Exec("UPDATE users SET age = 10 + ?", user.Age); err != nil {
            log.Println("update error", err)
            if err := tx.Rollback(); err != nil {
                log.Println("rollback error", err)
            }
            return
        }

        if err := tx.Commit(); err != nil {
            fmt.Println("commit error", err)
        } else {
            // fmt.Println("commit SELECT users")
        }
    }()
}

wg.Wait()

user := User{}
if err := db.Get(&user, "SELECT * FROM users WHERE name = ?", "ngoctd"); err != nil {
    log.Println("select error", err)
    return
}
log.Println(user)
```
## Nonrepeatable reads

Occur when two SELECT statement that try to get the same data get different values because another transaction has updated the data in the time between the two statements. When transaction A selects the same row again, the data is different. For example, transaction A selects a row, transaction B then updates the row. When transaction A selects the same row again, the data is different.

```go
wg := sync.WaitGroup{}
updateUser := make(chan struct{}, 1)

go func() {
    wg.Add(1)
    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelRepeatableRead,
    })
    if err != nil {
        log.Fatal(err)
    }

    user := User{}
    err = tx.Get(&user, "SELECT * FROM users WHERE name = ?", "ngoctd")
    if err != nil {
        fmt.Println("SELECT user error", err)
        if err := tx.Rollback(); err != nil {
            fmt.Println("rollback error", err)
        }
        return
    }
    fmt.Println(user)
    // block to update user
    <-updateUser
    err = tx.Get(&user, "SELECT * FROM users WHERE name = ?", "ngoctd")
    if err != nil {
        fmt.Println("SELECT user error", err)
        if err := tx.Rollback(); err != nil {
            fmt.Println("rollback error", err)
        }
        return
    }
    fmt.Println(user)

    if err := tx.Commit(); err != nil {
        fmt.Println("commit error", err)
    }

}()

go func() {
    wg.Add(1)

    tx, err := db.BeginTxx(ctx, &sql.TxOptions{
        Isolation: sql.LevelRepeatableRead})
    if err != nil {
        log.Fatal(err)
    }

    if _, err := tx.Exec("UPDATE users SET age = 18 WHERE name = ?", "ngoctd"); err != nil {
        fmt.Println("UPDATE error", err)
        if err := tx.Rollback(); err != nil {
            fmt.Println("rollback error", err)
        }
        return
    }

    if err := tx.Commit(); err != nil {
        fmt.Println("commit error", err)
    }
    updateUser <- struct{}{}

}()

wg.Wait()
```

## Phantom reads

Occur when you perform an update or delete on set of rows at the same time that another transaction if performing an insert or delete that affects one or more rows in that same set of rows. For example, transaction a updates the payment total for each invoice that has a balance due, but transaction B inserts a new, unpaid, invoice while transaction A is still running. After transaction A finishes, there is still an invoice with a balance due.

**The concurrency problems prevented by each transaction isolation level**

|Isolation level|Dirty reads|Lost updateds|Nonrepeatable reads|Phantom reads|
|-|-|-|-|-|
|READ UNCOMMITTED|Allows|Allows|Allows|Allows|
|READ COMMITTED|Prevents|Allows|||
|REPEATABLE READ|Prevents|Prevents|||
|SERIALIZABLE|Prevents|Prevents|||
