package service

import (
	"context"
	"reflect"
	"testing"

	pb "github.com/nehulsukralia/go-gRPC/user"
)

func TestGetUsers(t *testing.T) {
	ctx := context.Background()
	s := UserService{}

	// Test case 1: Get users by ID list
	t.Run("GetUsersByIDList", func(t *testing.T) {
		input := &pb.GetUsersParam{
			UserIdList: []int32{1, 2},
		}
		want := []*pb.User{
			{
				Id:        1,
				FirstName: "Manoj",
				City:      "Meerut",
				Phone:     "9897234572",
				Height:    5.6,
				Married:   true,
			},
			{
				Id:        2,
				FirstName: "Mohit",
				City:      "Lucknow",
				Phone:     "9892579545",
				Height:    5.8,
				Married:   true,
			},
		}

		got, err := s.GetUsers(ctx, input)
		if err != nil {
			t.Errorf("GetUsers returned error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetUsers = %v, want %v", got, want)
		}
	})

	// Test case 2: Get all users
	t.Run("GetAllUsers", func(t *testing.T) {
		input := &pb.GetUsersParam{
			UserIdList: []int32{},
		}
		want := UserData

		got, err := s.GetUsers(ctx, input)
		if err != nil {
			t.Errorf("GetUsers returned error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetUsers = %v, want %v", got, want)
		}
	})

	// Test case 3: Error Case, Get users with non-existent IDs
	t.Run("GetUsersWithNonExistentIDs", func(t *testing.T) {
		input := &pb.GetUsersParam{
			UserIdList: []int32{59, 34},
		}
		expError := "users with the input IDs not found"
		_, err := s.GetUsers(ctx, input)
		if err == nil || err.Error() != expError {
			t.Errorf("GetUsers returned unexpected error: got %v, want %v", err, expError)
		}
	})
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	s := UserService{}

	// Test case 1: Get user by ID
	t.Run("GetUserByID", func(t *testing.T) {
		input := &pb.GetUserParam{
			UserId: 4,
		}

		want := &pb.User{
			Id:        4,
			FirstName: "Rajan",
			City:      "New Delhi",
			Phone:     "9883945622",
			Height:    5.6,
			Married:   true,
		}

		got, err := s.GetUser(ctx, input)
		if err != nil {
			t.Errorf("GetUser returned error: %v", err)
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("GetUser = %v, want %v", got, want)
		}
	})

	// Test case 3: Error Case, Get users with non-existent IDs
	t.Run("GetUserWithNonExistentID", func(t *testing.T) {
		input := &pb.GetUserParam{
			UserId: 8,
		}
		expError := "user with the input ID not found"
		_, err := s.GetUser(ctx, input)
		if err == nil || err.Error() != expError {
			t.Errorf("GetUsers returned unexpected error: got %v, want %v", err, expError)
		}
	})
}
