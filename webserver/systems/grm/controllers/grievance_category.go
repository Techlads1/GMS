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


const grievanceCategoryViewPath = "/grm/views/grievance_category/"

var GrievanceCategory grievanceCategoryHandler

type grievanceCategoryHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceCategoryHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_categories []models.GrievanceCategory

	err := json.Unmarshal(resp.Body, &grievance_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_categories,
		"total": len(grievance_categories),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceCategoryViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceCategoryHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance Category",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceCategoryViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceCategoryHandler) Store(c echo.Context) error {

	grievance_category := models.GrievanceCategory{}

	if err := c.Bind(&grievance_category); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_categories/store"

	params := map[string]string{
		"name":       grievance_category.Name,
		"code_name":       grievance_category.CodeName,
		"description": grievance_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_categories")

}


func (handler *grievanceCategoryHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_category models.GrievanceCategory

	if err := c.Bind(&grievance_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_category.Id

	endPoint := fmt.Sprintf("/grievance_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_category,
		"title": "Show Grievance Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceCategoryViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceCategoryHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_category models.GrievanceCategory

	if err := c.Bind(&grievance_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_category.Id

	endPoint := fmt.Sprintf("/grievance_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_category)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_category,
		"title": "Edit Grievance Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceCategoryViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceCategoryHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_category := models.GrievanceCategory{}
	
	if err := c.Bind(&grievance_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_category_id := fmt.Sprintf("%v", grievance_category.Id)

	endPoint := "/grievance_categories/update"

	params := map[string]string{
		"id":                     grievance_category_id,
		"name":       						grievance_category.Name,
		"code_name":       				grievance_category.CodeName,
		"description": 						grievance_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_categories")

}


func (handler *grievanceCategoryHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_category := models.GrievanceCategory{}

	if err := c.Bind(&grievance_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_category.Id)

	endPoint := "/grievance_categories/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_categories")

}

