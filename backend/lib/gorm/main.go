package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := db.Clauses(clause.OnConflict{DoUpdates: clause.Assignments(
		map[string]interface{}{
			"c": 12,
		})},
	).Create(&Tbl{A: 2, B: 1, C: 10}).Error; err != nil {
		log.Println(err)
	}

}

type Tbl struct {
	ID int `gorm:"primaryKey"`
	A  int
	B  int
	C  int
}

func (Tbl) TableName() string {
	return "tbl"
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
