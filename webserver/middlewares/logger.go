package middlewares

import (
	"gateway/package/config"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Logger Middleware
func Logger(debug bool) echo.MiddlewareFunc {
	//TODO(jwm) ensure this is logged in the correct log director

	path := config.LoggerPath()
	out, err := os.Create(path + "/requests.log")

	if err != nil || debug {
		out = os.Stdout
	}

	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "Method=${method}, Url=\"${uri}\", Status=${status}, Latency:${latency_human} \n",
		Output: out,
	})
}
