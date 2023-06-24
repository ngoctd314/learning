package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	crud()
}
func crud() {
	dsn := "ghtk:secret@tcp(192.168.49.2:30200)/iam_db?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	r := repository{db}
	// batchSize := 100
	// listUserName := make([]AccountUsernameModel, 0, batchSize)
	// for i := 0; i < batchSize; i++ {
	// 	username := AccountUsernameModel{
	// 		AccountID:   strconv.FormatInt(time.Now().Unix(), 10),
	// 		Username:    fmt.Sprintf("ngoctd_%d_%d", i, time.Now().UnixNano()),
	// 		AccountType: "staff",
	// 	}
	// 	listUserName = append(listUserName, username)
	// }
	// r.BatchedInsertAccountUsername(listUserName, 10)
	// acc, err := r.FirstUsername(&AccountUsernameModel{
	// 	Username:    "ngoctd",
	// 	AccountType: "staff",
	// })
	acc, err := r.FindIAMClient()

	if err != nil {
		log.Println(err)
	}

	if acc != nil {
		fmt.Println(acc)
	}
}
