package server

import (
	"golang-bootcamp/config"
	"golang-bootcamp/infrastructure/router"
)

func Init() {
	config := config.GetConfig()
	r := router.NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
