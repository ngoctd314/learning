package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/RoaringBitmap/roaring"
	_ "github.com/go-sql-driver/mysql"
)

type repository struct {
	db        *sql.DB
	stmtChunk *sql.Stmt
	chunk     int
}

func newRepostiory() *repository {
	db, err := sql.Open("mysql", "root:secret@(192.168.49.2:30300)/warehouse?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(200)
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("ping error", pingErr)
	}
	chunk := 100
	q := fmt.Sprintf("SELECT bitmap FROM relate WHERE id IN (?%s)", strings.Repeat(",?", chunk-1))
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal("stmt error", err)
	}

	return &repository{
		db:        db,
		stmtChunk: stmt,
		chunk:     chunk,
	}
}

func (r *repository) getOrBitmap(params []any) []byte {
	list := make([]*roaring.Bitmap, 0, len(params))
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	n := len(params)/r.chunk + 1
	if len(params) == r.chunk || len(params)%r.chunk == 0 {
		n--
	}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			lower, upper := i*r.chunk, (i+1)*r.chunk
			if i+1 == n {
				upper = len(params)
			}

			var (
				rows     *sql.Rows
				queryErr error
			)
			if upper-lower == r.chunk {
				rows, queryErr = r.stmtChunk.Query(params[lower:upper]...)
			} else {
				q := fmt.Sprintf("SELECT bitmap FROM relate WHERE id IN (?%s)", strings.Repeat(",?", upper-lower-1))
				rows, queryErr = r.db.Query(q, params[lower:upper]...)
			}
			if queryErr != nil {
				log.Printf("error occur when SELECT bitmap (%v)", queryErr)
				return
			}

			listData := make([][]byte, 0, upper-lower)
			for rows.Next() {
				// var bitmap sql.RawBytes
				var bitmap []byte
				if err := rows.Scan(&bitmap); err != nil {
					log.Printf("error occur when Scan bitmap (%v)", err)
					return
				}
				listData = append(listData, bitmap)
			}
			rows.Close()
			fmt.Println("len listData", len(listData))

			listBitmap := make([]*roaring.Bitmap, 0, upper-lower)
			for j := 0; j < len(listData); j++ {
				newrb := roaring.New()
				_, err := newrb.FromBuffer(listData[j])
				if err != nil {
					log.Printf("error occur when ReadFrom bitmap (%v)", err)
					return
				}

				listBitmap = append(listBitmap, newrb)
			}

			mu.Lock()
			list = append(list, listBitmap...)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	result, err := roaring.ParOr(6, list...).ToBytes()
	if err != nil {
		log.Println(err)
	}
	return result
}
