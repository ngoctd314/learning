# SQL Builder

## Raw SQL

Query Raw SQL with Scan

```go
type Result struct {
    ID int
    Name string
    Age int
}
var result Result
db.Raw("SELECT id, name, age FROM users WHERE id = ?", 3).Scan(&result)
db.Raw("SELECT id, name, age FROM users WHERE name = ?", "jinzhu").Scan(&result)

var age int
db.Raw("SELECT SUM(age) FROM users WHERE role = ?", "admin").Scan(&age)
```
