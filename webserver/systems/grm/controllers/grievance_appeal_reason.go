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


const grievanceAppealReasonViewPath = "/grm/views/grievance_appeal_reason/"

var GrievanceAppealReason grievanceAppealReasonHandler

type grievanceAppealReasonHandler struct{
	
}

//Index this is a landing page
func (handler *grievanceAppealReasonHandler) Index(c echo.Context) error {

	pp.Printf("in the index file...\n")

	endPoint := "/grievance_appeal_reasons/list"

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var grievance_appeal_reasons []models.GrievanceAppealReason

	err := json.Unmarshal(resp.Body, &grievance_appeal_reasons)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data":  grievance_appeal_reasons,
		"total": len(grievance_appeal_reasons),
		"error": err,
	}

	return c.Render(http.StatusOK, grievanceAppealReasonViewPath+"index", services.Serve(c, data))

}


func (handler *grievanceAppealReasonHandler) Create(c echo.Context) error {

	pp.Println("in the create file...\n")

	data := services.Map{
		"title": "Create New Grievance AppealReason",
		"new":   true,
	}

	return c.Render(http.StatusOK, grievanceAppealReasonViewPath+"create", services.Serve(c, data))

}


func (handler *grievanceAppealReasonHandler) Store(c echo.Context) error {

	grievance_appeal_reason := models.GrievanceAppealReason{}

	if err := c.Bind(&grievance_appeal_reason); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	
	endPoint := "/grievance_appeal_reasons/store"

	params := map[string]string{
		"name":       grievance_appeal_reason.Name,
		"description": grievance_appeal_reason.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance AppealReason created successfully!")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_appeal_reasons")

}


func (handler *grievanceAppealReasonHandler) Show(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_appeal_reason models.GrievanceAppealReason

	if err := c.Bind(&grievance_appeal_reason); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_appeal_reason.Id

	endPoint := fmt.Sprintf("/grievance_appeal_reasons/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	err := json.Unmarshal(resp.Body, &grievance_appeal_reason)

	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_appeal_reason,
		"title": "Show Grievance AppealReason",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceAppealReasonViewPath+"show", services.Serve(c, data))

}


func (handler *grievanceAppealReasonHandler) Edit(c echo.Context) error {

	pp.Println("in the update file...")

	var grievance_appeal_reason models.GrievanceAppealReason

	if err := c.Bind(&grievance_appeal_reason); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := grievance_appeal_reason.Id

	endPoint := fmt.Sprintf("/grievance_appeal_reasons/show/%d", Id)

	resp := systems.GRMAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}
	
	err := json.Unmarshal(resp.Body, &grievance_appeal_reason)
	
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}

	data := services.Map{
		"data": grievance_appeal_reason,
		"title": "Edit Grievance AppealReason",
		"new":   false,
	}

	return c.Render(http.StatusOK, grievanceAppealReasonViewPath+"edit", services.Serve(c, data))

}

func (handler *grievanceAppealReasonHandler) Update(c echo.Context) error {
	pp.Println("in the update file...")
	
	grievance_appeal_reason := models.GrievanceAppealReason{}
	
	if err := c.Bind(&grievance_appeal_reason); err != nil {
		log.Errorf("%s\n", err)
	}

	grievance_appeal_reason_id := fmt.Sprintf("%v", grievance_appeal_reason.Id)

	endPoint := "/grievance_appeal_reasons/update"

	params := map[string]string{
		"id":                     grievance_appeal_reason_id,
		"name":       						grievance_appeal_reason.Name,
		"description": 						grievance_appeal_reason.Description,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance AppealReason updated successfully")

	return c.Redirect(http.StatusSeeOther, "/grm/grievance_appeal_reasons")

}


func (handler *grievanceAppealReasonHandler) Delete(c echo.Context) error {

	pp.Println("in the delete file...")

	grievance_appeal_reason := models.GrievanceAppealReason{}

	if err := c.Bind(&grievance_appeal_reason); err != nil {
		log.Errorf("%s\n", err)
	}

	Id := fmt.Sprintf("%v", grievance_appeal_reason.Id)

	endPoint := "/grievance_appeal_reasons/delete"

	params := map[string]string{
		"id":         Id,
	}

	resp := systems.GRMAPI.Send(endPoint, params, true)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Grievance AppealReason deleted successfully")

	return c.Redirect(http.StatusSeeOther,  "/grm/grievance_appeal_reasons")

}

