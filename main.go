package main

import (
	"project/config"
	"project/routes"
)

func main() {
	config.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
}
