package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const complainantViewPath = "/grm/views/complainant/"

var Complainant complainants

type complainants struct{}

//Index this is a landing page
func (r *complainants) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"index.html", services.Serve(c, data))
}

func (r *complainants) Academic(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"Academic_categories.html", services.Serve(c, data))
}

func (r *complainants) Extra_curricular(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"etra_curricular.html", services.Serve(c, data))
}
func (r *complainants) Placements(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"placementsAndinternship.html", services.Serve(c, data))
}

func (r *complainants) Maintanance(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"maintanance.html", services.Serve(c, data))
}

func (r *complainants) OtherRelatedIssues(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"otherRelatedIssues.html", services.Serve(c, data))
}

func (r *complainants) General(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"GeneralAdministration.html", services.Serve(c, data))
}

func (r *complainants) Grievance(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"Complaint_Form_Annonymus.html", services.Serve(c, data))
}

func (r *complainants) Complaints(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantViewPath+"Complaint_Form_ForRegistered.html", services.Serve(c, data))
}

const followupViewPath = "/grm/views/followUp/"

var Followup followups

type followups struct{}

//Index this is a landing page
func (r *followups) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, followupViewPath+"followUp_registered.html", services.Serve(c, data))
}
