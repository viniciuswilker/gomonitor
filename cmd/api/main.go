package main

import (
	"gomonitor/database"
	"gomonitor/handlers"
	"gomonitor/monitor"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Conectar()

	monitor.Iniciar()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	api := r.Group("/api")
	{

		api.POST("/servicos", handlers.CriarServico)
		api.POST("/servicos/:id/checar", handlers.TestarServico)
		api.GET("/servicos", handlers.ListarServicos)
		api.DELETE("/servicos/:id", handlers.DeletarServico)
	}

	r.POST("/servicos", handlers.CriarServico)
	r.GET("/servicos", handlers.ListarServicos)

	porta := os.Getenv("PORT")
	if porta == "" {
		porta = "8080"
	}

	r.Run(":" + porta)
}
