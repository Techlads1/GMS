package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const ForgotViewPath = "/grm/views/account_management/"

var Forgot Forgotpage

type Forgotpage struct{}

//Index this is a landing page
func (r *Forgotpage) Index(c echo.Context) error {
	pp.Printf("in the forgot file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, ForgotViewPath+"forgot.html", services.Serve(c, data))
}
