package controllers

import (
	"fmt"
	"minicommerce/models"
	"minicommerce/services"
	"net/http"

	"github.com/kataras/go-sessions/v3"
	"github.com/labstack/echo/v4"
)

var adminServices = services.NewAdminServices()

func Register(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "daftar.html", nil)
	} else if c.Request().Method == http.MethodPost {
		input := new(models.InputAdmin)
		c.Bind(input)
		input.Nama = c.FormValue("name")
		input.Email = c.FormValue("email")
		input.Password = c.FormValue("password")
		data := adminServices.Register(*input)
		response := map[string]interface{}{
			"response": data,
		}
		return c.Render(http.StatusOK, "login.html", response)
	} else {
		return c.Render(http.StatusBadRequest, "login.html", "Method Tidak Ditemukan")
	}
}

func Login(c echo.Context) error {
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "login.html", nil)
	} else if c.Request().Method == http.MethodPost {
		input := new(models.InputAdmin)
		c.Bind(input)
		input.Email = c.FormValue("email")
		input.Password = c.FormValue("password")
		data, token := adminServices.Login(*input)
		if data.ID == 0 {
			data := map[string]interface{}{
				"response": "Login gagal",
			}
			return c.Render(http.StatusNotFound, "daftar.html", data)
		}
		session := sessions.Start(c.Response(), c.Request())
		id := fmt.Sprintf("%d", data.ID)
		session.Set("token", token)
		session.Set("id", id)
		return c.Redirect(http.StatusFound, "/dashboard")
	} else {
		return c.Render(http.StatusBadRequest, "login.html", "Method Tidak Ditemukan")
	}

}
