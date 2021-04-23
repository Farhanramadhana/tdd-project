package entity

// RegistrationUserEntity is entity for registration
type RegistrationUserEntity struct {
	FullName string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=10"`
	Role     string `validate:"required"`
}

// DataUserEntity is entity for database column
type DataUserEntity struct {
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
