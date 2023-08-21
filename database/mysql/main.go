package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var ctx = context.Background()

// User ...
type User struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func init() {
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

func main() {
	now := time.Now()
	_, _ = db.Exec("SELECT count(name) FROM persons")
	fmt.Printf("since: %s\n", time.Since(now))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		log.Println("exit")
	}
}

type person struct {
	Name     string `db:"name"`
	Age      int8   `db:"age"`
	Address  string `db:"address"`
	Birthday string `db:"birthday"`
}

func seed() {
	var in [][]person
	for i := 0; i < 2000; i++ {
		var tmp []person
		for j := 0; j < 2000; j++ {
			tmp = append(tmp, person{
				Name:     fmt.Sprintf("name-%d-%d", i, j),
				Age:      int8(rand.Intn(100)),
				Address:  fmt.Sprintf("address-%d", rand.Intn(1000)),
				Birthday: fmt.Sprintf("birthday-%d", rand.Intn(365*100)),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := db.NamedExec("INSERT INTO persons (name, age, address, birthday) VALUES (:name, :age, :address, :birthday) ", v)
		if err != nil {
			log.Println(err)
		}
	}
}
