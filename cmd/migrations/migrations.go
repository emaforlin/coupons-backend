package main

import (
	"log"

	"github.com/emaforlin/coupons-app/internal/config"
	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
)

func AutoMigrate(db database.Database) {
	db.GetDb().AutoMigrate(entities.InsertUserDto{})
	db.GetDb().AutoMigrate(entities.InsertPersonDto{})
	db.GetDb().AutoMigrate(entities.InsertFoodPlaceDto{})
	db.GetDb().AutoMigrate(entities.InsertCouponDto{})
}

func main() {
	config.InitViper("config")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	log.Println("Migrating db...")
	AutoMigrate(db)
	log.Println("Migration ended.")
}
