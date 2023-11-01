package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Address  string `json:"address"`
	Birthday string `json:"birthday"`
}

func main() {
	db, err := sqlx.Open("mysql", "root:secret@tcp(192.168.49.2:30300)/learn")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)

	rows, _ := db.Query("SELECT name, address FROM persons WHERE id = 4000001")
	defer rows.Close()
	for rows.Next() {
		var name, address sql.NullString
		err := rows.Scan(&name, &address)
		if err != nil {
			log.Println(err)
		}
		log.Println(name, address)
	}

	if err := rows.Err(); err != nil {
		return "", 0, err
	}

	now := time.Now()

	fmt.Printf("Since %fs", time.Since(now).Seconds())
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))

}
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
