package main

import (
	"fmt"
)

type Alumno struct {
	ID, Nombre string
	Notas      []float32
	Faltas     int
}

type Profesor struct {
	ID, Nombre string
	Antiguedad int
}

type Curso struct {
	ID                     string
	Nombre                 string
	Aula                   int
	CantidadMaxEstudiantes int
	Alumno                 []Alumno
	Profesor               Profesor
}

func (curso Curso) logAlumnos() {
	for _, each := range curso.Alumno {
		fmt.Print(each.Nombre, " - ")
		fmt.Print(each.ID, " - ")
		fmt.Print(each.Notas, " - ")
		fmt.Print(each.Faltas)
		fmt.Println()
	}
	fmt.Println()
}

func (curso *Curso) altaAlumno() bool {

	var nombre, id string
	var notas []float32
	var faltas int
	var nota float32

	fmt.Println("Ingrese los datos para el nuevo alumno...")
	fmt.Print("Nombre: ")
	fmt.Scan(&nombre)
	fmt.Print("ID: ")
	fmt.Scan(&id)
	fmt.Print("Notas: ")
	fmt.Scan(&nota)
	notas = append(notas, nota)
	fmt.Scan(&nota)
	notas = append(notas, nota)
	fmt.Scan(&nota)
	notas = append(notas, nota)
	fmt.Print("Cantidad de faltas: ")
	fmt.Scan(&faltas)

	curso.Alumno = append(curso.Alumno, Alumno{id, nombre, notas, faltas})
	return true
}

func RemoveIndex(s []Alumno, index int) []Alumno {
	return append(s[:index], s[index+1:]...)
}

func (curso *Curso) bajaAlumno(IDAlumno string) bool {

	cont := 0
	for _, each := range curso.Alumno {
		if each.ID == IDAlumno {
			curso.Alumno = RemoveIndex(curso.Alumno, cont)
			return true
		}
		cont += 1
	}
	return false
}

func (curso *Curso) modificarAlumno(IDAlumno string) {

	var nombre, id string
	var faltas int
	var addNota float32
	var notas []float32

	fmt.Println("Ingrese las modificaciones...")
	fmt.Print("Nombre:")
	fmt.Scan(&nombre)
	fmt.Print("ID: ")
	fmt.Scan(&id)
	fmt.Print("Agregar nota: ")
	fmt.Scan(&addNota)
	fmt.Print("Cantidad de faltas: ")
	fmt.Scan(&faltas)

	cont := 0
	for _, each := range curso.Alumno {

		if each.ID == IDAlumno {
			notas = each.Notas
			fmt.Print(each)
			if addNota != -1 {
				notas = append(notas, addNota)
			}

			curso.Alumno = RemoveIndex(curso.Alumno, cont)
			curso.Alumno = append(curso.Alumno, Alumno{id, nombre, notas, faltas})

			fmt.Println("El alumno de ID", each.ID, "fue modificado!")
			return
		}
		cont += 1
	}
}

func (curso Curso) verPromedioBajo() {

	var promedio float32
	promedios := make(map[string]float32)

	for _, each := range curso.Alumno {

		promedio = 0
		for _, nota := range each.Notas {
			promedio += nota
		}
		promedio = promedio / float32(len(each.Notas))
		if promedio < 6 {
			promedios[each.Nombre] = promedio
		}
	}

	if len(promedios) == 0 {
		fmt.Println("No existen alumnos con promedio inferior a 6!")
	} else {
		for key, element := range promedios {
			fmt.Println("Alumno:", key, "- Promedio:", element)
		}
	}
}

func (curso Curso) verLibres() {
	for _, each := range curso.Alumno {
		if each.Faltas > 3 {
			fmt.Println("Alumno:", each.Nombre, " - Cantidad de faltas: ", each.Faltas)
		}
	}
}

func (curso Curso) verMejoresPromedios() {

}

func main() {

	curso1 := Curso{
		ID:                     "001",
		Nombre:                 "Curso 1",
		Aula:                   1,
		CantidadMaxEstudiantes: 30,
		Alumno: []Alumno{
			{
				ID:     "001-1",
				Nombre: "Juan",
				Notas:  []float32{8, 2, 10, 6},
				Faltas: 2,
			}, {
				ID:     "001-2",
				Nombre: "Marcos",
				Notas:  []float32{5, 6, 7, 4},
				Faltas: 2,
			}, {
				ID:     "001-3",
				Nombre: "Valeria",
				Notas:  []float32{3, 2, 2, 1},
				Faltas: 7,
			}, {
				ID:     "001-4",
				Nombre: "Florencia",
				Notas:  []float32{8, 8, 9, 10},
				Faltas: 0,
			}, {
				ID:     "001-5",
				Nombre: "Esteban",
				Notas:  []float32{3, 10, 6, 1},
				Faltas: 4,
			},
		},
		Profesor: Profesor{
			ID:         "100-M",
			Nombre:     "Mauricio",
			Antiguedad: 8,
		},
	}
	var opcion int
	var IDAlumno string
	flagSwitch := true

	for flagSwitch == true {
		fmt.Println("1 - Ver alumnos.")
		fmt.Println("2 - Dar ALTA a alumno.")
		fmt.Println("3 - Dar BAJA a alumno.")
		fmt.Println("4 - Modificar alumno.")
		fmt.Println("5 - Ver alumnos con promedio inferior a 6.")
		fmt.Println("6 - Ver alumnos en condicion LIBRE.")
		fmt.Println("7 - Ver top 3 alumnos.")
		fmt.Println("8 - Ver capacidad de curso.")
		fmt.Println("0 - Salir.")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			curso1.logAlumnos()
		case 2:
			flag := curso1.altaAlumno()
			if flag == true {
				fmt.Println("El alumno fue dado de alta exitosamente!")
			} else {
				fmt.Println("No fue posible dar el alta al alumno!")
			}
		case 3:
			fmt.Print("Ingrese el ID del alumno: ")
			fmt.Scan(&IDAlumno)
			flag := curso1.bajaAlumno(IDAlumno)
			if flag == true {
				fmt.Println("El alumno de ID", IDAlumno, "fue dado de baja!")
			} else {
				fmt.Println("No existe ningun alumno con ID", IDAlumno)
			}
			fmt.Println()
		case 4:
			fmt.Print("Ingrese el ID del alumno: ")
			fmt.Scan(&IDAlumno)
			curso1.modificarAlumno(IDAlumno)
		case 5:
			curso1.verPromedioBajo()
		case 6:
			curso1.verLibres()
		case 7:
			curso1.verMejoresPromedios()
		case 0:
			flagSwitch = false
		}
	}
}
