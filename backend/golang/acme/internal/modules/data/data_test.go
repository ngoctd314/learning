package data

import (
	"database/sql"
	"errors"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSave_happyPath(t *testing.T) {
	// define a mock db
	testDB, mockDB, err := sqlmock.New()
	defer testDB.Close()
	require.NoError(t, err)

	queryRegex := convertSQLToRegex(sqlInsert)
	mockDB.ExpectExec(queryRegex).WillReturnResult(sqlmock.NewResult(2, 1))

	// monkey patching starts here
	defer func(originalDB *sql.DB) {
		db = originalDB
	}(db)
	// replace db for this test
	db = testDB
	// end of monkey patching

	in := &Person{
		FullName: "NgocTD",
		Phone:    "01234",
		Currency: "AUD",
		Price:    123.45,
	}

	// call function
	resultID, err := Save(in)

	// validate result
	require.NoError(t, err)
	assert.Equal(t, 2, resultID)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestSave_insertError(t *testing.T) {
	// define a mock
	testDB, mockDB, _ := sqlmock.New()
	defer testDB.Close()

	mockDB.ExpectExec(convertSQLToRegex(sqlInsert)).WillReturnError(errors.New("failed to insert"))

	// monkey patching starts here
	defer func(original *sql.DB) {
		// restore original DB (after test)
		db = original
	}(db)

	// replace db fo this test
	db = testDB
	// end of monkey patching

	// inputs
	in := &Person{
		FullName: "Author",
		Phone:    "01234",
		Currency: "USD",
		Price:    123.45,
	}

	// call function
	resultID, err := Save(in)

	// validate result
	require.Error(t, err)
	assert.Equal(t, defaultPersonID, resultID)
	assert.NoError(t, mockDB.ExpectationsWereMet())
}

func TestSave_getDBError(t *testing.T) {
	defer func(original func() (*sql.DB, error)) {
		getDB = original
	}(getDB)

	// replace getDB() function for this test
	getDB = func() (*sql.DB, error) {
		return nil, errors.New("getDB() failed")
	}
	// end of monkey patching

	// inputs
	in := &Person{}

	// call function
	resultID, err := Save(in)
	require.Error(t, err)
	assert.Equal(t, defaultPersonID, resultID)
}

func TestLoadAll_tableDrivenTest(t *testing.T) {
	tests := []struct {
		desc            string
		configureMockDB func(sqlmock.Sqlmock)
		expectedResults []*Person
		expectError     bool
	}{
		{
			desc: "happy path",
			configureMockDB: func(mock sqlmock.Sqlmock) {
				queryRegex := convertSQLToRegex(sqlLoadAll)

				mock.ExpectQuery(queryRegex).WillReturnRows(
					sqlmock.
						NewRows(strings.Split(sqlAllColumns, ", ")).
						AddRow(1, "John", "01234", "USD", 123.45))
			},

			expectedResults: []*Person{{
				ID:       1,
				FullName: "John",
				Phone:    "01234",
				Currency: "USD",
				Price:    123.45,
			}},
			expectError: false,
		},
	}

	for _, tt := range tests {
		// define a mock db
		testDB, mock, _ := sqlmock.New()

		tt.configureMockDB(mock)

		// monkey patching starts here
		defer func(originalDB *sql.DB) {
			db = originalDB
		}(db)

		db = testDB

		results, err := LoadAll()

		assert.Equal(t, tt.expectedResults, results, tt.desc)
		assert.Equal(t, tt.expectError, err != nil, tt.desc)
		assert.NoError(t, mock.ExpectationsWereMet())
	}
}

// convertSQLToRegex convert SQL string to regex by treating the entire query as a literal
func convertSQLToRegex(in string) string {
	return `\Q` + in + `\E`
}
