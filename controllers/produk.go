package controllers

import (
	"fmt"
	"io"

	"minicommerce/app"
	"minicommerce/helper"
	"minicommerce/models"
	"minicommerce/services"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/go-sessions/v3"
	"github.com/labstack/echo/v4"
)

var produkServices = services.NewProdukServices()

func GetAllProduk(c echo.Context) error {
	ss := sessions.Start(c.Response(), c.Request())
	if len(ss.GetString("token")) == 0 {
		data := map[string]interface{}{
			"response": "gagal login",
		}
		return c.Render(http.StatusFound, "login.html", data)
	}
	produk := produkServices.GetAllProduk()
	data := map[string]interface{}{
		"produk": produk,
	}
	return c.Render(http.StatusFound, "produk.html", data)
}
func Dashboard(c echo.Context) error {
	ss := sessions.Start(c.Response(), c.Request())
	if len(ss.GetString("token")) == 0 {
		data := map[string]interface{}{
			"response": "gagal login",
		}
		return c.Render(http.StatusFound, "login.html", data)
	}
	return c.Render(http.StatusFound, "index.html", nil)
}
func GetByIDProduk(c echo.Context) error {
	ss := sessions.Start(c.Response(), c.Request())
	if len(ss.GetString("token")) == 0 {
		data := map[string]interface{}{
			"response": "gagal login",
		}
		return c.Render(http.StatusFound, "login.html", data)
	}
	queryString := c.Request().URL.Query()
	id, _ := strconv.Atoi(queryString.Get("id"))
	produk := produkServices.GetByIDProduk(id)
	data := map[string]interface{}{
		"produk": produk,
	}
	return c.Render(http.StatusFound, "produk-detail.html", data)
}

func AddProduk(c echo.Context) error {
	ss := sessions.Start(c.Response(), c.Request())
	if len(ss.GetString("token")) == 0 {
		data := map[string]interface{}{
			"response": "gagal login",
		}
		return c.Render(http.StatusFound, "login.html", data)
	}
	if c.Request().Method == http.MethodGet {
		return c.Render(http.StatusOK, "tambah-produk.html", nil)
	} else if c.Request().Method == http.MethodPost {
		input := new(models.InputProduk)
		c.Bind(input)
		c.Request().ParseForm()
		nama := c.FormValue("nama")
		kategori := c.FormValue("kategori")
		deskripsi := c.FormValue("deskripsi")
		harga := c.FormValue("harga")
		file, _ := c.FormFile("foto")
		src, _ := file.Open()
		filename := "views/img/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
		dst, _ := os.Create(filename)
		defer dst.Close()
		io.Copy(dst, src)
		basepath := helper.GetConfig("STORAGE")

		image := fmt.Sprintf(`%s/%s`, basepath, filename)

		input.Nama = nama
		input.Kategori = kategori
		input.Deskripsi = deskripsi
		input.Harga = harga
		input.Foto = image

		produk := produkServices.AddProduk(*input)
		data := map[string]interface{}{
			"response": produk,
		}
		return c.Render(http.StatusAccepted, "tambah-produk.html", data)
	} else {
		return c.Render(http.StatusBadRequest, "tambah-produk.html", "Method tak ditemukan")

	}
}

func EditProduk(c echo.Context) error {
	ss := sessions.Start(c.Response(), c.Request())
	if len(ss.GetString("token")) == 0 {
		data := map[string]interface{}{
			"response": "gagal login",
		}
		return c.Render(http.StatusFound, "login.html", data)
	}
	if c.Request().Method == http.MethodGet {
		queryString := c.Request().URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
		produk := produkServices.GetByIDProduk(int(id))

		data := map[string]interface{}{
			"produk": produk,
		}
		return c.Render(http.StatusOK, "edit-produk.html", data)
	} else if c.Request().Method == http.MethodPost {
		id, _ := strconv.ParseInt(c.FormValue("id"), 10, 64)
		fmt.Println(id)
		var produk models.Produk
		app.DB.First(&produk, "id=?", id)
		basepath, _ := os.Getwd()
		imageDelete2 := strings.Replace(produk.Foto, helper.GetConfig("STORAGE"), basepath, 1)
		os.Remove(imageDelete2)

		input := new(models.InputProduk)
		c.Bind(input)
		c.Request().ParseForm()
		file, _ := c.FormFile("foto")
		src, err := file.Open()
		if err != nil {
			return err
		}
		filename := "views/img/" + strconv.FormatInt(time.Now().Unix(), 10) + ".png"
		dst, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			return err
		}
		image := fmt.Sprintf(`%s/%s`, helper.GetConfig("STORAGE"), filename)
		nama := c.FormValue("nama")
		kategori := c.FormValue("kategori")
		description := c.FormValue("deskripsi")
		harga := c.FormValue("harga")

		input.Nama = nama
		input.Kategori = kategori
		input.Deskripsi = description
		input.Harga = harga
		input.Foto = image
		fmt.Println(produk.ID)
		produkServices.EditProduk(int(produk.ID), *input)
		data := map[string]interface{}{
			"success": "berhasil ubah data",
		}
		return c.Render(http.StatusAccepted, "edit-produk.html", data)

	} else {
		return c.Render(http.StatusOK, "edit-produk.html", "method tidak ditemukan")
	}
}

func DeleteProduk(c echo.Context) error {
	queryString := c.Request().URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)
	produk := produkServices.GetByIDProduk(int(id))
	basepath, _ := os.Getwd()
	imageDelete2 := strings.Replace(produk.Foto, basepath, basepath, 1)
	fmt.Println(imageDelete2)
	os.Remove(imageDelete2)
	produkServices.DeleteProduk(int(id))
	return c.Redirect(http.StatusMovedPermanently, "/index_produk")
}
