package usecases

import (
	"github.com/emaforlin/coupons-app/pkg/entities"
	"github.com/emaforlin/coupons-app/pkg/models"
)

type CouponUsecase interface {
	Create(in *models.AddCoupon) (int, error)
	GetAll() []*entities.Coupon
	Delete(id int) error
	Update(in *models.AddCoupon) error

	Use(couponCode string) (string, error)

	Claim(id int) (string, error)
}

type couponUsecaseImpl struct{}

// Claim implements CouponUsecase.
func (c *couponUsecaseImpl) Claim(id int) (string, error) {
	panic("unimplemented")
}

// Create implements CouponUsecase.
func (c *couponUsecaseImpl) Create(in *models.AddCoupon) (int, error) {
	panic("unimplemented")
}

// Delete implements CouponUsecase.
func (c *couponUsecaseImpl) Delete(id int) error {
	panic("unimplemented")
}

// GetAll implements CouponUsecase.
func (c *couponUsecaseImpl) GetAll() []*entities.Coupon {
	panic("unimplemented")
}

// Update implements CouponUsecase.
func (c *couponUsecaseImpl) Update(in *models.AddCoupon) error {
	panic("unimplemented")
}

// Use implements CouponUsecase.
func (c *couponUsecaseImpl) Use(couponCode string) (string, error) {
	panic("unimplemented")
}

func NewCouponUsecase() CouponUsecase {
	return &couponUsecaseImpl{}
}
