package repositories

import (
	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
	"golang.org/x/crypto/bcrypt"
)

type accountsMysqlRepositoryImpl struct {
	db database.Database
}

func (u *accountsMysqlRepositoryImpl) SelectFoodPlace(in *entities.GetFoodPlaceDto) (*entities.FoodPlace, error) {
	result := &entities.FoodPlace{}
	response := u.db.GetDb().Model(entities.GetFoodPlaceDto{}).First(&result, in)
	if response.Error != nil {
		return nil, response.Error
	}
	return result, nil
}

func (u *accountsMysqlRepositoryImpl) SelectPerson(in *entities.GetPersonDto) (*entities.Person, error) {
	result := &entities.Person{}
	response := u.db.GetDb().Model(entities.GetPersonDto{}).First(&result, in)
	if response.Error != nil {
		return nil, response.Error
	}
	return result, nil
}

func (u *accountsMysqlRepositoryImpl) GetAccountPasswordHash(in *entities.GetUserDto) (string, error) {
	type getPasswd struct {
		Password string `gorm:"not null" json:"password"`
	}

	result := getPasswd{}
	response := u.db.GetDb().Table("users").First(&result, in)
	if response.Error != nil {
		return "", response.Error
	}
	return result.Password, nil
}

func (u *accountsMysqlRepositoryImpl) UpdateFoodPlace(id uint32, in *entities.InsertFoodPlaceDto) error {
	response := u.db.GetDb().Where("user_id = ?", id).Updates(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) InsertFoodPlace(in *entities.InsertFoodPlaceDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
	response := u.db.GetDb().Create(in)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) DeleteAccount(in *entities.GetUserDto) error {
	res := u.db.GetDb().Delete(in)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (u *accountsMysqlRepositoryImpl) SelectAccount(in *entities.GetUserDto) (*entities.User, error) {
	result := &entities.User{}
	response := u.db.GetDb().Model(entities.GetUserDto{}).Find(&result, in)
	if response.Error != nil {
		return nil, response.Error
	}
	return result, nil
}

func (u *accountsMysqlRepositoryImpl) InsertPerson(in *entities.InsertPersonDto) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(in.User.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	// hash password
	in.User.Password = string(hash)
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
