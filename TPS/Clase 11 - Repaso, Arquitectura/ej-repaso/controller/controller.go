package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"repaso.com/errors"
	"repaso.com/models"
	"repaso.com/service"
)

// validaciones son hechas en el handler
func GetCursos( ctx *gin.Context ) {

	var data models.Cursos

	if err := ctx.ShouldBindJSON(&data); err != nil {

		ctx.IndentedJSON(http.StatusBadRequest, errors.ERROR_BAD_REQUEST)
	}
	
	response, err := (&service.Service{}).GetCursos( data )

	if err != nil {

		ctx.IndentedJSON(http.StatusBadRequest, err)
	}
	ctx.IndentedJSON(http.StatusOK, response)

}

func Saludar( ctx *gin.Context ) {
	
	response := (&service.Service{}).Saludar()
	ctx.IndentedJSON(http.StatusOK, response)
}
