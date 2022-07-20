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


const grievanceStateActionViewPath = "/grm/views/grievance_state_action/"

var GrievanceStateAction grievanceStateActionHandler

type grievanceStateActionHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceStateActionHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_state_actions/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)
	

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_state_actions []models.GrievanceStateAction

	err := json.Unmarshal(resp.Body, &grievance_state_actions)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_state_actions,
		"total": len(grievance_state_actions),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceStateActionViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceStateActionHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	grievance_states,_ := handler.FetchAllGrievanceStates()

	data := services.Map{
		"title": "Create New Grievance State Action",
		"new":   true,
		"grievance_states": grievance_states,
	}

	return c.Render(http.StatusOK, grievanceStateActionViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceStateActionHandler) Store(c echo.Context) error {

	grievance_state_action := models.GrievanceStateAction{}

	if err := c.Bind(&grievance_state_action); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_state_actions/store"

	params := map[string]string{
		"name":       grievance_state_action.Name,
		"role_perform_action":       grievance_state_action.RolePerformAction,
		"state_id":       		fmt.Sprintf("%v", grievance_state_action.StateId),
	}


	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Action created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_state_actions")

}


func (handler *grievanceStateActionHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state_action models.GrievanceStateAction

	if err := c.Bind(&grievance_state_action); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state_action.Id

	endPoint := fmt.Sprintf("/grievance_state_actions/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_state_action)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_state_action,
		"title": "Show Grievance State Action",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateActionViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceStateActionHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state_action models.GrievanceStateAction

	if err := c.Bind(&grievance_state_action); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state_action.Id

	endPoint := fmt.Sprintf("/grievance_state_actions/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_state_action)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_states,_ := handler.FetchAllGrievanceStates()
	grievance_state,_ := handler.FetchGrievanceState(grievance_state_action.StateId)

	data := services.Map{
		"grievance_states": grievance_states,
		"grievance_state": grievance_state,
		"data": grievance_state_action,
		"title": "Edit Grievance State Action",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateActionViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceStateActionHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_state_action := models.GrievanceStateAction{}
	
	if err := c.Bind(&grievance_state_action); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_state_action_id := fmt.Sprintf("%v", grievance_state_action.Id)

	endPoint := "/grievance_state_actions/update"

	params := map[string]string{
		"id":                     grievance_state_action_id,
		"name":       						grievance_state_action.Name,
		"role_perform_action":       				grievance_state_action.RolePerformAction,
		"state_id":       				    fmt.Sprintf("%v", grievance_state_action.StateId),
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Action updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_state_actions")

}


func (handler *grievanceStateActionHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_state_action := models.GrievanceStateAction{}

	if err := c.Bind(&grievance_state_action); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_state_action.Id)

	endPoint := "/grievance_state_actions/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Action deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_state_actions")

}

func (handler *grievanceStateActionHandler) FetchAllGrievanceStates() ([]models.GrievanceState, error) {
	endPoint := "/grievance_states/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_state []models.GrievanceState

	err := json.Unmarshal(resp.Body, &grievance_state)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievance_state, nil

 }

 func (handler *grievanceStateActionHandler) FetchGrievanceState(Id int) (models.GrievanceState, error) {
	var grievance_state models.GrievanceState

	endPoint := fmt.Sprintf("/grievance_states/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return grievance_state, errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_state)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return grievance_state, err
	}

	return grievance_state, nil

 }