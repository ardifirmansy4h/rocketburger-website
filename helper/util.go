package helper

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(k string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error Config: %s", err)
	}
	return viper.GetString(k)
}
