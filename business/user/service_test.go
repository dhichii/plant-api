package user_test

import (
	"errors"
	"os"
	"plant-api/business"
	userBusiness "plant-api/business/user"
	"plant-api/business/user/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	ID       uint = 1
	NAME          = "name"
	EMAIL         = "email"
	PASSWORD      = "password"
	ROLE          = "super"
)

var (
	mockRepo     mocks.Repository
	us           = userBusiness.NewService(&mockRepo)
	mockUser     userBusiness.User
	mockListUser []userBusiness.User
	mockNewUser  userBusiness.User
)

func setup() {
	mockUser = userBusiness.User{
		ID:       ID,
		Name:     NAME,
		Email:    EMAIL,
		Password: PASSWORD,
		Role:     ROLE,
	}

	mockListUser = append(mockListUser, mockUser)

	mockNewUser = userBusiness.User{
		Name:     NAME,
		Email:    EMAIL,
		Password: PASSWORD,
		Role:     ROLE,
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(nil).Once()
		err := us.Create(mockNewUser)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(errors.New("test error")).Once()
		err := us.Create(mockNewUser)
		assert.Error(t, err)
	})

	t.Run("failed email already used", func(t *testing.T) {
		mockRepo.On("Create", mock.Anything).Return(errors.New("Error 1062")).Once()
		err := us.Create(mockNewUser)
		assert.Error(t, err)
		assert.Equal(t, business.ErrConflict, err)
	})
}

func TestGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything).Return(mockListUser, nil).Once()
		users, err := us.GetAll()
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(users))
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("GetAll", mock.Anything).Return(nil, errors.New("error")).Once()
		users, err := us.GetAll()
		assert.Error(t, err)
		assert.Equal(t, 0, len(users))
	})
}

func TestGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(&mockUser, nil).Once()
		user, err := us.Get(int(ID))
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, ID, mockUser.ID)
		assert.Equal(t, NAME, mockUser.Name)
		assert.Equal(t, EMAIL, mockUser.Email)
		assert.Equal(t, PASSWORD, mockUser.Password)
		assert.Equal(t, ROLE, mockUser.Role)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(nil, errors.New("error")).Once()
		user, err := us.Get(int(ID))
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(nil, errors.New("record not found")).Once()
		user, err := us.Get(int(ID))
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, business.ErrNotFound, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(&mockUser, nil).Once()
		mockRepo.On("Update", mock.AnythingOfType("int"), mock.AnythingOfType("user.User")).Return(nil).Once()
		err := us.Update(int(ID), mockNewUser)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(nil, errors.New("error")).Once()
		err := us.Update(int(ID), mockNewUser)
		assert.Error(t, err)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("Get", mock.AnythingOfType("int")).Return(nil, errors.New("record not found")).Once()
		err := us.Update(int(ID), mockNewUser)
		assert.Error(t, err)
		assert.Equal(t, business.ErrNotFound, err)
	})
}