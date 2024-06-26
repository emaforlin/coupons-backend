package main

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/pkg/config"
	"github.com/emaforlin/coupons-app/pkg/database"
	"github.com/emaforlin/coupons-app/pkg/server"
)

func main() {
	config.InitViper("config.yaml")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	fmt.Printf("Time: %s", time.Now().Format(time.DateTime))
	server.NewEchoServer(conf, db).Start()
}
