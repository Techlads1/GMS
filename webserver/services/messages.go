package services

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

//const sessionName = "flash_message" //name of the session
const errorMessageKey = "errors"
const infoMessageKey = "infos"

// SetMessage adds a new message into the session storage.
func SetMessage(c echo.Context, name, value string) {

	sess, _ := session.Get("session", c)
	//session, _ := getCookieStore().Get(c.Request(), sessionName)
	sess.AddFlash(value, name)
	sess.Save(c.Request(), c.Response())
}

// SetErrorMessage adds a new error message into the session storage.
func SetErrorMessage(c echo.Context, message string) {
	SetMessage(c, errorMessageKey, message)
}

// GetErrorMessage adds a new error message into the session storage.
func GetErrorMessage(c echo.Context) []string {
	return GetMessage(c, errorMessageKey)
}

// SetInfoMessage adds a new error message into the session storage.
func SetInfoMessage(c echo.Context, message string) {
	SetMessage(c, infoMessageKey, message)
}

// GetInfoMessage adds a new error message into the session storage.
func GetInfoMessage(c echo.Context) []string {
	return GetMessage(c, infoMessageKey)
}

// GetMessage gets flash messages from the session storage.
func GetMessage(c echo.Context, name string) []string {
	//session, _ := getCookieStore().Get(c.Request(), sessionName)
	sess, _ := session.Get("session", c)
	fm := sess.Flashes(name)
	// If we have some messages.
	if len(fm) > 0 {
		sess.Save(c.Request(), c.Response())
		// Initiate a strings slice to return messages.
		var flashes []string
		for _, fl := range fm {
			// Add message to the slice.
			flashes = append(flashes, fl.(string))
		}
		return flashes
	}
	return nil
}
