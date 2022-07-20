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

const grievanceTimeExtensionViewPath = "/grm/views/grievance_time_extension/"

var GrievanceTimeExtension grievanceTimeExtensionHandler

type grievanceTimeExtensionHandler struct{}


//Index this is a landing page
func (handler *grievanceTimeExtensionHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_time_extension/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_time_extension []models.GrievanceTimeExtension

	err := json.Unmarshal(resp.Body, &grievance_time_extension)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_time_extension,
		"total": len(grievance_time_extension),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceTimeExtensionHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance Time Extension",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceTimeExtensionHandler) Store(c echo.Context) error {

	grievance_time_extension := models.GrievanceTimeExtension{}

	if err := c.Bind(&grievance_time_extension); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_time_extension/store"

	var grievance_id int = 1
	var gfu_id int = 1

	params := map[string]string{
		"grievance_id": strconv.Itoa(grievance_id),
		"gfu_id": strconv.Itoa(gfu_id),
		"state":       "1",
		"description": grievance_time_extension.Description,
		"comment":"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Time Extension created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_time_extension")

}


func (handler *grievanceTimeExtensionHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_time_extension models.GrievanceTimeExtension

	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_time_extension.Id

	endPoint := fmt.Sprintf("/grievance_time_extension/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_time_extension)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_time_extension,
		"title": "Show Grievance TimeExtension",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceTimeExtensionHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_time_extension models.GrievanceTimeExtension

	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_time_extension.Id

	endPoint := fmt.Sprintf("/grievance_time_extension/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_time_extension)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_time_extension,
		"title": "Edit Grievance Time Extension",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceTimeExtensionHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_time_extension := models.GrievanceTimeExtension{}
	
	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_time_extension_id := fmt.Sprintf("%v", grievance_time_extension.Id)

	endPoint := "/grievance_time_extension/update"

	params := map[string]string{
		"id":                     grievance_time_extension_id,
		"grievance_id":       	  fmt.Sprintf("%v", grievance_time_extension.GrievanceId),
		"gfu_id":       		  fmt.Sprintf("%v", grievance_time_extension.GFUId),
		"state":       			  grievance_time_extension.State,
		"description": 			  grievance_time_extension.Description,
		"comment":"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance TimeExtension updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_time_extension")

}

func (handler *grievanceTimeExtensionHandler) Approve(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_time_extension := models.GrievanceTimeExtension{}
	
	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_time_extension_id := fmt.Sprintf("%v", grievance_time_extension.Id)

	endPoint := "/grievance_time_extension/update"

	params := map[string]string{
		"id":               grievance_time_extension_id,
		"grievance_id":     fmt.Sprintf("%v", grievance_time_extension.GrievanceId),
		"gfu_id":       	fmt.Sprintf("%v", grievance_time_extension.GFUId),
		"state":       		grievance_time_extension.State,
		"description": 		grievance_time_extension.Description,
		"comment":			grievance_time_extension.Comment,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance TimeExtension updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_time_extension")

}


func (handler *grievanceTimeExtensionHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_time_extension := models.GrievanceTimeExtension{}

	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_time_extension.Id)

	endPoint := "/grievance_time_extension/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance TimeExtension deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_time_extension")

}


func (handler *grievanceTimeExtensionHandler) ChangeState(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_time_extension models.GrievanceTimeExtension

	if err := c.Bind(&grievance_time_extension); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_time_extension.Id

	endPoint := fmt.Sprintf("/grievance_time_extension/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_time_extension)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_time_extension,
		"title": "Edit Grievance Time Extension",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"grievance_approve", services.Serve(c, data))

}

func (handler *grievanceTimeExtensionHandler) ListNew(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_time_extension/list_new"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_time_extension []models.GrievanceTimeExtension

	err := json.Unmarshal(resp.Body, &grievance_time_extension)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_time_extension,
		"total": len(grievance_time_extension),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"grievance_new", services.Serve(c, data))

}

func (handler *grievanceTimeExtensionHandler) ListApproved(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_time_extension/list_approved"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_time_extension []models.GrievanceTimeExtension

	err := json.Unmarshal(resp.Body, &grievance_time_extension)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_time_extension,
		"total": len(grievance_time_extension),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"grievance_approved", services.Serve(c, data))

}

func (handler *grievanceTimeExtensionHandler) ListDenied(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_time_extension/list_denied"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_time_extension []models.GrievanceTimeExtension

	err := json.Unmarshal(resp.Body, &grievance_time_extension)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_time_extension,
		"total": len(grievance_time_extension),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceTimeExtensionViewPath+"grievance_denied", services.Serve(c, data))

}

func (handler *grievanceTimeExtensionHandler) FetchAllGrievanceResolutionStates() ([]models.GrievanceResolutionState, error) {
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


