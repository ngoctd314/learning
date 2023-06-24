# GORM Guides

## Declaring models

Models are normal structs with basic Go types, pointers/alias as them or custom types implementing Scanner and Valuer interfaces.

## Conventions

GORM prefers convention over configuration. 

## Advanced

Exported fields have all permissions when doing CRUD with GORM.

Creating/Updating Time/Unix (Milli/Nano) Seconds Tracking

Embedded Struct

```go
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
    ID int64
    Name string
    Email string
    Upvotes int32
}
```

You can use tag embeddedPrefix to add prefix to embedded field's db name

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

## Fields Tags

Tags are optional to use when declaring models  Tags are case insensitive, however camelCase is preferred.

|Tag Name|Description|
|-|-|
|column|column db name|
|serializer|specifies serializer for how to serialize and deserialize data into db, serializer:json/gob/unixtime|
|embedded|embed the field|
|embeddedPrefix|column name prefix fox embedded fields|

