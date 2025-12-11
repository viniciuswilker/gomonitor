package monitor

import (
	"gomonitor/database"
	"gomonitor/models"
	"log"
	"net/http"
	"time"
)

func Iniciar() {
	go func() {
		for {
			log.Println("Rodando ciclo de verificação")
			verificarServicos()

			time.Sleep(10 * time.Second)
		}
	}()
}

func verificarServicos() {

	var servicos []models.Servico

	database.DB.Find(&servicos)

	for _, servico := range servicos {

		go testarConectividade(servico)

	}

}

func testarConectividade(s models.Servico) {

	cliente := http.Client{
		Timeout: 10 * time.Second,
	}

	inicio := time.Now()

	resp, erro := cliente.Get(s.URL)

	latencia := time.Since(inicio).Milliseconds()
	novoStatus := "DOWN"

	if erro == nil && resp.StatusCode >= 200 && resp.StatusCode < 400 {
		novoStatus = "UP"
		defer resp.Body.Close()
	}

	atualizarStatusNoBanco(s, novoStatus, latencia)

}

func atualizarStatusNoBanco(s models.Servico, status string, latencia int64) {

	s.Status = status
	s.UltimaChecagem = time.Now()

	database.DB.Save(&s)

	logServico := models.LogServico{
		ServicoID:    s.ID,
		Status:       status,
		Latencia:     latencia,
		DataChecagem: time.Now(),
	}

	database.DB.Create(&logServico)

}

func TestarUnico(id uint) (*models.Servico, error) {
	var s models.Servico

	if err := database.DB.First(&s, id).Error; err != nil {
		return nil, err
	}

	testarConectividade(s)

	database.DB.First(&s, id)
	return &s, nil

}

func DeletarUnico(id uint) (*models.Servico, error) {
	var s models.Servico

	if err := database.DB.First(&s, id).Error; err != nil {
		return nil, err
	}

	if err := database.DB.Delete(&s).Error; err != nil {
		return nil, err
	}

	return &s, nil
}
