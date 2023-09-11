package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// LogMode(LogLevel) Interface
// Info(context.Context, string, ...interface{})
// Warn(context.Context, string, ...interface{})
// Error(context.Context, string, ...interface{})
// Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)

func main() {
	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		LogLevel: logger.Info,
		Colorful: true,
	})
	dsn := "root:secret@tcp(127.0.0.1:3306)/ap?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	r := repository{db}

	start, _ := time.Parse("2006-01-02", "2013-06-01")
	end, _ := time.Parse("2006-01-02", "2018-06-30")
	for _, v := range r.getInvoiceBetweenDates(ctx, start, end) {
		fmt.Println(v.CreditTotal, v.InvoiceDate, v.InvoiceID)
	}
}
