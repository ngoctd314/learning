# Many to Many

Many to Many add a join table between two models.

```go
type User struct {
    gorm.Model
    Languages []Language `gorm:"many2many:user_languages"`
}

type Language struct {
    gorm.Model
    Name string
}
```

## Back-Reference

**Declare**

```go
type User struct {
    gorm.Model
    Languages []*Language `gorm:"many2many:user_languages"`
}

type Language struct {
    gorm.Model
    Name string
    Users []*User `gorm:"many2many:user_languages"`
}
```

## Override Foreign Key

For a `many2many` relationship, the join table owns the fk which references two models, for example:

```go
type User struct {
    gorm.Model
    Languages []Language `gorm:"many2many:user_languages;foreignKey:UserID;references:ID"`
}

type Language struct {
    gorm.Model
    Name string
    Users []User `gorm:"many2many:user_languages;foreignKey:LanguageID;references:ID"`
}

type UserLanguage struct {
    gorm.Model
    UserID uint // fk user_id references user (id)
    LanguageID uint // fk language_id references language (id)
}
```
