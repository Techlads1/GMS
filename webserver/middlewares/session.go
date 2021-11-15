package middlewares

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//Session middleware
func Session() echo.MiddlewareFunc {
	sessionKey := "1B1CDF6F75CE839A977F12C8168B7"
	store := sessions.NewCookieStore([]byte(sessionKey))
	return session.Middleware(store)
}
