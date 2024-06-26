package models

// food place related models

// person related models
type AddPersonAccountData struct {
	Username    string `json:"username" validate:"required,min=4,max=20"`
	FirstName   string `json:"first_name" validate:"required,min=1,max=20"`
	LastName    string `json:"last_name" validate:"required,min=1,max=20"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"min=8,max=64"`
}

// account related modelspo
type GetAccountData struct {
	Username    string
	Email       string
	PhoneNumber string
}

type DeleteAccountData struct {
	Id uint32
}
