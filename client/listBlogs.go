package main

import (
	"context"
	"io"
	"log"
	pb "microservice/blog/protoc"

	"google.golang.org/protobuf/types/known/emptypb"
)

func ListBlogs(c pb.BlogServiceClient) {
	log.Println("----ListBlogs was invoked-----")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Error while list the blogs")
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("the error while recv the Blog from the server ")
		}
		log.Println(res)
	}
}
