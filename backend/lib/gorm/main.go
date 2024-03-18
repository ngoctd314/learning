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

	u := []User{}
	db.Select("users.id", "address.id AS Address__id").
		Joins("LEFT JOIN address ON address.id = users.address_id").
		Joins("INNER JOIN address ON address.id = users.address_id").
		Find(&u, "name = ?", "test1")
	for _, v := range u {
		fmt.Println(v.Address)
	}
}

type User struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"default:abc"`
	Email      string `gorm:"embedded"`
	AddressID  int
	AddressID1 int
	Address    *Address `gorm:"foreignKey:AddressID"`
	Address1   *Address `gorm:"foreignKey:AddressID1"`
}

type Address struct {
	ID   int
	City string
}

func (Address) TableName() string {
	return "address"
}
