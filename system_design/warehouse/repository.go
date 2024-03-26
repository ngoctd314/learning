package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/RoaringBitmap/roaring"
	_ "github.com/go-sql-driver/mysql"
)

type repoV3 struct {
	db        *sql.DB
	stmtChunk *sql.Stmt
	chunk     int
}

func newRepoV3() *repoV3 {
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

	return &repoV3{
		db:        db,
		stmtChunk: stmt,
		chunk:     chunk,
	}
}

func repov3Insert() {
	// insert 2K record
	// 1e6 => 1M
	records := 1
	total := 10
	chunk := total / records

	// r := newRepoV3()
	// r.truncate()

	for i := 0; i < chunk; i++ {
		lRecords := make([]any, records)
		for j := 0; j < records; j++ {
			rb := roaring.New()
			relateItemList := make([]uint32, 5000)
			for k := 0; k < 5000; k++ {
				relateItemList[k] = uint32(rand.Intn(1e8) + 1)
			}
			rb.AddMany(relateItemList)
			b, err := rb.ToBytes()
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("HasRunCompression", rb.HasRunCompression())
			lRecords[j] = b
		}
		// r.insertMany(lRecords)
	}

	// f, err := os.Open("data.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()
	//
	// data, _ := io.ReadAll(f)
	// strs := strings.Split(string(data), "\n")
	//
	// mt := make([][]uint32, 0)
	// for _, a := range strs {
	// 	vs := strings.Split(a, " ")
	// 	tmp := make([]uint32, 0)
	// 	for _, v := range vs {
	// 		vi, convErr := strconv.Atoi(v)
	// 		if convErr == nil {
	// 			tmp = append(tmp, uint32(vi))
	// 		}
	// 	}
	// 	mt = append(mt, tmp)
	// }

	// for i, v := range mt {
	// 	rb := roaring.New()
	// 	rb.AddMany(v)
	// 	b, err := rb.ToBytes()
	// 	if err != nil {
	// 		log.Fatalf("error occur when rb.ToBytes (%v)", err)
	// 	}
	// 	r.insert(uint32(i+1), b)
	// }
}
func repov3Count() {
	r := newRepoV3()
	s := make(map[uint32]struct{})
	items, j := 1000, 0
	for j < items {
		// for i := 1; i <= 100; i++ {
		n := uint32(rand.Intn(1e7))
		for n == 0 {
			n = uint32(rand.Intn(1e7))
		}
		s[n] = struct{}{}
		j++
		// s[uint32(j)] = struct{}{}
	}
	paramsInt := make([]uint32, 0, items)
	for k := range s {
		paramsInt = append(paramsInt, k)
	}
	sort.Slice(paramsInt, func(i, j int) bool { return paramsInt[i] < paramsInt[j] })

	params := make([]any, 0, items)
	for _, k := range paramsInt {
		params = append(params, k)
	}

	r.countDistinctRelate(params...)
	// r.countDistinctRelate(params...)
}

func (r *repoV3) insert(itemID uint32, relate []byte) error {
	q := "INSERT INTO relate (id, bitmap) VALUES (?, ?)"
	rs, err := r.db.Exec(q, itemID, relate)
	if err != nil {
		log.Printf("error when insert %v\n", err)
		return err
	}
	_ = rs

	return nil
}

func (r *repoV3) insertMany(relates []any) {
	q := fmt.Sprintf("INSERT INTO relate (bitmap) VALUES (?)%s", strings.Repeat(",(?)", len(relates)-1))
	_, err := r.db.Exec(q, relates...)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func (r *repoV3) truncate() {
	// r.db.Exec("TRUNCATE TABLE relate")
}

func (r *repoV3) countDistinctRelate(params ...any) int {
	totalTime := time.Now()

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

			ti := time.Now()
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
			fmt.Println("scan: ", time.Since(ti))

			listBitmap := make([]*roaring.Bitmap, 0, upper-lower)
			// ti1 := time.Now()
			for j := 0; j < len(listData); j++ {
				newrb := roaring.New()
				_, err := newrb.FromBuffer(listData[j])
				if err != nil {
					log.Printf("error occur when ReadFrom bitmap (%v)", err)
					return
				}

				listBitmap = append(listBitmap, newrb)
			}
			// println("FromBuffer ", time.Since(ti1).Seconds(), len(listData[0]))

			mu.Lock()
			list = append(list, listBitmap...)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	ti := time.Now()
	rs := roaring.ParOr(6, list...)
	fmt.Println("ParOr", time.Since(ti))
	fmt.Println("result: ", rs.GetCardinality())

	fmt.Println("total time: ", time.Since(totalTime))
	return 0
}
