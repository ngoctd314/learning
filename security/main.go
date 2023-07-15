package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"security/sqlinjection"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	ctx context.Context
)

func init() {
	conn, err := sqlx.Connect("mysql", "admin:secret@(192.168.49.2:30300)/db")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	db.Exec("DELETE FROM users")
	ctx = context.Background()
}

func main() {
	go sqlinjection.APIServer(db)

	seed()
	// time.Sleep(time.Second)
	// sqlinjection.APICaller()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		fmt.Println("Bye")
	}
}

// User model
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func seed() {
	for i := 0; i < 10; i++ {
		db.Exec("INSERT INTO users VALUES(?, ?)", fmt.Sprintf("name_%d", i), i+10)
	}
	db.Exec("INSERT INTO users VALUES(?, ?)", "ngoctd OR '1' = '1'", 23)
}
