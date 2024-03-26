package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/rand"
	"sort"
	"time"
	"warehouse/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// dial workerConn
	opts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(math.MaxInt32),
			grpc.MaxCallRecvMsgSize(math.MaxInt32),
		),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	workerConn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal("can't dial worker: ", err)
	}
	ctx := context.Background()

	// init client stub
	worker := pb.NewWorkerClient(workerConn)

	aggregator := NewAggregator(worker)

	s := make(map[uint32]struct{})
	items, j := uint32(10000), uint32(0)
	for j < items {
		n := uint32(rand.Intn(1e7))
		for n == 0 {
			n = uint32(rand.Intn(1e7))
		}
		s[j] = struct{}{}
		j++
	}
	paramsInt := make([]uint32, 0, items)
	for k := range s {
		paramsInt = append(paramsInt, k)
	}
	sort.Slice(paramsInt, func(i, j int) bool { return paramsInt[i] < paramsInt[j] })

	now := time.Now()
	aggregator.CountDistinct(ctx, paramsInt)
	fmt.Println("since: ", time.Since(now))
}
