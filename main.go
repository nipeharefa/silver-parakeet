package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	gateway "github.com/nipeharefa/silver-parakeet/api_gateway"
	search "github.com/nipeharefa/silver-parakeet/search_service"
	"github.com/spf13/viper"
)

type (
	Service interface {
		Run() error
	}
)

func init() {
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	viper.SetConfigFile(`config.yml`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Error(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("add server or consumer parameter.")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "api-gateway":

		apiGateway := gateway.NewAPIGateway()
		apiGateway.Run()

	case "movie-service":
		search.New().Run()
	}
}
