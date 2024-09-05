package main

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/internal/config"
	"github.com/emaforlin/coupons-app/internal/database"
	"github.com/emaforlin/coupons-app/internal/server"
)

func main() {
	config.InitViper("config")
	conf := config.LoadConfig()
	db := database.NewMySQLDatabase(conf)
	fmt.Printf("Time: %s", time.Now().Format(time.DateTime))
	server.NewEchoServer(conf, db).Start()
}
