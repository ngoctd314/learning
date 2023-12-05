package main

import (
	"context"
	"database/sql"
	"log"
	"math/rand"
	"mysql/srv1"

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
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

func main() {
	seed()
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
