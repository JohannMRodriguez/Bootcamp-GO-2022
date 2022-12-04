package interfaces

import (
	"github.com/gin-gonic/gin"
	"repaso.com/models"
)

type ServiceInterface interface {
	Saludar() string
	GetCursos(models.Cursos) (models.Cursos, gin.H)
}
