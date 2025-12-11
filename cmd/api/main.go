package main

import (
	"gomonitor/database"
	"gomonitor/handlers"
	"gomonitor/monitor"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func main() {

// 	database.Conectar()

// 	r := gin.Default()

// 	r.Use(func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}
// 		c.Next()
// 	})

// 	r.POST("/servicos", handlers.CriarServico)
// 	r.GET("/servicos", handlers.ListarServicos)

// 	r.Run(":8080")

// }

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

	r.Run(":8080")
}
