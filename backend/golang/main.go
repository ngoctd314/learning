package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
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

type slice []int

func (s *slice) add(element int) {
	*s = append(*s, element)
}

type customer struct {
	data *data
}

type data struct {
	balance float64
}

func (c customer) add(operation float64) {
	c.data.balance += operation
}

func main() {
	c := customer{
		data: &data{
			balance: 0,
		},
	}
	c.add(50)
	fmt.Println(c.data.balance)
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

func getBytes(reader io.Reader) ([]byte, error) {
	b, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	// call sanitize
	return bytes.TrimSpace(b), nil
}

func concat(values ...string) string {
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
}

func concat1(values ...string) string {
	s := ""
	for _, value := range values {
		s += value
	}

	return s
}

func concat2(values ...string) string {
	sb := strings.Builder{}
	for _, value := range values {
		_, _ = sb.WriteString(value)
	}

	return sb.String()
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
