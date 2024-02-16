package main

import (
	"context"
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	conn *sqlx.DB
	ctx  = context.Background()
)

func init() {
	var err error
	// conn, err = sql.Open()
	conn, err = sqlx.Open("mysql", "root:secret@(192.168.49.2:30300)/learn_lock?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

type cart struct {
	id             int    `db:"id"`
	cartID         string `db:"cart_id"`
	cartType       int    `db:"cart_type"`
	noHoldPackages int    `db:"no_hold_packages"`
	codID          int    `db:"cod_id"`
}

func main() {
	fmt.Println(string('a' + 12))
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
