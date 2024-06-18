# Context

GORM's context support, enabled by the `WithContext` method, is a powerful feature that enhances the flexibility and control of database operations in Go applications, it allows for context management across different operational models, timeout settings, and even integration into hooks/callbacks and middlewares.

**Single Session Mode**

Single session model is appropriate for executing individual operations. It ensures that the specific operation is executed within the context's scope, allowing for better control and monitoring.

```go
db.WithContext(ctx).Find(&users)
```

**Continuous Session Mode**

Continuous session mode is ideal for performing a series of related operations. It maintains the context across these operations, which is particularly useful in scenarios like transactions.

```go
tx := db.WithContext(ctx)

tx.First(&user, 1)
tx.Model(&user).Update("role", "admin")
```

**Context timeout**

Setting a timeout on the context passed to db.WithContext can control the duration of long-running queries. This is crucial for maintaining performance and avoiding resource lock-ups in database interactions.

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()

db.WithContext(ctx).Find(&user)
```


