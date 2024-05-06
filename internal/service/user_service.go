package service

import (
	"context"

	"user-service/internal/model"
	"user-service/internal/repository"
)

// UserService implements methods for managing user details
type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService creates a new instance of UserService
func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

// GetUserByID fetches user details by ID
func (s *UserService) GetUserByID(ctx context.Context, userID uint64) (*model.User, error) {
	return s.userRepository.GetUserByID(ctx, userID)
}

// GetUserListByID fetches a list of user details by IDs
func (s *UserService) GetUserListByID(ctx context.Context, userIDs []uint64) ([]*model.User, error) {
	return s.userRepository.GetUserListByID(ctx, userIDs)
}

// SearchUsers searches for users based on criteria
func (s *UserService) SearchUsers(ctx context.Context, criteria string) ([]*model.User, error) {
	return s.userRepository.SearchUsers(ctx, criteria)
}
