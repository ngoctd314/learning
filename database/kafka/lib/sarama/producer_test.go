package sarama

import (
	"sync"
	"testing"

	"github.com/IBM/sarama"
)

func Benchmark_sequenceProduce(b *testing.B) {
	w, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.SendMessage(&sarama.ProducerMessage{
			Topic: "test",
			Value: sarama.StringEncoder("test"),
		})
	}
}

func Benchmark_concurrencyProduce(b *testing.B) {
	w, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrency := func() {
			level := 6
			wg := sync.WaitGroup{}
			wg.Add(level)
			for c := 0; c < level; c++ {
				go func() {
					defer wg.Done()
					w.SendMessage(&sarama.ProducerMessage{
						Topic: "test",
						Value: sarama.StringEncoder("test"),
					})
				}()
			}
			wg.Wait()
		}
		concurrency()
	}
}

func Benchmark_sequentialBatch(b *testing.B) {
	w, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		listMsg := []*sarama.ProducerMessage{}
		for i := 0; i < 100; i++ {
			listMsg = append(listMsg, &sarama.ProducerMessage{
				Topic: "test",
				Value: sarama.StringEncoder("test"),
			})
		}
		w.SendMessages(listMsg)
	}
}
