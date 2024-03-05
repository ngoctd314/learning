package main

import (
	"fmt"
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
	var user User
	db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&user)
	fmt.Println(user)
	db.Model(nil).Pluck()
}

type User struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Email Email `gorm:"embedded"`
}

type Email struct {
	ID     int `gorm:"primaryKey"`
	UserID int
	Email  string
}

type result struct {
	Name  string
	Email string
}
