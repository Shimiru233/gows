package main

import (
	"exchangeapp/config"
	"fmt"
)

func main() {
	config.Init()
	fmt.Println(config.AppConfig.App.Name)
	fmt.Println(config.AppConfig.App.Port)
}
