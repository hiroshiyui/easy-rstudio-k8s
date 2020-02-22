package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var appName = "easy-rstudio-k8s"

func main() {
	fmt.Println("Hello?")

	// Setup configuration file
	viper.SetConfigName(fmt.Sprintf("%s", appName))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(fmt.Sprintf("/etc/%s/", appName))
	viper.AddConfigPath(fmt.Sprintf("$HOME/.config/%s/", appName))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}
