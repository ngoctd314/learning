package pubsub

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Subscribe ...
func Subscribe() {
	log.Println("Subscribe")
	nc, _ := nats.Connect(nats.DefaultURL)
	sub, err := nc.SubscribeSync(">")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msg, err := sub.NextMsg(time.Second)
		if err != nil {
			log.Println(err)
		}
		if msg != nil {
			log.Printf("Receive msg %s\n", string(msg.Data))
		}
	}
}
