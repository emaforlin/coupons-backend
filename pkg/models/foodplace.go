package models

type AddFoodPlaceData struct {
	Username     string   `json:"username" validate:"required,min=4,max=20"`
	PhoneNumber  string   `json:"phone_number" validate:"required,e164"`
	Email        string   `json:"email" validate:"required,email"`
	Password     string   `json:"password" validate:"required,min=8,max=64"`
	BusinessName string   `json:"business_name" validate:"required"`
	Location     string   `json:"location" validate:"required"`
	Tags         []string `json:"tags" validate:"required"`
}

type VerifyFoodPlace struct {
	Email string `json:"email" validate:"required,email"`
}
