package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	ctx = context.Background()
)

// User ...
type User struct {
	Name string `json:"name" db:"name"`
	Age  int    `db:"age" json:"age"`
}

func init() {
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/learn")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	// seedPeople()
	// seedTestHash()
}

func main() {
	var rs []people
	query := "SELECT * FROM people LIMIT 4000000"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var tmp people

		if err := rows.Scan(&tmp.LastName, &tmp.FirstName, &tmp.Dob, &tmp.Gender); err != nil {
			log.Println(err)
		}

		rs = append(rs, tmp)
	}
	for _, v := range rs {
		_ = v
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		log.Println("exit")
	}
}

func seedCityDemo() {}

type person struct {
	Name     string `db:"name"`
	Age      int8   `db:"age"`
	Address  string `db:"address"`
	Birthday string `db:"birthday"`
}

func seedPerson() {
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

type gender string

const (
	male   gender = "m"
	female gender = "f"
)

type people struct {
	LastName  string `db:"last_name"`
	FirstName string `db:"first_name"`
	Dob       string `db:"dob"`
	Gender    gender `db:"gender"`
}

func seedPeople() {
	var in [][]people
	for i := 0; i < 2000; i++ {
		var tmp []people
		for j := 0; j < 2000; j++ {
			tmp = append(tmp, people{
				LastName:  fmt.Sprintf("last_name-%d-%d", i, j),
				FirstName: fmt.Sprintf("first_name-%d-%d", i, j),
				// Dob:       time.Now().Add(-time.Hour * 24 * 365 * time.Duration(rand.Intn(100)+1)),
				Gender: func() gender {
					if j%2 == 0 {
						return male
					}
					return female
				}(),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := db.NamedExec(`INSERT INTO people (last_name, first_name, dob, gender) 
         VALUES (:last_name, :first_name, :dob, :gender)`, v)
		if err != nil {
			log.Println(err)
		}
	}
}

type testhash struct {
	FName string `db:"fname"`
	LNAME string `db:"lname"`
}

func seedTestHash() {
	var in [][]testhash

	for i := 0; i < 2000; i++ {
		var tmp []testhash
		for j := 0; j < 2000; j++ {
			tmp = append(tmp, testhash{
				FName: fmt.Sprintf("fname-%d-%d", i, j),
				LNAME: fmt.Sprintf("lname-%d-%d", i, j),
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := db.NamedExec(`INSERT INTO testhash (fname, lname) 
         VALUES (:fname, :lname)`, v)
		if err != nil {
			log.Println(err)
		}
	}
}

func add(a, b int) int {
	return a + b
}
