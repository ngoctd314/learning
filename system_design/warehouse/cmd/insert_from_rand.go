package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/RoaringBitmap/roaring"
)

func insertFromRand() {
	insertTimes := os.Getenv("INSERT_TIMES")
	relates := os.Getenv("RELATE")
	recordPerInsert := os.Getenv("RECORDS_FOR_INSERT")

	insertTimesNum, err := strconv.Atoi(insertTimes)
	if err != nil {
		log.Fatal("records must be integer", insertTimes)
	}

	relatesNum, err := strconv.Atoi(relates)
	if err != nil {
		log.Fatal("relates must be integer", insertTimes)
	}

	recordPerInsertNum, err := strconv.Atoi(recordPerInsert)
	if err != nil {
		log.Fatal("relates must be integer", insertTimes)
	}

	r := NewRepository()
	for i := 0; i < insertTimesNum; i++ {
		lRecords := make([]any, recordPerInsertNum)
		for j := 0; j < recordPerInsertNum; j++ {
			rb := roaring.New()
			relateItemList := make([]uint32, relatesNum)
			for k := 0; k < relatesNum; k++ {
				relateItemList[k] = uint32(rand.Intn(1e8) + 1)
			}
			rb.AddMany(relateItemList)
			b, err := rb.ToBytes()
			if err != nil {
				log.Fatal(err)
				return
			}
			lRecords[j] = b
		}
		r.InsertMany(lRecords)
		log.Printf("complete insert: %d records, each record  is related to %d items\n", recordPerInsertNum, relatesNum)
	}
}
