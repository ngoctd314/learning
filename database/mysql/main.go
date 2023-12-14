package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"mysql/srv1"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	mysqlConn *sqlx.DB
	ctx       = context.Background()
)

type nameCharTable struct {
	NameChar    sql.NullString `db:"name_char"`
	NameVarchar sql.NullString `db:"name_varchar"`
}

func init() {
	srv1.Fn()
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/auth?parseTime=true")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

type Session struct {
	ID            string    `gorm:"column:id;primaryKey" json:"id" db:"id"`
	SessionData   string    `gorm:"column:session_data" json:"session_data" db:"session_data"`
	AccountID     string    `gorm:"column:account_id" json:"account_id" db:"account_id"`
	SessionSource string    `gorm:"column:session_source" json:"session_source" db:"session_source"`
	CustomInfo    string    `gorm:"column:custom_info" json:"custom_info" db:"custom_info"`
	Status        int       `gorm:"column:status" json:"status" db:"status"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"updated_at" db:"updated_at"`
	ExpiresOn     time.Time `gorm:"column:expires_on" json:"expires_on" db:"expires_on"`
}

func main() {
	var session []Session
	err := mysqlConn.Select(&session, "SELECT * FROM sessions")
	fmt.Println(err, len(session))
	runtime.KeepAlive(session)
	printAlloc()

	// 2882 KB / 1206 sessions
	// 394  KB / 1206 session
}

type Data struct {
	StoreID int `db:"store_id"`
	FilmID  int `db:"film_id"`
}

func seed() {
	var in [][]Data
	for i := 0; i < 100; i++ {
		var tmp []Data
		for j := 0; j < 100; j++ {
			tmp = append(tmp, Data{
				StoreID: rand.Intn(10000),
				FilmID:  rand.Intn(10000),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := mysqlConn.NamedExec("INSERT INTO inventory (store_id, film_id) VALUES (:store_id, :film_id) ", v)
		if err != nil {
			log.Println(err)
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
