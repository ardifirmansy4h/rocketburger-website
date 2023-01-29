package app


import (
	helper "minicommerce/helper"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	var dtb string = fmt.Sprintf("%s:%s@/%s?charset=utf8mb4&parseTime=True&loc=Local",
		helper.GetConfig("DB_USERNAME"),
		helper.GetConfig("DB_PASSWORD"),
		helper.GetConfig("DB_NAME"),
	)
	DB, err = gorm.Open(mysql.Open(dtb), &gorm.Config{})

	if err != nil {
		fmt.Println("Database Tidak Terkoneksi")
	}
	DBMigrate()
	fmt.Println("Database Terkoneksi")

}