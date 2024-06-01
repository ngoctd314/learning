# Introduction

Gin is a web framework in Go.

## Features

**Fast**

Radix tree based routing, small memory foot print. No reflection.
Predictable API performance.

**Middleware support**

Logger, Authorization, GZIP and finally post a message in the DB.

**Crash-free**

**JSON validation**

**Routes grouping**

Organize your routes better. Authorization required vs non required, different API versions... In addition, the groups can be nested unlimitedly without degrading performance.

**Error management**

Gin provides a convenient way to collect all the errors occurred during a HTTP request. Eventually, a middleware can write them to a log file, to a database and send them through the network.

**Rendering built-in**

Gin provides an easy to use API for JSON, XML and HTML rendering

**Extendable**

Creating a new middleware is so easy, just check out the sample codes.

## Bind form

```go
type query struct {
	FieldA string `form:"field_a" json:"field_a"`
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
	var q query
		if err := c.Bind(&q); err != nil {
			c.JSON(200, gin.H{
				"err": err,
			})
		}
		c.JSON(200, gin.H{
			"query": q,
		})
	}

	r.Run()
}
```

**Goroutines inside a middleware**

When starting new Goroutines inside a middleware or handler, you SHOULD NOT use the original context inside it, you have to use a read-only copy.

```go
cCp := c.Copy()
go func() {
    log.Println("Done! in path " + cCp.Request.URL.Path)
}()
```
