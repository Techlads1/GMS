package webserver

import (
	"gateway/webserver/routes"
	"gateway/webserver/services"
	"gateway/webserver/systems"

	"github.com/labstack/echo/v4"
)

//StartWebserver starts a webserver
func StartWebserver() {
	// Echo instance
	e := echo.New()
	//Define renderer
	e.Renderer = Renderer()

	//Disable echo banner
	e.HideBanner = true

	// Routes
	routes.Routers(e)

	//init cache
	services.Init() //check if this solves the problem
	systems.Init()  //check if this solves the problem
	// Start server
	e.Logger.Fatal(e.Start(":4322"))
}
