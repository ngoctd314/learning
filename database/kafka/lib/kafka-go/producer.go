package kafkago

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func sequentialProduce(w *kafka.Writer) {

	// now := time.Now()
	w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("test"),
		},
	)

	// log.Printf("since %d ns", time.Since(now).Nanoseconds())
}
