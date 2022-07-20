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


const grievanceFillingModeViewPath = "/grm/views/grievance_filling_mode/"

var GrievanceFillingMode grievanceFillingModeHandler

type grievanceFillingModeHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceFillingModeHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_filling_modes/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_filling_modes []models.GrievanceFillingMode

	err := json.Unmarshal(resp.Body, &grievance_filling_modes)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_filling_modes,
		"total": len(grievance_filling_modes),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceFillingModeViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceFillingModeHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance FillingMode",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceFillingModeViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceFillingModeHandler) Store(c echo.Context) error {

	grievance_filling_mode := models.GrievanceFillingMode{}

	if err := c.Bind(&grievance_filling_mode); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_filling_modes/store"

	params := map[string]string{
		"name":       grievance_filling_mode.Name,
		"code_name":       grievance_filling_mode.CodeName,
		"description": grievance_filling_mode.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance FillingMode created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_filling_modes")

}


func (handler *grievanceFillingModeHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_filling_mode models.GrievanceFillingMode

	if err := c.Bind(&grievance_filling_mode); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_filling_mode.Id

	endPoint := fmt.Sprintf("/grievance_filling_modes/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_filling_mode)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_filling_mode,
		"title": "Show Grievance FillingMode",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceFillingModeViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceFillingModeHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_filling_mode models.GrievanceFillingMode

	if err := c.Bind(&grievance_filling_mode); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_filling_mode.Id

	endPoint := fmt.Sprintf("/grievance_filling_modes/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_filling_mode)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_filling_mode,
		"title": "Edit Grievance FillingMode",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceFillingModeViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceFillingModeHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_filling_mode := models.GrievanceFillingMode{}
	
	if err := c.Bind(&grievance_filling_mode); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_filling_mode_id := fmt.Sprintf("%v", grievance_filling_mode.Id)

	endPoint := "/grievance_filling_modes/update"

	params := map[string]string{
		"id":                     grievance_filling_mode_id,
		"name":       						grievance_filling_mode.Name,
		"code_name":       				grievance_filling_mode.CodeName,
		"description": 						grievance_filling_mode.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance FillingMode updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_filling_modes")

}


func (handler *grievanceFillingModeHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_filling_mode := models.GrievanceFillingMode{}

	if err := c.Bind(&grievance_filling_mode); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_filling_mode.Id)

	endPoint := "/grievance_filling_modes/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance FillingMode deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_filling_modes")

}

