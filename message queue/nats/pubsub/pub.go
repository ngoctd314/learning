package pubsub

import (
	"log"

	"github.com/nats-io/nats.go"
)

// Publish ...
func Publish() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	nc.Subscribe("test", func(msg *nats.Msg) {
		err = nc.Publish("test", []byte("hello world"))
		if err != nil {
			log.Fatal(err)
		}
	})
}
