# Associations

## Belongs To

A `belongs to` association sets up a one-to-one connection with another model, such that each instance of the declaring model "belongs" to one instance of the other model.

On the `User` object, there is both a `CompanyID` as well as a `Company`. By default, the `CompanyID` is implicitly used to create a foreign key relationship between the User and Company tables, and thus must be included in the `User` struct in order to fill the `Company` inner struct.

```go
// `User` belongs to `Company`, `CompanyID` is the foreign key
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
	Name string
}
```

**Override Foreign Key**

To define a belongs to relationship, the foreign key must exist, the default foreign key uses the owner's type name plus its primary field name.

For the above example, to define the `User` model that belongs to `Company`, the foreign key should be `CompanyID` by convention.

GORM provides a ways to customize the foreign key, for example:

```go
type User struct {
    gorm.Model
    Name string
    CompanyRefer int
    // use CompanyRefer as foreign key
    Company Company `gorm:"foreignKey:CompanyRefer"`
}

type Company struct {
    ID int 
    Name string
}
```

**Override References**

For a belongs to relationship, GORM usually uses the owner's primary field as the foreign key's value, for the above example, it is `Company` field ID.

When you assign a user to a company, GORM will save he company's ID into the user's CompanyID field.

You are able to change it with tag `references`, e.g:

```go
type User struct {
    gorm.Model
    Name string
    CompanyID string
    Company Company `gorm:"foreignKey:CompanyID;references:Code"` // use Code as references
}

type Company struct {
    ID int
    Code string
    Name string
}
```

## CRUD with Belongs To

GORM automates the saving of associations and their references when creating or updating records using an upsert technique that primarily updates foreign key references for existing associations.

**Auto-Saving Associations on Create**

When you create a new record, GORM will automatically save its associated data. This includes inserting data into related tables and managing foreign key references. 

```go

```

**Association Tags**

Association tags in GORM are used to specify how associations between models are handled. These tags define the relationship's details, such as foreign keys references, and constraintsl. Understanding these tags is essential for setting up and managing relationships effectively.

|Tag|Description|
|-|-|
|foreignKey|Specifies the column name of the current model used as a foreignKey key in the join table.|
|references|Indicated the column name in the references table that the foreign key of the join table maps to.|
|polymorphic|Defines the polymorphic type, typically the model name|

## Eager Loading

GORM allows eager loading belongs to associations with `Preload` or `Joins`, refer Preloading for details.

### Preload

```go
// `User` belongs to `Company`, `CompanyID` is the foreign key
// User -> Order => foreignKey = ref model field, references: current model field
type User struct {
	ID       int
	Username string
	Orders   []Order `gorm:"foreignKey:UserID;references:id"`
}

func (User) TableName() string {
	return "users"
}

type Order struct {
	ID     int
	UserID uint
	Price  int
}

func (Order) TableName() string {
	return "orders"
}

func main() {
	db, _ := gorm.Open(mysql.Open("root:secret@(192.168.49.2:30300)/learn_mysql?parseTime=true"), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	db.Exec("DROP TABLE IF EXISTS users")
	db.Exec("CREATE TABLE users (id int auto_increment primary key, username varchar(255))")
	db.Exec("INSERT INTO users (username) VALUES ('test1')")
	db.Exec("DROP TABLE IF EXISTS orders")
	db.Exec("CREATE TABLE orders (id int auto_increment primary key, user_id int, price int)")
	db.Exec("INSERT INTO orders (user_id, price) VALUES (1, 1), (1,2)")

	var users []User
	db.Preload("Orders").Find(&users)
	fmt.Println(users)
}
```

### Joins Preloading

`Preload` loads the associations data in a separate query, `Join Preload` will loads association data using left join, for example:

```sql
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, "users.name = ?", "jinzhu")
db.Joins("Company").Joins("Manager").Joins("Account").Find(&user,"users.id IN ?", []int{1,2,3,4,5})
```

Join with conditions

```sql
db.Joins("Company", db.Where(&Company{Alive: true})).Find(&users)
// SELECT `users`.`id`, `users`.`name`, `users`.`age`, `Company`.`id` AS `Company__id`....
```

Join nested model

```sql
db.Joins("Manager").Joins("Manager.Company").Find(&users)
```

`Join Preload` works with one-to-one relationship, e.g: `has one`, `belongs to`

### Preload All

`clause.Associations` can work with Preload similar like `Select` when creating/updating, you can use it to `Preload` all associations, for example:

```go
type User struct {
    gorm.Model
    Name string
    CompanyID uint
}
```

```sql
db.Preload(clause.Associations).Find(&users)
```

`clause.Associations` won't preload nested associations, but you can use it with Nested Preloading together.

```go
db.Preload("Orders.OrderItems.Product").Preload(clause.Associations).Find(&users)
```

### Preload with conditions

GORM allows Preload associations with conditions, it works similar to Inline Conditions

```go
db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('cancelled');

db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users WHERE state = 'active';
// SELECT * FROM orders WHERE user_id IN (1,2) AND state NOT IN ('cancelled');
```


