package config

import (
	"log"

	"github.com/spf13/viper"
)

func InitViper(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigType("json")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Failed to load configuration: ", err)
	}
}
