package main

import (
	"fmt"
	"time"
)

var sharedRsc = false

type Address struct {
	City string `json:"city,omitempty"`
}

// person struct Helo World
// http://localhost
type person struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Address Address `json:"address,omitempty"`
}

func fn() error {
	return nil
}

func fn1() int {
	return 1
}

func main() {
	fmt.Println("Hello World")
}

func event() {
	for conditionTrue() == false {
		// cần 1 function hoặc cách gì đó để goroutine có thể sleep cho đến khi có 1 tín hiệu thực thi
		time.Sleep(time.Millisecond)
		fmt.Println("RUN")
	}
}

var cnt = 0

func conditionTrue() bool {
	defer func() {
		cnt++
	}()

	return cnt == 10
}

// https://www.sobyte.net/post/2022-07/go-sync-cond/
