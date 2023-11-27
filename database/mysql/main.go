package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"

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
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn")
	// conn, err := sqlx.Connect("postgres", "user=admin password=secret host=192.168.49.2 port=30303 dbname=db sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	mysqlConn = conn
}

type dtText struct {
	Name string `db:"name"`
}

func fn() {
	print("fn")
}

/*
CREATE TABLE `varchar_go` (

	`name` varchar(10) DEFAULT NULL

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
func main() {
	rows, _ := mysqlConn.Query("SELECT name from varchar_go")
	defer rows.Close()
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(len(name))
	}
}

type Data struct {
	Meta []byte `db:"meta"`
}

type Meta struct {
	ID   int            `db:"id" json:"id"`
	From string         `db:"from" json:"from"`
	To   string         `db:"to" json:"to"`
	Data map[string]int `db:"data" json:"data"`
}

func seed() {
	var in [][]Data
	for i := 0; i < 1000; i++ {
		var tmp []Data
		for j := 0; j < 1000; j++ {
			meta, _ := json.Marshal(Meta{
				ID:   i*j + j,
				From: fmt.Sprintf("from-%d-%d", i, j),
				To:   fmt.Sprintf("to-%d-%d", i, j),
				Data: map[string]int{
					"rand1000": rand.Intn(1000),
					"rand2000": rand.Intn(2000),
				},
			})
			tmp = append(tmp, Data{
				Meta: meta,
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := mysqlConn.NamedExec("INSERT INTO test_jsons (meta) VALUES (:meta) ", v)
		if err != nil {
			log.Println(err)
		}
	}
}
