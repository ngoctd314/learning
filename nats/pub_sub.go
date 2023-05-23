package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// core NATS publish-subscribe behavior. This is the fundamental pattern that all other NATS patterns and higher-level
// APIs build upon. There are a few takeaways from this example:
//
// Delivery is at most one
// There are two circumstances when a published message won't be delivered to a subscriber
// - The subscriber does not have an active connection to the server
// - There is a network interruption where the message is ultimately dropped
func pubSub() {
	// Create an unauthentication connection to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	// Drain is a safe way to ensure all buffered messages that were published
	// are sent and all buffered messages received on a subscription are processed being
	// closing the connection.
	// Drain will put a connection into a drain state. All subscriptions will inmediately
	// be put into a drain state. Upon completion, the publishers will be drained and can
	// not publish any additional messages. Upon draining of the publishers, the connection
	// will be closed.
	defer nc.Drain()

	// SubscribeSync will express interest on the given subject. Messages will be
	// received synchronously using Subscription.NextMsg()
	// Create a subscription on the greet.* wildcard
	sub, _ := nc.SubscribeSync("greet.*")

	// Publish publishes the data argument to the given subject. The data argument is left
	// untouched and needs to be correctly interpreted on the receiver.
	// Messages are published to subjects. Although there are no subscribers
	// this will be published successfully.
	nc.Publish("greet.joe", []byte("hello"))

	// NextMsg will return the next message available to a synchronous subscriber
	// or block until one is available. An error is returned if the subscription is invalid (ErrBadSubscription)
	// the connection is closed (ErrConnectionClosed), the timeout is reached (ErrTimeout),
	// or if there were no responders (ErrNoResponders) when used in the context of a request/reply.
	for {
		msg, err := sub.NextMsg(time.Second * 5)
		if err != nil {
			log.Println("sub error", err)
		}
		if msg != nil {
			log.Println("receive msg", string(msg.Data))
		}
	}
}
