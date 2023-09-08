# Query

## Retrieving a single object

GORM provides First, Take, Last methods to retrieve a single object from the database, it adds LIMIT 1 condition when querying the database, and it will return the error ErrRecordNotFound if no record is found.

```go
person := Person{
    ID: 1,
}
// Get the first record ordered by primary key
// SELECT * FROM persons WHERE id = ? ORDER BY id LIMIT 1
tx = db.First(&person)
if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
    log.Fatal(tx.Error)
}
log.Println(tx.RowsAffected, person)
```

```go
person := Person{
    ID: 2,
}
// Get the first record (no order)
// SELECT * FROM persons WHERE id = ? LIMIT 1
tx = db.Take(&person)
if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
    log.Fatal(tx.Error)
}
log.Println(tx.RowsAffected, person)
```

```go
person := Person{
    ID: 5,
}
// Get the last record
// SELECT * FROM persons WHERE id = ? ORDER BY id DESC LIMIT 1
tx = db.Last(&person)
if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
    log.Fatal(tx.Error)
}
log.Println(tx.RowsAffected, person)
```

If you want to avoid the ErrRecordNotFound error, you could use Find like db.Limit(1).Find(&user), the Find method accepts both struct and slice data.

Using Find without a limit for single object db.Find(&user) will query the full table and return only the first object which is not performant and nondeterministic.

```go
person := Person{
    ID: 3,
}
// SELECT * FROM persons
tx = db.Find(&person)
if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
    log.Fatal(tx.Error)
}
log.Println(tx.RowsAffected, person)
```

The First and Last methods will find the first and last record (respectively) as orderd by primary key. They only work when a pointer to the destination struct is passed to the methods as argument or when the model is specified using db.Model()
