package main

import (
	"fmt"
	"github.com/spf13/viper"
	"role-catalog-go/inits"
	"role-catalog-go/router"
)

func main() {
	viperConfig := viper.New()
	inits.LoadConfig("role-catalog-dev.yaml", viperConfig)
	inits.LoadDBConfig()
	// getting env variables DB.PORT
	// viper.Get() returns an empty interface{}
	// so we have to do the type assertion, to get the value
	port := viperConfig.Get("SERVER.PORT")
	strPort := fmt.Sprintf("%v", port)

	fmt.Printf("viper : %s = %s \n", "server", strPort)
	if port != nil {
		routes := router.InitializeRoleCatalogRoute()
		routes.Run(strPort)
	}

}
