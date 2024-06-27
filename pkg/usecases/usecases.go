package usecases

import (
	"github.com/emaforlin/coupons-app/pkg/entities"
	"github.com/emaforlin/coupons-app/pkg/models"
)

type AccountUsecase interface {
	// Accounts in general related
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	DeleteAccount(in *models.DeleteAccountData) error

	// Accounts of type Person related
	AddPersonAccount(in *models.AddPersonAccountData) error

	// Accounts of type FoodPlace
	AddFoodPlaceAccount(in *models.AddFoodPlaceData) error
}
