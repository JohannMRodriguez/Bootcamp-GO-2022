package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type datos struct {
	Nombre       string
	Edad         int
	Nacionalidad string
}

func main() {

	jsonData := `{
		"Nombre": "Johann",
		"Edad": 21,
		"Nacionalidad": "Brasilera"
	}`

	var alumno datos

	err := json.Unmarshal([]byte(jsonData), &alumno)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(alumno)
	}
}
