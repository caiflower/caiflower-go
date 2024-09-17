package service

import (
	"caiflower.com/caiflower-go/pkg/grpc/server/api/pb"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type HelloImpl struct {
	pb.UnimplementedIServiceServer
}

func (h *HelloImpl) Search(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	fmt.Printf("req is %v\n", req)
	if req.Query == "1" {
		if len(req.Hobby) >= int(req.GetPageNumber()) {
			return &pb.SearchResponse{Code: 1, Message: req.Hobby[req.PageNumber-1]}, nil
		} else {
			return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("pageNumber %v out of range", req.GetPageNumber()))
		}
	} else if req.Query == "2" {
		return &pb.SearchResponse{Code: 1, Message: strings.Join(req.Hobby, ",")}, nil
	}
	return nil, status.Errorf(codes.OutOfRange, fmt.Sprintf("query %v is not impl", req.Query))
}
