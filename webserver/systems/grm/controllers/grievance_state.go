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


const grievanceStateViewPath = "/grm/views/grievance_state/"

var GrievanceState grievanceStateHandler

type grievanceStateHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceStateHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_states/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)
	

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_states []models.GrievanceState

	err := json.Unmarshal(resp.Body, &grievance_states)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_states,
		"total": len(grievance_states),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceStateViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceStateHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance State",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceStateViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceStateHandler) Store(c echo.Context) error {

	grievance_state := models.GrievanceState{}

	if err := c.Bind(&grievance_state); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_states/store"

	params := map[string]string{
		"name":       grievance_state.Name,
		"code_name":       grievance_state.CodeName,
		"days":       		fmt.Sprintf("%v", grievance_state.Days),
		"description": grievance_state.Description,
	}


	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_states")

}


func (handler *grievanceStateHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state models.GrievanceState

	if err := c.Bind(&grievance_state); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state.Id

	endPoint := fmt.Sprintf("/grievance_states/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_state)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_state,
		"title": "Show Grievance State",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceStateHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state models.GrievanceState

	if err := c.Bind(&grievance_state); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state.Id

	endPoint := fmt.Sprintf("/grievance_states/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_state)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_state,
		"title": "Edit Grievance State",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceStateHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_state := models.GrievanceState{}
	
	if err := c.Bind(&grievance_state); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_state_id := fmt.Sprintf("%v", grievance_state.Id)

	endPoint := "/grievance_states/update"

	params := map[string]string{
		"id":                     grievance_state_id,
		"name":       						grievance_state.Name,
		"code_name":       				grievance_state.CodeName,
		"days":       				    fmt.Sprintf("%v", grievance_state.Days),
		"description": 						grievance_state.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_states")

}


func (handler *grievanceStateHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_state := models.GrievanceState{}

	if err := c.Bind(&grievance_state); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_state.Id)

	endPoint := "/grievance_states/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_states")

}

