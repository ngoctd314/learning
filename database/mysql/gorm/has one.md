# Has one

A `HasOne` association is used when a model has exactly one related model. This indicates that the current model "hasOne" of another model. For example, if you have a `User` model and each `User` has one `Profile`, you would use a `HasOne` association.

```go
type User struct {
    ID uint
    Name string
    Profile Profile
}

type Profile struct {
    ID uint
    UserID uint
    Bio string
}
```

A has one association sets up a one-to-one connection with another model.

For example, if your application includes users and credit cards, and each user can only have one credit card.

```go
type User struct {
    gorm.Model
    CreditCard CreditCard `foreignKey:UserID;references:ID`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint
}
```

## Override Foreign Key

For a `has one` relationship, a foreign key field must also exist, the owner will save the primary key of the model belongs to it into this field.

The field's name is usually generated with has one model's type plus its primary key, for the above example it is UserID.

If you want to use another field to save the relationship, you can change it with tag foreignKey.

```go
type User struct {
    gorm.Model
    Name string
    CreditCard CreditCard `gorm:"foreignKey:Username;references:Name"`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserName string
}
// FK: fk_username (username) ON user (name)
```

## Override References

```go
type User struct {
    gorm.Model
    Name string
    CreditCard CreditCard `gorm:"foreignKey:UserName;references:Name"`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserName string
    // fk_username (username) on user.name
}
```
