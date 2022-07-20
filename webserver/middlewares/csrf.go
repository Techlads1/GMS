package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//CSRF setup basic CSRF configuration
func CSRF() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(
		middleware.CSRFConfig{
			ContextKey:  "csrf",
			TokenLookup: "form:csrf",
		})
}
