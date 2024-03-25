package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/RoaringBitmap/roaring"
)

func insertFromFile() {
	r := NewRepository()

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	strs := strings.Split(string(data), "\n")

	for _, str := range strs {
		records := strings.Split(str, " ")
		if len(records) == 0 {
			return
		}
		if len(records[0]) == 0 {
			return
		}
		id, err := strconv.Atoi(records[0])
		if err != nil {
			log.Fatalf("invalid format, first column must be integer id, got: %v", records[0])
			return
		}
		relateItems := []uint32{}
		for i := 1; i < len(records); i++ {
			v, err := strconv.Atoi(strings.TrimSpace(records[i]))
			if err != nil {
				log.Fatalf("invalid format, relate item must be integer, got: %v", records[i])
				return
			}
			relateItems = append(relateItems, uint32(v))
		}
		rb := roaring.New()
		rb.AddMany(relateItems)
		b, err := rb.ToBytes()
		if err != nil {
			log.Fatalf("error occur when convert relate item to bytes: %v", err)
			return
		}
		r.Insert(uint32(id), b)
	}
}
