# Go 100 mistakes

[#36. Not understanding the concept of a rune](./5.%20Strings.md#36-not-understanding-the-concept-of-a-rune)

- A charset is a set of characters(a,b,c...), whereas an encoding describes how to translate a charset into binary(97, 98, 99 -> base2).
- In Go, a string references an immutable slice of arbitrary bytes.
- Go source code is encoded using UTF-8. Hence, all string literals are UTF-8 strings. But because a string can contain arbitrary bytes, if it's obtained from somewhere else (not from the source code). It isn't guaranteed to be based on the UTF-8 encoding.
- A rune corresponds to the concept of a Unicode code point, meaning an item represented by a single value.
- Using UTF-8, a Unicode code point can be encoded into 1 to 4 bytes.
- Using len on a string in Go returns the number of bytes, not the number of runes.

[#37. Inaccurate string iteration]("./5.%20Strings.md#37-inaccurate-string-iteration")

[#40. ]("./5.%20Strings.md#40-useless-string-conversions")

[#41. Substrings and memory leaks]("./5.%20Strings.md#41-substrings-and-memory-leaks")

- Understanding that a rune corresponds to the concept of a Unicode code point and that it can be composed of multiple bytes should be part of the Go developer's core knowledge to work accurately with strings.
- Iterating on a string with the range operator iterates on the runes with the index corresponding to the starting index of the rune's byte sequence. To access a specific rune index (such as the third rune), convert the string into a []rune.
- Strings.TrimRight/strings. TrimLeft removes all the trailing/leading runes contained in a given set, whereas strings. Trimsuffix/strings TrimPrefix returns a string without a provided suffix/prefix.
- Concatenating a list of strings should be done with strings.Builder to prevent allocating a new string during each iteration.
- Remembering that the bytes package offers the same operations as the strings package can help avoid extra byte/string conversions.
- Using copies instead of subtrings can prevent memory leaks, as the string returned by a substring operation will be backed by the same byte array.
