package main

import "fmt"

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

func (curso Curso) altaAlumno() bool {

	var nombre, id string
	var notas []float32
	var faltas int
	var nota float32

	fmt.Println("Ingrese los datos para el nuevo alumno...")
	fmt.Print("Nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("ID: ")
	fmt.Scanln(&id)
	fmt.Print("Notas: ")
	notas = append(notas, fmt.Scan(&nota))
	notas = append(notas, fmt.Scan(&nota))
	notas = append(notas, fmt.Scan(&nota))
	fmt.Print("Cantidad de faltas: ")
	fmt.Scanln(&faltas)

	curso.Curso.Alumno.Alumno = append(curso.Alumno.[]Alumno, id, nombre, notas, faltas)
	return true
}

// func (curso Curso) bajaAlumno(IDAlumno int) bool {
// 	curso.Curso.IDAlumno
// }

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
				Faltas: 9,
			}, {
				ID:     "001-3",
				Nombre: "Valeria",
				Notas:  []float32{8, 7, 8, 9},
				Faltas: 0,
			},
		},
		Profesor: Profesor{
			ID:         "100-M",
			Nombre:     "Mauricio",
			Antiguedad: 8,
		},
	}
	var opcion, IDAlumno int

	for {
		fmt.Println("1 - Ver alumnos.")
		fmt.Println("2 - Dar ALTA a alumno.")
		fmt.Println("3 - Ver BAJA a alumno.")
		fmt.Println("4 - Modificar alumno.")
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
			// curso1.bajaAlumno(IDAlumno)
			// case 4:
			// 	modificarAlumno()
		}
	}
}
