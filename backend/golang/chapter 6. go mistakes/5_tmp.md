## 38. Misusing trim functions

One common mistake made by Go developers when using the strings package is to mix TrimRight and TrimSuffix. Both functions serve a similar purpose, and it can be fairly easy to consufuse them.

In the following example, we use TrimRight. What should be the output of this code?

```go
fmt.Println(strings.TrimRight("123oxo", "xo"))
```

The answer is 123. Is that what you expected? If not, you were probably expecting the result of TrimSuffix, instead. Let's review both functions. 

TrimRight removes all the trailing runes contained in a given set. In our example, we passed as a set xo, which contains two runes: x and o.

The answer is 123. It that what you expected? If not, you were probably expecting the result of TrimSuffix, instead. Let's review both functions.

TrimRight iterates backward over each rune. If a rune is part of the provided set, the function removes it. If not, the function stops its iteration and returns the remaining string. This is why our example returns 123.

On the other hand, TrimSuffix returns a string without a provided trailing suffix:

```go
fmt.Println(strings.TrimSuffix("123oxo", "xo"))
```

Because 123oxo ends with xo, this code prints 123o. Also removing the trailing suffix isn't a repeating operation, so TrimSuffix("123xoxo", "xo") returns 123xo.

```go
fmt.Println(strings.TrimRight("123oxo", "xo"))
fmt.Println(strings.TrimLeft("oxo123", "xo"))
fmt.Println(strings.TrimPrefix("oxo123", "xo"))
fmt.Println(strings.TrimSuffix("123oxo", "xo"))
```

One last note related to this topic: Trim applies both TrimLeft and TrimRight on a string. So, it removes all the leading and trailing runes contained in a set:

```go
fmt.Println(strings.Trim("xox1ox23oxo", "xo")) // 1ox23
```

In summary, we have to make sure we understand the difference between TrimRight/TrimLeft, and TrimSuffix/TrimPrefix:

- TrimRight/TrimLeft removes the trailing, leading runes in a set.
- TrimSuffix/TrimPrefix removes a given suffix/prefix.

## 39. Under-optimized string concatenation

When it comes to concatenation strings, there are two main approaches in Go, and one of them can be really inefficient in some conditions.

```go
func concat(values ...string) string {
	s := ""
	for _, value := range values {
		s += value
	}
	return s
}
```

During each iteration, the += operator concatenates s with the value string. At first sight, this function may not look wrong. But with this implementation, we forget one of the core characteristics of a string: its immutability. Therefore, each iteration doesn't update s; it reallocates a new string in memory, which significantly impacts the performance of this function.

Fortunately, there is a solution to dead with this problem, using the strings package and the Buidler struct:

```go
func concat(values ...string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}
	return sb.String()
}
```

First, we created a strings.Builder struct using its zero value. During each iteration, we constructed the resulting string by calling the WriteString method that appends the content of value to its internal buffer, hence minimizing memory copying.

Using strings.Builder, we can also append

- A byte slice using Write
- A single byte using WriteByte 
- A single rune using WriteRune 

Internally, strings.Builder holds a byte slice. Each call to WriteString results in a call to append on this slice. There are two impacts. First, this struct shouldn't be used concurrently, as the call to append would lead to race conditions. The second impact is something that we saw in mistake #21, "Inefficient slice initialization": if the future length of a slice is already known, we should preallocate it. For that purpose, strings.Builder exposes a method Grow(n int) to guarantee space for another n bytes. 

Let's write another version of the concat method by calling Grow with the total number of bytes:

```go
total := 0
for i := 0; i < len(values); i++ {
    total += len(values[i])
}

sb := strings.Builder{}
sb.Grow(total)
for _, value := range values {
    _, _ = sb.WriteString(value)
}

return sb.String()
```

```txt
Benchmark_concat-grow+builder    1246920	      1254 ns/op	    3456 B/op	       1 allocs/op

Benchmark_concat1-+string    	   31044	     39378 ns/op	  176864 B/op	     101 allocs/op

Benchmark_concat2-builder    	  447038	      3411 ns/op	   12512 B/op	      10 allocs/op
```

If a slice isn't allocated with a given length or capacity, the slice will keep growing each time it becomes full, resulting in additional allocations and copies. Hence, iterating twice is the most efficient option in this case.

strings.Builder is the recommended solution to concanate a list of strings. Usually, this solution should be used within a loop. Indeed, if we just have to concatenate a few strings (such as name and a surname), using strings.Builder is not recommended as doing so will make the code a bit less readable than using the += operator or fmt.Sprintf.

## 40 Useless string conversions

When choosing to work with a string or a []byte, most programmers tend to favor strings for convenience. But most I/O is actually done with []byte. For example, io.Reader, io.Writer, and io.ReadAll work with []byte, not strings. Hence, working with strings means extra conversions, although the bytes package contains many of the same operations as the strings package.

```go
func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// call sanitize
	santized := strings.TrimSpace(string(b))

	return []byte(santized), nil
}
```

We have to pay the extra price of converting a []byte into a string and then converting a string into a []byte. Memory-wise, each of these conversions requires an extra allocation. Indeed, even though a string is backed by a []byte, converting a []byte into a string requries a copy of the bytes slice. It means a new memory allocation and a copy of all the bytes.

**String immutability**

```go
b := []byte{'a', 'b', 'c'}
s := string(b)
b[1] = 'x'
fmt.Println(s)
```

So, how should we implement the sanitize function? Instead of accepting and returning a string, we should manipulate a byte slice:

```go
func santize(b []byte) []byte {
    return bytes.TrimSpace(b)
}
```

The bytes package also has a TrimSpace function to trim all the leading and trailing while spaces. Then, calling the santize function doesn't require any extra conversions:

```go
func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// call sanitize
	return bytes.TrimSpace(b), nil
}
```

Most I/O is done with []byte, not strings. When we're wondering whether we should work with strings or []byte, let's recall that working with []byte isn't necessarily less convenient. Indeed, all the exported functions of the strings package also have alternatives in the bytes package: Split, Count, Contains, Index, and so on.

```go
func main() {
	sub := []byte{'a', 'b'}
	s := []byte{'x', 'y', 'a', 'b', 'z'}
	println(bytes.Contains(s, sub))
}
```

Hence, whether we're doing I/O or not, we should first check whether we would implement a whole workflow using bytes instead of strings and avoid the price of additional conversions.

## 41. Substrings and memory leaks

```go
s1 := "Hello, World!"
s2 := s1[:5] // Hello
```

s2 is constructed as a substring of s1. This example creates a string from the first five bytes, not the first five runes. Hence, we shouldn't use this syntax in the case of runes encoded with multiple bytes. Instead, we should convert the input string into a []rune byte first:

```go
s1 := "Hêllo, World!"
s2 := string([]rune(s1)[:5]) // Hêllo
```

Now that we have refreshed our minds regarding the substring operation, let's look at a concrete problem to illustrate possible meomory leaks.

We will receive log messages as strings. Each log will first be formatted with a universally unique identifier (UUID; 46 characters) followed by the message itself. We want to store these UUIDs in memory.

```go
func main() {
	s := store{}
	s.handleLog()
	runtime.GC()
	printAlloc()

	runtime.KeepAlive(s)
}

type store struct {
	data []string
}

func (s *store) handleLog() error {
	log := make([]byte, 1024*1024*1024)
	logStr := string(log)

	if len(logStr) < 36 {
		return errors.New("log is not correctly formatted")
	}
	s.data = append(s.data, logStr[:1])

	return nil
}
```

When doing a substring operation, the Go specification doesn't specify whether the resulting string and the one involved in the substring operation should share the same data. However, the standard Go compiler does let them share the same backing array, which is probably the best solution memory-wise and performance-wise as it prevents a new allocation and a copy.

We mentioned that log messages can be quite heavy. logStr[:1] will create a new string referencing the same backing array. Therefore, each uuid string that we store in memory will contain not just 36 bytes but the number of bytes in the initial log string: potentially, thousands of bytes.

How can we fix this? By making a deep copy of the substring so that the internal byte slice of uuid references a new backing array of only 36 bytes:

```go
func main() {
	s := store{}
	s.handleLog()
	runtime.GC()
	printAlloc()

	runtime.KeepAlive(s)
}

type store struct {
	data []string
}

func (s *store) handleLog() error {
	log := make([]byte, 1024*1024*1024)
	logStr := string(log)

	if len(logStr) < 36 {
		return errors.New("log is not correctly formatted")
	}
	s.data = append(s.data, string([]byte(logStr[:1])))

	return nil
}
```

The copy is performed by convertion the substring into a []byte first and then into a string again. By doing this, we prevent a memory leak from occuring. The uuid string is backed by an array consisting of only 36 bytes.

Note that some IDEs or linters may warn that the string([]byte(s)) conversions aren't necessary. For example, GoLand, the Go JetBrains IDE, warns about a reduntdant type conversion. This is true in the sense that we convert a string into a string, but this operation has an actual effect. As discussed, it prevents the new string from being backed by the same array as uuid. We need to aware that the warnings raised by IDEs or linters may sometimes be inaccurate.

**NOTE** Because a string is mostly a pointer, calling a function to pass a string doesn't result in a deep copy of the bytes. The copied string will still reference the same backing array.

```go
func main() {
	s := store{}
	s.handleLog()
	runtime.GC()
	printAlloc()

	runtime.KeepAlive(s)
}

type store struct {
	data []string
}

func (s *store) handleLog() error {
	log := make([]byte, 1024*1024*1024)
	logStr := string(log)

	if len(logStr) < 36 {
		return errors.New("log is not correctly formatted")
	}
	s.data = append(s.data, strings.Clone(logStr[:1]))

	return nil
}
```

As of Go 1.18, the standard library also includes a solution with strings.Clone that returns a fresh copy of a string:

```go
strings.Clone(log[:36])
```

Calling strings.Clone makes a copy of log[:36] into a new allocation, preventing a memory leak.

We need to keep two things in mind while using the substring operation in Go. First, the interval provided is based on the number of bytes, not the number of runes. Second, a substring operation may lead to a memory leak as the resulting substring will share the same backing array as the intial string.

### Summary

- Understanding that a rune corresponds to the concept of a Unicode code point and that it can be composed of multiple bytes should be part of the Go developer's core knowledge to work accurately with strings.
- Iterating on a string with the range operator iterates on the runes with the index corresponding to the starting index of the rune's byte sequence. To access a specific rune index (such as the third rune), convert the string into a []rune.
- Strings.TrimRight/strings. TrimLeft removes all the trailing/leading runes contained in a given set, whereas strings. Trimsuffix/strings TrimPrefix returns a string without a provided suffix/prefix.
- Concatenating a list of strings should be done with strings.Builder to prevent allocating a new string during each iteration.
- Remembering that the bytes package offers the same operations as the strings package can help avoid extra byte/string conversions.
- Using copies instead of subtrings can prevent memory leaks, as the string returned by a substring operation will be backed by the same byte array.
