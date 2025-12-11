package database

import (
	"gomonitor/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Conectar() {
	var erro error

	DB, erro = gorm.Open(sqlite.Open("monitoramento.db"), &gorm.Config{})

	if erro != nil {
		log.Fatal("Erro ao conectar com o banco de dados", erro)
	}

	erro = DB.AutoMigrate(&models.Servico{}, &models.LogServico{})

	if erro != nil {
		log.Fatal("Erro na migração do banco: ", erro)
	}

	log.Println("Banco de dados conectado e tabelas criadas!")

}
