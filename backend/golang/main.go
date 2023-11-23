package main

import (
	"fmt"
	"runtime"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type Input struct {
	a int64
	b int64
}

type Result struct {
	sumA int64
	sumB int64
}

func count(inputs []Input) Result {
	wg := sync.WaitGroup{}
	wg.Add(2)

	result := Result{}
	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumA += inputs[i].a
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < len(inputs); i++ {
			result.sumB += inputs[i].b
		}
		wg.Done()
	}()

	wg.Wait()

	return result
}

// 202311211415
func main() {
	n := 10000000
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			rs := count([]Input{
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
				{a: 1, b: 1},
			})
			if rs.sumA != rs.sumB || rs.sumA != 60 || rs.sumB != 60 {
				fmt.Println(rs.sumA, rs.sumB)
			}
		}()
	}
	wg.Wait()
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
