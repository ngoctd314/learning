package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type person struct {
	Name     string `json:"name"`
	Address  string `json:"address"`
	Birthday string `json:"birthday"`
	ID       int    `json:"id"`
	Age      int    `json:"age"`
}

type handler struct {
	client http.Client
	url    string
}

func (h handler) getBody(delay time.Duration) (string, error) {
	resp, err := h.client.Get(h.url)
	_ = resp
	if err != nil {
		return "", err
	}
	// defer resp.Body.Close()
	// readAll := func(r io.Reader) ([]byte, error) {
	// 	b := make([]byte, 0, 5)
	// 	for {
	// 		fmt.Println("b", string(b))
	// 		if len(b) == cap(b) {
	// 			// Add more capacity (let append pick how much).
	// 			b = append(b, 0)[:len(b)]
	// 		}
	// 		n, err := r.Read(b[len(b):cap(b)])
	// 		b = b[:len(b)+n]
	// 		if err != nil {
	// 			if err == io.EOF {
	// 				err = nil
	// 			}
	// 			return b, err
	// 		}
	// 	}
	// }
	// body, _ := readAll(resp.Body)
	dst := bytes.NewBuffer(nil)
	fmt.Println(io.Copy(dst, resp.Body))

	fmt.Println(dst.String())

	return string(""), nil
}

func main() {
	httpclient := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network string, addr string) (net.Conn, error) {
				log.Println("re DialContext")
				return net.Dial(network, addr)
			},
			// Dial: func(network string, addr string) (net.Conn, error) {
			// 	log.Println("re Dial")
			// 	return net.Dial(network, addr)
			// },
			DisableKeepAlives:   false,
			MaxIdleConns:        5,
			MaxIdleConnsPerHost: 5,
			MaxConnsPerHost:     5,
			IdleConnTimeout:     time.Millisecond,
		},
		// Timeout: time.Second,
	}
	_ = httpclient
	h := handler{
		client: httpclient,
		url:    "http://localhost:8080",
	}
	now := time.Now()
	n := 1
	var wg sync.WaitGroup
	wg.Add(n * 2)

	go func() {
		fmt.Println(strings.Repeat("~", 20))
		for i := 0; i < n*2; i++ {
			go func() {
				defer wg.Done()
				h.getBody(time.Second * 1)
			}()
		}
		fmt.Println(strings.Repeat("~", 20))
	}()

	wg.Wait()
	fmt.Printf("since %fs", time.Since(now).Seconds())
}

func printAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB\n", bToMb(m.TotalAlloc))

}
func bToMb(b uint64) uint64 {
	return b / 1024
}
