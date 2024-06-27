package usecases

import (
	"fmt"

	"github.com/emaforlin/coupons-app/pkg/entities"
	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/emaforlin/coupons-app/pkg/repositories"
	"github.com/rs/zerolog/log"
)

type accountUsecaseImpl struct {
	repository repositories.AccountsRepository
}

// AddFoodPlaceAccount implements AccountUsecase.
func (u *accountUsecaseImpl) AddFoodPlaceAccount(in *models.AddFoodPlaceData) error {
	_, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		return err
	}

	dto := entities.InsertFoodPlaceDto{
		BusinessName: in.BusinessName,
		Location:     in.Location,
		User: entities.InsertUserDto{
			Role:        "FoodPlace",
			Username:    in.Username,
			Email:       in.Email,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
		},
	}
	err = u.repository.InsertFoodPlace(&dto)
	if err != nil {
		log.Err(err)
		return fmt.Errorf("error creating account")
	}

	return nil
}

func (u *accountUsecaseImpl) DeleteAccount(in *models.DeleteAccountData) error {
	return u.repository.DeleteUser(&entities.GetUserDto{
		ID: in.Id,
	})
}

// AddPersonAccountDetails implements AccountsUsecase.
func (u *accountUsecaseImpl) AddPersonAccount(in *models.AddPersonAccountData) error {
	_, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})
	if err != nil {
		return err
	}

	dto := entities.InsertPersonDto{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		User: entities.InsertUserDto{
			Role:        "Customer",
			Username:    in.Username,
			Email:       in.Email,
			PhoneNumber: in.PhoneNumber,
			Password:    in.Password,
		},
	}

	err = u.repository.InsertPerson(&dto)
	if err != nil {
		log.Err(err)
		return fmt.Errorf("error creating account")
	}
	return nil
}

// GetUserDetails implements UserUsecase.
func (u *accountUsecaseImpl) GetAccountDetails(in *models.GetAccountData) (*entities.User, error) {
	found, err := u.repository.SelectUser(&entities.GetUserDto{
		Username:    in.Username,
		Email:       in.Email,
		PhoneNumber: in.PhoneNumber,
	})

	if err != nil {
		log.Err(err)
		return nil, fmt.Errorf("error user not found")
	}

	return found, nil
}

func NewAccountUsecaseImpl(repo repositories.AccountsRepository) AccountUsecase {
	return &accountUsecaseImpl{repository: repo}
}
