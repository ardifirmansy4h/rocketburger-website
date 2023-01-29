package repository

import (
	"minicommerce/app"
	"minicommerce/models"
)

type ProdukRepoIMPL struct{}

func (pr *ProdukRepoIMPL) GetAllProduk() []models.Produk {
	var produk []models.Produk
	app.DB.Find(&produk)
	return produk
}

func (pr *ProdukRepoIMPL) GetByIDProduk(id int) models.Produk {
	var produk models.Produk
	app.DB.First(&produk, "id = ?", id)
	return produk
}

func (pr *ProdukRepoIMPL) AddProduk(i models.InputProduk) string {
	var produk models.Produk = models.Produk{
		Nama:      i.Nama,
		Kategori:  i.Kategori,
		Deskripsi: i.Deskripsi,
		Harga:     i.Harga,
		Foto:      i.Foto,
	}
	app.DB.Create(&produk)

	return "Data berhasil ditambah"
}

func (pr *ProdukRepoIMPL) EditProduk(id int, i models.InputProduk) string {
	var produk models.Produk
	app.DB.First(&produk, "id = ?", id)
	produk.Nama = i.Nama
	produk.Kategori = i.Kategori
	produk.Deskripsi = i.Deskripsi
	produk.Harga = i.Harga
	produk.Foto = i.Foto

	app.DB.Save(&produk)

	return "Data berhasil diubah"
}

func (pr *ProdukRepoIMPL) DeleteProduk(id int) string {
	var produk models.Produk
	app.DB.First(&produk, "id=?", id)

	res := app.DB.Delete(produk)
	if res.RowsAffected == 0 {
		return "Data tidak ada"
	}

	return "Data berhasil dihapus"

}
