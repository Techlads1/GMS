package middlewares

import (
	"gateway/package/log"
	"gateway/webserver/auth"
	"gateway/webserver/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// JWT Middleware
func JWT() echo.MiddlewareFunc {
	return auth.JWT()
}

// CheckLogin middleware
func CheckAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if auth.SkipperLoginCheck(c) {
				return next(c)
			}
			u := c.Get("user").(*jwt.Token)

			if u == nil {
				return c.String(http.StatusUnauthorized, "unable to get token")
			}

			claims := u.Claims.(*auth.JWTCustomClaims)
			if claims == nil {
				return c.String(http.StatusUnauthorized, "unable to read claims")
			}

			url := c.Request().URL
			p, err := services.HasPermission(claims.Email, url.String())
			if p && err == nil {
				//log.Infoln("Access Granted")
				return next(c)
			} else {
				log.Infoln("Access Denied!")
				services.ClearCache(claims.Email)
				//return c.Redirect(http.StatusUnauthorized, "/")
				return c.Redirect(http.StatusFound, "/login")
			}
			//return c.Redirect(http.StatusFound, "/login")
		}
	}
}
