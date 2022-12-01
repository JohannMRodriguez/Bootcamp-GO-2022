package main

import (
	"github.com/gin-gonic/gin"
	"repaso.com/controller"
)

func main() {

	server := gin.Default()
	server.GET("/saludar", controller.Saludar)
	server.POST("/getCursos", controller.GetCursos)
	server.Run(":8080")
}
