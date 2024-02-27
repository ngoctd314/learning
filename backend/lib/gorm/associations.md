# Associations

## Belongs To

A belongs to associations sets up a one-to-one connection with another model, such that each instance of the declaring model "belongs to" one instance of the other model.

For example, if your application includes users and companies, and each user can be assigned to exactly one company, the following types represent that relationship. Notice here that, on the User object, there is both a CompanyID as well as a Company. By default, the CompanyID is implicitly used to create a foreign key relationship between the User and Company tables, and thus must be included in the User struct in order to fill the Company inner struct.

```go
// `User` belongs to `Company`, `CompanyID` is the foreign key
type User struct {
    gorm.Model
    Name string
    CompanyID int
    Company Company
}

type Company struct {
    ID int
    Name string
}
```

## Override Foreign Key

To define a belongs to relationship, the foreign key must exist, the default foreign key uses the owner's type name plus its primary field name.

For the above example, to define the User model that belongs to Company, the foreign key should be CompanyID by convention.

GORM provides a way to customize the foreign key.

```go
type User struct {
    gorm.Model
    Name string
    CompanyRefer  int
    Company Company `gorm:"foreignKey:CompanyRefer"`
}
```

## Override References

For a belongs to relationship, Gorm usually uses the owner's primary field as the foreign key's value, for the above example, it is Company's field ID.

```go
type User struct {
    gorm.Model
    Name string
    CompanyID string
    Company Company `gorm:"references:Code"`
}

type Company struct {
    ID int
    Code string
    Name string
}
```

```go
type User struct {
    gorm.Model
    Name string
    CompanyID string
    Company Company `gorm:"references:CompanyID"`
}

type Company struct {
    CompanyID int
    Code string
    Name string
}
```
