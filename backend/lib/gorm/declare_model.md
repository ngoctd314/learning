# Declare model

For a normal struct field, you can embed it with the tag embedded

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

And you can use tag embeddedPrefix to add prefix to embedded field's db name, for example:

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
