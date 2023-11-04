# Go 100 mistakes

[#36. Not understanding the concept of a rune](./5.%20Strings.md#36-not-understanding-the-concept-of-a-rune)

- A charset is a set of characters(a,b,c...), whereas an encoding describes how to translate a charset into binary(97, 98, 99 -> base2).
- In Go, a string references an immutable slice of arbitrary bytes.
- Go source code is encoded using UTF-8. Hence, all string literals are UTF-8 strings. But because a string can contain arbitrary bytes, if it's obtained from somewhere else (not from the source code). It isn't guaranteed to be based on the UTF-8 encoding.
- A rune corresponds to the concept of a Unicode code point, meaning an item represented by a single value.
- Using UTF-8, a Unicode code point can be encoded into 1 to 4 bytes.
- Using len on a string in Go returns the number of bytes, not the number of runes.

[#37. Inaccurate string iteration]("./5. Strings.md# 37. Inaccurate string iteration")

Iterating on a string is a common operation for developers.
