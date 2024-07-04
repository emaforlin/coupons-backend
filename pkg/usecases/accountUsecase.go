package usecases

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/entities"
	"github.com/emaforlin/coupons-app/pkg/helpers"
	"github.com/emaforlin/coupons-app/pkg/models"
	"github.com/emaforlin/coupons-app/pkg/repositories"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

type AccountUsecase interface {
	Authenticate(in *models.Login) (*entities.User, error)
	Authorize(in *models.Login) (string, error)
	GetAccountDetails(in *models.GetAccountData) (*entities.User, error)
	DeleteAccount(in *models.DeleteAccountData) error

	AddPersonAccount(in *models.AddPersonAccountData) error

	VerifyFoodPlace(in *models.VerifyFoodPlace) error
	AddFoodPlaceAccount(in *models.AddFoodPlaceData) error
}

type accountUsecaseImpl struct {
	repository repositories.AccountsRepository
	jwtConfig  config.Jwt
}

func (u *accountUsecaseImpl) Authorize(in *models.Login) (string, error) {
	claims := helpers.CustomJWTClaims{
		Role: in.AccountType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(u.jwtConfig.TTL))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.jwtConfig.Secret)
}

func (u *accountUsecaseImpl) Authenticate(in *models.Login) (*entities.User, error) {
	account, err := u.repository.SelectAccount(&entities.GetUserDto{Email: in.Email})
	if err != nil {
		return nil, err
	}
	if in.AccountType != account.Role {
		return nil, fmt.Errorf("authentication error")
	}
	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(in.Password))

	if err != nil {
		return nil, err
	}

	return account, nil
}

func (u *accountUsecaseImpl) VerifyFoodPlace(in *models.VerifyFoodPlace) error {
	// find user id
	user, err := u.repository.SelectAccount(&entities.GetUserDto{Email: in.Email})
	if err != nil {
		return err
	}

	found, err := u.repository.SelectFoodPlace(&entities.GetFoodPlaceDto{UserId: user.ID})
	if err != nil {
		return err
	}
	err = u.repository.UpdateFoodPlace(found.UserId, &entities.InsertFoodPlaceDto{Verified: true})
	if err != nil {
		return err
	}

	return nil
}

func (u *accountUsecaseImpl) AddFoodPlaceAccount(in *models.AddFoodPlaceData) error {
	_, err := u.repository.SelectAccount(&entities.GetUserDto{
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
	return u.repository.DeleteAccount(&entities.GetUserDto{
		ID: in.Id,
	})
}

func (u *accountUsecaseImpl) AddPersonAccount(in *models.AddPersonAccountData) error {
	_, err := u.repository.SelectAccount(&entities.GetUserDto{
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

func (u *accountUsecaseImpl) GetAccountDetails(in *models.GetAccountData) (*entities.User, error) {
	found, err := u.repository.SelectAccount(&entities.GetUserDto{
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

func NewAccountUsecaseImpl(repo repositories.AccountsRepository, c config.Jwt) AccountUsecase {
	return &accountUsecaseImpl{
		repository: repo,
		jwtConfig:  c,
	}
}
