package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gateway/package/log"
	"strconv"

	//"fmt"
	"gateway/webserver/services"
	"gateway/webserver/systems"
	"gateway/webserver/systems/grm/models"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const grievanceForwardViewPath = "/grm/views/grievance_forward/"

var GrievanceForward grievanceForwardHandler

type grievanceForwardHandler struct{}


//Index this is a landing page
func (handler *grievanceForwardHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_forward/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_forward []models.GrievanceForward

	err := json.Unmarshal(resp.Body, &grievance_forward)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_forward,
		"total": len(grievance_forward),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceForwardHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance Forward",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceForwardHandler) Store(c echo.Context) error {

	grievance_forward := models.GrievanceForward{}

	if err := c.Bind(&grievance_forward); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_forward/store"

	var grievance_id int = 1
	var fromgfu_id = 1

	params := map[string]string{
		"grievance_id": strconv.Itoa(grievance_id),
		"state":       "1",
		"fromgfu_id":	strconv.Itoa(fromgfu_id),
		"togfu_id":		fmt.Sprintf("%v", grievance_forward.ToGFUId),
		"description": 	grievance_forward.Description,
		"comment":		"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Forward created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_forward")

}


func (handler *grievanceForwardHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_forward models.GrievanceForward

	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_forward.Id

	endPoint := fmt.Sprintf("/grievance_forward/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_forward)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_forward,
		"title": "Show Grievance Forward",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceForwardHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_forward models.GrievanceForward

	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_forward.Id

	endPoint := fmt.Sprintf("/grievance_forward/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_forward)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_forward,
		"title": "Edit Grievance Forward",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceForwardHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_forward := models.GrievanceForward{}
	
	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_forward_id := fmt.Sprintf("%v", grievance_forward.Id)

	endPoint := "/grievance_forward/update"

	params := map[string]string{
		"id":               grievance_forward_id,
		"grievance_id":     fmt.Sprintf("%v", grievance_forward.GrievanceId),
		"state":       		grievance_forward.State,
		"fromgfu_id":       fmt.Sprintf("%v", grievance_forward.FromGFUId),
		"togfu_id":       	fmt.Sprintf("%v", grievance_forward.ToGFUId),
		"description": 		grievance_forward.Description,
		"comment":"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Forward updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_forward")

}

func (handler *grievanceForwardHandler) Approve(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_forward := models.GrievanceForward{}
	
	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_forward_id := fmt.Sprintf("%v", grievance_forward.Id)

	endPoint := "/grievance_forward/update"

	params := map[string]string{
		"id":               grievance_forward_id,
		"grievance_id":     fmt.Sprintf("%v", grievance_forward.GrievanceId),
		"state":       		grievance_forward.State,
		"fromgfu_id":       fmt.Sprintf("%v", grievance_forward.FromGFUId),
		"togfu_id":       	fmt.Sprintf("%v", grievance_forward.ToGFUId),
		"description": 		grievance_forward.Description,
		"comment":			grievance_forward.Comment,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Forward updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_forward")

}


func (handler *grievanceForwardHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_forward := models.GrievanceForward{}

	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_forward.Id)

	endPoint := "/grievance_forward/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Forward deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_forward")

}


func (handler *grievanceForwardHandler) ChangeState(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_forward models.GrievanceForward

	if err := c.Bind(&grievance_forward); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_forward.Id

	endPoint := fmt.Sprintf("/grievance_forward/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_forward)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_forward,
		"title": "Edit Grievance Forward",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"grievance_approve", services.Serve(c, data))

}

func (handler *grievanceForwardHandler) ListNew(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_forward/list_new"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_forward []models.GrievanceForward

	err := json.Unmarshal(resp.Body, &grievance_forward)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_forward,
		"total": len(grievance_forward),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"grievance_new", services.Serve(c, data))

}

func (handler *grievanceForwardHandler) ListApproved(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_forward/list_approved"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_forward []models.GrievanceForward

	err := json.Unmarshal(resp.Body, &grievance_forward)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_forward,
		"total": len(grievance_forward),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"grievance_approved", services.Serve(c, data))

}

func (handler *grievanceForwardHandler) ListDenied(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_forward/list_denied"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_forward []models.GrievanceForward

	err := json.Unmarshal(resp.Body, &grievance_forward)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_forward,
		"total": len(grievance_forward),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceForwardViewPath+"grievance_denied", services.Serve(c, data))

}

func (handler *grievanceForwardHandler) FetchAllGrievanceResolutionStates() ([]models.GrievanceResolutionState, error) {
	endPoint := "/grievance_resolution_state/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return nil, errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_resolution_states []models.GrievanceResolutionState

	err := json.Unmarshal(resp.Body, &grievance_resolution_states)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return nil, err
	}

	return grievance_resolution_states, nil

 }


