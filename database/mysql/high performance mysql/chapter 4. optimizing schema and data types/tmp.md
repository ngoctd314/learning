## String Types

MySQL supports quite a few string data types, with many variations on each. These data types changed greatly in versions 4.1 and 5.0, which makes them even more complicated. This can impact performance greatly.

**VARCHAR and CHAR types**

The two major string types are VARCHAR and CHAR, which store character values. Unfortunately, it's hard to explain exactly those these values are stored on disk and in memory, because the implementations are storage engine-dependent. We assume you are using InnoDB and /or MyISAM. If not, you should read the documentation for your storage engine.

**VARCHAR**

VARCHAR stores variable-length character strings and is the most common string data type. It can require less storage space than fixed-length types, it uses only as much space as it needs.

VARCHAR uses 1 or 2 extra bytes to record the value's length: 1 byte if the column's maximum length is 255 bytes or less, and 2 bytes if it's more. Assuming the latin1 character set, a VARCHAR(10) will be use up to 11 bytes of storage space. A VARCHAR(1000) can use up to 1002 bytes, because it needs 2 bytes to store length information.

VARCHAR helps performance because it saves space. However, because the rows are variable-length, they can grow when you update them, which can cause extra work. If a row grows and no longer fits its original location, the behavior is storage engine-dependent. For example, MyISAM may fragment the row, and InnoDB may need to split the page to fit the row into it.

In version 5.0 and newer, MySQL preserves trailing spaces when you store and retrieve values. In version 4.1 and older, MySQL strips trailing spaces. It's trickier with InnoDB, which can store long VARCHAR values as BLOBs. 

**CHAR**

CHAR is fixed-length: MySQL always allocates enough space for the specified number of characters. When storing a CHAR value, MySQL removes any trailing spaces (This is was also true of VARCHAR in MySQL 4.1 and older version). Values are padded with spaces as needed for comparisons.

```go
func main() {
	_, err := mysqlConn.Exec("INSERT INTO datatype (name_char, name_varchar) VALUES (?, ?)", "  ngoctd   ", "  ngoctd     ")
	if err != nil {
		log.Fatal(err)
	}

	var result []nameCharTable
	if err := mysqlConn.Select(&result, "SELECT * FROM datatype"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

// -> select * from datatype;
// +-----------+---------------+
// | name_char | name_varchar  |
// +-----------+---------------+
// |   ngoctd  |   ngoctd      |
// +-----------+---------------+
```

CHAR is useful if you want to store very short strings, of if all values are nearly the same length. For example, CHAR is a good choice for MD5 values for user passwords (hashed), which are always the same length. Char is also better than VARCHAR for data that's changed frequently, because a fixed-length row is not prone to fragmentation. For very short columns, CHAR is also more efficient than VARCHAR; a CHAR(1) designed to hold only Y and N values will use only one byte in single-byte character set.

**(1)** Remember that the length is specified in characters, not bytes. A multibyte character set can require more than one byte to store each character.
