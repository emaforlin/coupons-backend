package models

type GetAccountData struct {
	Username    string `validate:"required"`
	Email       string `validate:"required"`
	PhoneNumber string `validate:"required"`
}

type DeleteAccountData struct {
	Id uint32 `validate:"required"`
}
