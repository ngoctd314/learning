package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/RoaringBitmap/roaring"
	_ "github.com/go-sql-driver/mysql"
)

type Repository struct {
	db        *sql.DB
	stmtChunk *sql.Stmt
	chunk     int
}

func NewRepository() *Repository {
	user := os.Getenv("MYSQL_USER")
	passwd := os.Getenv("MYSQL_PASSWD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=true", user, passwd, host, port, dbName))
	if err != nil {
		log.Fatal(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatal("ping error", pingErr)
	}
	chunk := 100
	q := fmt.Sprintf("SELECT bitmap FROM relate WHERE id IN (?%s)", strings.Repeat(",?", chunk-1))
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal("stmt error", err)
	}

	return &Repository{
		db:        db,
		stmtChunk: stmt,
		chunk:     chunk,
	}
}

func (r *Repository) Insert(itemID uint32, relate []byte) error {
	q := "INSERT INTO relate (id, bitmap) VALUES (?, ?)"
	rs, err := r.db.Exec(q, itemID, relate)
	if err != nil {
		log.Printf("error when insert %v\n", err)
		return err
	}
	_ = rs

	return nil
}

func (r *Repository) InsertMany(relates []any) {
	q := fmt.Sprintf("INSERT INTO relate (bitmap) VALUES (?)%s", strings.Repeat(",(?)", len(relates)-1))
	_, err := r.db.Exec(q, relates...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (r *Repository) CountDistinctRelate(params ...any) uint64 {
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

	rs := roaring.ParOr(6, list...)

	return rs.GetCardinality()
}
