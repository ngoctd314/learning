# CRUD interface

## Create

```go
user := User{Name: "ngoctd", Age: 23, Birthday: time.Now()}

result := db.Create(&user) // pass pointer of data to Create

user.ID // returns inserted data's primary key
result.Error // returns error
result.RowsAffected // returns inserted records connt
```

**1. Create Record**

**2. Create Record With Selected, Omit Fields**

**3. Batch Insert**

To efficiently insert large number of records, pass a slice to the Create method. GORM will generate a single SQL statement to insert all the data and backfill primary key values.

## Query

**1. Retrieving a single object**

GORM provides First, Take, Last methods to retrieve a single object from the database

## Advanced Query

**1. Smart Select Fields**

GORM allows selecting specific fields with Select, if you often use this in your application.

```go
type User struct {
    ID uint
    Name string
    Age int
    Gender string
}

type APIUser struct {
    ID int
    Name string
}

db.Model(&User{}).Limit(10).Find(&APIUser{})
// SELECT `id`, `name` FROM `users` LIMIT 10
```

**2. Locking (FOR UPDATE)**