# Dummy questions

### 1. Make main goroutine Deadlock after t seconds

```go
var empty = struct{}{}

func main() {
	ch := make(chan struct{}, 1)
	// Goexit terminates the goroutine that calls it. No other goroutine is affected.
	t := time.Second * 2
	now := time.Now()
	go func() {
		defer func() {
			fmt.Printf("runtime.Goexit() after %s\n", time.Since(now).String())
		}()
		time.Sleep(t)
		// Main goroutine deadlock after t seconds
		runtime.Goexit()
		ch <- empty
	}()

	<-ch
	fmt.Println("DEADLOCK")
}

```

### 2. How to make main goroutine exist but child goroutine still run

```go
func main() {
	go func() {
		time.Sleep(time.Second * 2)
		fmt.Println("Continue running")
	}()

	runtime.Goexit()
}
```

### 3. When child goroutine exist, return

### 4. Result of this program

```go
func main() {
	fmt.Println(baz())
}

func baz() (x int) {
	defer func() {
		fmt.Println("RUN defer")
		x = 10
	}()

	return foo()
}
func foo() int {
	fmt.Println("RUN foo")
	return 1
}
```

### 5. Result of this program

```go
func main() {
    var (
        b *bytes.Buffer
        w io.Writer
    )

    w = b
    fmt.Println(w == nil)
}
```

### 6. Result or this program

```go
func main() {
	ch := foo1()
	go func() {
		for v := range ch {
			_ = v
		}
		fmt.Println("unreachable")
	}()
	go func() {
		time.Sleep(time.Second * 100)
	}()

	select {}
}
func foo1() chan int {
	ch := make(chan int)
	return ch
}
```

### 7. Result of this program

```go
func main() {
	a := make([]int, 0, 5)
	b := a

	b = append(b, 1, 2, 3, 4, 5)
	fmt.Println(a, b)
}
```

### 8. Result of this program

```go
type person struct {
	id int
}

func (p person) print() {
	fmt.Println(p.id)
}

func main() {
	ar := []person{}
	for i := 0; i < 5; i++ {
		ar = append(ar, person{
			id: i,
		})
	}

	for _, v := range ar {
		go v.print()
	}
	runtime.Goexit()
}
```

### 9. Result of this program

**For loop**
```go
type person struct {
	id int
}

func (p *person) print() {
	fmt.Println(p.id)
}

func main() {
	ar := []person{}
	for i := 0; i < 5; i++ {
		ar = append(ar, person{
			id: i,
		})
	}

	for _, v := range ar {
		go v.print()
	}
	runtime.Goexit()
}
```

**Data race**
```go
var data int // 1

go func() { // 2
    data++ // 3
}() // 4
if data == 0 { // 5
    fmt.Printf("the value is %v.\n", data)
}
```

**Slice len, cap**
```go
func main() {
	s := make([]int, 0, 1)

	_ = append(s, 10)
	fmt.Println(s)
}
```

**This is data race, why it is also race condition**
```go
func main() {
	s := make([]int, 0, 1)

	go func() {
		s1 := append(s, 1)
		fmt.Println("s1:", s1)
	}()

	go func() {
		s2 := append(s, 2)
		fmt.Println("s2:", s2)
	}()

	time.Sleep(time.Millisecond)
}
```

**Improve data race**
```go
func main() {
	s := make([]int, 0, 1)

	go func() {
		s1 := append(s, 1)
		fmt.Println("s1:", s1)
	}()

	go func() {
		s2 := append(s, 2)
		runtime.Gosched()
		fmt.Println("s2:", s2)
	}()

	time.Sleep(time.Millisecond * 2)
}
```

## Column in MySQL name VARCHAR(10), query this column in Golang, max len(name) can be?

40 or more

```go
/*
CREATE TABLE `varchar_go` (

	`name` varchar(10) DEFAULT NULL

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
*/
func main() {
	rows, _ := mysqlConn.Query("SELECT name from varchar_go")
	defer rows.Close()
	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(len(name))
	}
}
```

## Scan sql query to map (without knowing about struct)?
