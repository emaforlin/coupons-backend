package entities

import (
	"time"

	"gorm.io/gorm"
)

type (
	InsertUserDto struct {
		ID          uint32         `gorm:"primaryKey;autoIncrement" json:"user_id"`
		Username    string         `gorm:"not null;unique" json:"username"`
		Role        string         `gorm:"type:enum('Customer','FoodPlace');default:'Customer'" json:"role"`
		Email       string         `gorm:"not null;unique" json:"email"`
		PhoneNumber string         `gorm:"not null;unique" json:"phone_number"`
		Password    string         `gorm:"not null" json:"password"`
		CreatedAt   time.Time      `gorm:"auusertoCreateTime" json:"created_at"`
		UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"" json:"deleted_at"`
	}

	GetUserDto struct {
		ID          uint32
		Username    string
		Role        string
		Email       string
		PhoneNumber string
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt
	}

	User struct {
		ID          uint32         `gorm:"primaryKey;autoIncrement" json:"user_id"`
		Username    string         `gorm:"not null;unique" json:"username"`
		Role        string         `gorm:"type:enum('Customer','FoodPlace');default:'Customer'" json:"role"`
		PhoneNumber string         `gorm:"not null;unique" json:"phone_number"`
		Email       string         `gorm:"not null;unique" json:"email"`
		Password    string         `gorm:"not null" json:"password"`
		CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
		UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
		DeletedAt   gorm.DeletedAt `gorm:"" json:"deleted_at"`
	}
)

func (InsertUserDto) TableName() string {
	return "users"
}

func (GetUserDto) TableName() string {
	return "users"
}
