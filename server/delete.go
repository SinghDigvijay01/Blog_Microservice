package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {

	log.Println("deleteBlog was invoked with %v", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse Id ")
	}
	res, err := Collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while deleting the blog with %v", oid)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Notfound the blog with %v", oid)
	}
	return &emptypb.Empty{}, nil
}
