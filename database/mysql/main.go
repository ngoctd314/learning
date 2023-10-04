package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db  *sqlx.DB
	ctx = context.Background()
)

type Shop struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	JoinDate    time.Time `db:"join_date"`
	JoinChannel string    `db:"join_channel"`
}

// User ...
type User struct {
	Name string `json:"name" db:"name"`
	Age  int    `db:"age" json:"age"`
}

func init() {
	conn, err := sqlx.Connect("mysql", "root:secret@(192.168.49.2:30300)/mysql")
	if err != nil {
		log.Fatal(err)
	}
	db = conn
}

type Order struct {
	ID      int `db:"id"`
	Product int `db:"product"`
}

type Bills struct {
	ID      int `db:"id"`
	OrderID int `db:"order_id"`
	UserID  int `db:"user_id"`
}

func main() {
	seedPerson()
}

type person struct {
	Name     string `db:"name"`
	Age      int8   `db:"age"`
	Address  string `db:"address"`
	Birthday string `db:"birthday"`
}

var e = [4]string{"facebook", "email", "ads", "cs_refer"}

func seedPerson() {
	var in [][]Shop
	for i := 0; i < 2000; i++ {
		var tmp []Shop
		for j := 0; j < 2000; j++ {
			tmp = append(tmp, Shop{
				Name:        fmt.Sprintf("name-%d-%d", i, j),
				JoinDate:    time.Now().Add(-1 * time.Minute * time.Duration(i+j)),
				JoinChannel: e[rand.Intn(4)],
			})
		}
		in = append(in, tmp)
	}

	for _, v := range in {
		_, err := db.NamedExec("INSERT INTO shops (name, join_date, join_channel) VALUES (:name, :join_date, :join_channel) ", v)
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
