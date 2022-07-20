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

const departmentViewPath = "/aim/views/department/"

var Department departments

type departments struct{}

//Index this is a landing page
func (r *departments) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	//replace this with the actual api to pull department data that will be used to render the index page
	endPoint := "/departments/list-full"

	resp := systems.AimAPI.Send(endPoint, nil, false)
	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var departments []models.DepartmentFull
	err := json.Unmarshal(resp.Body, &departments)
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}
	data := services.Map{
		"data":  departments,
		"total": len(departments),
		"error": err,
	}

	//services.

	return c.Render(http.StatusOK, departmentViewPath+"index", services.Serve(c, data))
}

//Show Single Record
func (r *departments) Show(c echo.Context) error {
	pp.Printf("in the show file...\n")

	department := models.DepartmentFull{}

	if err := c.Bind(&department); err != nil {
		log.Errorf("%s\n", err)
	}

	rID := department.Id

	endPoint := fmt.Sprintf("/departments/get-full/%d", rID)

	resp := systems.AimAPI.Send(endPoint, nil, false)

	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	var departments models.DepartmentFull
	err := json.Unmarshal(resp.Body, &departments)
	if err != nil {
		pp.Printf("error decoding json data: %v\n", err)
		return err
	}
	data := services.Map{
		"data":  departments,
		"error": err,
	}

	return c.Render(http.StatusOK, departmentViewPath+"show", services.Serve(c, data))
}

//Create Record
func (r *departments) Create(c echo.Context) error {
	pp.Println("create department file...")

	//For campuses
	/*var campuses []models.Campus
	err := r.FetchAll("/campuses/list", &campuses)
	if err != nil {
		pp.Printf("error retrieving campuses: %v\n", err)
		return err
	}*/

	data := services.Map{
		//"camps": campuses,
		"title": "Create New Department",
		"new":   true,
	}

	return c.Render(http.StatusOK, departmentViewPath+"create", services.Serve(c, data))

}

//CreateDepartment create a department
func (r *departments) CreateDepartment(c echo.Context) error {

	department := models.Department{}

	if err := c.Bind(&department); err != nil {
		services.SetErrorMessage(c, err.Error())
		log.Errorf("%s\n", err)
	}
	department.CreatedBy = 1

	endPoint := "/departments/create"

	params := map[string]string{
		"department_title":       department.DepartmentTitle,
		"department_description": department.DepartmentDescription,
		"department_size":        fmt.Sprintf("%d", department.DepartmentSize),
		"campus_id":              fmt.Sprintf("%d", department.CampusId),
		"created_by":             fmt.Sprintf("%d", department.CreatedBy),
	}

	resp := systems.AimAPI.Send(endPoint, params, true)
	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Department created successfully")

	return c.Redirect(http.StatusSeeOther, "/aim/department")
}

//Update Record
func (r *departments) Update(c echo.Context) error {
	pp.Println("update department file...")

	//For campuses
	/*var campuses []models.Campus
	err := r.FetchAll("/campuses/list", &campuses)
	if err != nil {
		pp.Printf("error retrieving campuses: %v\n", err)
		return err
	}*/

	department := models.Department{}

	if err := c.Bind(&department); err != nil {
		log.Errorf("%s\n", err)
	}

	rID := department.Id

	var departments models.Department
	endPoint := fmt.Sprintf("/departments/get/%d", rID)
	err := r.FetchAll(endPoint, &departments)
	if err != nil {
		pp.Printf("error retrieving campuses: %v\n", err)
		return err
	}

	/*var campusSingle models.Campus
	endPointSingle := fmt.Sprintf("/campuses/get/%d", departments.CampusId)
	err = r.FetchAll(endPointSingle, &campusSingle)
	if err != nil {
		pp.Printf("error retrieving campuses: %v\n", err)
		return err
	}*/

	data := services.Map{
		//"campselected": campusSingle,
		//"camps":        campuses,
		"data":  departments,
		"error": err,
	}
	return c.Render(http.StatusOK, departmentViewPath+"update", services.Serve(c, data))

}

//UpdateDepartment update a
func (r *departments) UpdateDepartment(c echo.Context) error {

	department := models.Department{}

	if err := c.Bind(&department); err != nil {
		log.Errorf("%s\n", err)
	}

	departmentId := fmt.Sprintf("%v", department.Id)
	updatedBy := fmt.Sprintf("%v", 1)

	endPoint := "/departments/update"

	params := map[string]string{
		"id":                     departmentId,
		"department_title":       department.DepartmentTitle,
		"department_description": department.DepartmentDescription,
		"department_size":        fmt.Sprintf("%d", department.DepartmentSize),
		"campus_id":              fmt.Sprintf("%d", department.CampusId),
		"updated_by":             updatedBy,
	}

	resp := systems.AimAPI.Send(endPoint, params, true)
	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Department updated successfully")

	return c.Redirect(http.StatusSeeOther, "/aim/department")
}

//Delete Record
func (r *departments) Delete(c echo.Context) error {
	pp.Println("delete department file...")

	department := models.Department{}

	if err := c.Bind(&department); err != nil {
		log.Errorf("%s\n", err)
	}

	rID := fmt.Sprintf("%v", department.Id)

	deletedBy := fmt.Sprintf("%v", 1)

	endPoint := "/departments/delete"

	params := map[string]string{
		"id":         rID,
		"deleted_by": deletedBy,
	}

	resp := systems.AimAPI.Send(endPoint, params, true)
	if resp == nil && resp.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", resp.StatusCode)
		return errors.New("error getting response")
	}

	services.SetInfoMessage(c, "Department deleted successfully")

	return c.Redirect(http.StatusSeeOther, "/aim/department")
}

func (r *departments) FetchAll(endPoint string, modelName interface{}) error {

	respList := systems.AimAPI.Send(endPoint, nil, false)
	if respList == nil && respList.StatusCode != http.StatusOK {
		pp.Printf("error getting data: %v\n", respList.StatusCode)
		return errors.New("error getting response")
	}

	// Fill the data with the data from the JSON
	//var campuses []models.Campus
	errC := json.Unmarshal(respList.Body, &modelName)
	if errC != nil {
		pp.Printf("error decoding json data: %v\n", errC)
		return errC
	}

	return nil
}
