package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// User ...
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
	for i := 0; i < 10; i++ {
		db.Exec("INSERT INTO users VALUES(?, ?)", "ngoctd", i)
	}

	ctx := context.Background()
	_ = ctx

	wg := sync.WaitGroup{}
	updateUser := make(chan struct{}, 1)
	_ = updateUser

	go func() {
		wg.Add(1)
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			log.Fatal(err)
		}

		query := ""
		now := time.Now()
		fmt.Println("INSERT users1")
		query = "INSERT INTO users VALUES (?, ?)"
		if _, err := tx.Exec(query, "ngoctd", 0); err != nil {
			fmt.Println("UPDATE error", err)
			if err := tx.Rollback(); err != nil {
				fmt.Println("rollback error", err)
			}
			return
		}
		time.Sleep(time.Second)

		fmt.Println("INSERT user_tests1 after", time.Since(now).Seconds())
		query = "INSERT INTO users VALUES (?, ?)"
		query = "INSERT INTO user_tests VALUES (?, ?)"
		if _, err := tx.Exec(query, "ngoctd", 0); err != nil {
			fmt.Println("UPDATE error", err)
			if err := tx.Rollback(); err != nil {
				fmt.Println("rollback error", err)
			}
			return
		}

		fmt.Println("commit insert user")
		if err := tx.Commit(); err != nil {
			fmt.Println("commit error", err)
		}
		// updateUser <- struct{}{}
	}()

	go func() {
		wg.Add(1)
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if err != nil {
			log.Fatal(err)
		}

		query := ""

		fmt.Println("INSERT user_tests2")
		query = "INSERT INTO user_tests VALUES (?, ?)"
		if _, err := tx.Exec(query, "ngoctd", 0); err != nil {
			fmt.Println("UPDATE error", err)
			if err := tx.Rollback(); err != nil {
				fmt.Println("rollback error", err)
			}
			return
		}

		time.Sleep(time.Second * 2)
		fmt.Println("INSERT users2")
		query = "INSERT INTO users VALUES (?, ?)"
		if _, err := tx.Exec(query, "ngoctd", 0); err != nil {
			fmt.Println("UPDATE error", err)
			if err := tx.Rollback(); err != nil {
				fmt.Println("rollback error", err)
			}
			return
		}

		fmt.Println("commit insert user")
		if err := tx.Commit(); err != nil {
			fmt.Println("commit error", err)
		}
		// updateUser <- struct{}{}
	}()

	wg.Wait()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("Bye")
	}
}
