package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/letanthang/echo_stackdriver/db"
	"github.com/letanthang/echo_stackdriver/route"
)

var psqlClient = db.GetDB()

func main() {
	e := echo.New()

	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10,
	}))

	e.Use(middleware.RequestID())
	route.Public(e)
	// if config.Config.Profiler.StatsdAddress != "" {
	// 	e.Use(profiler.ProfilerWithConfig(profiler.ProfilerConfig{Address: config.Config.Profiler.StatsdAddress, Service: config.Config.Profiler.Service}))
	// }

	port := "9090"
	log.Println("Starting at port: " + port)
	err := e.Start(":" + port)
	if err != nil {
		log.Println(err)
	}
	defer psqlClient.Close()
}
