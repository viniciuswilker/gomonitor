package models

import (
	"time"

	"gorm.io/gorm"
)

type Servico struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	Nome           string         `json:"nome"`
	URL            string         `json:"url"`
	Intervalo      int            `json:"intervalo"`
	Status         string         `json:"status"`
	UltimaChecagem time.Time      `json:"ultima_checagem"`
	CriadoEm       time.Time      `json:"criado_em"`
	AtualizadoEm   time.Time      `json:"atualizado_em"`
	DeletadoEm     gorm.DeletedAt `gorm:"index" json:"-"`
}

type LogServico struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ServicoID    uint      `json:"servico_id"`
	Status       string    `json:"status"`
	Latencia     int64     `json:"latencia_ms"`
	DataChecagem time.Time `json:"data_checagem"`
}
