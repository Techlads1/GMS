package grm

import (
	"log"

	
	"gateway/webserver/systems/grm/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// WebRouters initialises web routes
func WebRouters(app *echo.Echo) {
	//put here routers for every controller
	log.Printf("GRM routers initialised....\n")

	app.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{Getter: middleware.MethodFromForm("_method")}))

	grmApp := app.Group("/grm")

	department := grmApp.Group("/department")
	{
		department.GET("", controllers.Department.Index)             //    /department/index
		department.GET("/show/:id", controllers.Department.Show)     //    /department/show/[id]
		department.GET("/edit/:id", controllers.Department.Update)   //  /department/edit/[id]
		department.GET("/delete/:id", controllers.Department.Delete) //  /department/delete/[id]
	}

	grievant_categories := grmApp.Group("/grievant_categories")
	{
		grievant_categories.GET("", controllers.GrievantCategory.Index)
		grievant_categories.GET("/create", controllers.GrievantCategory.Create)  
		grievant_categories.POST("/store", controllers.GrievantCategory.Store)
		grievant_categories.GET("/show/:id", controllers.GrievantCategory.Show) 
		grievant_categories.GET("/edit/:id", controllers.GrievantCategory.Edit)  
		grievant_categories.PUT("/update/:id", controllers.GrievantCategory.Update) 
		grievant_categories.DELETE("/delete/:id", controllers.GrievantCategory.Delete) 
	}


	//put here all your web routes

}
