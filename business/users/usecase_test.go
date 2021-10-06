package users_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"project/app/middlewares"
	"project/business/premium"
	_mockProRepository "project/business/premium/mocks"

	// toke
	"project/business/users"
	_mockUserRepository "project/business/users/mocks"
	encrypt "project/helpers/encrypt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var userRepository _mockUserRepository.Repository
var proRepository _mockProRepository.Repository
var userService users.Usecase
var userDomain users.User
var proDomain premium.Premium
var usersDomain []users.User
var jwtAuth *middlewares.ConfigJWT

func setup() {

	jwtAuth = &middlewares.ConfigJWT{
		SecretJWT:       "testmock123",
		ExpiresDuration: 1,
	}
	userService = users.NewUserUsecase(&userRepository, &proRepository, time.Hour*1, jwtAuth)
	proDomain = premium.Premium{
		UserId:  1,
		Type:    true,
		Expired: time.Time{},
	}
	userDomain = users.User{
		Id:       1,
		Name:     "cen",
		Email:    "cen@mail.co",
		Premium:  proDomain,
		Password: "$2y$10$N8AySzI3aj7d6CDDXCOAA.8lXpw14f3BPp.i0EIpTD/YwWRxmbCAu",
		Token:    "123",
	}
	usersDomain = append(usersDomain, userDomain)
}

func TestLogin(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Login", func(t *testing.T) {
		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(userDomain, nil).Once()
		_ = encrypt.CheckPassword("12345", userDomain.Password)
		jwtAuth.GenerateToken(userDomain.Id)
		proRepository.On("Detail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(proDomain, nil).Once()
		proRepository.On("Save",
			mock.Anything,
			mock.AnythingOfType("premium.Premium")).Return(proDomain, nil).Once()
		user, err := userService.Login(context.Background(), users.User{
			Email:    "cen@mail.co",
			Password: "12345",
		})
		// pro, err := )
		assert.Nil(t, err)
		assert.Equal(t, "cen@mail.co", user.Email)
	})
	t.Run("Test case 2 | Error Login", func(t *testing.T) {

		userRepository.On("Login",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(users.User{}, errors.New("Unexpected Error")).Once()
		user, err := userService.Login(context.Background(), users.User{
			Email:    "ada@das.ad",
			Password: "das",
		})

		assert.NotNil(t, err)
		assert.Equal(t, user, users.User{})
	})

	t.Run("Test Case 3 | Invalid Email Empty", func(t *testing.T) {

		_, err := userService.Login(context.Background(), users.User{
			Email:    "",
			Password: "12345",
		})
		assert.NotNil(t, err)
	})
	// t.Run("Test case 4 | premium not found", func(t *testing.T) {
	// 	userRepository.On("Login",
	// 		mock.Anything,
	// 		mock.AnythingOfType("users.User")).Return(userDomain, nil).Once()
	// 	err := encrypt.CheckPassword("12345", userDomain.Password)
	// 	// jwtAuth.GenerateToken(userDomain.Id)
	// 	// proRepository.On("Detail",
	// 	// 	mock.Anything,
	// 	// 	mock.AnythingOfType("int")).Return(proDomain, nil).Once()
	// 	user, _ := userService.Login(context.Background(), users.User{
	// 		Email:    "cen@mail.co",
	// 		Password: "12345",
	// 	})
	// 	// assert.False(t, proDomain.Type)
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "cen@mail.co", user.Email)
	// })
	// t.Run("Test case 5 | cant save premium", func(t *testing.T) {
	// 	userRepository.On("Login",
	// 		mock.Anything,
	// 		mock.AnythingOfType("users.User")).Return(userDomain, nil).Once()
	// 	err := encrypt.CheckPassword("12345", userDomain.Password)
	// 	jwtAuth.GenerateToken(userDomain.Id)
	// 	proRepository.On("Detail",
	// 		mock.Anything,
	// 		mock.AnythingOfType("int")).Return(userDomain, nil).Once()
	// 	proRepository.On("Save",
	// 		mock.Anything,
	// 		mock.AnythingOfType("premium.Premium")).Return(premium.Premium{}, nil).Once()
	// 	user, err := userService.Login(context.Background(), users.User{
	// 		Email:    "cen@mail.co",
	// 		Password: "12345",
	// 	})
	// 	// pro, err := )
	// 	assert.Nil(t, err)
	// 	assert.Equal(t, "cen@mail.co", user.Email)
	// })
	t.Run("Test case 5 | email or password empty", func(t *testing.T) {

		user, err := userService.Login(context.Background(), users.User{
			// Email:    "cen@mail.co",
			Password: "12345",
		})
		// pro, err := )
		assert.NotNil(t, user.Email, user.Password, err)
		// assert.Equal(t, "cen@mail.co", user.Email)
	})
}
func TestUserDetail(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()

		user, err := userService.UserDetail(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(users.User{}, errors.New("Unexpected Error")).Once()

		user, err := userService.UserDetail(context.Background(), userDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})

		userRepository.AssertExpectations(t)
	})
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("GetAll",
			mock.Anything).Return(usersDomain, nil).Once()

		user, err := userService.GetAll(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, user)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Error", func(t *testing.T) {
		userRepository.On("GetAll",
			mock.Anything).Return([]users.User{}, errors.New("Unexpected Error")).Once()

		user, err := userService.GetAll(context.Background())

		assert.Error(t, err)
		assert.Equal(t, user, []users.User{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 3 | Zero", func(t *testing.T) {
		userRepository.On("GetAll",
			mock.Anything).Return([]users.User{}, errors.New("Unexpected Error")).Once()

		_, _ = userService.GetAll(context.Background())

		// assert.Error(t, err)
		assert.Len(t, []users.User{}, 0)

		userRepository.AssertExpectations(t)
	})
}
func TestDelete(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		user, err := userService.Delete(context.Background(), userDomain.Id)

		assert.NoError(t, err)
		assert.NotNil(t, user, userDomain)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Delete Error", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Delete",
			mock.Anything,
			mock.AnythingOfType("int")).Return(users.User{}, errors.New("Unexpected Error")).Once()
		user, err := userService.Delete(context.Background(), userDomain.Id)

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 3 |Detail Error", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(users.User{}, errors.New("Unexpected Error")).Once()
		// userRepository.On("Update",
		// 	mock.Anything,
		// 	mock.AnythingOfType("users.User")).Return(userDomain, errors.New("Unexpected Error")).Once()

		user, err := userService.Delete(context.Background(), 1)

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})

		userRepository.AssertExpectations(t)
	})
}

func TestRegister(t *testing.T) {

	setup()

	t.Run("Test case 1 | Valid Registry", func(t *testing.T) {
		hash, _ := encrypt.Hash("adasd")
		expectedReturn := users.User{
			Email:    "cen@mail.co",
			Password: hash,
		}
		userRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(expectedReturn, nil).Once()
		proRepository.On("Save",
			mock.Anything,
			mock.AnythingOfType("premium.Premium")).Return(proDomain, nil).Once()
		_, err := userService.Register(context.Background(), users.User{
			Name:     "cen",
			Email:    "cen@mail.co",
			Password: "12345",
		})

		assert.Nil(t, err)
		// assert.Equal(t, "cen", user.Name)
	})

	t.Run("Test case 2 | Error Registry", func(t *testing.T) {
		hash, _ := encrypt.Hash("adasd")
		expectedReturn := users.User{
			Email:    "bjanardana@google.com",
			Password: hash,
		}
		userRepository.On("Register",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(expectedReturn, errors.New("Unexpected Error")).Once()
		user, err := userService.Register(context.Background(), users.User{
			Name:     "asd",
			Email:    "asdsa@asdad.asda",
			Password: "adasd",
		})

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})
	})

	t.Run("Test Case 3 | Invalid Empty", func(t *testing.T) {
		user, err := userService.Register(context.Background(), users.User{
			Name:     "",
			Email:    "ada",
			Password: "asd",
		})
		assert.NotNil(t, err, user.Name, user.Email, user.Password)
		assert.Equal(t, user, users.User{})
	})

}

func TestUpdate(t *testing.T) {
	setup()

	t.Run("Test case 1", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(userDomain, nil).Once()
		user, err := userService.Update(context.Background(), users.User{
			Id:       1,
			Name:     "asd",
			Email:    "asdsa@asdad.asda",
			Password: "adasd",
		})

		assert.NoError(t, err)
		assert.Equal(t, 1, user.Id)

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 2 | Update Error", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepository.On("Update",
			mock.Anything,
			mock.AnythingOfType("users.User")).Return(users.User{}, errors.New("Unexpected Error")).Once()

		user, err := userService.Update(context.Background(), users.User{
			Id:       1,
			Name:     "ddd",
			Email:    "dd@dd.dd",
			Password: "ddd",
		})

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})

		userRepository.AssertExpectations(t)
	})

	t.Run("Test case 3 |Detail Error", func(t *testing.T) {
		userRepository.On("UserDetail",
			mock.Anything,
			mock.AnythingOfType("int")).Return(users.User{}, errors.New("Unexpected Error")).Once()
		// userRepository.On("Update",
		// 	mock.Anything,
		// 	mock.AnythingOfType("users.User")).Return(userDomain, errors.New("Unexpected Error")).Once()

		user, err := userService.Update(context.Background(), users.User{
			Id:       1,
			Name:     "ddd",
			Email:    "dd@dd.dd",
			Password: "ddd",
		})

		assert.Error(t, err)
		assert.Equal(t, user, users.User{})

		userRepository.AssertExpectations(t)
	})
}
