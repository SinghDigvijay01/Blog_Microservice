package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"
)

func CreateBlog(c pb.BlogServiceClient) string {

	log.Println("----CreateBlog was invoked-----")

	blog := &pb.Blog{
		AuthorId: "Digvijay Singh ",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected Error %v", err)
	}

	log.Printf("The Blod Id is %v", res.Id)

	return res.Id
}
