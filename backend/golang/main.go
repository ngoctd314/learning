package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

func main() {
	ctx := context.Background()
	var m sync.Mutex
	key := "key"
	wg := sync.WaitGroup{}
	ctx = context.WithValue(ctx, key, 0)

	n := 1000
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			m.Lock()
			ctx = context.WithValue(ctx, key, ctx.Value(key).(int)+1)
			m.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(ctx.Value(key))
}

func cancelCtx(ctx context.Context) {
	fmt.Printf("ctx.Err() %v\n", ctx.Err())
	select {
	case <-ctx.Done():
		fmt.Println("<-ctx.Done()")
	}
}

func convPointer(i *int) {
	fmt.Printf("addr1 %p\n", i)
	ii := *i
	foobyval(ii)
}

func foobyval(n int) {
	fmt.Println()
	// println(n)
	fmt.Printf("addr2 %p\n", &n)
}

func m() {
	x := 2
	fmt.Printf("%p\n", &x)
	defer func() {
		fmt.Printf("%p\n", &x)
	}()
}

func fn() (string, error) {
	rs := "ngoctd"
	err := errors.New("err")
	defer func() {
		rs = "xyz"
		err = errors.New("invalid")
	}()
	return rs, err

}

type Person struct {
	Name string
}

func sequentialVer() (int64, float64) {
	return 0, 0
}

// func printAlloc() {
// var m runtime.MemStats
// ReadMemStats populates m with memory allocator statistic
// The returned memory allocator statistics are up to date as of the
// call to ReadMemStats. This is in constrast with a heap profileff
// which is a snapshot as of the most recently completed garbage
// collection cycle.
// runtime.ReadMemStats(&m)
// fmt.Printf("%d KB\n", m.Alloc/1024)
// }
