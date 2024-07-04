package models

type Login struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	AccountType string `json:"role" validate:"required,oneof=FoodPlace Customer Adm"`
}

type AuthorizationData struct {
	Name        string `json:"name" validate:"required,min=1"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	AccountType string `json:"role" validate:"required"`
}
