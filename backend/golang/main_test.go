package main

import (
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestA(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
	fmt.Println("TestA")
}

func TestB(t *testing.T) {
	t.Parallel()
	time.Sleep(time.Second)
	fmt.Println("TestB")
}

func TestC(t *testing.T) {
	time.Sleep(time.Second)
	fmt.Println("TestC")
}
