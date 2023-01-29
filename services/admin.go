package services

import (
	"minicommerce/models"
	"minicommerce/repository"
)

type AdminServices struct {
	adminRepo repository.AdminRepo
}

func NewAdminServices() AdminServices{
	return AdminServices{
		adminRepo: &repository.AdminRepoIMPL{},
	}
}

func (as *AdminServices) Register(i models.InputAdmin)string{
	return as.adminRepo.Register(i)
}

func (as *AdminServices) Login(i models.InputAdmin)(models.Admin, string){
	return as.adminRepo.Login(i)
}