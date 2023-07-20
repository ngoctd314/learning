package main

import (
	"context"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var ctx = context.Background()

// User ...
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func init() {
	conn, err := sqlx.Connect("mysql", "admin:secret@(192.168.49.2:30300)/db")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

func main() {
	// execPhantomRead()
	seedDeadlock()
	execDeadLock()
}

// Test ...
type Test struct {
	ID          int    `db:"id,omitempty"`
	Name        string `db:"name"`
	NameNoIndex string `db:"name_no_index"`
	Age         int    `db:"age"`
}

type User1 struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (User1) Table() string {
	return "user1"
}

type User2 struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (User2) Table() string {
	return "user2"
}

type User3 struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (User3) Table() string {
	return "user3"
}

func seed() {
	db.Exec("DROP TABLE IF EXISTS user1")
	db.Exec("DROP TABLE IF EXISTS user2")
	db.Exec("DROP TABLE IF EXISTS user3")

	db.Exec("CREATE TABLE user1 (name VARCHAR(255), age INT)")
	db.Exec("CREATE TABLE user2 (name VARCHAR(255), age INT)")
	db.Exec("CREATE TABLE user3 (name VARCHAR(255), age INT)")

	db.Exec("INSERT INTO user1 VALUES ('admin', 18)")
	db.Exec("INSERT INTO user2 VALUES ('admin', 18)")
	db.Exec("INSERT INTO user3 VALUES ('admin', 18)")
}
