package data

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSave_happyPath(t *testing.T) {
	// define a mock db
	testDB, mockDB, err := sqlmock.New()
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

	resultID, err := Save(in)
	require.NoError(t, err)
	assert.Equal(t, 2, resultID)
}

// convertSQLToRegex convert SQL string to regex by treating the entire query as a literal
func convertSQLToRegex(in string) string {
	return `\Q` + in + `\E`
}
