package service

import (
	"context"
	"errors"

	pb "github.com/nehulsukralia/go-gRPC/user"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetUsers(ctx context.Context, in *pb.GetUsersParam) ([]*pb.User, error) {
	var users []*pb.User
	if len(in.UserIdList) > 0 {
		for _, v := range in.GetUserIdList() {
			for _, user := range UserData {
				if v == user.Id {
					users = append(users, user)
					break
				}
			}
		}
		if len(users) > 0 {
			return users, nil
		} else {
			return nil, errors.New("users with the input IDs not found")
		}
	} else {
		users = UserData
		return users, nil
	}
}

func (s *UserService) GetUser(ctx context.Context, in *pb.GetUserParam) (*pb.User, error) {
	for _, v := range UserData {
		if in.GetUserId() == v.GetId() {
			return v, nil
		}
	}
	return nil, errors.New("user with the input ID not found")
}
