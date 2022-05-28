package main

import "fmt"

type Direccion struct {
	Ciudad       string
	CodigoPostal int
}

type DNI struct {
	Nombre, Apellido, Historial, numeroCuenta string
	Direccion                                 Direccion
}

func main() {

	var dni DNI

	dni = DNI{
		Nombre:   "Johann",
		Apellido: "Rodriguez",
		Direccion: Direccion{
			Ciudad:       "Obera",
			CodigoPostal: 3360,
		},
		Historial:    "-",
		numeroCuenta: "00012084347547854856",
	}

	fmt.Println("\nDatos ingresados...")
	fmt.Println(dni.Nombre, dni.Apellido)
	fmt.Println(dni.Direccion)
	fmt.Println(dni.Historial)
	fmt.Println(dni.numeroCuenta)
}
