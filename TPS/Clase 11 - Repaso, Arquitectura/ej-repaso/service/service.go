package service

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"repaso.com/errors"
	"repaso.com/interfaces"
	"repaso.com/models"
)

type Service struct {
	service interfaces.ServiceInterface
}

func checkOrderLevel(order int, cursos []models.CursoResponse) int {

	for _, each := range(cursos) {

		if each.Order == order + 1 {

			return order + 2
		}
	}
	return order + 1
}

func toCursoResponse(toConvert map[string]interface{}) (models.CursoResponse, error) {

	// - primero pasamos de map a json
	jsonbody, err := json.Marshal(toConvert)
	if err != nil {
		return models.CursoResponse{}, err
	}
	// - y luego de json a struct
	cursoResponse := models.CursoResponse{}
	if err := json.Unmarshal(jsonbody, &cursoResponse); err != nil {
		return models.CursoResponse{}, err
	}
	return cursoResponse, nil
}

func validateIfExist(curso string, listaCursos []string) bool {

	for _, each := range listaCursos {

		if each == curso {
			return true
		}
	}
	return false

}

func (s *Service) GetCursos(data models.Cursos) (models.CursosResponse, gin.H) {

	var cursosListDesired, cursosListRequired, cursosListAdd []string //almacenan los cursos desired y required
	var response models.CursosResponse                                //la respuesta al servicio
	var cursosResponse []models.CursoResponse                         //almacena el campo Course para la respuesta
	cursoMap := make(map[string]interface{})                          //almacena temporalmente un mapa con 'curso':orden

	//TODO validaciones en la request
	if data.UserId == "" || len(data.Courses) == 0 {

		return models.CursosResponse{}, gin.H{"error": errors.ERROR_REQUIRED_PARAMETER_NOT_FOUND}
	}

	//TODO validado el body, trabajamos con la info pasada
	// obtener la request y almacenar desiredCourse en un array y requiredCourse en otro
	for _, curso := range data.Courses {
		cursosListDesired = append(cursosListDesired, curso.DesiredCourse)
		cursosListRequired = append(cursosListRequired, curso.RequiredCourse)
	}

	//TODO agregamos los cursos por nivel al mapa temporal, que luego se usa para pasar a una instancia de la struct
	// validar si hay un curso de menor nivel (orden 0), que no es deseado, o q es deseado sin requerido
	for index, curso := range cursosListRequired {

		if !validateIfExist(curso, cursosListDesired) {

			if curso == "" {

				cursoMap["course"] = cursosListDesired[index]
				cursoMap["order"] = 0

			} else {

				cursoMap["course"] = curso
				cursoMap["order"] = 0

			}

			// convertir el mapa en una struct y agregar a la lista
			convert, err := toCursoResponse(cursoMap)
			if err != nil {
				return models.CursosResponse{}, gin.H{"error": errors.ERROR_UNPROCESSABLE_ENTITY}
			}
			cursosResponse = append(cursosResponse, convert)
			cursosListAdd = append(cursosListAdd, convert.Course)
		}
	}

	// buscamos los cursos de niveles intermediarios (ni 0 ni el mayor)
	for index := 0; index < len(cursosListDesired); index++ {
	// ! problema encontrado -> al agregar curso, vuelve al index 0, entonces, no respeta el orden
	// ! el primero que encuentra ya le agrega a la lista con orden + 1. 
	// ! solucion -> controlar que no hayan + de 1 curso deseado con un mismo curso requerido

		if validateIfExist(cursosListRequired[index], cursosListAdd) && !validateIfExist(cursosListDesired[index], cursosListAdd){

			cursoMap["course"] = cursosListDesired[index]

			for _, newCurso := range cursosResponse {

				if newCurso.Course == cursosListRequired[index] {

					order := checkOrderLevel(newCurso.Order, cursosResponse)

					cursoMap["order"] = order

					// convertir el mapa en una struct y agregar a la lista
					convert, err := toCursoResponse(cursoMap)
					if err != nil {
						return models.CursosResponse{}, gin.H{"error": errors.ERROR_SORTING_COURSES}
					}
					cursosResponse = append(cursosResponse, convert)
					cursosListAdd = append(cursosListAdd, convert.Course)
				}
			}
			index = 0
		}
	}

	// agregamos los cursos de mayor nivel, aquellos q no son requeridos
	for index, curso := range cursosListDesired {

		if !validateIfExist(curso, cursosListRequired) && !validateIfExist(curso, cursosListAdd) {

			cursoMap["course"] = curso

			for _, newCurso := range cursosResponse {

				if newCurso.Course == cursosListRequired[index] {

					order := checkOrderLevel(newCurso.Order, cursosResponse)

					cursoMap["order"] = order

					// convertir el mapa en una struct y agregar a la lista
					convert, err := toCursoResponse(cursoMap)
					if err != nil {
						return models.CursosResponse{}, gin.H{"error": errors.ERROR_UNPROCESSABLE_ENTITY}
					}
					cursosResponse = append(cursosResponse, convert)
					cursosListAdd = append(cursosListAdd, convert.Course)
				}
			}
		}
	}

	response.UserId = data.UserId
	response.Courses = cursosResponse

	return response, nil
}

func (s *Service) Saludar() string {

	return "Bien Venido al sitio web!"
}
