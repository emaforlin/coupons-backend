package repositories

import "github.com/emaforlin/coupons-app/pkg/entities"

type AccountsRepository interface {
	SelectUser(in *entities.GetUserDto) (*entities.User, error)
	DeleteUser(in *entities.GetUserDto) error

	InsertPerson(in *entities.InsertPersonDto) error
	InsertFoodPlace(in *entities.InsertFoodPlaceDto) error
}
