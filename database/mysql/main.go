package main

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// `User` belongs to `Company`, `CompanyID` is the foreign key
// User -> Order => foreignKey = ref model field, references: current model field
type User struct {
	ID int
	// Username string
	// Order    Order `gorm:"foreignKey:UserID;references:ID"`
	// Orders []Order `gorm:"foreignKey:UserID;references:ID"`
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID;references:ID"`
}

func (User) TableName() string {
	return "users"
}

type Company struct {
	ID   int
	Name string
}

func (Company) TableName() string {
	return "companies"
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

// func main() {
// 	db, _ := gorm.Open(mysql.Open("root:secret@(192.168.49.2:30300)/learn_mysql?parseTime=true"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
// 	db.Exec("DROP TABLE IF EXISTS users")
// 	db.Exec("CREATE TABLE users (id int auto_increment primary key, company_id int)")
// 	db.Exec("INSERT INTO users (company_id) VALUES (1)")
//
// 	db.Exec("DROP TABLE IF EXISTS companies")
// 	db.Exec("CREATE TABLE companies (id int auto_increment primary key, name varchar(255))")
// 	db.Exec("INSERT INTO companies (id, name) VALUES (1, 'company-1'), (2,'company-2')")
//
// 	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
// 	var users []User
// 	db.Preload("Company").Find(&users)
//
// 	// db.Joins("Order", db.Select("price").Where(Order{ID: 1})).Find(&users)
// 	fmt.Println(users)
// }

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

type HTTPError struct{}

func (*HTTPError) Error() string {
	return "error"
}

func create() *HTTPError {
	return nil
}

func main() {
	var err error
	err = create()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("never run")
}
