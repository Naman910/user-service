package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user-service/internal/api"
	"user-service/internal/model"
	"user-service/internal/service"
)

// Define MockUserRepository struct and methods
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) GetUserByID(ctx context.Context, userID uint64) (*model.User, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) GetUserListByID(ctx context.Context, userIDs []uint64) ([]*model.User, error) {
	args := m.Called(ctx, userIDs)
	return args.Get(0).([]*model.User), args.Error(1)
}

// SearchUsers implements repository.UserRepository.
func (m *MockUserRepository) SearchUsers(ctx context.Context, criteria string) ([]*model.User, error) {
	args := m.Called(ctx, criteria)
	return args.Get(0).([]*model.User), args.Error(1)
}

func TestGetUserByID(t *testing.T) {
	// Initialize mock repository
	userRepository := &MockUserRepository{}

	// Create test user
	testUser := &model.User{ID: 1, FirstName: "John", City: "New York"}

	// Mock repository method to return test user
	userRepository.On("GetUserByID", 1).Return(testUser, nil)

	// Create test service
	userService := service.NewUserService(userRepository)

	// Create UserServiceServer instance
	userServiceServer := api.NewUserServiceServer(userService)

	// Create GetUserByID request
	req := &api.GetUserByIDRequest{UserId: 1}

	// Call GetUserByID API method
	resp, err := userServiceServer.GetUserByID(context.Background(), req)

	// Assert response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, testUser.ID, resp.User.Id)
	assert.Equal(t, testUser.FirstName, resp.User.Fname)
	assert.Equal(t, testUser.City, resp.User.City)

	// Verify mock repository method call
	userRepository.AssertCalled(t, "GetUserByID", 1)
}
