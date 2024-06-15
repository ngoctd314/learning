package main

import (
	"database/sql"
	"fmt"
	"log"
	"mysql/gsql"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// `User` belongs to `Company`, `CompanyID` is the foreign key
// User -> Order => foreignKey = ref model field, references: current model field
type User struct {
	ID       int
	Username string
	// Order    Order `gorm:"foreignKey:UserID;references:ID"`
	Orders []Order `gorm:"foreignKey:UserID;references:ID"`
}

func (User) TableName() string {
	return "users"
}

type Order struct {
	ID        int
	UserID    uint
	Price     int
	AddressID int
	Address   Address `gorm:"foreignKey:ID;references:AddressID"`
}

func (Order) TableName() string {
	return "orders"
}

type Address struct {
	ID   int
	Name string
}

func (Address) TableName() string {
	return "addresses"
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:secret@(192.168.49.2:30300)/learn_mysql?parseTime=true"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("CREATE TABLE users (id int auto_increment primary key, username varchar(255))")
	db.Exec("INSERT INTO users (username) VALUES ('test1')")

	db.Exec("DROP TABLE IF EXISTS orders")
	db.Exec("CREATE TABLE orders (id int auto_increment primary key, user_id int, price int, address_id int)")
	db.Exec("INSERT INTO orders (user_id, price, address_id) VALUES (1, 1, 1), (1,2,1)")

	db.Exec("DROP TABLE IF EXISTS addresses")
	db.Exec("CREATE TABLE addresses (id int auto_increment primary key, name varchar(255))")
	db.Exec("INSERT INTO addresses (name) VALUES ('address-1')")
	db.Exec("INSERT INTO addresses (name) VALUES ('address-2')")
	db.Exec("INSERT INTO addresses (name) VALUES ('address-3')")
	db.Exec("INSERT INTO addresses (name) VALUES ('address-4')")
	db.Exec("INSERT INTO addresses (name) VALUES ('address-5')")

	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
	var users []User
	// db.Model(User{}).Select("*").Joins("LEFT JOIN orders ON orders.user_id = users.id").Scan(&users).Scan(users.Orders)
	// db.Preload()

	// [0.283ms] [rows:1] SELECT * FROM `orders` WHERE id=1 AND `orders`.`user_id` = 1
	// db.Preload("Orders", func(tx *gorm.DB) *gorm.DB {
	// 	return tx.Where("id=?", 1)
	// }).Find(&users)
	// [0.125ms] [rows:1] SELECT * FROM `orders` WHERE `orders`.`user_id` = 1 AND `orders`.`id` = 1

	// db.Preload("Orders.Address").Preload("Orders", func(tx *gorm.DB) *gorm.DB { return tx.Select("price", "address_id", "user_id") }).Omit("username").Where("id=?", 1).Limit(1).Find(&users)
	condOrder := gsql.Equal("id", 1)
	builder := gsql.NewBuilder(gsql.WithPreload("Orders.Address", nil), gsql.WithPreload("Orders", condOrder))
	tx := db.Omit("username").Where("id=?", 1).Limit(1)
	tx = builder.Build(tx)
	tx.Find(&users)

	// db.Joins("Order", db.Select("price").Where(Order{ID: 1})).Find(&users)
	fmt.Println(users)
}

func readItems(tx *sql.Tx) {
	rows, err := tx.Query("SELECT id, name FROM items")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Item: ID = %d, Name = %s\n", id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func count(tx *sql.Tx) int {
	row := tx.QueryRow("SELECT COUNT(*) as cnt FROM items")

	var cnt int
	row.Scan(&cnt)

	return cnt
}

func printAlloc() {
	var m runtime.MemStats

	// Read and print memory statistics
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %d KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc: %d KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys: %d KB\n", m.Sys/1024)
	fmt.Printf("NumGC: %d\n", m.NumGC)
}
