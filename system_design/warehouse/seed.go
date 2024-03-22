package main

import (
	"context"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

func genData() {
	f, err := os.OpenFile("data.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i := 0; i < lines; i++ {
		s := make([]string, 0, lenOnLine)
		for i := 0; i < lenOnLine; i++ {
			rn := rand.Intn(chunk * 63 * 63)
			s = append(s, strconv.Itoa(rn))
		}
		_, err := f.Write([]byte(strings.Join(s, " ") + "\n"))
		if err != nil {
			log.Fatal(err)
		}
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
		vs := strings.Split(a, " ")
		tmp := make([]uint32, 0)
		for _, v := range vs {
			vi, convErr := strconv.Atoi(v)
			if convErr == nil {
				tmp = append(tmp, uint32(vi))
			}
		}
		if err := r.insertRelate(context.Background(), uint32(i+1), tmp...); err != nil {
			log.Printf("insertRelate error (%v)\n", err)
		}
	}
}
