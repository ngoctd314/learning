# Go 100 mistakes

[#36. Not understanding the concept of a rune](./5.%20Strings.md#36-not-understanding-the-concept-of-a-rune)

- A charset is a set of characters(a,b,c...), whereas an encoding describes how to translate a charset into binary(97, 98, 99 -> base2).
- In Go, a string references an immutable slice of arbitrary bytes.
- Go source code is encoded using UTF-8. Hence, all string literals are UTF-8 strings. But because a string can contain arbitrary bytes, if it's obtained from somewhere else (not from the source code). It isn't guaranteed to be based on the UTF-8 encoding.
- A rune corresponds to the concept of a Unicode code point, meaning an item represented by a single value.
- Using UTF-8, a Unicode code point can be encoded into 1 to 4 bytes.
- Using len on a string in Go returns the number of bytes, not the number of runes.

[#37. Inaccurate string iteration]("./5.%20Strings.md#37-inaccurate-string-iteration")

Iterating on a string is a common operation for developers.

Let's look at a concrete example. Here, we want to print the different runes in a string and their corresponds positions:

```go
s := "hêllo"
for i := range s {
    fmt.Printf("position %d: %c\n", i, s[c])
}
```

```txt
position 0: h
position 1: Ã
position 3: l
position 4: l
position 5: o
len=6
```

This code doesn't do what we want. Let's highlight three points:

- The second rune is Ã in the output instead of ê.
- We jumped from position 1 to position 3: what is at position 2?.
- len returns a count of 6, whereas s contains only 5 runes.

Len returns the number of bytes in a string, not the number of runes. Because we assigned a string literal to s, s is a UTF-8 string. Meanwhile, the special character ê isn't encoded in a single byte; it requires 2 bytes. Therefore, calling len(s) returns 6.

```go
fmt.Println(len([]rune(s)))
fmt.Println(utf8.RuneCountInString(s))
```

When convert a slice of runes using []rune(s). Introduces a run-time overhead compared to the previous one. Indeed, converting a string into a slice of runes requires allocating an additional slice and converting the bytes into runes: an O(n) time complexity with n the number of bytes in the string. Therefore if we want to iterate over all the runes, we should use: 

```go
for i, v := range s {
    fmt.Println(i, v) // v will be rune
}
```

**A possible optimization to access a specific rune**

One optimization is possible if a string is composed of single-byte runes:

