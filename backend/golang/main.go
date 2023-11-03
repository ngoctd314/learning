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

type h struct {
}

func (h) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 2) // exceed time.Second
	defer r.Body.Close()
	// only load from meme, not from network
	data, _ := io.ReadAll(r.Body)

	w.Write([]byte(fmt.Sprintf("forward body: %s", string(data))))
}

func main() {
	httpclient := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				log.Println("re-dial")
				return net.Dial(network, addr)
			},
			DisableKeepAlives: false,
			MaxConnsPerHost:   100,
			MaxIdleConns:      1000,
			IdleConnTimeout:   time.Second * 10,
		},
	}
	for i := 0; i < 10; i++ {
		go httpclient.Get("http://localhost:8080")

	}
	time.Sleep(time.Second * 5)
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
