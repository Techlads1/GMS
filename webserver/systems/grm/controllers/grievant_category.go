package controllers

import (
	"gateway/package/log"
	"gateway/package/util"
	"time"

	//"fmt"
	"gateway/webserver/services"
	"gateway/webserver/systems/grm/models"
	"gateway/webserver/systems/grm/repositories"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)


const grievantCategoryViewPath = "/grm/views/grievant_category/"

var GrievantCategory grievantCategoryHandler

type grievantCategoryHandler struct{
	
}

//Index this is a landing page
func (handler *grievantCategoryHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	service := repositories.NewGrievantCategoryRepository()

	grievant_categories, err := service.All()

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievant_categories,
		"total": len(grievant_categories),
		"error": err,
	}

	// return c.JSON(http.StatusOK,data)
	return c.Render(http.StatusOK, grievantCategoryViewPath+"index", services.Serve(c, data))

}


func (handler *grievantCategoryHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievant Category",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievantCategoryViewPath+"create", services.Serve(c, data))

}


func (handler *grievantCategoryHandler) Store(c echo.Context) error {

	grievant_category := models.GrievantCategory{}

	service := repositories.NewGrievantCategoryRepository()

	if err := c.Bind(&grievant_category); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}

	grievant_category.CreatedAt = time.Now()

	grievant_category.UpdatedAt = time.Now()

	service.Create(grievant_category)

	services.SetInfoMessage(c, "Grievant Category created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_categories")

}


func (handler *grievantCategoryHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievant_category models.GrievantCategory

	service := repositories.NewGrievantCategoryRepository()

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievant_category, err := service.Get(grievant_category.Id)
	
	if err != nil {
		pp.Printf("error retrieving grievant category: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievant_category,
		"title": "Show Grievant Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievantCategoryViewPath+"show", services.Serve(c, data))

}


func (handler *grievantCategoryHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievant_category models.GrievantCategory

	service := repositories.NewGrievantCategoryRepository()

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievant_category, err := service.Get(grievant_category.Id)
	
	if err != nil {
		pp.Printf("error retrieving grievant category: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievant_category,
		"title": "Edit Grievant Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievantCategoryViewPath+"edit", services.Serve(c, data))

}

func (handler *grievantCategoryHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	grievant_category := models.GrievantCategory{}

	service := repositories.NewGrievantCategoryRepository()
	
	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievant_category.UpdatedAt = time.Now()

	data, err := service.Get(grievant_category.Id)

	if util.CheckError(err) {
		return c.JSON(http.StatusInternalServerError, "error retrieving grievant category")
	}

	data.Id = grievant_category.Id
	data.Name = grievant_category.Name
	data.Description = grievant_category.Description
	data.UpdatedAt = time.Now()

	_, err = service.Update(&data)
	pp.Println(err)
	services.SetInfoMessage(c, "Grievant Category updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_categories")

}


func (handler *grievantCategoryHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievant_category := models.GrievantCategory{}

	service := repositories.NewGrievantCategoryRepository()

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	service.Delete(grievant_category.Id)

	services.SetInfoMessage(c, "Grievant Category deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievant_categories")

}

