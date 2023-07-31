package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

// The request-reply pattern allows a client to send a message and expect a reply
// of some kind. In practice, the request message will either be a command,
// Unlike request-reply constrained protocols like HTTP, NATS is not limited to
// a strict point-to-point interfaction between a client and server. The request-reply
// pattern is built on top of the core publish-subscribe model.
// By default, this means any one of subscribers could be a responder and reply to the client
// However, because NATS is not limited to point-to-point interactions, the client
// could indicate to NATS that multiple replies shoule be allowed.
func requestReply() {
	// create an unauthenticated connection to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	// In addition to vanilla publish-request, NATS supports request-reply
	// interactions as well. Under the cover, this is just an optimized
	// pair of publish-subscribe operations.
	// The request handler is just a subscription that responds to a message
	// sent to it. This kind of subscription is called a service.
	sub, _ := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		name := msg.Subject[6:]
		// Respond allows a convenient way to respond to requests in service based subscriptions.
		fmt.Println("receive", string(msg.Data), msg.Reply, msg)
		msg.Respond([]byte("hello, " + name))
	})

	// New we can use the built-in Request method to do the service request.
	// We simply pass a nil body since that is being used right now. In addition,
	// we need to specify a timeout since with a request we are waiting for the reply
	// and we likely don't want to wait forever.
	req, _ := nc.Request("greet.joe", []byte("i am joe"), time.Second)
	fmt.Println(string(req.Data), req.Subject)

	req, _ = nc.Request("greet.bob", nil, time.Second)
	fmt.Println(string(req.Data), req.Reply)

	sub.Unsubscribe()

	_, err := nc.Request("greet.joe", nil, time.Second)
	fmt.Println(err)
}
