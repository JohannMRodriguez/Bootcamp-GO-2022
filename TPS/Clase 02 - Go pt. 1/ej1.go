package main

import "fmt"

func main() {

	nombre, edad := "Johann", 22

	fmt.Printf(nombre + " ")
	fmt.Printf("%d\n", edad)

	fmt.Printf("\n" + nombre + " ")
	fmt.Println(edad)
}
