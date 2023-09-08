package main

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	ID        uint64    `gorm:"id,primaryKey"`
	Code      string    `gorm:"code"`
	Price     uint64    `gorm:"price"`
	CreatedAt time.Time `gorm:"created_at"`
	UpdatedAt time.Time `gorm:"updated_at"`
}

type Person struct {
	ID        int `gorm:"primaryKey"`
	Age       int
	Name      string
	Birthday  *time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Person) TableName() string {
	return "persons"
}

func (p Person) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Age < 18 {
		return errors.New("child is not allow")
	}

	return nil
}

func main() {
	db, err := gorm.Open(mysql.Open("root:secret@tcp(192.168.49.2:30300)/gorm?parseTime=true"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db = db.Session(&gorm.Session{
		SkipHooks: false,
	})

	if err != nil {
		return
	}
	var tx *gorm.DB
	person := Person{
		ID: 3,
	}
	tx = db.Find(&person)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		log.Fatal(tx.Error)
	}
	log.Println(tx.RowsAffected, person)
}
