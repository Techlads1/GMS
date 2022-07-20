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
	"strconv"
)

const grievanceResolutionViewPath = "/grm/views/grievance_resolution/"

var GrievanceResolution grievanceResolutionHandler

type grievanceResolutionHandler struct{}


//Index this is a landing page
func (handler *grievanceResolutionHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_resolution/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_resolution []models.GrievanceResolution

	err := json.Unmarshal(resp.Body, &grievance_resolution)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_resolution,
		"total": len(grievance_resolution),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceResolutionHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance Resolution",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceResolutionHandler) Store(c echo.Context) error {

	grievance_resolution := models.GrievanceResolution{}

	if err := c.Bind(&grievance_resolution); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_resolution/store"

	var grievance_id int = 1
	var gfu_id int = 1

	params := map[string]string{
		"grievance_id": strconv.Itoa(grievance_id),
		"gfu_id": strconv.Itoa(gfu_id),
		"state":       "1",
		"description": grievance_resolution.Description,
		"comment":"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Resolution created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_resolution")

}


func (handler *grievanceResolutionHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_resolution models.GrievanceResolution

	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_resolution.Id

	endPoint := fmt.Sprintf("/grievance_resolution/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_resolution)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_resolution,
		"title": "Show Grievance Resolution",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceResolutionHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_resolution models.GrievanceResolution

	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_resolution.Id

	endPoint := fmt.Sprintf("/grievance_resolution/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_resolution)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_resolution,
		"title": "Edit Grievance Resolution",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceResolutionHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_resolution := models.GrievanceResolution{}
	
	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_resolution_id := fmt.Sprintf("%v", grievance_resolution.Id)
	endPoint := "/grievance_resolution/update"

	params := map[string]string{
		"id":                     grievance_resolution_id,
		"grievance_id":       fmt.Sprintf("%v", grievance_resolution.GrievanceId),
		"gfu_id":       fmt.Sprintf("%v", grievance_resolution.GFUId),
		"state":       grievance_resolution.State,
		"description": 						grievance_resolution.Description,
		"comment":"Null",
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Resolution updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_resolution")

}

func (handler *grievanceResolutionHandler) Approve(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_resolution := models.GrievanceResolution{}
	
	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_resolution_id := fmt.Sprintf("%v", grievance_resolution.Id)

	endPoint := "/grievance_resolution/update"

	params := map[string]string{
		"id":                   grievance_resolution_id,
		"grievance_id":         fmt.Sprintf("%v", grievance_resolution.GrievanceId),
		"gfu_id":       		fmt.Sprintf("%v", grievance_resolution.GFUId),
		"state":       			grievance_resolution.State,
		"description": 			grievance_resolution.Description,
		"comment":				grievance_resolution.Comment,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Resolution updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_resolution")

}


func (handler *grievanceResolutionHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_resolution := models.GrievanceResolution{}

	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_resolution.Id)

	endPoint := "/grievance_resolution/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance Resolution deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_resolution")

}


func (handler *grievanceResolutionHandler) ChangeState(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_resolution models.GrievanceResolution

	if err := c.Bind(&grievance_resolution); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_resolution.Id

	endPoint := fmt.Sprintf("/grievance_resolution/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_resolution)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	grievance_resolution_states,_ := handler.FetchAllGrievanceResolutionStates()

	data := services.Map{
		"grievance_resolution_states": grievance_resolution_states,
		"data": grievance_resolution,
		"title": "Edit Grievance Resolution",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"grievance_approve", services.Serve(c, data))

}

func (handler *grievanceResolutionHandler) ListNew(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_resolution/list_new"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_resolution []models.GrievanceResolution

	err := json.Unmarshal(resp.Body, &grievance_resolution)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_resolution,
		"total": len(grievance_resolution),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"grievance_new", services.Serve(c, data))

}

func (handler *grievanceResolutionHandler) ListApproved(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_resolution/list_approved"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_resolution []models.GrievanceResolution

	err := json.Unmarshal(resp.Body, &grievance_resolution)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_resolution,
		"total": len(grievance_resolution),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"grievance_approved", services.Serve(c, data))

}

func (handler *grievanceResolutionHandler) ListDenied(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_resolution/list_denied"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_resolution []models.GrievanceResolution

	err := json.Unmarshal(resp.Body, &grievance_resolution)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_resolution,
		"total": len(grievance_resolution),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceResolutionViewPath+"grievance_denied", services.Serve(c, data))

}

func (handler *grievanceResolutionHandler) FetchAllGrievanceResolutionStates() ([]models.GrievanceResolutionState, error) {
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


