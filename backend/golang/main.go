package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
)

const (
	StatusSuccess  = "success"
	StatusErrorFoo = "error_foo"
	StatusErrorBar = "error_bar"
)

type person struct {
	id string
}

func (p *person) print() {
	fmt.Println(p.id)
}

type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionID, err)
	}

	return amount, nil
}

var tmperr = transientError{err: errors.New("record not found")}

func getTransactionAmountFromDB(transitionID string) (float32, error) {
	if len(transitionID) == 5 {
		return 0, tmperr
	}
	return 5.0, nil
}

func f() {
	notify()
}

func notify() error {
	return nil
}

func main() {
	var r *sql.Rows
	r.Close()

}

type Route struct{}

func GetRoute(srcLat, srcLng, dstLat, dstLng float32) (Route, error) {
	err := validateCoordinates(srcLat, srcLng)
	if err != nil {
		log.Println("failed to validate source coordinates")
		return Route{}, err
	}
	err = validateCoordinates(dstLat, dstLng)
	if err != nil {
		log.Println("failed to validate target coordinates")
		return Route{}, err
	}

	return Route{}, nil
}

func validateCoordinates(lat, lng float32) error {
	if lat > 90.0 || lat < -90.0 {
		log.Printf("invalid latitude: %f", lat)
		return fmt.Errorf("invalid latitude: %f", lat)
	}
	if lng > 180.0 || lng < -180.0 {
		log.Printf("invalid longitude: %f", lng)
		return fmt.Errorf("invalid longitude: %f", lng)
	}

	return nil
}

func baz() error {
	var status string
	defer notify(status)
	defer incrementCounter(status)

	if err := foo(); err != nil {
		status = StatusErrorFoo
		return err
	}
	if err := bar(); err != nil {
		status = StatusErrorBar
		return err
	}
	status = StatusSuccess

	return nil
}

func bar() error {
	return errors.New("is bar")
}

func foo() error {
	return errors.New("is foo")
}

func incrementCounter(status string) {
	fmt.Println("incrementCounter", status)
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
