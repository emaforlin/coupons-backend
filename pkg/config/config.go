package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	ApiVersion string
	Ports      map[string]uint16
}

type Db struct {
	Uri    string
	Name   string
	User   string
	Passwd string
	Host   string
}

func LoadConfig() *Config {
	var ports map[string]uint16
	err := viper.UnmarshalKey("service.ports", &ports)
	if err != nil {
		panic("error loading config file")
	}
	return &Config{
		App: App{
			ApiVersion: viper.GetString("service.api"),
			Ports:      ports,
		},
		Db: Db{
			Name:   viper.GetString("database.name"),
			User:   viper.GetString("database.user"),
			Passwd: viper.GetString("database.password"),
			Host:   viper.GetString("database.host"),
		},
	}
}

func InitViper(filename string) {
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}
}
