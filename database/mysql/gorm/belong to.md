# Belong to

A `BelongsTo` association is used when a model contains a fk that references another model. This typically indicates that the current model "belongs to" another model. For example, if you have a `Book` model and an `Author` model, each `Book` belongs to an `Author`.

Relationship: Author - Book : 1 - n, If you want detail of author in Book, you must use belong to (1 - to - many) 

```go
type Author struct {
    ID uint
    Name string
}

type Book struct {
    ID uint
    Title string
    AuthorID uint
    Author Author `gorm:"references:ID,foreignKey:AuthorID"`
}
```

A `belongs to` association sets up a one-to-one connection with another model, such that each instance of the declaring model "belongs to" one instance of the other model.

For example, if your application includes uses and companies, and each user can be assigned to exactly one company, the following types represent that relationship. Notice here that, on the User object, there is both a CompanyID as well as a Company. By default, the `CompanyID` is implicitly used to create a FK relationship between the User and Company table, and thus must be included in the User struct in order to fill the Company inner struct.

```go
// User belongs to `Company`, `CompanyID` is the fk
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

For the above example, to define the User model that belongs to `Company`, the foreign key should be `CompanyID` by convention.

GORM provides a ways to customize the foreign key, for example:

```go
type User struct {
    gorm.Model
    Name string
    CompanyRefer int
    Company Company `gorm:"foreignKey:CompanyRefer;references:ID"`
}

type Company struct {
    ID int
    Name string
}
```

## Override References

For a belongs to relationship, GORM usually uses the owner's primary field as the foreign key's value, for the above example, it is Company's field ID.

When you assign a user to a company, GORM will save the company's ID into the user's `CompanyID` field.

You are able to change it with tag `references`:

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

**NOTE** GORM usually the relationship as has one if override key name already exist it owner's type, we need to specify references in the belongs to relationship.


|Association|ForeignKey|References|
|-|-|-|
|Belong to|Entity + ID|Entity.ID|
|Has one|||
|Has many|||
