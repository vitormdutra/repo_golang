package database

import (
	"gin-api-rest/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectarComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=123 dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic(err)
	}
	DB.AutoMigrate(&models.Aluno{})
}
