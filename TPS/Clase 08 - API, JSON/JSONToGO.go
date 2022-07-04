package main

import (
	"encoding/json"
	"fmt"
)

type datos struct {
	Nombre       string
	Edad         int
	Nacionalidad string
}

func main() {
	alumno := datos{
		Nombre:       "Johann",
		Edad:         21,
		Nacionalidad: "Brasilera",
	}

	jsonData, err := json.Marshal(alumno)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonData))
	}
}
