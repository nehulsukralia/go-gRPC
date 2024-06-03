package main

import (
	"context"
	"log"
	"net"

	"github.com/nehulsukralia/go-gRPC/service"
	pb "github.com/nehulsukralia/go-gRPC/user"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserServer struct {
	userService service.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserServer() *UserServer {
	return &UserServer{}
}

func (server *UserServer) Run() error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	return s.Serve(lis)
}

func (s *UserServer) GetUsers(ctx context.Context, in *pb.GetUsersParam) (*pb.UserList, error) {
	users, err := s.userService.GetUsers(ctx, in);
	if err != nil {
		return nil, err
	}

	userList := &pb.UserList{
		Users: users,
	}

	return userList, nil
}

func (s *UserServer) GetUser(ctx context.Context, in *pb.GetUserParam) (*pb.User, error) {
	user, err := s.userService.GetUser(ctx, in);
	if err != nil {
		return nil, err
	}

	return user, err
}

func main() {
	userServer := NewUserServer()
	if err := userServer.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
