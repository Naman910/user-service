package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"user-service/internal/api"
	"user-service/internal/api/user-service/proto"
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
	userRepository.On("GetUserByID", mock.Anything, uint64(1)).Return(testUser, nil)

	// Create test service
	userService := service.NewUserService(userRepository)

	// Create UserServiceServer instance
	userServiceServer := api.NewUserServiceServer(userService)

	// Create GetUserByID request
	req := &proto.GetUserByIDRequest{Id: 1}

	// Call GetUserByID API method
	resp, err := userServiceServer.GetUserByID(context.Background(), req)

	// Assert response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, testUser.ID, resp.User.Id)
	assert.Equal(t, testUser.FirstName, resp.User.Fname)
	assert.Equal(t, testUser.City, resp.User.City)

	// Verify mock repository method call
	userRepository.AssertCalled(t, "GetUserByID", mock.Anything, uint64(1))
}

func TestGetUserListByID(t *testing.T) {
	// Initialize mock repository
	userRepository := &MockUserRepository{}

	// Create test users
	testUsers := []*model.User{
		{ID: 1, FirstName: "John", City: "New York"},
		{ID: 2, FirstName: "Alice", City: "Los Angeles"},
	}

	// Mock repository method to return test users
	userRepository.On("GetUserListByID", mock.Anything, []uint64{1, 2}).Return(testUsers, nil)

	// Create test service
	userService := service.NewUserService(userRepository)

	// Create UserServiceServer instance
	userServiceServer := api.NewUserServiceServer(userService)

	// Create GetUserListByID request
	req := &proto.GetUserListByIDRequest{Ids: []uint64{1, 2}}

	// Call GetUserListByID API method
	resp, err := userServiceServer.GetUserListByID(context.Background(), req)

	// Assert response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, len(testUsers), len(resp.Users))
	for i, user := range resp.Users {
		assert.Equal(t, testUsers[i].ID, user.Id)
		assert.Equal(t, testUsers[i].FirstName, user.Fname)
		assert.Equal(t, testUsers[i].City, user.City)
	}

	// Verify mock repository method call
	userRepository.AssertCalled(t, "GetUserListByID", mock.Anything, []uint64{1, 2})
}

func TestSearchUsers(t *testing.T) {
	// Initialize mock repository
	userRepository := &MockUserRepository{}

	// Create test users
	testUsers := []*model.User{
		{ID: 1, FirstName: "John", City: "New York"},
		{ID: 2, FirstName: "Alice", City: "Los Angeles"},
	}

	// Mock repository method to return test users
	userRepository.On("SearchUsers", mock.Anything, "New York").Return(testUsers, nil)

	// Create test service
	userService := service.NewUserService(userRepository)

	// Create UserServiceServer instance
	userServiceServer := api.NewUserServiceServer(userService)

	// Create SearchUsers request
	req := &proto.SearchUsersRequest{Criteria: "New York"}

	// Call SearchUsers API method
	resp, err := userServiceServer.SearchUsers(context.Background(), req)

	// Assert response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, len(testUsers), len(resp.Users))
	for i, user := range resp.Users {
		assert.Equal(t, testUsers[i].ID, user.Id)
		assert.Equal(t, testUsers[i].FirstName, user.Fname)
		assert.Equal(t, testUsers[i].City, user.City)
	}

	// Verify mock repository method call
	userRepository.AssertCalled(t, "SearchUsers", mock.Anything, "New York")
}
