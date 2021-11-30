package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/package/log"

	//"fmt"
	"gateway/webserver/services"
	"gateway/webserver/systems"
	"gateway/webserver/systems/grm/models"
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

	endPoint := "/grievant_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievant_categories []models.GrievantCategory

	err := json.Unmarshal(resp.Body, &grievant_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievant_categories,
		"total": len(grievant_categories),
		"error": err,
	}

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

	if err := c.Bind(&grievant_category); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievant_categories/store"

	params := map[string]string{
		"name":       grievant_category.Name,
		"description": grievant_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Category created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_categories")

}


func (handler *grievantCategoryHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievant_category models.GrievantCategory

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievant_category.Id

	endPoint := fmt.Sprintf("/grievant_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievant_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
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

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievant_category.Id

	endPoint := fmt.Sprintf("/grievant_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievant_category)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
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
	
	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievant_category_id := fmt.Sprintf("%v", grievant_category.Id)

	endPoint := "/grievant_categories/update"

	params := map[string]string{
		"id":                     grievant_category_id,
		"name":       						grievant_category.Name,
		"description": 						grievant_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Category updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_categories")

}


func (handler *grievantCategoryHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievant_category := models.GrievantCategory{}

	if err := c.Bind(&grievant_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievant_category.Id)

	endPoint := "/grievant_categories/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Category deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievant_categories")

}

