# Compile-time Dependency Injection With Go  Cloud's Wire

## What problem does Wire solve?

Dependency injection is a standard technique for producing flexible and loosely coupled code, be expliciting providing components with all of the dependencies they need to work. In Go, this often takes the form of passing dependencies to constructors:

```go
func NewUserStore(cfg *Config, db *mysql.DB) (*UserStore, error) {}
```

This technique works great at small code, but larger applications can have a complex graph of dependencies. Replace one implementation of a service with another can be painful because it involves modifying the dependency graph by adding a whole new set of dependencies (and their dependencies...), and removing unused old ones. In practice, making changes to initialization code in applications with large dependency graphs is tedious and slow.

