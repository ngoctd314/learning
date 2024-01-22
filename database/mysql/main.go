package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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
	var listUUID []string
	for i := 0; i < 100; i++ {
		listUUID = append(listUUID, uuid.NewString())
	}

	for i := 0; i < 5000; i++ {
		data := []map[string]any{}
		for j := 0; j < 1000; j++ {
			data = append(data, map[string]any{
				"cart_id":          listUUID[rand.Intn(100)],
				"cart_type":        rand.Intn(5),
				"no_hold_packages": rand.Intn(10000),
				"cod_id":           i*1000 + j + 1000000,
			})
		}
		_, err := conn.NamedExec("INSERT INTO carts (cart_id, cart_type, no_hold_packages, cod_id) VALUES (:cart_id, :cart_type, :no_hold_packages, :cod_id)", data)
		if err != nil {
			log.Fatal(err)
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
