package auth

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWT() echo.MiddlewareFunc {
	if accessTokenMiddleware != nil {
		return accessTokenMiddleware
	}
	accessTokenMiddleware = middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:                  &JWTCustomClaims{},
		SigningKey:              []byte(GetJWTSecret()),
		TokenLookup:             "cookie:" + accessTokenCookieName,
		ErrorHandlerWithContext: JWTErrorChecker,
		Skipper:                 SkipperLoginCheck,
	})
	return accessTokenMiddleware
}

//SkipperLoginCheck register all routes that do not need login
func SkipperLoginCheck(c echo.Context) bool {
	if strings.HasSuffix(c.Path(), "/login") ||
		strings.HasSuffix(c.Path(), "/register") ||
		strings.HasSuffix(c.Path(), "/register/create") ||
		strings.HasSuffix(c.Path(), "/forgot-password") ||
		strings.HasPrefix(c.Path(), "/images") ||
		strings.HasPrefix(c.Path(), "/favicon.ico") ||
		strings.Contains(c.Path(), "adminlte") {
		return true
	}
	return false
}

// JWTErrorChecker will be executed when user try to access a protected path.
func JWTErrorChecker(err error, c echo.Context) error {
	//log.Error(err)
	return c.Redirect(http.StatusSeeOther, "/login")
	//return c.Redirect(http.StatusSeeOther, c.Echo().Reverse("signin"))
}
