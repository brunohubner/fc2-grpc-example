package services

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/brunohubner/fc2-grpc-example/pb"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	fmt.Println(req.Name)

	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {
	stream.Send(&pb.UserResultStrem{
		Status: "Init",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStrem{
		Status: "Inserting",
		User:   &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStrem{
		Status: "User has been inserted",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStrem{
		Status: "Completed",
		User: &pb.User{
			Id:    "123",
			Name:  req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)

	return nil
}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Users{
				User: users,
			})
		}
		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id:    req.GetId(),
			Name:  req.Name,
			Email: req.Email,
		})

		fmt.Println("Adding", req.GetName())
	}
}

func (*UserService) AddUserStreamBoth(stream pb.UserService_AddUserStreamBothServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error receiving stream from the client: %v\n", err)
		}

		if err = stream.Send(&pb.UserResultStrem{
			Status: "Added",
			User:   req,
		}); err != nil {
			log.Fatalf("Error receiving stream from the client: %v\n", err)
		}
	}
}
