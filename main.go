package main

import "fileshare/app"

func main() {
	server := app.InitServer()
	server.Start()
}
