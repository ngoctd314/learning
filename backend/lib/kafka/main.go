package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func main() {
	// reader()
	mechanism, err := scram.Mechanism(scram.SHA256, "ghtk_testing_rw", "Te1kgHb1UlqeJDvtPRjqPsx")
	if err != nil {
		panic(err)
	}
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}

	w := &kafka.Writer{
		Addr:  kafka.TCP("10.110.69.50:9092", "10.110.69.51:9092", "10.110.69.52:9092", "10.110.69.53:9092", "10.110.69.54:9092"),
		Topic: "db_slow_query",
		// Topic:     "ecom_slow_query",
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}

	f, _ := os.Open("./data.txt")
	data, _ := io.ReadAll(f)
	f.Close()

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Value: data,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}

func reader() {
	mechanism, err := scram.Mechanism(scram.SHA256, "ghtk_testing_rw", "Te1kgHb1UlqeJDvtPRjqPsx")
	if err != nil {
		panic(err)
	}

	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"10.110.69.50:9092", "10.110.69.51:9092", "10.110.69.52:9092", "10.110.69.53:9092", "10.110.69.54:9092"},
		GroupID:     "groupid4",
		Topic:       "db_slow_query",
		StartOffset: 0,
		// Partition: 3,
		Dialer: dialer,
	})
	r.SetOffset(0)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("ReadMessage", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

// Brokers:     []string{"10.110.69.50:9092", "10.110.69.51:9092", "10.110.69.52:9092", "10.110.69.53:9092", "10.110.69.54:9092"},

func writer() {
	mechanism, err := scram.Mechanism(scram.SHA256, "ghtk_testing_rw", "Te1kgHb1UlqeJDvtPRjqPsx")
	if err != nil {
		panic(err)
	}
	sharedTransport := &kafka.Transport{
		SASL: mechanism,
	}

	w := &kafka.Writer{
		Addr: kafka.TCP("10.110.69.50:9092", "10.110.69.51:9092", "10.110.69.52:9092", "10.110.69.53:9092", "10.110.69.54:9092"),
		// Topic: "db_slow_query",
		Topic:     "ecom_slow_query",
		Balancer:  &kafka.Hash{},
		Transport: sharedTransport,
	}

	f, _ := os.Open("./data.txt")
	data, _ := io.ReadAll(f)
	f.Close()

	err = w.WriteMessages(context.Background(),
		kafka.Message{
			Value: data,
		},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := w.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
