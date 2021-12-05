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


const grievanceStateTransitionViewPath = "/grm/views/grievance_state_transition/"

var GrievanceStateTransition grievanceStateTransitionHandler

type grievanceStateTransitionHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceStateTransitionHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_state_transitions/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)
	

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_state_transitions []models.GrievanceStateTransition

	err := json.Unmarshal(resp.Body, &grievance_state_transitions)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_state_transitions,
		"total": len(grievance_state_transitions),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceStateTransitionViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceStateTransitionHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	grievance_states,_ := handler.FetchAllGrievanceStates()

	data := services.Map{
		"title": "Create New Grievance State Transition",
		"new":   true,
		"grievance_states": grievance_states,
	}

	return c.Render(http.StatusOK, grievanceStateTransitionViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceStateTransitionHandler) Store(c echo.Context) error {

	grievance_state_transition := models.GrievanceStateTransition{}

	if err := c.Bind(&grievance_state_transition); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_state_transitions/store"

	params := map[string]string{
		"description":       grievance_state_transition.Description,
		"from_state_id":       fmt.Sprintf("%v", grievance_state_transition.FromStateId),
		"to_state_id":       		fmt.Sprintf("%v", grievance_state_transition.ToStateId),
	}


	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Transition created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_state_transitions")

}


func (handler *grievanceStateTransitionHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state_transition models.GrievanceStateTransition

	if err := c.Bind(&grievance_state_transition); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state_transition.Id

	endPoint := fmt.Sprintf("/grievance_state_transitions/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_state_transition)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_state_transition,
		"title": "Show Grievance State Transition",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateTransitionViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceStateTransitionHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_state_transition models.GrievanceStateTransition

	if err := c.Bind(&grievance_state_transition); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_state_transition.Id

	endPoint := fmt.Sprintf("/grievance_state_transitions/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_state_transition)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_states,_ := handler.FetchAllGrievanceStates()
	grievance_to_state,_ := handler.FetchGrievanceState(grievance_state_transition.ToStateId)
	grievance_from_state,_ := handler.FetchGrievanceState(grievance_state_transition.FromStateId)

	data := services.Map{
		"grievance_states": grievance_states,
		"grievance_to_state": grievance_to_state,
		"grievance_from_state": grievance_from_state,
		"data": grievance_state_transition,
		"title": "Edit Grievance State Transition",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceStateTransitionViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceStateTransitionHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_state_transition := models.GrievanceStateTransition{}
	
	if err := c.Bind(&grievance_state_transition); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_state_transition_id := fmt.Sprintf("%v", grievance_state_transition.Id)

	endPoint := "/grievance_state_transitions/update"

	params := map[string]string{
		"id":                     grievance_state_transition_id,
		"description":       grievance_state_transition.Description,
		"from_state_id":       fmt.Sprintf("%v", grievance_state_transition.FromStateId),
		"to_state_id":       		fmt.Sprintf("%v", grievance_state_transition.ToStateId),
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Transition updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_state_transitions")

}


func (handler *grievanceStateTransitionHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_state_transition := models.GrievanceStateTransition{}

	if err := c.Bind(&grievance_state_transition); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_state_transition.Id)

	endPoint := "/grievance_state_transitions/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance State Transition deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_state_transitions")

}

func (handler *grievanceStateTransitionHandler) FetchAllGrievanceStates() ([]models.GrievanceState, error) {
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

 func (handler *grievanceStateTransitionHandler) FetchGrievanceState(Id int) (models.GrievanceState, error) {
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