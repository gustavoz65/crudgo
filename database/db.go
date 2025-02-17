package database

import (
	"log"

	"github.com/gustavoz65/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

// realizar conexao com o banco
func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=postgres password=postgres dbname=mydb port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar ao com o Banco de Dados")
	}
	DB.AutoMigrate(&models.Aluno{})
}
