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
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn?parseTime=true")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

func main() {
	seed()
}

type Data struct {
	Indexed    int `db:"indexed"`
	NonIndexed int `db:"non_indexed"`
}

func seed() {
	var in [][]Data
	for i := 0; i < 100; i++ {
		var tmp []Data
		for j := 0; j < 100; j++ {
			tmp = append(tmp, Data{
				Indexed:    rand.Intn(10000),
				NonIndexed: rand.Intn(100),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := mysqlConn.NamedExec("INSERT INTO test_index (indexed, non_indexed) VALUES (:indexed, :non_indexed) ", v)
		if err != nil {
			log.Println(err)
		}
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
