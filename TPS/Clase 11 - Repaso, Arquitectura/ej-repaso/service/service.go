package service

import (
	"github.com/gin-gonic/gin"
	"repaso.com/interfaces"
	"repaso.com/messages"
	"repaso.com/models"
)

type Service struct {

	service interfaces.ServiceInterface
}

func ( s *Service ) GetCursos( ctx *gin.Context )  ( models.Cursos, gin.H ) {

	var data models.Cursos

	if err := ctx.ShouldBindJSON(&data); err != nil {
		return models.Cursos{}, gin.H{"error": err.Error()}
	}

	if ( data.UserId == "" || len(data.Courses) == 0 ) {

		return models.Cursos{}, gin.H{"error":messages.ERROR_REQUIRED_PARAMETER_NOT_FOUND}
	}

	return data, nil
}

func ( s *Service ) Saludar( ctx *gin.Context ) string {

	return "Bien Venido al sitio web!"
}
