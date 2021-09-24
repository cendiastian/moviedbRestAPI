package main

import (
	"project/mvc/config"
	oM "project/mvc/middlewares"
	"project/mvc/routes"
)

func main() {
	config.InitDB()
	e := routes.NewRoute()
	e.Start(":8000")
	oM.LogMiddlewareInit(e)
}
