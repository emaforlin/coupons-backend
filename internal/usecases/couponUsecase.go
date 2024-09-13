package usecases

import (
	"log"

	"github.com/emaforlin/coupons-app/internal/repositories"
	"github.com/emaforlin/coupons-app/pkg/entities"
	"github.com/emaforlin/coupons-app/pkg/models"
)

type CouponUsecase interface {
	Create(in *models.AddCoupon) (int, error)
	Get(in *models.GetCoupons) ([]entities.Coupon, error)
	Delete(id int) error
	Update(in *models.AddCoupon) error
	Use(couponCode string) (string, error)
	Claim(id int) (string, error)
}

type couponUsecaseImpl struct {
	repository repositories.CouponsRepository
}

// Claim implements CouponUsecase.
func (c *couponUsecaseImpl) Claim(id int) (string, error) {
	panic("unimplemented")
}

// Create implements CouponUsecase.
func (c *couponUsecaseImpl) Create(in *models.AddCoupon) (int, error) {
	id, err := c.repository.InsertCoupon(&entities.InsertCouponDto{
		OwnerID:   in.OwnerID,
		Code:      in.Code,
		Title:     in.Title,
		Discount:  in.Discount,
		Remaining: in.Remaining,
	})
	if err != nil {
		log.Println("cannot create new coupon", err)
		return -1, err
	}
	return id, nil
}

// Delete implements CouponUsecase.
func (c *couponUsecaseImpl) Delete(id int) error {
	return c.repository.DeleteCoupon(id)
}

// GetAll implements CouponUsecase.
func (c *couponUsecaseImpl) Get(in *models.GetCoupons) ([]entities.Coupon, error) {
	coupons, err := c.repository.SelectCoupons(&entities.GetCouponDto{ID: in.ID, BatchSize: in.BatchSize})

	if err != nil {
		log.Println("error retrieving coupons", err)
		return nil, err
	}
	return coupons, nil
}

// Update implements CouponUsecase.
func (c *couponUsecaseImpl) Update(in *models.AddCoupon) error {
	err := c.repository.UpdateCoupon(&entities.InsertCouponDto{
		OwnerID:   in.OwnerID,
		Code:      in.Code,
		Title:     in.Title,
		Discount:  in.Discount,
		Remaining: in.Remaining,
	})
	if err != nil {
		log.Println("cannot create new coupon", err)
		return err
	}
	return nil
}

// Use implements CouponUsecase.
func (c *couponUsecaseImpl) Use(couponCode string) (string, error) {
	panic("unimplemented")
}

func NewCouponUsecase(mysqlRepo repositories.CouponsRepository) CouponUsecase {
	return &couponUsecaseImpl{
		repository: mysqlRepo,
	}
}
