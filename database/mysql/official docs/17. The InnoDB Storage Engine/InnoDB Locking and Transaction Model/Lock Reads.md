# Locking Reads

If you query data and then insert or update related data within the same transaction, the regular SELECT statement does not give enough protection. Other transactions can update or delete the same rows you just queried. InnoDB support two types of locking reads that offer extra safety:

## SELECT ... FOR SHARE

Set a shared mode lock on any rows that are read. Other sessions can read the rows, but cannot modify them until your transaction commits. If any of these rows were changed by another transaction that has not yet committed, your query waits until your transactions commits. If any of these rows were changed by another transaction that has not yet committed, your query waits until that transaction ends and then uses the latest values.

```go
func main() {
	db.Exec("DROP TABLE IF EXISTS test")
	db.Exec("CREATE TABLE test (id int auto_increment primary key, a int)")
	db.Exec("INSERT INTO test (id, a) VALUES (1, 0)")

	txA, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})

	txA.Exec("UPDATE test SET a = 1 WHERE id = 1")

	go func() {
		now := time.Now()
		fmt.Println("block txB")
		txB, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelSerializable})
		row := txB.QueryRow("SELECT a FROM test WHERE id = 1")
		var b int
		row.Scan(&b)
		fmt.Printf("txB: a: %d, since %d ms\n", b, time.Since(now).Milliseconds())
		txB.Commit()
	}()

	time.Sleep(time.Second * 2)
	txA.Commit()
	select {}
}
```

SELECT ... FOR SHARE required the SELECT privilege.   

SELECT ... FOR SHARE statements do not acquire read locks on MySQL grant tables. 

## SELECT FOR UPDATE
