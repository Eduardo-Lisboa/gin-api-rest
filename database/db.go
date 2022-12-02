package database

import (
	"log"

	"github.com/Eduardo-Lisboa/api-go-gin/config"
	"github.com/Eduardo-Lisboa/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectDatabase() {
	dns := config.DbHost + "" + config.DbUser + "" + config.DbPassword + "" + config.DBName + "" + config.PortDatabase + "" + config.SSLMode
	DB, err = gorm.Open(postgres.Open(dns))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}
	DB.AutoMigrate(&models.Product{})
}
