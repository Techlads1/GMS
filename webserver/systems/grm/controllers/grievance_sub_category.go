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


const grievanceSubCategoryViewPath = "/grm/views/grievance_sub_category/"

var GrievanceSubCategory grievanceSubCategoryHandler

type grievanceSubCategoryHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceSubCategoryHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_sub_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_sub_categories []models.GrievanceSubCategory

	err := json.Unmarshal(resp.Body, &grievance_sub_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_sub_categories,
		"total": len(grievance_sub_categories),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceSubCategoryViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceSubCategoryHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	grievance_categories,_ := handler.FetchAllGrievanceCategories()

	data := services.Map{
		"title": "Create New Grievance Category",
		"new":   true,
		"grievance_categories": grievance_categories,
	}

	return c.Render(http.StatusOK, grievanceSubCategoryViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceSubCategoryHandler) Store(c echo.Context) error {

	grievance_sub_category := models.GrievanceSubCategory{}

	if err := c.Bind(&grievance_sub_category); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_sub_categories/store"

	params := map[string]string{
		"name":      							 grievance_sub_category.Name,
		"code_name":       				 grievance_sub_category.CodeName,
		"grievance_category_id":   fmt.Sprintf("%v", grievance_sub_category.GrievanceCategoryId),
		"description":  					 grievance_sub_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_sub_categories")

}


func (handler *grievanceSubCategoryHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_sub_category models.GrievanceSubCategory

	if err := c.Bind(&grievance_sub_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_sub_category.Id

	endPoint := fmt.Sprintf("/grievance_sub_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_sub_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_sub_category,
		"title": "Show Grievance Sub Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceSubCategoryViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceSubCategoryHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_sub_category models.GrievanceSubCategory

	if err := c.Bind(&grievance_sub_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_sub_category.Id

	endPoint := fmt.Sprintf("/grievance_sub_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_sub_category)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_categories,_ := handler.FetchAllGrievanceCategories()
	grievance_category,_ := handler.FetchGrievanceCategory(grievance_sub_category.GrievanceCategoryId)

	data := services.Map{
		"data": grievance_sub_category,
		"grievance_categories": grievance_categories,
		"grievance_category": grievance_category,
		"title": "Edit Grievance Sub Category",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceSubCategoryViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceSubCategoryHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_sub_category := models.GrievanceSubCategory{}
	
	if err := c.Bind(&grievance_sub_category); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_sub_category_id := fmt.Sprintf("%v", grievance_sub_category.Id)

	endPoint := "/grievance_sub_categories/update"

	params := map[string]string{
		"id":                     grievance_sub_category_id,
		"name":       						grievance_sub_category.Name,
		"code_name":       				grievance_sub_category.CodeName,
		"grievance_category_id":    fmt.Sprintf("%v", grievance_sub_category.GrievanceCategoryId),
		"description": 						grievance_sub_category.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_sub_categories")

}


func (handler *grievanceSubCategoryHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_sub_category := models.GrievanceSubCategory{}

	if err := c.Bind(&grievance_sub_category); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_sub_category.Id)

	endPoint := "/grievance_sub_categories/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Category deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_sub_categories")

}

func (handler *grievanceSubCategoryHandler) FetchAllGrievanceCategories() ([]models.GrievanceCategory, error) {
	endPoint := "/grievance_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_categories []models.GrievanceCategory

	err := json.Unmarshal(resp.Body, &grievance_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievance_categories, nil

 }

 func (handler *grievanceSubCategoryHandler) FetchGrievanceCategory(Id int) (models.GrievanceCategory, error) {
	var grievance_category models.GrievanceCategory

	endPoint := fmt.Sprintf("/grievance_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievance_category, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievance_category, err
	}

	return grievance_category, nil

 }