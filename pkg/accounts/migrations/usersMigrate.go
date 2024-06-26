package main

import (
	"fmt"

	"github.com/emaforlin/coupons-app/pkg/accounts/entities"
	"github.com/emaforlin/coupons-app/pkg/accounts/models"
	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/database"
)

var records = []models.AddPersonAccountData{
	{
		Username:    "TheBoss",
		FirstName:   "Emanuel",
		LastName:    "Forlin",
		PhoneNumber: "+54 3496413921",
		Email:       "emaa.forlin@gmail.com",
		Password:    "BWsafEJNnws",
	}, {
		Username:    "matox",
		FirstName:   "Matias",
		LastName:    "Alegre",
		PhoneNumber: "+54 123678",
		Email:       "mati.triste@gmail.com",
		Password:    "estoytriste",
	}, {
		Username:    "nacho",
		FirstName:   "Ignacio",
		LastName:    "Burgi",
		PhoneNumber: "+54 3496555667",
		Email:       "nacho.nn@gmail.com",
		Password:    "hvwoeg423ty",
	},
}

func main() {
	config.InitViper("config.yaml")
	cfg := config.LoadConfig()
	db := database.NewMySQLDatabase(cfg)
	usersMigrate(db, records)
}

func usersMigrate(db database.Database, data []models.AddPersonAccountData) {
	for _, r := range data {
		user := entities.InsertUserDto{
			Username:    r.Username,
			Email:       r.Email,
			PhoneNumber: r.PhoneNumber,
			Password:    r.Password,
			Role:        "Customer",
		}
		cust := entities.InsertPersonDto{
			FirstName: r.FirstName,
			LastName:  r.LastName,
			User:      user,
		}
		res2 := db.GetDb().Omit("created_at", "updated_at", "deleted_at").Create(&cust)
		// res2 := db.GetDb().Create(&cust)
		fmt.Printf("Rows affected: %d\n", res2.RowsAffected)
	}
}
