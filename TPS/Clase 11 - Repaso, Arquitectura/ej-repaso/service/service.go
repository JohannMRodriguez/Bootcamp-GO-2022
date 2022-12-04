package service

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"repaso.com/errors"
	"repaso.com/interfaces"
	"repaso.com/models"
)

type Service struct {
	service interfaces.ServiceInterface
}

func (s *Service) GetCursos( data models.Cursos ) (models.CursosResponse, gin.H) {

	var cursosListDesired, cursosListRequired []string
	var response models.CursosResponse					//la respuesta al servicio
	var cursosResponse []models.CursoResponse			//almacena el campo Course para la respuesta
	cursosList := make(map[string]int)					//usado para relacionar los cursos y su nivel

	//validaciones en la request
	if data.UserId == "" || len(data.Courses) == 0 {

		return models.CursosResponse{}, gin.H{"error": errors.ERROR_REQUIRED_PARAMETER_NOT_FOUND}
	}

	//validado el body, trabajamos con la info pasada
	response.UserId = data.UserId

	// obtener la request y almacenar desiredCourse en un array y requiredCourse en otro
	for _, curso := range(data.Courses) {
		cursosListDesired = append(cursosListDesired, curso.DesiredCourse)
		cursosListRequired = append(cursosListRequired, curso.RequiredCourse)
	}

	// lista de cursos deseados inicia el ciclo for de control
	for index := range(cursosListDesired) {

		if (cursosListRequired[index] == "") {
			cursosList[cursosListDesired[index]] = 0
		}
	}

	// convertir el mapa con los cursos y ordenes en una struct del tipo CursoResponse{}
	// - primero pasamos de map a json
	jsonbody, err := json.Marshal(cursosList)
    if err != nil {
		fmt.Println("error ---", err)
        return models.CursosResponse{}, gin.H{"error": errors.ERROR_UNPROCESSABLE_ENTITY}
    }
	// - y luego de json a struct
	cursoResponse := models.CursoResponse{}
	if err := json.Unmarshal(jsonbody, &cursoResponse); err != nil {
		fmt.Println("error ---", err)
        return models.CursosResponse{}, gin.H{"error": errors.ERROR_UNPROCESSABLE_ENTITY}
    }
	cursosResponse = append(cursosResponse, cursoResponse)
	fmt.Println("cursosList - ", cursosList)
	fmt.Println("cursoResponse -", cursoResponse)

	response.Courses = cursosResponse

	return response, nil
}

func (s *Service) Saludar() string {

	return "Bien Venido al sitio web!"
}
