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
)

var (
	conn *sql.DB
	ctx  = context.Background()
)

func init() {
	var err error
	conn, err = sql.Open("mysql", "root:secret@(192.168.49.2:30300)/learn_lock?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	guard1, guard2 := make(chan struct{}, 1), make(chan struct{}, 1)

	releaseProcess1 := func() {
		guard1 <- struct{}{}
	}
	blockProcess1 := func() {
		<-guard1
	}
	releaseProcess2 := func() {
		guard2 <- struct{}{}
	}
	blockProcess2 := func() {
		<-guard2
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// process 1
	go func() {
		blockProcess1()

		defer wg.Done()
		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
		_, err := tx.Exec("INSERT INTO tbl (created_at) VALUES (?)", time.Now())
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return
		}
		tx.Commit()

		releaseProcess2()
	}()

	// process 2
	go func() {
		defer wg.Done()
		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
		rows, err := tx.Query("SELECT created_at FROM tbl")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return
		}
		for rows.Next() {
			var dest time.Time
			if err := rows.Scan(&dest); err != nil {
				log.Println(err)
			}
			fmt.Println("phase 1", dest)
		}
		releaseProcess1()

		blockProcess2()
		rows, err = tx.Query("SELECT created_at FROM tbl")
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return
		}
		for rows.Next() {
			var dest time.Time
			if err := rows.Scan(&dest); err != nil {
				log.Println(err)
			}
			fmt.Println("phase 2", dest)
		}

		tx.Commit()
	}()

	wg.Wait()

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
