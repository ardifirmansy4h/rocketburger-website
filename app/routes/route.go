package routes

import (
	"minicommerce/controllers"

	"github.com/labstack/echo/v4"
)

func Routes(server *echo.Echo) {
	server.GET("/index_daftar", controllers.Register)
	server.POST("/daftar", controllers.Register)
	server.GET("/index_login", controllers.Login)
	server.POST("/login", controllers.Login)
	server.GET("index", controllers.Login)
	server.GET("/dashboard", controllers.Dashboard)
	server.GET("/index_produk", controllers.GetAllProduk)
	server.GET("/index_tambah", controllers.AddProduk)
	server.POST("/tambah", controllers.AddProduk)
	server.GET("/delete", controllers.DeleteProduk)
	server.GET("/index_edit", controllers.EditProduk)
	server.POST("/edit", controllers.EditProduk)
	server.GET("/home", controllers.GetHome)
	server.GET("/lokasi", controllers.GetKontak)
	server.GET("/menu", controllers.GetMenu)
	server.GET("/detail", controllers.GetDetail)
	server.GET("/", controllers.GetHome)
}
