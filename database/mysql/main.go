package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

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

type sales struct {
	OrderDate string `db:"order_date"`
}

func main() {
	seed()
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
	var in [][]sales
	for i := 0; i < 1000; i++ {
		var tmp []sales
		for i := 0; i < 1000; i++ {
			tmp = append(tmp, sales{
				OrderDate: fmt.Sprintf("2009-01-%d", rand.Intn(30)+1),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := db.NamedExec("INSERT INTO sales (order_date) VALUES (:order_date) ", v)
		if err != nil {
			log.Println(err)
		}
	}
}
