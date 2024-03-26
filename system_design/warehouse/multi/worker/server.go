package main

import (
	"context"
	"warehouse/pb"
)

type workerServer struct {
	pb.UnimplementedWorkerServer
	r *repository
}

func newWorkerServer(r *repository) *workerServer {
	return &workerServer{
		r: r,
	}
}

func (s *workerServer) ParOr(ctx context.Context, req *pb.ParOrRequest) (*pb.ParOrResponse, error) {
	params := make([]any, len(req.Ids))
	for i := range req.Ids {
		params[i] = req.Ids[i]
	}

	bitmap := s.r.getOrBitmap(params)

	return &pb.ParOrResponse{
		Bitmap: bitmap,
	}, nil
}
