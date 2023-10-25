package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type Event struct {
	ID int `json:"id"`
	time.Time
}

type foo struct{}

func (foo) MarshalJSON() ([]byte, error) {
	return []byte(`"foo"`), nil
}

func main() {
	b, err := json.Marshal(foo{})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	e := Event{
		ID:   1234,
		Time: time.Now(),
	}
	data, _ := json.Marshal(e)
	fmt.Println(string(data))
}

func (e Event) MarshalJSON() ([]byte, error) {
	return []byte(`"ABC"`), nil
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
