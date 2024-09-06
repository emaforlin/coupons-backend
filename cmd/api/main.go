package main

import (
	"fmt"
	"time"

	"github.com/emaforlin/coupons-app/internal/config"
	"github.com/emaforlin/coupons-app/internal/database"
)

func main() {
	config.InitViper("config")
	conf := config.LoadConfig()
	_ = database.NewMySQLDatabase(conf)
	fmt.Printf("Time: %s", time.Now().Format(time.DateTime))

	// server.NewEchoServer(conf, db).Start()
}
