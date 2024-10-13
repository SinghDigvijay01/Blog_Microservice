package main

import (
	"context"
	"fmt"
	"log"
	pb "microservice/blog/protoc"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {

	log.Printf("CreateBlog was invoked with %v", in)

	data := BlogItem{
		AuthorID: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := Collection.InsertOne(ctx, data)

	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error : %v", err))
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot convert to OID"))
	}
	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
