package main

import "fmt"

type Nombre struct {
	nombre string
	edad   int
}

type Nombres struct {
	Nombre []Nombre
}

func RemoveIndex(s []Nombre, index int) []Nombre {
	return append(s[:index], s[index+1:]...)
}

func main() {

	m := make(map[string]int)

	m["juan"] = 2
	fmt.Print(m)
}
