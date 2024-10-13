package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) ListBlogs(in *emptypb.Empty, stream grpc.ServerStreamingServer[pb.Blog]) error {
	log.Println("The listBlog was invoked ")

	curr, err := Collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(codes.Internal, "Internal Server Error")
	}

	defer curr.Close(context.Background())

	for curr.Next(context.Background()) {
		data := &BlogItem{}
		err := curr.Decode(data)

		if err != nil {
			return status.Errorf(codes.Internal, "Error while decoding the data from the mongoDb")
		}
		stream.Send(documentToBlog(data))
	}

	if curr.Err(); err != nil {
		return status.Errorf(codes.Internal, "Unknown internal Error")
	}
	return nil
}
