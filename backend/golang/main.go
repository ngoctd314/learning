package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

type publisher interface {
	Publish(ctx context.Context, position int) error
}

type publishHandler struct {
	pub publisher
}

func (h publishHandler) publishPosition(position int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	// What's the rationale for calling the cancel function as a defer function?
	// Internally, context.WithTimeout creates a goroutine that will be retained in memory for 4 seconds or until cancel
	// is called. Therefore, calling cancel as a defer function means that when we exit the parent function, the context will
	// be canceled, and the goroutine created will be stopped. It's a safeguard so that when we return, we don't leave retianed
	// objects in memory.
	defer cancel()

	return h.pub.Publish(ctx, position)
}

func callByReference(n1, n2 *int, s1, s2 *string) {
	*n1 = 3
	*n2 = 7
	fmt.Printf("n1(%p)=%v n2(%p)=%v\n", n1, *n1, n2, *n2)

	*s1 = "world"
	*s2 = "universe"

	fmt.Printf("s1(%p)=%v s2(%p)=%v\n", s1, *s1, s2, *s2)

}

func callByValue(n1, n2 int, s1, s2 string) {
	n1 = 3
	n2 = 7
	fmt.Printf("n1(%p)=%v n2(%p)=%v\n", &n1, n1, &n2, n2)

	s1 = "world"
	s2 = "universe"

	fmt.Printf("s1(%p)=%v s2(%p)=%v\n", &s1, s1, &s2, s2)
}

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

func baz() (x int) {
	defer func() {
		fmt.Println("RUN after")
		x = 10
	}()

	return foo()
}
func foo() int {
	fmt.Println("RUN")
	return 1
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

type detactContext struct {
	context.Context
}

func (d detactContext) Deadline() (time.Time, bool) {
	return time.Time{}, false
}

func (d detactContext) Done() <-chan struct{} {
	return nil
}

func (d detactContext) Err() error {
	return nil
}
