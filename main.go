package main

import (
	"github.com/maikmacs/golang-bootcamp-2020/config"

	"github.com/maikmacs/golang-bootcamp-2020/infrastructure/router"
)

func main() {
	config.InitConfig()
	router.InitRouter()
}
