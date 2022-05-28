package main

import "fmt"

func multiplicar(num1 float32, num2 float32) float32 {
	return num1 * num2
}

func encontrarMayor(enteros ...int) int {
	mayor := enteros[0]
	for _, num := range enteros {
		if num > mayor {
			mayor = num
		}
	}
	return mayor
}

type Animales struct {
	sexo  string
	color string
	edad  int
}

type Animal struct {
	especie string
	datos   Animales
}

func (animal Animal) addAnimal(addEspecie, addSexo, addColor string, addEdad int) *Animal {
	animal = Animal{
		especie: addEspecie,
		datos: Animales{
			sexo:  addSexo,
			color: addColor,
			edad:  addEdad,
		},
	}
	pointerToAnimal := &animal
	return pointerToAnimal //retorna la direccion de la struct instanciada por medio de un puntero
}

func factorial(numero int) int {
	if numero <= 1 {
		return 1
	} else {
		return numero * factorial(numero-1)
	}
}

func main() {
	// punto A
	var num1 float32 = 9
	var num2 float32 = 5

	resultadoMultiplicacion := multiplicar(num1, num2)
	fmt.Println("El resultado de", num1, "x", num2, "es", resultadoMultiplicacion)

	// punto B
	enteros := []int{1, 2, 3, 4, 2, 21, 209, 554, 5, 6, 7, 8, 9}
	mayor := encontrarMayor(enteros...)
	fmt.Println("El mayor numero entre los numeros", enteros, "es", mayor)

	// punto C
	animal1 := Animal{}
	punteroAnimal1 := animal1.addAnimal("Perro", "Macho", "Negro", 2) //puntero recibe la direccion de la estructura
	fmt.Println(*punteroAnimal1)

	// punto D
	num3 := 5
	fmt.Println("El factorial de", num3, "es", factorial(num3))
}
