# sql

## func(*Rows) Scan

```go
func (rs *Rows) Scan(dest ...any) error {}
```

Scan copies the columns in the current row into the values pointed at by dest. The number of values in dest must be the same as the number of columns in Rows.

Scan converts columns read from the database into the following common Go types and special types provided by the sql package:

```txt
*string
*[]byte
*int, *int8, *int16, *int32, *int64
*float32, *float64
*interface{}
*RawBytes
*Rows (cursor value)
any type implementing Scanner (see Scanner docs)
```
