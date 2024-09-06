package repositories

import (
	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
)

type couponsMysqlRepositoryImpl struct {
	db database.Database
}

// DeleteCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) DeleteCoupon(id int) error {
	panic("unimplemented")
}

// InsertCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) InsertCoupon(in *entities.InsertCouponDto) (int, error) {
	panic("unimplemented")
}

// SelectAllCoupons implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) SelectAllCoupons() ([]*entities.Coupon, error) {
	panic("unimplemented")
}

// UpdateCoupon implements CouponsRepository.
func (c *couponsMysqlRepositoryImpl) UpdateCoupon(in *entities.InsertCouponDto) error {
	panic("unimplemented")
}

func NewCouponMysqlRepositoryImpl(d database.Database) CouponsRepository {
	return &couponsMysqlRepositoryImpl{
		db: d,
	}
}
