package main

import (
	"go-code-generator/utils"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigType("properties")
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	utils.CreateRouter()
	utils.CreateController()
	utils.CreateService()
}
