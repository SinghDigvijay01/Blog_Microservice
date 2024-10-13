package main

import (
	"log"
	pb "microservice/blog/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:5001"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connect : %v", err)
	}
	defer conn.Close()
	c := pb.NewBlogServiceClient(conn)
	id := CreateBlog(c)
	ReadBlog(c, id)
	UpdateBlog(c, id)
	ListBlogs(c)
	deleteBlog(c, id)
}
