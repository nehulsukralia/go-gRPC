package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/nehulsukralia/go-gRPC/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	params1 := &pb.GetUsersParam{
		UserIdList: []int32{2, 3},
	}
	r1, err := c.GetUsers(ctx, params1)
	if err != nil {
		log.Fatalf("could not retrieve users: %v", err)
	}
	log.Print("\nUser List: \n")
	fmt.Printf("r.GetUsers(): %v\n", r1.GetUsers())

	param2 := &pb.GetUserParam{
		UserId: 2,
	}
	r2, err := c.GetUser(ctx, param2)
	if err != nil {
		log.Fatalf("could not retrieve user: %v", err)
	}

	log.Print("\nUser: \n")
	fmt.Printf("r.GetUser(): %v\n", r2)
}
