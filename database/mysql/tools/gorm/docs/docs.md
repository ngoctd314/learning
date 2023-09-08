# GORM

## Getting Started

### Overview

### Declaring Models

Models are normal structs with basic Go types, pointers/alias of them or custom types implementing Scanner and Valuer interfaces

GORM prefers convention over configuration. By default, GORM uses ID as primary key, pluralizes struct name to snake_cases as table name, snake_case as column name, and uses CreatedAt, UpdatedAt to track creating/updating time.

**gorm.Model**

GORM defined a gorm.Model struct which includes fields ID, CreatedAt, UpdatedAt, DeletedAt

```go
type Model struct {
    ID uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}
```

**Advanced**

**Field-level Permission**

Exported fields have all permissions when doing CRUD with GORM, and GORM allows you to change the field-level permission with tag, you can make a field to be read-only, create-only, update-only or ignored.

```go
type User struct {
    Name string `gorm:"<-:create"`
    Name string `gorm:"<-:update"`
    Name string `gorm:"<-"`
    Name string `gorm:"<-:false"`
    Name string `gorm:"->"`
    Name string `gorm:"->;<-:create"`
    Name string `gorm:"->:false;<-:create"`
    Name string `gorm:"-"`
}
```

**Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking**

GORM use CreatedAt, UpdatedAt to track creating/updating time by convention, and GORM will set the current time when creating/updating if the fields are defined.

To use fields with different name, you can configure those fields with tag autoCreateTime, autoUpdateTime.

If you prefer to save UNIX (milli/nano) seconds instead of time, you can use int

```go
type User struct {
    CreatedAt time.Time // Set to current time if it is zero on creating
    UpdatedAt int
    Updated int64 `gorm:"autoUpdateTime:nano"`
    Updated int64 `gorm:"autoUpdateTime:milli"`
    Created int64 `gorm:"autoCreateTime"`
}
```

**Embedded Struct**

For anonymous fields, GORM will include its fields into its parent struct

```go
type User struct {
    gorm.Model
    Name string
}
// equals
type User struct {
    ID uint `gorm:"primaryKey"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
    Name string
}

type Author struct {
    Name string
    Email string
}

type Blog struct {
    ID int
    Author Author `gorm:"embedded"`
    Upvotes int32
}

// equals
type Blog struct {
    ID int
    Name string
    Email string
    Upvotes int32
}
```

And you can use tag embeddedPrefix to add prefix to embedded field's db name

```go
type Blog struct {
    ID int
    Author Author `gorm:"embedded;embeddedPrefix:author_"`
    Upvotes int32
}
// equals
type Blog struct {
    ID int64
    AuthorName string
    AuthorEmail string
    Upvotes int32
}
```

### Connecting to Database

## CRUD Inteface

### Create

```sql
CREATE TABLE persons (
    id INT AUTO_INCREMENT,
    age INT NOT NULL,
    birthday DATE NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    KEY(id),
);
```

**Single row**

```go
user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}

result := db.Create(&user) // pass pointer of data to Create

user.ID // returns inserted data's primary key
result.Error // returns error
result.RowsAffected // returns inserted records count
```

**Multiple rows**

```go
users := []*Person{
    {
        Age:      19,
        Birthday: time.Now(),
    },
    {
        Age:      20,
        Birthday: time.Now(),
    },
}
tx = db.Create(&users) // pass a slice to insert multiple row
log.Println(tx.Error, tx.RowsAffected)
```

**Create Record With Selected Fields**

```go
person := Person{
    Age:      20,
    Birthday: time.Now(),
}
tx = db.Select("age", "birthday").Create(&person)
log.Println(tx.Error, tx.RowsAffected, person)
```

**Create a record and ignore the values for fields passed to omit**

```go
person := &Person{
    Age: 18,
}
// INSERT INTO persons (age, created_at, updated_at) VALUES (...)
tx = db.Omit("name", "birthday").Create(&person)
log.Println(tx.Error, tx.RowsAffected, person)
```

**Batch Insert**

To efficient insert large number of records, pass a slice to the Create method. GORM will generate a single SQL statement to insert all the data and backfill primary key values, hook methods will be invoked too. It will begin a transaction when records can be splited into multiple batches.

```go
persons := []*Person{
    {
        Age:      18,
        Name:     "name-18",
        Birthday: time.Now(),
    },
    {
        Age:  19,
        Name: "name-19",
    },
    {
        Age:  20,
        Name: "name-20",
    },
}
tx = db.Create(&persons)
log.Println(tx.Error, tx.RowsAffected, persons)
```

**Create Hooks**

GORM allows user defined hooks to be implemented for BeforeSave, BeforeCreate, AfterSave, AfterCreate. These hook method will be called creating a record

```go
func (Person) TableName() string {
	return "persons"
}

func (p Person) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Age < 18 {
		return errors.New("child is not allow")
	}

	return nil
}
```

**Create From Map**

GORM supports create from map[string]interface{} and []map[string]interface{}{}

When creating from map, hooks won't be invoked, associations won't be saved and primary key value won't be back filled

```go
var person Person
tx = db.Model(person).Create(map[string]interface{}{
    "age":        20,
    "name":       "NgocTD",
    "created_at": time.Now(),
    "updated_at": time.Now(),
})
log.Println(tx.Error, tx.RowsAffected, person)
```
