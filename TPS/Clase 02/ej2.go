package main

import "fmt"

func main() {

	var temperatura, humedad, presion float32

	fmt.Printf("Ingrese temperatura: ")
	fmt.Scanln(&temperatura)
	fmt.Printf("Ingrese humedad: ")
	fmt.Scanln(&humedad)
	fmt.Printf("Ingrese presion atmosferica: ")
	fmt.Scanln(&presion)

	fmt.Printf("\n")
	fmt.Printf("Temperatura: ")
	fmt.Println(temperatura)
	fmt.Printf("Humedad: ")
	fmt.Println(humedad)
	fmt.Printf("Presion atmosferica: ")
	fmt.Println(presion)
}
