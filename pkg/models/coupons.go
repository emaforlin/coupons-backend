package models

type (
	AddCoupon struct {
		ID        uint    `json:"id,omitempty" validate:"required"`
		OwnerID   uint    `json:"owner_id,omitempty" validate:"required"`
		Code      string  `json:"code,omitempty" validate:"omitempty,min=8,max=16"`
		Title     string  `json:"title" validate:"max=128"`
		Discount  float32 `json:"discount" validate:"gt=0;lte=100"`
		Remaining int     `json:"remaining" validate:"gt=0"`
	}
	GetCoupons struct {
		ID        uint `query:"id"`
		BatchSize int  `query:"batch_size"`
	}
)
