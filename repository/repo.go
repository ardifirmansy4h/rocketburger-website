package repository

import "minicommerce/models"

type AdminRepo interface {
	Register(i models.InputAdmin) string
	Login(i models.InputAdmin) (models.Admin, string)
}

type ProdukRepo interface {
	GetAllProduk()[]models.Produk
	GetByIDProduk(id int)models.Produk
	AddProduk(i models.InputProduk) string
	EditProduk(id int, i models.InputProduk) string
	DeleteProduk(id int) string
}