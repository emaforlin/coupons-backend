package entities

import "gorm.io/gorm"

type (
	InsertCouponDto struct {
		ID          uint32         `gorm:"primaryKey;autoIncrement" json:"coupon_id"`
		Description string         `gorm:"not null" json:"description"`
		Discount    int8           `gorm:"not null" json:"discount"`
		Remaining   uint16         `gorm:"not null" json:"remaining"`
		DeletedAt   gorm.DeletedAt `gorm:"" json:"deleted_at"`
	}

	GetCouponDto struct {
		ID uint32
	}

	Coupon struct {
		ID          uint32         `gorm:"primaryKey;autoIncrement" json:"coupon_id"`
		Description string         `gorm:"not null" json:"description"`
		Discount    int8           `gorm:"not null" json:"discount"`
		Remaining   uint16         `gorm:"not null" json:"remaining"`
		DeletedAt   gorm.DeletedAt `gorm:"" json:"deleted_at"`
	}
)
