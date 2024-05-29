package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	// make a writer that produces to topic-A, using the least-bytes distribution
	w := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "test",
		BatchSize: 1,
	}

	now := time.Now()
	wg := sync.WaitGroup{}
	concurrency := 6
	workload := 100
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < workload; i++ {
				w.WriteMessages(context.Background(),
					kafka.Message{
						Value: []byte("test"),
					},
				)
			}
		}()

	}
	wg.Wait()

	// 2024/05/28 13:51:19 since 90677697 ns
	// 2024/05/28 13:53:19 since 20200483 ns
	log.Printf("since %d ns", time.Since(now).Nanoseconds())
}

// 2024/05/28 11:34:22 since 952098 ns
// 2024/05/28 11:34:21 since 1007838571 ns
