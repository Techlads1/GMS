package webserver

import (
	"html/template"
	"time"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/echoview-v4"
)

//Renderer fetches the template render
func Renderer() *echoview.ViewEngine {
	gvc := goview.Config{
		
		Root:      "webserver/systems",
		Extension: ".html",
		Master:    "layouts/master",
		
		Funcs: template.FuncMap{
			"sub": func(a, b int) int {
				return a - b
			},
			"add": func(a, b int) int {
				return a + b
			},
			"inc": func(i int) int {
				return i + 1
			},
			"copy": func() string {
				return time.Now().Format("2006")
			},
		},
		DisableCache: true,
	}
	return echoview.New(gvc)
}
