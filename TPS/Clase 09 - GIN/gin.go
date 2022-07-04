package main

import "github.com/gin-gonic/gin"

var alumnos = map[string]string{
	"101": "Johann",
	"102": "Micaela",
	"202": "Cristopher",
}

func main() {
	server := gin.Default()
	server.GET("/", PaginaPrincipal)
	server.GET("/alumnos/:id", BuscarAlumno)
	server.GET("/mostrar/", MostrarAlumnos)
	server.GET("/registrar/:nuevoid/:nombre", RegistrarNuevoAlumno)
	server.Run(":8085")
}

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "Bienvenido al Bootcamp!")
}

func BuscarAlumno(ctx *gin.Context) {
	alumno, ok := alumnos[ctx.Param("id")]
	if ok {
		ctx.String(200, "Informacion del alumno %s, nombre %s", ctx.Param("id"), alumno)
	} else {
		ctx.String(600, "Informacion del alumno no existe!")
	}
}

func RegistrarNuevoAlumno(ctx *gin.Context) {
	_, ok := alumnos[ctx.Param("nuevoid")]
	if ok {
		ctx.String(404, "El alumno ya existe!")
	} else {
		alumnos[ctx.Param("nuevoid")] = ctx.Param("nombre")
		ctx.String(200, "Informacion del alumno %s, %s registrada!", ctx.Param("nuevoid"), ctx.Param("nombre"))
	}
}

func MostrarAlumnos(ctx *gin.Context) {
	if len(alumnos) == 0 {
		ctx.String(200, "No existen alumnos registrados!")
	} else {
		var display string
		for k, v := range alumnos {
			display += k + " - " + v + "\n"
		}
		ctx.String(200, "Lista de alumnos registrados...\n"+display)
	}
}
