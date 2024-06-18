# Method chaining

GORM's method chaining feature allows for a smooth and fluent style of coding.

```go
db.Where("name=?","jinzhu").Where("age=?",18).First(&user)
```

## Method Categories

GORM organized methods into three primary categories: Chain Methods, Finisher Methods, and New Session methods.

**Chain Methods**

Chain methods are used to modify or append Clauses to the current statement. Some common chain methods include:

- Where
- Select
- Omit
- Joins
- Scopes
- Preload
- Raw (Raw cannot be used in conjunction with other chainable methods to build SQL)
