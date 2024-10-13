package main

import (
	"context"
	"log"
	pb "microservice/blog/protoc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {

	log.Println("readBlog function  was invoked %v", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Cannot parse Id ")
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := Collection.FindOne(ctx, filter)

	if err = res.Decode(data); err != nil {
		return nil, status.Errorf(codes.NotFound, "Cannot find the Blog with provided Id ")
	}
	return documentToBlog(data), nil
}
