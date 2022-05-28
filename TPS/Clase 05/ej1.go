package main

import "fmt"

func multiplicar(num1 float32, num2 float32) float32 {
	return num1 * num2
}

func main() {
	var num1 float32 = 9
	var num2 float32 = 5

	resultadoMultiplicacion := multiplicar(num1, num2)
	fmt.Println("El resultado de", num1, "x", num2, "es", resultadoMultiplicacion)
}
