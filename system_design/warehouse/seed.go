package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func genData() {
	f, err := os.OpenFile("data.txt", os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	ln := 2000
	line := 2
	for i := 0; i < line; i++ {
		s := make([]string, 0, ln)
		for i := 0; i < ln; i++ {
			rn := rand.Intn(24 * 63 * 63)
			s = append(s, strconv.Itoa(rn))
		}
		n, err := f.Write([]byte(strings.Join(s, " ") + "\n"))
		fmt.Println(n, err)
	}
}

func insertData(coll *mongo.Collection) {
	f, err := os.OpenFile("data.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := relateItemRepository{collection: coll}
	r.drop()

	data, _ := io.ReadAll(f)
	ar := strings.Split(string(data), "\n")
	for i, a := range ar {
		r.createItem(context.Background(), uint32(i+1))
		vs := strings.Split(a, " ")
		for _, v := range vs {
			vi, convErr := strconv.Atoi(v)
			if convErr == nil {
				if err := r.insertRelate(context.Background(), uint32(i+1), uint32(vi)); err != nil {
					log.Printf("insertRelate error (%v)\n", err)
				}
			}
		}
	}
}
