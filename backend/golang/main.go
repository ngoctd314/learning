package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-API-VERSION", "1.0")
	b, _ := io.ReadAll(r.Body)
	_, _ = w.Write(append([]byte("hello "), b...))
	w.WriteHeader(http.StatusCreated)
}

func foo() {
	acc := 0
	for i := 0; i < 1e6; i++ {
		acc += i
	}
}

type Foo struct {
	b1 byte
	_  [7]byte // Added by compiler
	i  int64
	b2 byte
	_  [7]byte // Added by the compiler
}

func main() {
	fmt.Println(fmt.Sprint(rune('a')))
	fmt.Println(string(rune('a')))
}

type Input struct {
	a int64
	b int64
}
type Result struct {
	sumA int64
	sumB int64
}
type Result2 struct {
	sumA int64
	_    [56]byte // padding
	sumB int64
}

func count1(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)
	result := Result{}

	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
	}()

	wg.Wait()
	return result
}
func count2(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)
	chA, chB := make(chan int64, 1), make(chan int64, 1)

	go func() {
		defer wg.Done()
		var sumA int64
		for i := 0; i < len(inputs); i++ {
			sumA += inputs[i].a
		}
		chA <- sumA
	}()

	go func() {
		defer wg.Done()
		var sumB int64
		for i := 0; i < len(inputs); i++ {
			sumB += inputs[i].b
		}
		chB <- sumB
	}()

	wg.Wait()
	return Result{
		sumA: <-chA,
		sumB: <-chB,
	}
}

func count3(inputs []Input) Result2 {
	wg := sync.WaitGroup{}
	wg.Add(2)
	rs := Result2{}

	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			rs.sumA += inputs[i].a
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < len(inputs); i++ {
			rs.sumB += inputs[i].b
		}
	}()

	wg.Wait()
	return rs
}

func expensiveSetup() {
	for i := 0; i < 1e10; i++ {
	}
}

func calculateSum512(s [][512]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}

	return sum
}

func calculateSum513(s [][513]int64) int64 {
	var sum int64
	for i := 0; i < len(s); i++ {
		for j := 0; j < 8; j++ {
			sum += s[i][j]
		}
	}

	return sum
}

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}
