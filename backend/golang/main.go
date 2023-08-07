package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func tmp() {
	return
}

var empty = struct{}{}

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

func fn2() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("RUN")
		}
	}()
}

func main() {
	fn2()
	time.Sleep(time.Second)
}

type goExit struct{}

func (goExit) GetDeferFunc() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()
	defer func() {
		panic("defer panic")
	}()

	runtime.Goexit()
}

// https://www.sobyte.net/post/2022-07/go-sync-cond/
