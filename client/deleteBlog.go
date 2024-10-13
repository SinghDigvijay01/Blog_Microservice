package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("----DeleteBlog was invoked-----")

	blogId := &pb.BlogId{Id: id}
	_, err := c.DeleteBlog(context.Background(), blogId)

	if err != nil {
		log.Fatalf("Error while deleting the blog with  %v", id)
	}

}
