package repositories

import (
	"errors"

	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
	"gorm.io/gorm"
)

type couponsMysqlRepositoryImpl struct {
	db database.Database
}

// DeleteCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) DeleteCoupon(id int) error {
	return c.db.GetDb().Delete(&entities.InsertCouponDto{}, id).Error
}

// InsertCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) InsertCoupon(in *entities.InsertCouponDto) (int, error) {
	err := c.db.GetDb().Create(in).Error

	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return -1, errors.New("coupon already exists")
	} else if err != nil {
		return -1, err
	}
	return in.ID, nil
}

// SelectAllCoupons implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) SelectAllCoupons() ([]entities.Coupon, error) {
	var coupons []entities.Coupon

	err := c.db.GetDb().Find(&coupons).Error
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

// UpdateCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) UpdateCoupon(in *entities.InsertCouponDto) error {
	return c.db.GetDb().Where("id = ?", in.ID).Updates(in).Error
}

func NewCouponMysqlRepositoryImpl(d database.Database) CouponsRepository {
	return &couponsMysqlRepositoryImpl{
		db: d,
	}
}
