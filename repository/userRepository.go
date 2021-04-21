package repository

type UserRepository interface {
	CreateUser() bool
	GetAllUser() []DataUser
	UpdateUser() bool
	GetUserById(id string) DataUser
}

type DataUser struct {
	ID          string
	FirstName   string
	MiddleName  string
	LastName    string
	Username    string
	Role        string
	InitialName string
	Email       string
	Password    string
	UpdatedAt   string
}

var data []DataUser

//CreateUser is to save data to database
func (d DataUser) CreateUser() bool {
	data = append(data, d)
	return true
}

// FindByID return specific user detail by id
func GetAllUser() []DataUser {
	return data
}

// FindByID return specific user detail by id
func GetUserById(id string) DataUser {
	return DataUser{
		ID:          "123",
		FirstName:   "farhan",
		MiddleName:  "stona",
		LastName:    "ramadhana",
		Username:    "frans",
		Role:        "admin",
		InitialName: "FSR",
		Email:       "frans@kata.ai",
		Password:    "pass",
		UpdatedAt:   "2020-02-02",
	}
}

func GetUserByName(firstName, lastName string) []DataUser {
	var user []DataUser
	for _, v := range data {
		if firstName == v.FirstName && lastName == v.LastName {
			user = append(user, v)
		}
	}

	return user
}
