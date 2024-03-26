package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func count() {
	r := NewRepository()
	if len(os.Args) < 2 {
		log.Fatal("ids is required with format 1,2,...,x")
	}
	ids := strings.Split(os.Args[1], ",")
	if len(ids) == 0 {
		log.Fatal("ids is required with format 1,2,...,x")
	}

	params := []any{}
	for _, id := range ids {
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("ids is required with format 1,2,...,x")
		}
		params = append(params, idInt)
	}

	fmt.Println("result: ", r.CountDistinctRelate(params...))
}
