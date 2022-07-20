package routes

import (
	"gateway/package/validator"
	"gateway/webserver/middlewares"
	"gateway/webserver/systems/grm"

	"github.com/labstack/echo/v4"
)

// Routers function
func Routers(app *echo.Echo) {
	//Common middleware for all type of routers
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger(true))
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	//app.Use(middlewares.CSRF())
	app.Use(middlewares.Session())

	app.Validator = validator.GetValidator() //initialize custom validator

	//register grm subsystem web routes
	grm.WebRouters(app)

}
