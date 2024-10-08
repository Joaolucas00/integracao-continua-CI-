package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	nameDB := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	stringDeConexao := "host=" + host + " user=" + usuario + " password=" + password + " dbname=" + nameDB + " port=" + dbPort + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}
