package main

import (
	"context"
	"database/sql"
	"log"
	"sync"

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
	// _, err := db.Exec("INSERT INTO stock_prices (stock_id, date) VALUES(4, '2002-05-01')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = db.Exec("INSERT INTO stock_prices (stock_id, date) VALUES(3, '2002-05-02')")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		_, err := tx.Exec("UPDATE stock_prices SET close = 10 WHERE stock_id = 4 AND date = '2002-05-01'")
		if err != nil {
			tx.Rollback()
			return
		}
		_, err = tx.Exec("UPDATE stock_prices SET close = 20 WHERE stock_id = 3 AND date = '2002-05-02'")
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	go func() {
		defer wg.Done()
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		_, err := tx.Exec("UPDATE stock_prices SET close = 21 WHERE stock_id = 3 AND date = '2002-05-02'")
		if err != nil {
			tx.Rollback()
			return
		}
		_, err = tx.Exec("UPDATE stock_prices SET close = 11 WHERE stock_id = 4 AND date = '2002-05-01'")
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()
	wg.Wait()
}

type stock struct {
	StockID int     `db:"stock_id"`
	Date    int     `db:"date"`
	Close   float64 `db:"close"`
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
