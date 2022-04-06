package main

import (
	"fileshare/app"
	"fileshare/bootstrap"
)

func main() {
	bootstrap.InitializeConfig()
	server := app.InitServer()
	server.Start()
}
