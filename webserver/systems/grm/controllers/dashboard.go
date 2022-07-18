package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const dashboardViewPath = "/grm/views/dashboard/"

var Dashboard dashboards

type dashboards struct{}

//Index this is a landing page
func (r *dashboards) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, dashboardViewPath+"index", services.Serve(c, data))
}

const gfudashboardViewPath = "/grm/views/gfu/"

var GFUdashboard Gfudashboards

type Gfudashboards struct{}

//Index this is a landing page
func (r *Gfudashboards) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, gfudashboardViewPath+"index", services.Serve(c, data))
}

const grcdashboardViewPath = "/grm/views/grc/"

var GRCdashboard Grcdashboards

type Grcdashboards struct{}

//Index this is a landing page
func (r *Grcdashboards) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, grcdashboardViewPath+"index", services.Serve(c, data))
}

const officerdashboardViewPath = "/grm/views/officers/"

var Officerdashboard officerdashboards

type officerdashboards struct{}

//Index this is a landing page
func (r *officerdashboards) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, officerdashboardViewPath+"index", services.Serve(c, data))
}

const complainantdashboardViewPath = "/grm/views/complainant/"

var Complainantdashboard complainantdashboards

type complainantdashboards struct{}

//Index this is a landing page
func (r *complainantdashboards) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, complainantdashboardViewPath+"index", services.Serve(c, data))
}
