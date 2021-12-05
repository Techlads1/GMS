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

	grievant_groups := grmApp.Group("/grievant_groups")
	{
		grievant_groups.GET("", controllers.GrievantGroup.Index)
		grievant_groups.GET("/create", controllers.GrievantGroup.Create)  
		grievant_groups.POST("/store", controllers.GrievantGroup.Store)
		grievant_groups.GET("/show/:id", controllers.GrievantGroup.Show) 
		grievant_groups.GET("/edit/:id", controllers.GrievantGroup.Edit)  
		grievant_groups.PUT("/update/:id", controllers.GrievantGroup.Update) 
		grievant_groups.DELETE("/delete/:id", controllers.GrievantGroup.Delete) 
	}

	grievance_categories := grmApp.Group("/grievance_categories")
	{
		grievance_categories.GET("", controllers.GrievanceCategory.Index)
		grievance_categories.GET("/create", controllers.GrievanceCategory.Create)  
		grievance_categories.POST("/store", controllers.GrievanceCategory.Store)
		grievance_categories.GET("/show/:id", controllers.GrievanceCategory.Show) 
		grievance_categories.GET("/edit/:id", controllers.GrievanceCategory.Edit)  
		grievance_categories.PUT("/update/:id", controllers.GrievanceCategory.Update) 
		grievance_categories.DELETE("/delete/:id", controllers.GrievanceCategory.Delete) 
	}

	grievance_sub_categories := grmApp.Group("/grievance_sub_categories")
	{
		grievance_sub_categories.GET("", controllers.GrievanceSubCategory.Index)
		grievance_sub_categories.GET("/create", controllers.GrievanceSubCategory.Create)  
		grievance_sub_categories.POST("/store", controllers.GrievanceSubCategory.Store)
		grievance_sub_categories.GET("/show/:id", controllers.GrievanceSubCategory.Show) 
		grievance_sub_categories.GET("/edit/:id", controllers.GrievanceSubCategory.Edit)  
		grievance_sub_categories.PUT("/update/:id", controllers.GrievanceSubCategory.Update) 
		grievance_sub_categories.DELETE("/delete/:id", controllers.GrievanceSubCategory.Delete) 
	}

	grievance_appeal_reasons := grmApp.Group("/grievance_appeal_reasons")
	{
		grievance_appeal_reasons.GET("", controllers.GrievanceAppealReason.Index)
		grievance_appeal_reasons.GET("/create", controllers.GrievanceAppealReason.Create)  
		grievance_appeal_reasons.POST("/store", controllers.GrievanceAppealReason.Store)
		grievance_appeal_reasons.GET("/show/:id", controllers.GrievanceAppealReason.Show) 
		grievance_appeal_reasons.GET("/edit/:id", controllers.GrievanceAppealReason.Edit)  
		grievance_appeal_reasons.PUT("/update/:id", controllers.GrievanceAppealReason.Update) 
		grievance_appeal_reasons.DELETE("/delete/:id", controllers.GrievanceAppealReason.Delete) 
	}

	grievance_filling_modes := grmApp.Group("/grievance_filling_modes")
	{
		grievance_filling_modes.GET("", controllers.GrievanceFillingMode.Index)
		grievance_filling_modes.GET("/create", controllers.GrievanceFillingMode.Create)  
		grievance_filling_modes.POST("/store", controllers.GrievanceFillingMode.Store)
		grievance_filling_modes.GET("/show/:id", controllers.GrievanceFillingMode.Show) 
		grievance_filling_modes.GET("/edit/:id", controllers.GrievanceFillingMode.Edit)  
		grievance_filling_modes.PUT("/update/:id", controllers.GrievanceFillingMode.Update) 
		grievance_filling_modes.DELETE("/delete/:id", controllers.GrievanceFillingMode.Delete) 
	}

	grievance_states := grmApp.Group("/grievance_states")
	{
		grievance_states.GET("", controllers.GrievanceState.Index)
		grievance_states.GET("/create", controllers.GrievanceState.Create)  
		grievance_states.POST("/store", controllers.GrievanceState.Store)
		grievance_states.GET("/show/:id", controllers.GrievanceState.Show) 
		grievance_states.GET("/edit/:id", controllers.GrievanceState.Edit)  
		grievance_states.PUT("/update/:id", controllers.GrievanceState.Update) 
		grievance_states.DELETE("/delete/:id", controllers.GrievanceState.Delete) 
	}

	grievance_state_actions := grmApp.Group("/grievance_state_actions")
	{
		grievance_state_actions.GET("", controllers.GrievanceStateAction.Index)
		grievance_state_actions.GET("/create", controllers.GrievanceStateAction.Create)  
		grievance_state_actions.POST("/store", controllers.GrievanceStateAction.Store)
		grievance_state_actions.GET("/show/:id", controllers.GrievanceStateAction.Show) 
		grievance_state_actions.GET("/edit/:id", controllers.GrievanceStateAction.Edit)  
		grievance_state_actions.PUT("/update/:id", controllers.GrievanceStateAction.Update) 
		grievance_state_actions.DELETE("/delete/:id", controllers.GrievanceStateAction.Delete) 
	}

	grievance_state_transitions := grmApp.Group("/grievance_state_transitions")
	{
		grievance_state_transitions.GET("", controllers.GrievanceStateTransition.Index)
		grievance_state_transitions.GET("/create", controllers.GrievanceStateTransition.Create)  
		grievance_state_transitions.POST("/store", controllers.GrievanceStateTransition.Store)
		grievance_state_transitions.GET("/show/:id", controllers.GrievanceStateTransition.Show) 
		grievance_state_transitions.GET("/edit/:id", controllers.GrievanceStateTransition.Edit)  
		grievance_state_transitions.PUT("/update/:id", controllers.GrievanceStateTransition.Update) 
		grievance_state_transitions.DELETE("/delete/:id", controllers.GrievanceStateTransition.Delete) 
	}

	//put here all your web routes

}
