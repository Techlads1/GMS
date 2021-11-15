package grm

import (
	"log"

	"gateway/webserver/systems/grm/controllers"

	"github.com/labstack/echo/v4"
)

// WebRouters initialises web routes
func WebRouters(app *echo.Echo) {
	//put here routers for every controller
	log.Printf("GRM routers initialised....\n")

	grmApp := app.Group("/grm")

	department := grmApp.Group("/department")
	{
		department.GET("", controllers.Department.Index)             //    /department/index
		department.GET("/show/:id", controllers.Department.Show)     //    /department/show/[id]
		department.GET("/edit/:id", controllers.Department.Edit)     //  /department/edit/[id]
		department.GET("/delete/:id", controllers.Department.Delete) //  /department/delete/[id]
	}

	//put here all your web routes

}
