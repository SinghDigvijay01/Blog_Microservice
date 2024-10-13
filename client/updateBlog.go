package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"
)

func UpdateBlog(c pb.BlogServiceClient, id string) {
	log.Println("------UpdateBlog was invoked------")
	blog := &pb.Blog{
		Id:       id,
		AuthorId: "digvijay2",
		Title:    "title2",
		Content:  "Content2",
	}
	_, err := c.UpdateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Error While Update the Blog %v", err)
	}
}
