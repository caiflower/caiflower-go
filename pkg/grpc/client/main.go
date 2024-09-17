package main

import (
	"caiflower.com/caiflower-go/pkg/grpc/client/api/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 远程调用
	client := pb.NewIServiceClient(conn)
	search, err := client.Search(context.Background(), &pb.SearchRequest{})
	printResponse(search, err)

	search, err = client.Search(context.Background(), &pb.SearchRequest{Query: "1", PageNumber: 2, Hobby: []string{"english", "math"}})
	printResponse(search, err)

	search, err = client.Search(context.Background(), &pb.SearchRequest{Query: "2", Hobby: []string{"english", "math"}})
	printResponse(search, err)
}

func printResponse(response *pb.SearchResponse, err error) {
	if err != nil {
		fmt.Printf("gRPC err = %v\n", err)
	} else {
		fmt.Printf("code = %v, message = %v\n", response.Code, response.Message)
	}
}
