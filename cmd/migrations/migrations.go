package main

import (
	"github.com/emaforlin/coupons-app/internal/config"
	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
)

func AutoMigrate(db database.Database) {
	db.GetDb().AutoMigrate(entities.InsertUserDto{})
	db.GetDb().AutoMigrate(entities.InsertPersonDto{})
	db.GetDb().AutoMigrate(entities.InsertFoodPlaceDto{})
}

func main() {
	config.InitViper("config")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	AutoMigrate(db)
}
