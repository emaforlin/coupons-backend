package database

import (
	"fmt"

	"github.com/emaforlin/coupons-app/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlDatabase struct {
	db *gorm.DB
}

func NewMySQLDatabase(cfg *config.Config) Database {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.User, cfg.Db.Passwd, cfg.Db.Host, cfg.Db.Name),
		SkipInitializeWithVersion: true,
	}))
	if err != nil {
		panic(err)
	}
	return &mysqlDatabase{db: db}
}

func (p *mysqlDatabase) GetDb() *gorm.DB {
	return p.db
}
