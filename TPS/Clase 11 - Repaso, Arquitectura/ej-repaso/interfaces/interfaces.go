package interfaces

import (
	"github.com/gin-gonic/gin"
	"repaso.com/models"
)

type ServiceInterface interface {
	GetCursos(models.Cursos) (models.Cursos, gin.H)
}
