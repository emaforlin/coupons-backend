package repositories

import "github.com/emaforlin/coupons-app/pkg/accounts/entities"

type AccountsRepository interface {
	SelectUser(in *entities.GetUserDto) (*entities.User, error)
	DeleteUser(in *entities.GetUserDto) error

	InsertPerson(in *entities.InsertPersonDto) error
}
