package sarama

import (
	"github.com/IBM/sarama"
)

func sequentialProduce(w sarama.SyncProducer) {
	// now := time.Now()
	w.SendMessage(&sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("test"),
	})

	// log.Printf("since %d ns", time.Since(now).Nanoseconds())
}
