package app

import (
	"minicommerce/models"
)

func DBMigrate(){
	DB.AutoMigrate(&models.Admin{}, &models.Produk{})
}