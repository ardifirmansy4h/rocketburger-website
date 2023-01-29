package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

func GetKontak(c echo.Context) error {
	return c.Render(http.StatusOK, "lokasi.html", nil)
}
func GetDetail(c echo.Context) error {
	queryString := c.Request().URL.Query()
	id, _ := strconv.Atoi(queryString.Get("id"))
	produk := produkServices.GetByIDProduk(id)
	data := map[string]interface{}{
		"produk": produk,
	}
	return c.Render(http.StatusFound, "detail.html", data)

}
func GetMenu(c echo.Context) error {
	produk := produkServices.GetAllProduk()
	data := map[string]interface{}{
		"produk": produk,
	}
	return c.Render(http.StatusFound, "menu.html", data)
}
