package main

import (
	"html/template"
	"io"
	"minicommerce/app"
	"minicommerce/app/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	app.ConnectDB()
	t := &Template{
		templates: template.Must(template.ParseGlob("views/*/*.html")),
	}
	server := echo.New()
	server.Use(middleware.Static("views"))
	server.Renderer = t
	routes.Routes(server)
	server.Logger.Fatal(server.Start(":2222"))
}
