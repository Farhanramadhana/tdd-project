package user

import (
	"bootcamp/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func (repository *userRepositoryMock) CreateUser(d entity.DataUserEntity) bool {
	arguments := repository.Called(d)
	return arguments.Bool(0)
}

func (repository *userRepositoryMock) GetAllUser() ([]entity.DataUserEntity, error) {
	arguments := repository.Called()
	user := arguments.Get(0).([]entity.DataUserEntity)
	return user, arguments.Error(1)
}

func (repository *userRepositoryMock) GetUserByID(id string) (entity.DataUserEntity, error) {
	arguments := repository.Called(id)
	if arguments.Get(0) == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}
	arg1, arg2 := arguments.Get(0).(entity.DataUserEntity), arguments.Error(1)
	return arg1, arg2
}

func (repository *userRepositoryMock) GetUserByUserName(username string) (entity.DataUserEntity, error) {
	arguments := repository.Called(username)
	if arguments.Get(0) == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}
	arg1, arg2 := arguments.Get(0).(entity.DataUserEntity), arguments.Error(1)
	return arg1, arg2
}

func (repository *userRepositoryMock) GetUserByEmail(email string) (entity.DataUserEntity, error) {
	arguments := repository.Called(email)
	if arguments.Get(0) == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}
	arg1, arg2 := arguments.Get(0).(entity.DataUserEntity), arguments.Error(1)
	return arg1, arg2
}

func (repository *userRepositoryMock) DeleteUserByID(id string) error {
	arguments := repository.Called(id)
	if arguments.Get(0) != nil {
		return errors.New("data doesn't exist")
	}
	return nil
}

func (repository *userRepositoryMock) UpdateUserByID(id string, data entity.DataUserEntity) (entity.DataUserEntity, error) {
	arguments := repository.Called(id, data)
	if arguments.Get(0) == nil {
		return entity.DataUserEntity{}, errors.New("data doesn't exist")
	}
	arg1, arg2 := arguments.Get(0).(entity.DataUserEntity), arguments.Error(1)
	return arg1, arg2
}

func TestCreateUser_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)

	input := entity.RegistrationUserEntity{
		FullName: "farhan ramadhana",
		Role:     "admin",
		Email:    "farhan@kata.ai",
		Password: "1234567890",
	}

	repoMock.On("CreateUser", mock.AnythingOfType("entity.DataUserEntity")).Return(true)
	u := UserService{repoMock}
	status, err := u.CreateUser(input)
	assert.Equal(t, true, status)
	assert.Nil(t, err)
}

func TestCreateUser_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}
	// missing fullname
	data1 := entity.RegistrationUserEntity{
		Role:     "admin",
		Email:    "farhan@kata.ai",
		Password: "1234567890",
	}

	// missing role
	data2 := entity.RegistrationUserEntity{
		FullName: "farhan ramadhana",
		Email:    "farhan@kata.ai",
		Password: "1234567890",
	}

	// missing email
	data3 := entity.RegistrationUserEntity{
		FullName: "farhan ramadhana",
		Role:     "admin",
		Password: "1234567890",
	}

	// missing password
	data4 := entity.RegistrationUserEntity{
		FullName: "farhan ramadhana",
		Role:     "admin",
		Email:    "farhan@kata.ai",
	}

	t.Run("one", func(t *testing.T) {
		repoMock.On("CreateUser", data1).Return(false)
		_, err := u.CreateUser(data1)
		assert.Equal(t, "FullName is required", err.Error())
	})
	t.Run("two", func(t *testing.T) {
		repoMock.On("CreateUser", data2).Return()
		_, err := u.CreateUser(data2)
		assert.Equal(t, "Role is required", err.Error())
	})
	t.Run("three", func(t *testing.T) {
		repoMock.On("CreateUser", data3).Return()
		_, err := u.CreateUser(data3)
		assert.Equal(t, "Email is required", err.Error())
	})
	t.Run("four", func(t *testing.T) {
		repoMock.On("CreateUser", data4).Return()
		_, err := u.CreateUser(data4)
		assert.Equal(t, "Password is required", err.Error())
	})
}

func TestGetAllUser_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{
		ID:          "zy95jC0t9rQU",
		FirstName:   "farhan",
		MiddleName:  "",
		LastName:    "ramadhana",
		Username:    "farhan.ramadhana",
		Role:        "user",
		InitialName: "FR",
		Email:       "farhan@kata.ai",
		Password:    "$2a$04$Vq2nAjigHrtOuy3Cw/i/IOiyjhTaGSeEduxjJjYjZ2ISosFE/SR4K",
		UpdatedAt:   "2021-04-28T06:08:28+07:00",
	}

	repoMock.On("GetAllUser").Return([]entity.DataUserEntity{data}, nil)

	userData, err := u.GetAllUser()

	assert.Nil(t, err)
	assert.Equal(t, "zy95jC0t9rQU", userData[0].ID)
}

func TestGetAllUser_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{}

	repoMock.On("GetAllUser").Return([]entity.DataUserEntity{data}, errors.New("data doesn't exist"))

	userData, err := u.GetAllUser()

	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error())
	assert.Equal(t, "", userData[0].ID)
}

func TestGetUserByID_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{
		ID:          "zy95jC0t9rQU",
		FirstName:   "farhan",
		MiddleName:  "",
		LastName:    "ramadhana",
		Username:    "farhan.ramadhana",
		Role:        "user",
		InitialName: "FR",
		Email:       "farhan@kata.ai",
		Password:    "$2a$04$Vq2nAjigHrtOuy3Cw/i/IOiyjhTaGSeEduxjJjYjZ2ISosFE/SR4K",
		UpdatedAt:   "2021-04-28T06:08:28+07:00",
	}

	repoMock.On("GetUserByID", "1").Return(data, nil)

	userData, err := u.GetUserByID("1")
	assert.Nil(t, err)
	assert.Equal(t, "zy95jC0t9rQU", userData.ID)
}

func TestGetUserByID_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{}

	repoMock.On("GetUserByID","1").Return(data, errors.New("data doesn't exist"))

	_, err := u.GetUserByID("1")
	assert.NotNil(t, err)
	assert.Equal(t, "data doesn't exist", err.Error())
}

func TestGetUserByUsername_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{
		ID:          "zy95jC0t9rQU",
		FirstName:   "farhan",
		MiddleName:  "",
		LastName:    "ramadhana",
		Username:    "farhan.ramadhana",
		Role:        "user",
		InitialName: "FR",
		Email:       "farhan@kata.ai",
		Password:    "$2a$04$Vq2nAjigHrtOuy3Cw/i/IOiyjhTaGSeEduxjJjYjZ2ISosFE/SR4K",
		UpdatedAt:   "2021-04-28T06:08:28+07:00",
	}

	repoMock.On("GetUserByUserName", "farhan.ramadhana").Return(data, nil)

	userData, err := u.GetUserByUserName("farhan.ramadhana")
	assert.Nil(t, err)
	assert.Equal(t, "farhan.ramadhana", userData.Username)
}

func TestGetUserByUsername_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{}

	repoMock.On("GetUserByUserName", "farhan.ramadhana").Return(data, errors.New("data doesn't exist"))

	userData, err := u.GetUserByUserName("farhan.ramadhana")
	assert.NotNil(t, err)
	assert.Equal(t, "", userData.Username)
}

func TestGetUserByEmail_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{
		ID:          "zy95jC0t9rQU",
		FirstName:   "farhan",
		MiddleName:  "",
		LastName:    "ramadhana",
		Username:    "farhan.ramadhana",
		Role:        "user",
		InitialName: "FR",
		Email:       "farhan@kata.ai",
		Password:    "$2a$04$Vq2nAjigHrtOuy3Cw/i/IOiyjhTaGSeEduxjJjYjZ2ISosFE/SR4K",
		UpdatedAt:   "2021-04-28T06:08:28+07:00",
	}

	repoMock.On("GetUserByEmail", "farhan@kata.ai").Return(data, nil)

	userData, err := u.GetUserByEmail("farhan@kata.ai")
	assert.Nil(t, err)
	assert.Equal(t, "farhan@kata.ai", userData.Email)
}

func TestGetUserByEmail_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	data := entity.DataUserEntity{}

	repoMock.On("GetUserByEmail", "farhan@kata.ai").Return(data, errors.New("data doesn't exist"))

	userData, err := u.GetUserByEmail("farhan@kata.ai")
	assert.NotNil(t, err)
	assert.Equal(t, "", userData.Email)
}

func TestDeleteUserByID_Success(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	repoMock.On("DeleteUserByID", "zy95jC0t9rQU").Return(nil)

	err := u.DeleteUserByID("zy95jC0t9rQU")
	assert.Nil(t, err)
}

func TestDeleteUserByID_Failed(t *testing.T) {
	repoMock := new(userRepositoryMock)
	u := UserService{repoMock}

	repoMock.On("DeleteUserByID", "zy95jC0t9rQU").Return(errors.New("data does'nt exist"))

	err := u.DeleteUserByID("zy95jC0t9rQU")
	assert.NotNil(t, err)
}

func TestSplitFullName(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		result := SplitFullName("farhan")
		assert.Equal(t, "farhan", result[0])
	})
	t.Run("two", func(t *testing.T) {
		result := SplitFullName("farhan ramadhana")
		assert.Equal(t, "farhan", result[0])
		assert.Equal(t, "", result[1])
		assert.Equal(t, "ramadhana", result[2])
	})
	t.Run("three", func(t *testing.T) {
		result := SplitFullName("farhan muhammad ramadhana")
		assert.Equal(t, "farhan", result[0])
		assert.Equal(t, "muhammad", result[1])
		assert.Equal(t, "ramadhana", result[2])
	})
	t.Run("four", func(t *testing.T) {
		result := SplitFullName("stona stoni stino stani")
		assert.Equal(t, "stona", result[0])
		assert.Equal(t, "stoni stino", result[1])
		assert.Equal(t, "stani", result[2])
	})
}

func TestCreateInitialName(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		result := CreateInitialName("farhan")
		assert.Equal(t, "F", result)
	})
	t.Run("two", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana")
		assert.Equal(t, "FR", result)
	})
	t.Run("three", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana ganteng")
		assert.Equal(t, "FRG", result)
	})
	t.Run("four", func(t *testing.T) {
		result := CreateInitialName("farhan ramadhana ganteng banget")
		assert.Equal(t, "FRGB", result)
	})
}

// func TestGenerateUserName(t *testing.T) {
// 	t.Run("one", func(t *testing.T) {
// 		name := []string{"stona", "", ""}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona", result)
// 	})
// 	t.Run("two", func(t *testing.T) {
// 		name := []string{"stona", "", "stoni"}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona.stoni", result)
// 	})
// 	t.Run("three", func(t *testing.T) {
// 		name := []string{"stona", "stoni", "stino"}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona.stino1", result)
// 	})
// 	t.Run("four", func(t *testing.T) {
// 		name := []string{"stona", "", "stoni"}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona.stoni2", result)
// 	})
// 	t.Run("five", func(t *testing.T) {
// 		name := []string{"stona", "", "stoni"}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona.stoni3", result)
// 	})
// 	t.Run("six", func(t *testing.T) {
// 		name := []string{"stona", "", "stoni"}
// 		result := GenerateUserName(name)
// 		assert.Equal(t, "stona.stoni4", result)
// 	})
// }
