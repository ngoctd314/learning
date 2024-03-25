package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
	"warehouse/pb"

	"github.com/RoaringBitmap/roaring"
)

type Aggregator struct {
	workers []pb.WorkerClient
}

func NewAggregator(workers ...pb.WorkerClient) *Aggregator {
	return &Aggregator{
		workers: workers,
	}
}

func (s *Aggregator) CountDistinct(ctx context.Context, itemIDs []uint32) {
	numberOfWorkers := len(s.workers)
	listBitmap := make([]*roaring.Bitmap, 0, numberOfWorkers)
	mu := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(numberOfWorkers)
	chunk := len(itemIDs) / numberOfWorkers
	for i, worker := range s.workers {
		cp := worker
		lower, upper := i*chunk, (i+1)*chunk
		if upper > len(itemIDs) {
			upper = len(itemIDs)
		}

		go func() {
			defer wg.Done()
			resp, err := cp.ParOr(ctx, &pb.ParOrRequest{
				Ids: itemIDs[lower:upper],
			})
			if err != nil {
				log.Printf("error when ParOr ('%v')\n", err)
				return
			}
			rb := roaring.New()
			if _, parseBitmapErr := rb.FromBuffer(resp.Bitmap); parseBitmapErr != nil {
				log.Println("parseBitmapErr", err)
				return
			}

			mu.Lock()
			listBitmap = append(listBitmap, rb)
			mu.Unlock()
		}()
	}
	wg.Wait()

	ti := time.Now()
	rs := roaring.ParOr(5, listBitmap...).GetCardinality()
	fmt.Println(rs)
	fmt.Println("ParOr time:", time.Since(ti))
}
