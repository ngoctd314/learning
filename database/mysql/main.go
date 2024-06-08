package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"log/slog"
	"runtime"
	"sync"

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
	waitForInsert := make(chan struct{}, 1)
	waitForSelect := make(chan struct{}, 1)
	waitForInsertCommit := make(chan struct{}, 1)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()

		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelReadCommitted,
		})
		<-waitForSelect
		{
			rs, _ := tx.Exec("INSERT INTO persons (name) VALUES (?)", "test3")
			slog.Info("exec: INSERT persons test1", "results", rs)

			rs, _ = tx.Exec("INSERT INTO persons (name) VALUES (?)", "test4")
			slog.Info("exec: INSERT persons test2", "results", rs)
		}
		waitForInsert <- struct{}{}
		<-waitForSelect
		tx.Commit()
		waitForInsertCommit <- struct{}{}
		slog.Info("COMMIT INSERT")
	}()
	go func() {
		defer wg.Done()

		tx, _ := conn.BeginTx(ctx, &sql.TxOptions{
			Isolation: sql.LevelRepeatableRead,
		})
		{
			rows, _ := tx.Query("SELECT name FROM persons WHERE id >=1")
			for rows.Next() {
				var name string
				_ = rows.Scan(&name)
				slog.Info("result 1: SELECT name", "value", name)
			}
		}
		waitForSelect <- struct{}{}

		<-waitForInsert
		{
			rows, _ := tx.Query("SELECT name FROM persons WHERE id >=1")
			for rows.Next() {
				var name string
				_ = rows.Scan(&name)
				slog.Info("result 1: SELECT name", "value", name)
			}
		}
		waitForSelect <- struct{}{}
		<-waitForInsertCommit
		{
			rows, _ := tx.Query("SELECT name FROM persons WHERE id >=1")
			for rows.Next() {
				var name string
				_ = rows.Scan(&name)
				slog.Info("result 2: SELECT name", "value", name)
			}
		}
		tx.Commit()
		slog.Info("COMMIT SELECT")
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
