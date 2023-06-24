# Gorm - ORM mapping for Go

## Query

### Retrieving a single object

GORM provides First, Take, Last methods to retrieve a single object from the database.

```go
// Get the first record ordered by primary key
db.First(&user)
// SELECT * FROM users ORDER BY id LIMIT 1; 

// Get one record, no specified order
db.Take(&user)
// SELECT * FROM users LIMIT 1;

// Get last record, ordered by primary key desc
db.Last(&user)
// SELECT * FROM users ORDER BY id desc LIMIT 1;

// check error ErrRecordNotFound
errors.Is(result.Error, gorm.ErrRecordNotFound)
```

If you want to avoid the ErrRecordNotFound error, you could use Find like db.Limit(1).Find(&user)

Using Find without a limit for single object db.Find(&user)  will query the full table and return only the first object which is not performant and nondeterministic.