package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	// dsn := "dev_ngoctd10:ZmmtijIZmuhepN7d6tQXYZpUlpcW@tcp(10.110.69.70:3306)/goku_be?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:secret@tcp(192.168.49.2:30300)/learn_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}

	// var result result
	name := "test3"
	u := User{Name: name, Email: "test@gmail.com", Cnt: 0}
	if err := db.FirstOrCreate(&u, "name = ?", name).Error; err != nil {
		log.Fatal(err)
	}
}

type User struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"embedded"`
	Cnt   int
}
