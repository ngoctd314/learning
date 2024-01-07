package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"strconv"

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

type tmp1 struct {
	Data any
}

func main() {
	seed(0)
}

type Data struct {
	ID       int    `db:"id"`
	Zipcode  string `db:"zipcode"`
	Lastname string `db:"lastname"`
	Address  string `db:"address"`
}

func seed(k int) {
	var in []Data
	for i := (1 + 1000*k); i <= (1000 + 1000*k); i++ {
		in = append(in, Data{
			Zipcode:  strconv.Itoa(i * 10000),
			Lastname: fmt.Sprintf("%detrunia%d", i, i),
			Address:  fmt.Sprintf("%dMain Street%d", i, i),
		})
	}

	_, err := mysqlConn.NamedExec("INSERT INTO people (zipcode, lastname, address) VALUES (:zipcode, :lastname, :address) ", in)
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
