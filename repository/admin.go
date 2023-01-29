package repository

import (
	"minicommerce/app"
	"minicommerce/models"
)

type AdminRepoIMPL struct{}

func (ar *AdminRepoIMPL) Register(i models.InputAdmin) string {
	var all []models.Admin
	app.DB.Find(&all)
	for _, data := range all {
		if data.Email == i.Email {
			return "Email sudah terdaftar"
		}
	}
	var admin models.Admin = models.Admin{
		Nama:     i.Nama,
		Email:    i.Email,
		Password: i.Password,
	}
	app.DB.Create(&admin)
	return "berhasil daftar"
}

func (ar *AdminRepoIMPL) Login(i models.InputAdmin) (models.Admin, string) {
	var admin models.Admin
	app.DB.Where("email = ?", i.Email).Find(&admin)
	if admin.ID == 0 {
		return models.Admin{}, ""
	}

	return admin, "ini adalah token"

}
