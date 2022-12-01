package interfaces

import (
	"github.com/gin-gonic/gin"
	"repaso.com/models"
)

type ServiceInterface interface {
	
	Saludar(ctx *gin.Context) string
	GetCursos(ctx *gin.Context) (models.Cursos, gin.H)
}
