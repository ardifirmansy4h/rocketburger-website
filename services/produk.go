package services

import (
	"minicommerce/models"
	"minicommerce/repository"
)

type ProdukServices struct {
	produkRepo repository.ProdukRepo
}

func NewProdukServices () ProdukServices {
	return ProdukServices{
		produkRepo: &repository.ProdukRepoIMPL{},
	}
}

func (ps *ProdukServices) GetAllProduk()[]models.Produk{
	return ps.produkRepo.GetAllProduk()
}

func (ps *ProdukServices) GetByIDProduk(id int)models.Produk{
	return ps.produkRepo.GetByIDProduk(id)
}

func (ps *ProdukServices) AddProduk(i models.InputProduk)string{
	return ps.produkRepo.AddProduk(i)
}

func (ps *ProdukServices) EditProduk(id int, i models.InputProduk)string{
	return ps.produkRepo.EditProduk(id, i)
}

func (ps *ProdukServices) DeleteProduk(id int) string {
	return ps.produkRepo.DeleteProduk(id)
}