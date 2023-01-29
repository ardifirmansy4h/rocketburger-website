package models

import "time"

type Produk struct {
	ID        int
	Nama      string
	Kategori  string
	Deskripsi string
	Harga     string
	Foto      string
	CreatedAt time.Time
}

type InputProduk struct {
	Nama      string
	Kategori  string
	Deskripsi string
	Harga     string
	Foto      string
}