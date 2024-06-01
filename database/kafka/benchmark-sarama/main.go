package main

import (
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	produceSequential()
}

func produceSequential() {
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)

	now := time.Now()
	wg := sync.WaitGroup{}
	concurrency := 6
	workload := 100
	wg.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < workload; i++ {
				producer.SendMessage(&sarama.ProducerMessage{
					Topic: "test",
					Value: sarama.StringEncoder("test"),
				})
			}
		}()
	}
	wg.Wait()

	log.Printf("since %d ns", time.Since(now).Nanoseconds())
}
