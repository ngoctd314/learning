package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "key", "val")
		// ctx = context.WithoutCancel(ctx)
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("ctx.Err() %s\n", ctx.Err())
				default:
					fmt.Println(ctx.Value("key"))
				}
			}
		}(ctx)
		time.Sleep(time.Second * 5)

		w.Write([]byte("OK"))
	})
	http.ListenAndServe(":8080", nil)
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
