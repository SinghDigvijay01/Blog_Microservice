# Blog Microservice

This repository contains a gRPC-based **Blog Microservice** written in **Protocol Buffers (proto3)**. The service allows users to create, read, update, delete, and list blog posts. The service is designed for scalability and can be integrated into larger systems for managing blog content.

## Features

- **Create Blog**: Add new blog posts with an author, title, and content.
- **Read Blog**: Fetch a specific blog post by its unique ID.
- **Update Blog**: Modify existing blog content.
- **Delete Blog**: Remove a blog post by its ID.
- **List Blogs**: Stream a list of all available blog posts.

## API Definition

The service follows the gRPC model, and the following RPCs are available:

1. `CreateBlog(Blog) returns (BlogId)` - Create a new blog post and return its unique ID.
2. `ReadBlog(BlogId) returns (Blog)` - Retrieve a blog post using its ID.
3. `UpdateBlog(Blog) returns (google.protobuf.Empty)` - Update the details of an existing blog post.
4. `DeleteBlog(BlogId) returns (google.protobuf.Empty)` - Delete a blog post using its ID.
5. `ListBlogs(google.protobuf.Empty) returns (stream Blog)` - Stream all blogs in the system.

## Blog Model

The `Blog` message contains the following fields:

- **id**: Unique identifier for the blog post.
- **author_id**: ID of the blog's author.
- **title**: Title of the blog post.
- **content**: The blog post content.

## Technology Stack

- **gRPC**: For defining the service and RPCs.
- **Protocol Buffers**: Used for data serialization.
- **Go**: The `go_package` option is set to `"microservice/blog/protoc"` for Go code generation.

## Getting Started

To generate the Go files, use the following command:

```bash
protoc --go_out=. --go-grpc_out=. blog.proto

