# Preloading (Eager Loading)

GORM allows eager loading relations in other SQL with `Preload`

```go
type User struct {
    gorm.Model
    Username string
    Order []Order
}

type Order struct {
    gorm.Model
    UserID uint
    Price float64
}
```

Preload Orders when find users

```go
db.Preload("Orders").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4);
```

```go
db.Preload("Orders").Preload("Profile").Preload("Role").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4); // has many
// SELECT * FROM profiles WHERE user_id IN (1,2,3,4);; // has one
```

## Joins Preloading

Preload loads the association data in a separate query, Join Preload will loads association data using left join, for example:

```go
db.Joins("Company").Joins("Manager").Joins("Account").First(&user, 1)
```

Join with conditions

```go
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id` AND `Company`.`alive` = true;
db.Joins("Company", DB.Where(&Company{Alive: true})).Find(&users)
```

NOTE: `Join Preload` works with one-to-one relation, e.g: `has one`, `belongs to`

## Preload All

`clause.Associations` can work with `Preload` similar like `Select` when create/update you can use it to `Preload` all associations, for example:

```go
type User struct {
    gorm.Model
    Name string
    CompanyID uint
    Company Company
    Role Role
    Orders []Order
}

db.Preload(clause.Associations).Find(&users)
```

clause.Associations won't preload nested associations, but you can use it with Nested Preloading together

```go
db.Preload("Orders.OrderItems.Product").Preload(clause.Associations).Find(&users)
```

## Preload with conditions

GORM allows Preload associations, it works similar to inline conditions

```go
// Preload Orders with conditions
db.Preload("Orders", "state NOT IN (?)", "cancelled").Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) AND state NOT IN ('cancelled')

db.Where("state = ?", "active").Preload("Orders", "state NOT IN (?)", "cancelled")
```

## Custom Preloading SQL

You are able to custom preload SQL by passing in `func (db *gorm.DB) *gorm.DB`, for example:

```go
db.Preload("Orders", func(db *gorm.DB) *gorm.DB {
    return db.Order("orders.amount DESC")
}).Find(&users)
// SELECT * FROM users;
// SELECT * FROM orders WHERE user_id IN (1,2,3,4) ORDER BY orders.amount DESC;
```

## Nesting Preloading

GORM supports nested preloading, for example:

```go
db.Preload("Orders.OrderItems.Product").Preload("CreditCard").Find(&users)
```

## Embedded Preloading
