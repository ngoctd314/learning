package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	mysqlConn *sqlx.DB
	ctx       = context.Background()
)

type nameCharTable struct {
	NameChar    sql.NullString `db:"name_char"`
	NameVarchar sql.NullString `db:"name_varchar"`
}

func init() {
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn_explain?parseTime=true")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

type Data struct {
	ID      int `db:"id"`
	OrderID int `db:"order_id"`
	UserID  int `db:"user_id"`
}

func main() {
	var in []Data
	for j := 1; j <= 1000; j++ {
		in = append(in, Data{
			OrderID: rand.Intn(100000),
			UserID:  rand.Intn(50),
		})
	}
	_, err := mysqlConn.NamedExec("INSERT INTO tbl_ref (order_id, user_id) VALUES (:order_id, :user_id) ", in)
	if err != nil {
		log.Println(err)
	}

}

func printAlloc() {
	var m runtime.MemStats

	// Read and print memory statistics
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys: %d KB\n", m.Sys/1024)
	fmt.Printf("NumGC: %d\n", m.NumGC)
}
