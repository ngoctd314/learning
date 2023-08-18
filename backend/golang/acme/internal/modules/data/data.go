package data

import (
	"database/sql"
	"errors"
	"go-learn/acme/internal/config"
	"go-learn/acme/internal/logging"
)

const (
	defaultPersonID = 0
)

var (
	db          *sql.DB
	ErrNotFound = errors.New("not found")
)

func getDB() (*sql.DB, error) {
	if db == nil {
		if config.App == nil {
			return nil, errors.New("config is not initialized")
		}

		var err error
		db, err = sql.Open("mysql", config.App.DSN)
		if err != nil {
			panic(err.Error())
		}
	}

	return db, nil
}

type Person struct {
	ID       int     `json:"id,omitempty"`
	FullName string  `json:"fullname,omitempty"`
	Phone    string  `json:"phone,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Price    float64 `json:"price,omitempty"`
}

func Save(in *Person) (int, error) {
	db, err := getDB()
	if err != nil {
		logging.L.Error("failed to get DB connection. err: %s", err)
		return defaultPersonID, err
	}

	// perform db insert
	query := "INSERT INTO person (fullname, phone, currency, price) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, in.FullName, in.Phone, in.Currency, in.Price)
	if err != nil {
		logging.L.Error("failed to save person into DB. err: %s", err)
		return defaultPersonID, err
	}

	// retrieve and return the ID of the person created
	id, err := result.LastInsertId()
	if err != nil {
		logging.L.Error("failed to retrieve id of last saved person. err: %s", err)
		return defaultPersonID, err
	}

	return int(id), nil
}

func LoadAll() ([]*Person, error) {
	db, err := getDB()
	if err != nil {
		logging.L.Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// perform DB select
	query := "SELECT id, fullname, phone, currency, price FROM person"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var out []*Person

	for rows.Next() {
		// retrieve columns and populate the person object
		record, err := populatePerson(rows.Scan)
		if err != nil {
			logging.L.Error("failed to convert query result. err: %s", err)
			return nil, err
		}

		out = append(out, record)
	}

	if len(out) == 0 {
		logging.L.Warn("no people found in the database.")
		return nil, ErrNotFound
	}

	return out, nil
}

// Load will attempt to load and return a person.
// It will return ErrNotFound when the requested person does not exist.
// Any other errors returned are caused by the underlying database or our connection to it.
func Load(ID int) (*Person, error) {
	db, err := getDB()
	if err != nil {
		logging.L.Error("failed to get DB connection. err: %s", err)
		return nil, err
	}

	// perform DB select
	query := "SELECT id, fullname, phone, currency, price FROM person WHERE id = ? LIMIT 1"
	row := db.QueryRow(query, ID)

	// retrieve columns and populate the person object
	out, err := populatePerson(row.Scan)
	if err != nil {
		if err == sql.ErrNoRows {
			logging.L.Warn("failed to load requested person '%d'. err: %s", ID, err)
			return nil, ErrNotFound
		}

		logging.L.Error("failed to convert query result. err: %s", err)
		return nil, err
	}
	return out, nil
}

// custom type so we can convert sql results to easily
type scanner func(dest ...interface{}) error

// reduce the duplication (and maintenance) between sql.Row and sql.Rows usage
func populatePerson(scanner scanner) (*Person, error) {
	out := &Person{}
	err := scanner(&out.ID, &out.FullName, &out.Phone, &out.Currency, &out.Price)
	return out, err
}
