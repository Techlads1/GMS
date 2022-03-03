package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const helpViewPath = "/grm/views/help/"

var Help help

type help struct{}

//Index this is a frequency asked questions page
func (r *help) faq(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, helpViewPath+"faq", services.Serve(c, data))

}

//Index this is a contact us page
func (r *help) contact_us(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, helpViewPath+"contact_us", services.Serve(c, data))
}
