package controllers

import (
	"gateway/webserver/services"
	"net/http"

	"github.com/k0kubun/pp"

	"github.com/labstack/echo/v4"
)

const LoginViewPath = "/grm/views/account_management/"

var Login Loginpage

type Loginpage struct{}

//Index this is a landing page
func (r *Loginpage) Index(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, LoginViewPath+"login.html", services.Serve(c, data))
}

func (r *Loginpage) Registration(c echo.Context) error {
	pp.Printf("in the index file...\n")

	data := services.Map{
		"error": nil,
	}

	return c.Render(http.StatusOK, LoginViewPath+"otherRegistration.html", services.Serve(c, data))
}
