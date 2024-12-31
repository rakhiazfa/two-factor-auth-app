package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rakhiazfa/gin-boilerplate/config"
	"github.com/rakhiazfa/gin-boilerplate/internal/wire"
	"github.com/spf13/viper"
)

func main() {
	config.InitViper(".")

	local, err := time.LoadLocation(viper.GetString("application.timezone"))
	if err != nil {
		log.Fatal("Failed to load location: ", err)
	}

	time.Local = local

	addr := fmt.Sprintf("%s:%d", viper.GetString("application.host"), viper.GetInt("application.port"))

	err = wire.NewApplication().Run(addr)
	if err != nil {
		log.Fatal("Failed to run application: ", err)
	}
}
