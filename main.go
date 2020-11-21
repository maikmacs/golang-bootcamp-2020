package main

import (
	"golang-bootcamp/config"

	"golang-bootcamp/infrastructure/router"
)

func main() {
	config.InitConfig()
	router.InitRouter()
}
