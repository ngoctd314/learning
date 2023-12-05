package main

import (
	"database/sql"
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(queryMatchRegexp("UPDATE products", "UPDATE products SET views = views + 1"))
}

func queryMatchRegexp(expectedSQL, actualSQL string) error {
	expect := (expectedSQL)
	actual := (actualSQL)
	re, err := regexp.Compile(expect)
	if err != nil {
		return err
	}
	if !re.MatchString(actual) {
		return fmt.Errorf(`could not match actual sql: "%s" with expected regexp "%s"`, actual, re.String())
	}
	return nil
}

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			_ = tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if _, err = tx.Exec(
		"INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)",
		userID, productID); err != nil {
		return
	}

	return
}
