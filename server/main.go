package main

import (
	"context"
	"log"
	"net"

	pb "microservice/blog/protoc"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5001"
var Collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

func main() {

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))

	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB: ", err)
	} else {
		log.Println("Connected to MongoDB successfully!")
	}

	if err != nil {
		log.Fatal(err)
	}
	Collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on : %v", err)
	}
	log.Printf("the server is running on %v", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on : %v", err)
	}

}
