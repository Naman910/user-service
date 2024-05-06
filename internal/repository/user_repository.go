package repository

import (
	"context"
	"errors"
	"sync"

	"user-service/internal/model"
)

// UserRepository defines methods for interacting with user data
type UserRepository interface {
	GetUserByID(ctx context.Context, userID uint64) (*model.User, error)
	GetUserListByID(ctx context.Context, userIDs []uint64) ([]*model.User, error)
	SearchUsers(ctx context.Context, criteria string) ([]*model.User, error)
}

// userRepositoryImpl implements UserRepository
type userRepositoryImpl struct {
	users []*model.User
	mu    sync.RWMutex
}

// NewUserRepository creates a new instance of UserRepository
func NewUserRepository() UserRepository {
	return &userRepositoryImpl{
		users: []*model.User{
			{ID: 1, FirstName: "John", City: "New York", Phone: 1234567890, Height: 6.0, Married: false},
			{ID: 2, FirstName: "Alice", City: "Los Angeles", Phone: 9876543210, Height: 5.5, Married: true},
		},
	}
}

// GetUserByID retrieves user details by ID
func (r *userRepositoryImpl) GetUserByID(ctx context.Context, userID uint64) (*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, user := range r.users {
		if user.ID == userID {
			return user, nil
		}
	}

	return nil, errors.New("user not found")
}

// GetUserListByID retrieves a list of user details by IDs
func (r *userRepositoryImpl) GetUserListByID(ctx context.Context, userIDs []uint64) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var userList []*model.User
	for _, userID := range userIDs {
		for _, user := range r.users {
			if user.ID == userID {
				userList = append(userList, user)
				break
			}
		}
	}

	return userList, nil
}

// SearchUsers searches for users based on criteria
func (r *userRepositoryImpl) SearchUsers(ctx context.Context, criteria string) ([]*model.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var userList []*model.User
	for _, user := range r.users {
		// Perform search based on criteria (e.g., city, phone number, marital status, etc.)
		// For simplicity, let's assume we're searching by city only
		if user.City == criteria {
			userList = append(userList, user)
		}
	}

	return userList, nil
}
