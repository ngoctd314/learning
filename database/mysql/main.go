package main

import (
	"fmt"
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

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
	// seed()
}

func main() {
}

// Test ...
type Test struct {
	ID      int    `db:"id,omitempty"`
	Name    string `db:"name"`
	Email   string `db:"email"`
	Age     int    `db:"age"`
	Address string `db:"address"`
}

func seed() {
	var batches [][]Test
	for i := 0; i < 1_000; i++ {
		var tmp []Test
		for j := 0; j < 1_000; j++ {
			tmp = append(tmp, Test{
				Name:    fmt.Sprintf("name_%d", i*1000+j),
				Email:   fmt.Sprintf("email_%d", i*1000+j),
				Age:     rand.Intn(30) + 10,
				Address: fmt.Sprintf("address_%d", i*1000+j),
			})
		}
		batches = append(batches, tmp)
	}

	for _, v := range batches {
		if _, err := db.NamedExec("INSERT INTO tests (id, name, email, age, address) VALUES (:id, :name, :email, :age, :address)", v); err != nil {
			log.Fatal(err)
		} else {
			log.Println("seed success")
		}
	}
}
