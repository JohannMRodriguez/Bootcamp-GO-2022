package main

import "fmt"

type shape interface {
	area() float32
}

type rectangle struct {
	width, height float32
}

type circle struct {
	radius float32
}

func (r rectangle) area() float32 {
	return r.height * r.width
}

func (c *circle) area() float32 {
	return c.radius * c.radius * 3.14
}

func main() {

	// rect1 := rectangle{3, 5}
	// circle1 := circle{3.25}

	// fmt.Printf("Area del rectangulo %.2f", rect1.area())
	// fmt.Println()
	// fmt.Printf("Area del circulo %.2f", circle1.area())

	// shapes := []shape{rectangle{3, 5}}
	// fmt.Printf("Area del rectangulo %.2f", shapes[0].area())
	// fmt.Println()
	// fmt.Printf("Area del circulo %.2f", shapes[1].area())
	shape := rectangle{3, 5}
	fmt.Println("Area del ractangulo: ", shape.area())
}
