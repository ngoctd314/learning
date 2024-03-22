package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_data(t *testing.T) {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, _ := io.ReadAll(f)
	ar := strings.Split(string(data), "\n")
	s := make(map[int]struct{})
	for _, a := range ar {
		vs := strings.Split(a, " ")
		for _, v := range vs {
			vi, convErr := strconv.Atoi(v)
			if convErr == nil {
				s[vi] = struct{}{}
			}
		}
	}
	fmt.Println(len(s))
}
