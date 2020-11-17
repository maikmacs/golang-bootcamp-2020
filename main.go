package main

import (
	"golang-bootcamp/config"

	"golang-bootcamp/infrastructure/server"
)

func main() {
	config.Init()
	server.Init()
}
