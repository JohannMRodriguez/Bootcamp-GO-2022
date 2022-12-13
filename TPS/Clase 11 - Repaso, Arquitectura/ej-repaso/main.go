package main

import (
	"github.com/gin-gonic/gin"
	"repaso.com/controller"
)

func main() {

	server := gin.Default()
	server.POST("/getCursos", controller.GetCursos)
	server.Run(":8081")
}
