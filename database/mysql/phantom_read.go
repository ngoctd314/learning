package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

type phantomRead struct {
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func execPhantomRead() {
	go func() {
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		fmt.Println("Run 1")
		var list1 []phantomRead
		if err := tx.Select(&list1, "SELECT * FROM phantom_read WHERE age > 5"); err != nil {
			tx.Rollback()
			return
		}
		fmt.Println("list1", len(list1))
		tx.Commit()
	}()

	go func() {
		time.Sleep(time.Second * 2)
		tx, _ := db.BeginTxx(ctx, &sql.TxOptions{
			Isolation: sql.LevelSerializable,
		})
		if _, err := tx.Exec("INSERT INTO phantom_read VALUES (?, ?)", "name_23", 23); err != nil {
			tx.Rollback()
			return
		}
		err := tx.Commit()
		if err != nil {
			log.Println("insert error", err)
		}
		fmt.Println("insert success")
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	select {
	case <-sig:
		log.Println("Bye bye")
	}
}

func init() {
	// seedPhantomRead()
}

func seedPhantomRead() {
	// _, err := db.Exec("DROP TABLE IF EXISTS phantom_read")
	// if err != nil {
	// 	log.Println(err)
	// }
	// _, err = db.Exec("CREATE TABLE phantom_read (name VARCHAR(255), age INT)")
	// if err != nil {
	// 	log.Println(err)
	// }
	var batches [][]phantomRead
	for i := 0; i < 1_000; i++ {
		var tmp []phantomRead
		for j := 0; j < 1_000; j++ {
			tmp = append(tmp, phantomRead{
				Name: fmt.Sprintf("name_%d", i*1000+j),
				Age:  i,
			})
		}
		batches = append(batches, tmp)
	}
	for _, v := range batches {
		if _, err := db.NamedExec("INSERT INTO phantom_read (name, age) VALUES (:name, :age)", v); err != nil {
			log.Fatal(err)
		} else {
			// log.Println("seed success")
		}
	}
	log.Println("seed success")

}
