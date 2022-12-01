package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"repaso.com/service"
)

func GetCursos( ctx *gin.Context ) {
	
	response, err := (&service.Service{}).GetCursos( ctx )

	if err != nil {

		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, response)

}

func Saludar( ctx *gin.Context ) {
	
	response := (&service.Service{}).Saludar( ctx )
	ctx.IndentedJSON(http.StatusOK, response)
}
