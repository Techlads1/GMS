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