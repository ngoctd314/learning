# Has Many

A `has many` association sets up a one-to-many connection with another model, unlike `has one`, the owner could have zero or many instances of models.

```go
type User struct {
    gorm.Model
    CreditCards []CreditCard `gorm:"foreignKey:UserID;references:ID"`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserID uint
}
// foreign key fk_user(user_id) references user (id)
```

## Override Foreign Key

To define a has many relationship, a foreign key must exist. The default foreign key's name is the owner's type name plus the name of ifs primary key field.

To use another field as foreign key, you can customize it with a foreignKey tag:

```go
type User struct {
    gorm.Model
    CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserRefer uint
}
```

## Override References

GORM usually uses the owner's primary key as the foreign key's value, for the above example, it is the User's ID.

```go
type User struct {
    gorm.Model
    MemberNumber string
    CreditCards []CreditCard `gorm:"foreignKey:UserNumber;references:MemberNumber"`
}

type CreditCard struct {
    gorm.Model
    Number string
    UserNumber string
}
// foreign key fk_user(user_number) on user (member_number)
```
