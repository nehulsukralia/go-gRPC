package service

import (
	pb "github.com/nehulsukralia/go-gRPC/user"
)

var UserData = []*pb.User{
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
	{
		Id:        3,
		FirstName: "Ravi",
		City:      "Noida",
		Phone:     "9897238562",
		Height:    5.5,
		Married:   false,
	},
	{
		Id:        4,
		FirstName: "Rajan",
		City:      "New Delhi",
		Phone:     "9883945622",
		Height:    5.6,
		Married:   true,
	},
	{
		Id:        5,
		FirstName: "Deepak",
		City:      "Gurugram",
		Phone:     "9834234272",
		Height:    5.9,
		Married:   false,
	},
}
