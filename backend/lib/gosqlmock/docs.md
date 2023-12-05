# Whoami

sqlmock is a mock library implementing sql/driver. Which has one and only purpose - to simulate any sql driver behavior in tests, without needing a real database connection. It helps to maintain correct TDD workflow.

## Customize SQL query matching

There were plenty of requests from users regarding SQL query string validation or different matching option. We have now implemented the QueryMatcher interface, which can be passed through an option when calling sqlmock.New or sqlmock.NewWithDSN.

This now allows to include some library, which would allow for example to parse and validate mysql SQL AST. And create a custom QueryMatcher in order to validate SQL in sophistiated ways. 

By default, sqlmock is preserving backward compatibility and default query matcher is sqlmock.QueryMatcherRegexp which uses expected SQL string as a regular expression to match incoming query string. There is an equality matcher: QueryMatcherEqual which will do a full case sensitive match.

sqlmock preserving backward compatibility and default query matcher is sqlmock.QueryMatcherRegexp which uses expected SQL string as a regular expression to match incoming query string. There is an equality matcher: QueryMatcherEqual which will do a full case sensitive match.

```go
db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
```

The query matcher can be fully customize based on user needs. sqlmock will not provide a standard sql parsing matchers, since various drivers may not follow the same SQL standard.

```go
func Test_main(t *testing.T) {
	sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
}
```

## Matching arguments like time.Time

There may be arguments which are of struct type and cannot be compared easily by value like time.Time. In this case sqlmock provides as Argument interface which can be used in more sophistiated matching.

```go
type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestAnyTimeArgument(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Errorf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectExec("INSERT INTO users").
		WithArgs("john", AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err = db.Exec("INSERT INTO users(name, created_at) VALUES (?, ?)", "john", time.Now())
	if err != nil {
		t.Errorf("error '%s' was not expected, while inserting a row", err)
	}

    // ExpectationsWereMet checks whether all queued expectations
    // were met in order. If any of them was not met - an error is returned.
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
```
