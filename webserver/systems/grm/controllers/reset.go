package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const ResetViewPath = "/grm/views/account_management/"

var Reset Resetpage

type Resetpage struct{}

//Index this is a landing page
func (r *Resetpage) Index(c echo.Context) error {
	pp.Printf("in the reset file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, ResetViewPath+"reset.html", services.Serve(c, data))
}
