package main

import (
	"github.com/NetSinx/go-restful-API/app"
)

func main() {
	app.ConnectDB()

	router := app.Router()

	router.Run(":8080")
}