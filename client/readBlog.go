package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"
)

func ReadBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Printf("----ReadBlog was invoked-----")

	blogId := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), blogId)

	if err != nil {
		log.Fatalf("Error occurs during the reading %v", err)
	}
	log.Println(res)
	return res
}
