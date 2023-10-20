package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
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

func fileHandle() {
	f, err := os.Open("profile.html")
	if err != nil {
		log.Fatal(err)
	}
	buf := bufio.NewReaderSize(f, 1)
	data := make([]byte, 114)
	n, err := buf.Read(data)
	fmt.Println(n, err)
	println(string(data))
	n, err = buf.Read(data)
	fmt.Println(n, err)
}

type person struct {
	id int
}

func (p *person) print() {
	fmt.Println(p.id)
}

type message struct {
	data       int
	disconnect bool
}

type emptyStruct struct{}

type emptyInterface interface{}

func main() {
	ch1 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch1 <- i
	}
	close(ch1)
	ch2 := make(chan int, 5)
	for i := 0; i < 5; i++ {
		ch2 <- i
	}
	close(ch2)
	ch := merge(ch1, ch2)
	for v := range ch {
		print(v, "\t")
	}
}

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break // break select
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break // break select
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

func printAr(v int) {
	fmt.Println(v)
}

func udpServer() {
	addr := net.UDPAddr{
		Port: 8080,
		IP:   net.ParseIP("127.0.0.1"),
	}
	s, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatal(err)
	}
	p := make([]byte, 2048)
	for {
		_, ar, err := s.ReadFromUDP(p)
		fmt.Printf("Read a message from %v %s \n", ar, p)
		if err != nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(s, ar)
	}
}
func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_, err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr)
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}

func udpClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	conn.Write([]byte("Hello"))

	p := make([]byte, 2048)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}

func tcpServer() {
	ls, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	defer ls.Close()

	for {
		conn, err := ls.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleIncomingRequest(conn)
	}

}

func handleIncomingRequest(conn net.Conn) {
	p := make([]byte, 1024)
	conn.Read(p)
	fmt.Println("read p: ", string(p))
	conn.Write([]byte("Response"))
	conn.Close()
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

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	if r <= 10 {
		for i := 0; i <= r; i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	var mid int
	if nums[l] > nums[r] {
		for l <= r {
			mid = (l + r) / 2
			if nums[mid] > nums[mid+1] {
				break
			}
			if nums[mid] > nums[0] {
				l = mid + 1
			} else if nums[mid] < nums[0] {
				r = mid - 1
			}
		}
	}

	l1, r1 := 0, mid
	l2, r2 := mid+1, len(nums)-1
	if mid > 0 && target >= nums[l1] && target <= nums[r1] {
		for l1 <= r1 {
			mid := (l1 + r1) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				l1 = mid + 1
			} else {
				r1 = mid - 1
			}
		}
	} else if target >= nums[l2] && target <= nums[r2] {
		for l2 <= r2 {
			mid := (l2 + r2) / 2
			if nums[mid] == target {
				return mid
			} else if nums[mid] < target {
				l2 = mid + 1
			} else {
				r2 = mid - 1
			}
		}
	}

	return -1
}
