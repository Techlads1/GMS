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

	grievances := grmApp.Group("/grievances")
	{
		grievances.GET("", controllers.Grievance.Index)
		grievances.GET("/create", controllers.Grievance.Create)  
		grievances.POST("/store", controllers.Grievance.Store)
		grievances.GET("/show/:id", controllers.Grievance.Show) 
		grievances.GET("/edit/:id", controllers.Grievance.Edit)  
		grievances.PUT("/update/:id", controllers.Grievance.Update) 
		grievances.DELETE("/delete/:id", controllers.Grievance.Delete) 
	}

	dashboards := grmApp.Group("/dashboard")
	{
		dashboards.GET("", controllers.Dashboard.Index)
		
	}

	grievance_resolution := grmApp.Group("/grievance_resolution")
	{
		grievance_resolution.GET("", controllers.GrievanceResolution.Index)
		grievance_resolution.GET("/create", controllers.GrievanceResolution.Create)  
		grievance_resolution.POST("/store", controllers.GrievanceResolution.Store)
		grievance_resolution.GET("/show/:id", controllers.GrievanceResolution.Show) 
		grievance_resolution.GET("/edit/:id", controllers.GrievanceResolution.Edit)  
		grievance_resolution.GET("/changestate/:id", controllers.GrievanceResolution.ChangeState)  
		grievance_resolution.PUT("/update/:id", controllers.GrievanceResolution.Update) 
		grievance_resolution.PUT("/approve/:id", controllers.GrievanceResolution.Approve) 
		grievance_resolution.DELETE("/delete/:id", controllers.GrievanceResolution.Delete)
		grievance_resolution.GET("/list_new", controllers.GrievanceResolution.ListNew)
		grievance_resolution.GET("/list_approved", controllers.GrievanceResolution.ListApproved)
		grievance_resolution.GET("/list_denied", controllers.GrievanceResolution.ListDenied) 
	}

	grievance_resolution_state := grmApp.Group("/grievance_resolution_state")
	{
		grievance_resolution_state.GET("", controllers.GrievanceResolutionState.Index)
		grievance_resolution_state.GET("/create", controllers.GrievanceResolutionState.Create)  
		grievance_resolution_state.POST("/store", controllers.GrievanceResolutionState.Store)
		grievance_resolution_state.GET("/show/:id", controllers.GrievanceResolutionState.Show) 
		grievance_resolution_state.GET("/edit/:id", controllers.GrievanceResolutionState.Edit)  
		grievance_resolution_state.PUT("/update/:id", controllers.GrievanceResolutionState.Update) 
		grievance_resolution_state.DELETE("/delete/:id", controllers.GrievanceResolutionState.Delete) 
	}

	grievance_forward := grmApp.Group("/grievance_forward")
	{
		grievance_forward.GET("", controllers.GrievanceForward.Index)
		grievance_forward.GET("/create", controllers.GrievanceForward.Create)  
		grievance_forward.POST("/store", controllers.GrievanceForward.Store)
		grievance_forward.GET("/show/:id", controllers.GrievanceForward.Show) 
		grievance_forward.GET("/edit/:id", controllers.GrievanceForward.Edit)  
		grievance_forward.GET("/changestate/:id", controllers.GrievanceForward.ChangeState)  
		grievance_forward.PUT("/update/:id", controllers.GrievanceForward.Update) 
		grievance_forward.PUT("/approve/:id", controllers.GrievanceForward.Approve) 
		grievance_forward.DELETE("/delete/:id", controllers.GrievanceForward.Delete)
		grievance_forward.GET("/list_new", controllers.GrievanceForward.ListNew)
		grievance_forward.GET("/list_approved", controllers.GrievanceForward.ListApproved)
		grievance_forward.GET("/list_denied", controllers.GrievanceForward.ListDenied) 
	}

	grievance_time_extension := grmApp.Group("/grievance_time_extension")
	{
		grievance_time_extension.GET("", controllers.GrievanceTimeExtension.Index)
		grievance_time_extension.GET("/create", controllers.GrievanceTimeExtension.Create)  
		grievance_time_extension.POST("/store", controllers.GrievanceTimeExtension.Store)
		grievance_time_extension.GET("/show/:id", controllers.GrievanceTimeExtension.Show) 
		grievance_time_extension.GET("/edit/:id", controllers.GrievanceTimeExtension.Edit)  
		grievance_time_extension.GET("/changestate/:id", controllers.GrievanceTimeExtension.ChangeState)  
		grievance_time_extension.PUT("/update/:id", controllers.GrievanceTimeExtension.Update) 
		grievance_time_extension.PUT("/approve/:id", controllers.GrievanceTimeExtension.Approve) 
		grievance_time_extension.DELETE("/delete/:id", controllers.GrievanceTimeExtension.Delete)
		grievance_time_extension.GET("/list_new", controllers.GrievanceTimeExtension.ListNew)
		grievance_time_extension.GET("/list_approved", controllers.GrievanceTimeExtension.ListApproved)
		grievance_time_extension.GET("/list_denied", controllers.GrievanceTimeExtension.ListDenied) 
	}
	//put here all your web routes

}
