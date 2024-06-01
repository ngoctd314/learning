package kafkago

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/segmentio/kafka-go"
)

func Benchmark_sequenceProduce(b *testing.B) {
	w := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "test",
		BatchSize: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.WriteMessages(context.Background(),
			kafka.Message{
				Value: []byte("test"),
			},
		)
	}
}

func Benchmark_concurrencyProduce(b *testing.B) {
	w := &kafka.Writer{
		Addr:      kafka.TCP("localhost:9092"),
		Topic:     "test",
		BatchSize: 1,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency := func() {
			level := 6
			wg := sync.WaitGroup{}
			wg.Add(level)
			for c := 0; c < level; c++ {
				go func() {
					defer wg.Done()
					w.WriteMessages(context.Background(),
						kafka.Message{
							Value: []byte("test"),
						},
					)
				}()
			}
			wg.Wait()
		}
		concurrency()
	}
}

func Benchmark_sequentialBatch(b *testing.B) {
	w := &kafka.Writer{
		Addr:         kafka.TCP("localhost:9092"),
		Topic:        "test",
		BatchSize:    100,
		BatchTimeout: time.Second,
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// listMsg := []kafka.Message{}
		// for i := 0; i < 100; i++ {
		// 	listMsg = append(listMsg, kafka.Message{
		// 		Value: []byte(fmt.Sprintf("test-%d", i)),
		// 	})
		// }
		w.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(fmt.Sprintf("test-%d", i)),
		})
	}
}
