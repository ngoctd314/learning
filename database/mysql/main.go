package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	db, _ := gorm.Open(mysql.Open("root:secret@(192.168.49.2:30300)/learn_mysql?parseTime=true"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db.Exec("DROP TABLE IF EXISTS test1")
	db.Exec("DROP TABLE IF EXISTS test2")
	db.Exec("DROP TABLE IF EXISTS test3")
	db.Exec("CREATE TABLE test1 (id int auto_increment primary key, name varchar(255))")
	db.Exec("CREATE TABLE test2 (id int auto_increment primary key, name varchar(255))")
	db.Exec("CREATE TABLE test3 (id int auto_increment primary key, name varchar(255))")
	db.Exec("INSERT INTO test1 (name) VALUES ('name1.0'),('name1.1'),('name1.2')")
	db.Exec("INSERT INTO test2 (name) VALUES ('name2.0'),('name2.1'),('name2.2')")
	db.Exec("INSERT INTO test3 (name) VALUES ('name3.0'),('name3.1'),('name3.2')")

	now := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		tx := db.WithContext(context.Background())
		tx = tx.Begin()
		fmt.Println(48)
		tx.Exec("SELECT * FROM test1 WHERE id = 1 FOR UPDATE")
		time.Sleep(time.Second * 2)
		tx.Exec("UPDATE test1 SET name = 'name1.0' WHERE id = 1")
		fmt.Println(52)
		tx.Commit()
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		defer wg.Done()
		tx := db.WithContext(context.Background())
		tx = tx.Begin()
		fmt.Println(58)
		tx.Exec("SELECT * FROM test1 WHERE id = 1 FOR UPDATE")
		fmt.Println(60)
		// time.Sleep(time.Second)
		// tx.Exec("UPDATE test1 SET name = 'name1.0' WHERE id = 1")
		tx.Commit()
	}()

	wg.Wait()
	fmt.Printf("since %dms", time.Since(now).Milliseconds())
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
