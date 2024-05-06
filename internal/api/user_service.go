// internal/api/user_service.go

package api

import (
	"context"
	pb "user-service/internal/api/user-service/proto"
	"user-service/internal/model"
	"user-service/internal/service"

	"google.golang.org/grpc"
)

// RegisterUserServiceServer registers the UserServiceServer implementation with the gRPC server.
func RegisterUserServiceServer(s *grpc.Server, srv pb.UserServiceServer) {
	pb.RegisterUserServiceServer(s, srv)
}

// UserServiceServer implements the UserService gRPC service
type UserServiceServer struct {
	service *service.UserService
	pb.UnimplementedUserServiceServer
}

// NewUserServiceServer creates a new UserServiceServer instance
func NewUserServiceServer(service *service.UserService) *UserServiceServer {
	return &UserServiceServer{service: service}
}

// GetUserByID fetches user details by ID
func (s *UserServiceServer) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := s.service.GetUserByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByIDResponse{User: userToProto(user)}, nil
}

// GetUserListByID fetches a list of user details by IDs
func (s *UserServiceServer) GetUserListByID(ctx context.Context, req *pb.GetUserListByIDRequest) (*pb.GetUserListByIDResponse, error) {
	userList, err := s.service.GetUserListByID(ctx, req.Ids)
	if err != nil {
		return nil, err
	}

	return &pb.GetUserListByIDResponse{Users: usersToProto(userList)}, nil
}

// SearchUsers searches for users based on criteria
func (s *UserServiceServer) SearchUsers(ctx context.Context, req *pb.SearchUsersRequest) (*pb.SearchUsersResponse, error) {
	userList, err := s.service.SearchUsers(ctx, req.Criteria)
	if err != nil {
		return nil, err
	}

	return &pb.SearchUsersResponse{Users: usersToProto(userList)}, nil
}

// userToProto converts a model.User to its corresponding protobuf representation
func userToProto(user *model.User) *pb.User {
	return &pb.User{
		Id:      user.ID,
		Fname:   user.FirstName,
		City:    user.City,
		Phone:   user.Phone,
		Height:  user.Height,
		Married: user.Married,
	}
}

// usersToProto converts a slice of model.User to a slice of its corresponding protobuf representation
func usersToProto(users []*model.User) []*pb.User {
	var protoUsers []*pb.User
	for _, user := range users {
		protoUsers = append(protoUsers, userToProto(user))
	}
	return protoUsers
}
