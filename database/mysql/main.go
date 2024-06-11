package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	ctx = context.Background()
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:secret@(192.168.49.2:30300)/learn_mysql?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db.Exec("DROP TABLE IF EXISTS isolations")
	db.Exec("CREATE TABLE isolations (id int auto_increment primary key, network_status int, volume_status int)")
	db.Exec("INSERT INTO isolations (id, network_status, volume_status) VALUES (1, 0, 0)")

	txA, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	txA.Exec("UPDATE isolations SET network_status = 1 WHERE id = 1")

	txB, _ := db.BeginTx(context.Background(), &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	txB.Exec("UPDATE isolations SET network_status = 1 WHERE id = 1")
	txB.Commit()

	row := txA.QueryRow("SELECT network_status, volume_status FROM isolations WHERE id = 1")
	var networkStatus, volumeStatus int
	row.Scan(&networkStatus, &volumeStatus)
	log.Printf("txA: network %d, volume: %d", networkStatus, volumeStatus)

	txA.Exec("UPDATE isolations SET network_status = 1 WHERE id = 1")

	txA.Commit()
}

func readItems(tx *sql.Tx) {
	rows, err := tx.Query("SELECT id, name FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Item: ID = %d, Name = %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func count(tx *sql.Tx) int {
	row := tx.QueryRow("SELECT COUNT(*) as cnt FROM items")

	var cnt int
	row.Scan(&cnt)

	return cnt
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
