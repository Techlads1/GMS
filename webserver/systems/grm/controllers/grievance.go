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

const grievanceViewPath = "/grm/views/grievance/"

var Grievance grievanceHandler

type grievanceHandler struct {
}

//Index this is a landing page
func (handler *grievanceHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievances/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievances []models.Grievance

	err := json.Unmarshal(resp.Body, &grievances)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}


	data := services.Map{
		"data":  grievances,
		"total": len(grievances),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceViewPath+"index", services.Serve(c, data))

}

func (handler *grievanceHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	grievance_sub_categories,_ := handler.FetchAllGrievanceSubCategories()
	filling_modes,_ := handler.FetchAllGrievanceFillingModes()
	grievant_groups,_ := handler.FetchAllGrievantGroups()


	data := services.Map{
		"title": "Create New Grievance ",
		"new":   true,
		"grievant_groups": grievant_groups,
		"filling_modes": filling_modes,
		"grievance_sub_categories": grievance_sub_categories,
	}

	return c.Render(http.StatusOK, grievanceViewPath+"create", services.Serve(c, data))

}

func (handler *grievanceHandler) Store(c echo.Context) error {

	grievance := models.Grievance{}

	if err := c.Bind(&grievance); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}

	endPoint := "/grievances/store"
pp.Print(grievance)
	params := map[string]string{
		"name":        grievance.Name,
		"description": grievance.Description,
		"location_occurred": grievance.LocationOccurred,
		"filling_mode_id": fmt.Sprintf("%v", grievance.FillingModeId),
		"grievance_sub_category_id": fmt.Sprintf("%v", grievance.GrievanceSubCategoryId),
		"grievant_group_id": fmt.Sprintf("%v", grievance.GrievantGroupId),
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance  created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievances")

}

func (handler *grievanceHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance models.Grievance

	if err := c.Bind(&grievance); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance.Id

	endPoint := fmt.Sprintf("/grievances/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance,
		"title": "Show Grievance ",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceViewPath+"show", services.Serve(c, data))

}

func (handler *grievanceHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance models.Grievance

	if err := c.Bind(&grievance); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance.Id

	endPoint := fmt.Sprintf("/grievances/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_sub_categories,_ := handler.FetchAllGrievanceSubCategories()
	filling_modes,_ := handler.FetchAllGrievanceFillingModes()
	grievant_groups,_ := handler.FetchAllGrievantGroups()

	grievance_sub_category,_ := handler.FetchGrievanceSubCategory(grievance.GrievanceSubCategoryId)
	filling_mode,_ := handler.FetchGrievanceFillingMode(grievance.FillingModeId)
	grievant_group,_ := handler.FetchGrievantGroup(grievance.GrievantGroupId)

	data := services.Map{
		"data":  grievance,
		"title": "Edit Grievance ",
		"new":   false,
		"grievant_groups": grievant_groups,
		"filling_modes": filling_modes,
		"grievance_sub_categories": grievance_sub_categories,
		"grievant_group": grievant_group,
		"filling_mode": filling_mode,
		"grievance_sub_category": grievance_sub_category,
	}



	return c.Render(http.StatusOK, grievanceViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")

	grievance := models.Grievance{}

	if err := c.Bind(&grievance); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_id := fmt.Sprintf("%v", grievance.Id)

	endPoint := "/grievances/update"

	params := map[string]string{
		"id":          grievance_id,
		"name":        grievance.Name,
		"description": grievance.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance  updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievances")

}

func (handler *grievanceHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance := models.Grievance{}

	if err := c.Bind(&grievance); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance.Id)

	endPoint := "/grievances/delete"

	params := map[string]string{
		"id": Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance  deleted successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievances")

}

func (handler *grievanceHandler) FetchAllGrievanceSubCategories() ([]models.GrievanceSubCategory, error) {
	endPoint := "/grievance_sub_categories/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_sub_categories []models.GrievanceSubCategory

	err := json.Unmarshal(resp.Body, &grievance_sub_categories)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievance_sub_categories, nil

 }

 func (handler *grievanceHandler) FetchGrievanceSubCategory(Id int) (models.GrievanceSubCategory, error) {
	var grievance_sub_category models.GrievanceSubCategory

	endPoint := fmt.Sprintf("/grievance_sub_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievance_sub_category, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_sub_category)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievance_sub_category, err
	}

	return grievance_sub_category, nil

 }

 func (handler *grievanceHandler) FetchAllGrievanceFillingModes() ([]models.GrievanceFillingMode, error) {
	endPoint := "/grievance_filling_modes/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_filling_modes []models.GrievanceFillingMode

	err := json.Unmarshal(resp.Body, &grievance_filling_modes)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievance_filling_modes, nil

 }

 func (handler *grievanceHandler) FetchGrievanceFillingMode(Id int) (models.GrievanceFillingMode, error) {
	var grievance_filling_mode models.GrievanceFillingMode

	endPoint := fmt.Sprintf("/grievance_sub_categories/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievance_filling_mode, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_filling_mode)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievance_filling_mode, err
	}

	return grievance_filling_mode, nil

 }

 func (handler *grievanceHandler) FetchAllGrievantGroups() ([]models.GrievantGroup, error) {
	endPoint := "/grievant_groups/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievant_groups []models.GrievantGroup

	err := json.Unmarshal(resp.Body, &grievant_groups)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievant_groups, nil

 }

 func (handler *grievanceHandler) FetchGrievantGroup(Id int) (models.GrievantGroup, error) {
	var grievant_group models.GrievantGroup

	endPoint := fmt.Sprintf("/grievant_groups/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievant_group, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievant_group)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievant_group, err
	}

	return grievant_group, nil

 }