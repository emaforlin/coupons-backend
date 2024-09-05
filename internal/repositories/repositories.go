package repositories

import "github.com/emaforlin/coupons-app/pkg/entities"

type AccountsRepository interface {
	SelectAccount(in *entities.GetUserDto) (*entities.User, error)
	DeleteAccount(in *entities.GetUserDto) error

	SelectPerson(in *entities.GetPersonDto) (*entities.Person, error)
	InsertPerson(in *entities.InsertPersonDto) error

	SelectFoodPlace(in *entities.GetFoodPlaceDto) (*entities.FoodPlace, error)
	InsertFoodPlace(in *entities.InsertFoodPlaceDto) error
	UpdateFoodPlace(userId uint32, in *entities.InsertFoodPlaceDto) error
}
