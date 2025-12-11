package handlers

import (
	"gomonitor/database"
	"gomonitor/models"
	"gomonitor/monitor"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CriarServico(c *gin.Context) {

	var novoServico models.Servico

	if erro := c.ShouldBindJSON(&novoServico); erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Dados inválidos " + erro.Error()})
		return
	}

	novoServico.Status = "PENDENTE"

	resultado := database.DB.Create(&novoServico)
	if resultado.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"erro": "erro ao salvar dados"})
		return
	}

	c.JSON(http.StatusCreated, novoServico)

}

func ListarServicos(c *gin.Context) {
	var servicos []models.Servico

	database.DB.Find(&servicos)

	c.JSON(http.StatusOK, servicos)
}

func TestarServico(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	servicoAtualizado, err := monitor.TestarUnico(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Serviço não encontrado"})
		return
	}

	c.JSON(http.StatusOK, servicoAtualizado)

}

func DeletarServico(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "ID inválido"})
		return
	}

	if id == 4 {
		c.JSON(http.StatusForbidden, gin.H{"erro": "Proibido deletar esse serviço"})
		return
	}

	servicoDeletado, err := monitor.DeletarUnico(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"erro": "Serviço não encontrado"})
		return
	}
	c.JSON(http.StatusOK, servicoDeletado)

}
