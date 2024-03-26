package main

import (
	"log"
	"net"
	"warehouse/pb"

	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	// init aggregator
	r := newRepostiory()
	worker := newWorkerServer(r)
	pb.RegisterWorkerServer(grpcServer, worker)

	// listen and serve
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("cannot create listener: ", err)
	}

	log.Printf("start gRPC server on %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot create grpc server: ", err)
	}
}
