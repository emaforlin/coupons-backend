package repositories

import (
	"fmt"

	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type accountsMysqlRepositoryImpl struct {
	db database.Database
}

// InsertFoodPlace implements AccountsRepository.
func (u *accountsMysqlRepositoryImpl) InsertFoodPlace(in *entities.InsertFoodPlaceDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), 12)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
	fmt.Printf("Insert FoodPlace: %#v", in)
	response := u.db.GetDb().Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

// DeleteUser implements AccountsRepository.
func (u *accountsMysqlRepositoryImpl) DeleteUser(in *entities.GetUserDto) error {
	res := u.db.GetDb().Delete(in)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// GetUser implements AccountsRepository.
func (u *accountsMysqlRepositoryImpl) SelectUser(in *entities.GetUserDto) (*entities.User, error) {
	result := &entities.User{}
	response := u.db.GetDb().Table("users").Find(&result, in) //Model(&in).Find(&result)

	if response.Error != nil {
		return nil, response.Error
	}

	return result, nil
}

// AddPerson implements AccountsRepository
func (u *accountsMysqlRepositoryImpl) InsertPerson(in *entities.InsertPersonDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), 12)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
	fmt.Printf("Insert Person: %#v", in)
	response := u.db.GetDb().Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func NewAccountMysqlRepositoryImpl(d database.Database) AccountsRepository {
	return &accountsMysqlRepositoryImpl{
		db: d,
	}
}
