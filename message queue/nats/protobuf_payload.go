package main

import (
	"fmt"
	"learning-nats/pb"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

// NATS message payloads are byte slices, so any kind of serialization
// strategy can be applied.

func protobufRequestReply() {
	// create an unauthenticated connection to NATS
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Drain()

	nc.Subscribe("greet", func(msg *nats.Msg) {
		var req pb.GreetRequest
		proto.Unmarshal(msg.Data, &req)

		rep := pb.GreetReply{
			Text: fmt.Sprintf("hello %q!", req.Name),
		}
		data, _ := proto.Marshal(&rep)
		fmt.Println("reply", msg.Reply)
		msg.Respond(data)
	})

	req := pb.GreetRequest{
		Name: "joe",
	}
	data, _ := proto.Marshal(&req)
	msg, _ := nc.Request("greet", data, time.Second)

	var rep pb.GreetReply
	proto.Unmarshal(msg.Data, &rep)

	fmt.Printf("reply: %s\n", rep.Text)

}
