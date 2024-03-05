# Query

GORM provides First, Take, Last methods to retrieve a single object from the database, it adds LIMIT 1 condition when querying the database, and it will return the error ErrRecordNotFound if no record is found.

Get the first record ordered by primary key
```go
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1;

db.Take(&user)
// SELECT * FROM users LIMIT 1;

db.Last(&user)
// SELECT * FROM users ORDER BY id DESC LIMIT 1;

result := db.First(&user)
result.RowsAffected
result.Error

// check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

## Joins
