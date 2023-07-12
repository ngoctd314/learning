package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

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

	ctx := context.Background()
	_ = ctx

	trigger := make(chan struct{}, 1)
	commitInsert := make(chan struct{}, 1)

	go func() {
		<-trigger
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

		commitInsert <- struct{}{}
		log.Println("list user read commited", listUser)

		if err := tx.Commit(); err != nil {
			log.Println("commit error", err)
		}

		log.Println(listUser)
	}()

	go func() {
		listUser := []User{{Name: fmt.Sprintf("name_%s", time.Now().Format(time.DateTime)), Age: 2023}}
		tx, err := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadUncommitted,
			ReadOnly:  false,
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
		<-commitInsert
		if err := tx.Commit(); err != nil {
			log.Println("commit error", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("Bye")
	}

}
