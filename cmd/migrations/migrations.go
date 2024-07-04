package main

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/entities"
)

func AutoMigrate(db database.Database) {
	db.GetDb().AutoMigrate(entities.InsertUserDto{})
	db.GetDb().AutoMigrate(entities.InsertPersonDto{})
	db.GetDb().AutoMigrate(entities.InsertFoodPlaceDto{})
}

func main() {
	config.InitViper("config.yaml")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	AutoMigrate(db)
	fmt.Printf("Time: %s", time.Now().Format(time.DateTime))

}
