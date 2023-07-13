package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func main() {
	conn, err := sqlx.Connect("mysql", "admin:secret@(192.168.49.2:30300)/db")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	db.Exec("DELETE FROM users")
	db.Exec("INSERT INTO users VALUES(?, ?)", "ngoctd", 0)

	ctx := context.Background()
	_ = ctx

	wg := sync.WaitGroup{}
	updateUser := make(chan struct{}, 1)

	go func() {
		wg.Add(1)
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
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
			Isolation: sql.LevelSerializable,
		})
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

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("Bye")
	}
}
