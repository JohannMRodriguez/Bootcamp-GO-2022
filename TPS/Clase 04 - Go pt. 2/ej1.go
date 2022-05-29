package main

import "fmt"

func promedio(sf ...int) {
	total := 0
	for _, num := range sf {
		total += num
	}
	fmt.Println("Promedio es: ", total/len(sf))

}

func main() {

	// logear solo impares del slice
	s := make([]int, 10)

	i := 0
	for i < 10 {
		s[i] = i
		i = i + 1
	}

	i = 0
	for i < len(s) {
		if s[i]%2 != 0 {
			fmt.Println(s[i])
		}
		i += 1
	}

	// string de [3:7) a partir del slice anterior
	s2 := s[3:7]
	fmt.Println(s2)

	// obtener promedio de slice con funcion
	promedio(s...)
}
