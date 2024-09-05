package config

import (
	"flag"

	"github.com/spf13/viper"
)

var (
	debug bool
)

type Config struct {
	App App
	Db  Db
	Jwt Jwt
}

type App struct {
	ApiVersion string
	Port       int
}

type Db struct {
	Uri    string
	Name   string
	User   string
	Passwd string
	Host   string
}

type Jwt struct {
	Secret []byte
	TTL    uint // minutes
}

func LoadConfig() *Config {
	var p = 80
	if debug {
		p = 8080
	}
	return &Config{
		App: App{
			ApiVersion: viper.GetString("service.api"),
			Port:       p,
		},
		Db: Db{
			Name:   viper.GetString("database.name"),
			User:   viper.GetString("database.user"),
			Passwd: viper.GetString("database.password"),
			Host:   viper.GetString("database.host"),
		},
		Jwt: Jwt{
			Secret: []byte(viper.GetString("service.jwt.secret")),
			TTL:    viper.GetUint("service.jwt.ttl"),
		},
	}
}

func InitViper(filename string) {
	flag.BoolVar(&debug, "debug", false, "enable debug mode")
	flag.Parse()

	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./.conf")

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}
}
