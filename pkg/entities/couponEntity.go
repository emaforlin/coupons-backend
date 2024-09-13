package entities

import (
	"time"

	"gorm.io/gorm"
)

type (
	InsertCouponDto struct {
		ID        uint           `gorm:"primaryKey;autoIncrement"`
		OwnerID   uint           `gorm:"not null"`
		Code      string         `gorm:"not null;unique"`
		Title     string         `gorm:"not null;unique"`
		Discount  float32        `gorm:"not null"`
		Remaining int            `gorm:"not null"`
		CreatedAt time.Time      `gorm:"autoCreateTime"`
		UpdatedAt time.Time      `gorm:"autoUpdateTime"`
		DeletedAt gorm.DeletedAt `gorm:""`
	}

	GetCouponDto struct {
		ID        uint
		BatchSize int
	}

	Coupon struct {
		ID        uint      `json:"id,omitempty"`
		OwnerID   uint      `json:"owner_id,omitempty"`
		Code      string    `json:"code,omitempty"`
		Title     string    `json:"title"`
		Discount  float32   `json:"discount"`
		Remaining int       `json:"remaining"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at,omitempty"`
		DeletedAt time.Time `json:"deleted_at,omitempty"`
	}
)

func (InsertCouponDto) TableName() string {
	return "coupons"
}

func (GetCouponDto) TableName() string {
	return "coupons"
}
