# Advanced query

In GORM, you can efficiently select specific fields using the Select method. This is particularly useful when dealing with large models but requiring only a subset of fields, especially in API responses.

```go
type User struct {
    ID uint
    Name string
    Age int
    Gender string
}

type APIUser struct {
    ID uint
    Name string
}

// GORM will automatically select `id`, `name` fields when querying
db.Model(&User{}).Limit(10).Find(&APIUser{})
// SQL: SELECT `id`, `name` FROM users LIMIT 10
```


## Pluck

The Pluck method in GORM is used to query a single column from the database and scan the result into a slice. This method is ideal for when you need to retrieve specific fields from a model.

If you need to query more than one column, you can use Select with Scan or Find instead.

```go
// Retrieving ages of all users
var ages []int64
db.Model(&User{}).Pluck("age", &ages)

// Retrieving names of all users
var names []string
db.Model(&User{}).Pluck("name", &names)

// Retrieving names from a different table
db.Table("deleted_users").Pluck("name", &names)

// Using Distinct with Pluck
db.Model(&User{}).Distinct().Pluck("Name", &names)
// SQL: SELECT DISTINCT `name` FROM `users`

// Querying multiple columns
db.Select("name", "age").Scan(&users)
db.Select("name", "age").Find(&users)
```

## Scopes

Scopes in GORM are a powerful feature that allows you to define commonly-used query conditions as reusable methods. The scopes can be easily referenced in your queries, making your code more modular and readable.
