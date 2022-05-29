package main

import "fmt"

type Datos struct {
	Temperaturas     []float32
	Humedad          []float32
	CantidadFrutas   map[string]int
	Nombre           string
	CantidadHectares float32
}

func (chacra *Datos) addData(newTemperatura float32, newHumedad float32) {
	chacra.Temperaturas = append(chacra.Temperaturas, newTemperatura)
	chacra.Humedad = append(chacra.Humedad, newHumedad)
}

func (chacra Datos) logInfo() {
	fmt.Println()
	fmt.Println("Nombre:", chacra.Nombre)
	fmt.Println("Temperaturas:", chacra.Temperaturas)
	fmt.Println("Humedad:", chacra.Humedad)
	fmt.Println("Cantidad Frutas:", chacra.CantidadFrutas)
	fmt.Println("Cantidad Hectares:", chacra.CantidadHectares)
}

func (chacra Datos) Promedio() (float32, float32) {
	var totalTemperaturas float32 = 0
	var totalHumedad float32 = 0
	for _, value := range chacra.Temperaturas {
		totalTemperaturas += value
	}
	for _, value := range chacra.Humedad {
		totalHumedad += value
	}

	return totalTemperaturas / float32(len(chacra.Temperaturas)), totalHumedad / float32(len(chacra.Humedad))
}

func (chacra Datos) checkIfExist(nombreFruta string) int {
	for key := range chacra.CantidadFrutas {
		if key == nombreFruta {
			return chacra.CantidadFrutas[key]
		}
	}
	return -1
}

func main() {
	chacra1 := Datos{
		[]float32{10, 10, 10, 10, 10, 10},
		[]float32{54, 54.6, 57, 80, 90, 12, 45},
		map[string]int{
			"Manzana": 200,
			"Pera":    200,
			"Ciruela": 500,
		},
		"Rodolpho Asis",
		3000}
	chacra2 := Datos{
		[]float32{22, 34, 29, 28, 31, 32, 25},
		[]float32{78, 89, 45.6, 53, 34, 76},
		map[string]int{
			"Manzana": 500,
			"Cereza":  300,
			"Durazno": 300,
		},
		"Juan Centurion",
		3600}

	chacra1.logInfo()
	chacra2.logInfo()

	var newTemperatura float32 = 23.4
	var newHumedad float32 = 67

	chacra2.addData(newTemperatura, newHumedad)

	PromedioTemperaturas, PromedioHumedades := chacra1.Promedio()
	fmt.Println()
	fmt.Println("Valores para chacra de", chacra1.Nombre)
	fmt.Println("Promedio Temperaturas: ", PromedioTemperaturas)
	fmt.Println("Promedio Humedades:", PromedioHumedades)

	frutaToCheck := "Manzana"
	checkFruta := chacra2.checkIfExist(frutaToCheck)
	if checkFruta == -1 {
		fmt.Println("La fruta NO existe!")
	} else {
		fmt.Println("Existen", checkFruta, frutaToCheck, "plantadas en la chacra!")
	}
}
