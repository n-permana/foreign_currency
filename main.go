package main

import (
	"foreign_currency/router"
)

func main() {

	router := router.Init()
	router.Logger.Fatal(router.Start(":7001"))
}
