package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
)

func execDeadLock() {
	go func() {
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if _, err := tx.Exec("UPDATE dead_locks SET name = ? WHERE age = ?", "admin18", 18); err != nil {
			tx.Rollback()
			return
		}
		if _, err := tx.Exec("UPDATE dead_locks SET name = ? WHERE age = ?", "admin19", 19); err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		fmt.Println("update success")

	}()
	go func() {
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if _, err := tx.Exec("UPDATE dead_locks SET name = ? WHERE age = ?", "admin19", 19); err != nil {
			tx.Rollback()
			return
		}
		if _, err := tx.Exec("UPDATE dead_locks SET name = ? WHERE age = ?", "admin18", 18); err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		fmt.Println("update success")
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("Bye bye")
	}
}

type deadLock struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func seedDeadlock() {
	db.Exec("DROP TABLE IF EXISTS dead_locks")
	db.Exec("CREATE TABLE dead_locks (name VARCHAR(255), age INT)")
	tmp := deadLock{
		Name: "admin",
		Age:  18,
	}
	db.Exec("INSERT INTO dead_locks (name, age) VALUES (:name, :age)", tmp)
	tmp.Age = 19
	db.Exec("INSERT INTO dead_locks (name, age) VALUES (:name, :age)", tmp)
}
