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


const grievantGroupViewPath = "/grm/views/grievant_group/"

var GrievantGroup grievantGroupHandler

type grievantGroupHandler struct{
	
}

//Index this is a landing page
func (handler *grievantGroupHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievant_groups/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievant_groups []models.GrievantGroup

	err := json.Unmarshal(resp.Body, &grievant_groups)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievant_groups,
		"total": len(grievant_groups),
		"error": err,
	}

	return c.Render(http.StatusOK, grievantGroupViewPath+"index", services.Serve(c, data))

}


func (handler *grievantGroupHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	grievant_categories,_ := handler.FetchAllGrievantCategories()

	data := services.Map{
		"title": "Create New Grievant Group",
		"new":   true,
		"grievant_categories": grievant_categories,
	}

	return c.Render(http.StatusOK, grievantGroupViewPath+"create", services.Serve(c, data))

}


func (handler *grievantGroupHandler) Store(c echo.Context) error {

	grievant_Group := models.GrievantGroup{}

	if err := c.Bind(&grievant_Group); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievant_groups/store"

	params := map[string]string{
		"name":       grievant_Group.Name,
		"description": grievant_Group.Description,
		"grievant_category_id":  fmt.Sprintf("%v", grievant_Group.GrievantCategoryId),
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Group created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_groups")

}


func (handler *grievantGroupHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievant_Group models.GrievantGroup

	if err := c.Bind(&grievant_Group); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievant_Group.Id

	endPoint := fmt.Sprintf("/grievant_groups/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievant_Group)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievant_Group,
		"title": "Show Grievant Group",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievantGroupViewPath+"show", services.Serve(c, data))

}


func (handler *grievantGroupHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievant_Group models.GrievantGroup

	if err := c.Bind(&grievant_Group); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievant_Group.Id

	endPoint := fmt.Sprintf("/grievant_groups/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievant_Group)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievant_categories,_ := handler.FetchAllGrievantCategories()
	grievant_category,_ := handler.FetchGrievantCategory(grievant_Group.GrievantCategoryId)

	data := services.Map{
		"grievant_categories": grievant_categories,
		"grievant_category": grievant_category,
		"data": grievant_Group,
		"title": "Edit Grievant Group",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievantGroupViewPath+"edit", services.Serve(c, data))

}

func (handler *grievantGroupHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievant_Group := models.GrievantGroup{}
	
	if err := c.Bind(&grievant_Group); err != nil {
		log.Errorf("%s\n", err)
	}

	grievant_Group_id := fmt.Sprintf("%v", grievant_Group.Id)

	endPoint := "/grievant_groups/update"

	params := map[string]string{
		"id":                     grievant_Group_id,
		"name":       						grievant_Group.Name,
		"grievant_category_id":    fmt.Sprintf("%v", grievant_Group.GrievantCategoryId),
 		"description": 						grievant_Group.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Group updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievant_groups")

}


func (handler *grievantGroupHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievant_Group := models.GrievantGroup{}

	if err := c.Bind(&grievant_Group); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievant_Group.Id)

	endPoint := "/grievant_groups/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievant Group deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievant_groups")

}

 func (handler *grievantGroupHandler) FetchAllGrievantCategories() ([]models.GrievantCategory, error) {
	endPoint := "/grievant_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievant_categories []models.GrievantCategory

	err := json.Unmarshal(resp.Body, &grievant_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievant_categories, nil

 }

 func (handler *grievantGroupHandler) FetchGrievantCategory(Id int) (models.GrievantCategory, error) {
	var grievant_category models.GrievantCategory

	endPoint := fmt.Sprintf("/grievant_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievant_category, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievant_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievant_category, err
	}

	return grievant_category, nil

 }