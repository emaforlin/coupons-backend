package usecases

import (
	"github.com/emaforlin/coupons-app/pkg/accounts/entities"
	"github.com/emaforlin/coupons-app/pkg/accounts/models"
)

type AccountUsecase interface {
	// Accounts in general related
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	DeleteAccount(in *models.DeleteAccountData) error

	// Accounts of type Person related
	AddPersonAccount(in *models.AddPersonAccountData) error

	// Accounts of type FoodPlace
}
