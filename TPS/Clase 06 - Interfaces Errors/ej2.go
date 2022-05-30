package main

import (
	"errors"
	"fmt"
)

var errorDivideByZero = errors.New("divide by zero")

// interface declaration
type calculator interface {
	getAdition(x float32, y float32) float32
	getSubtract(x float32, y float32) float32
	getMultiply(x float32, y float32) float32
	getDivision(x float32, y float32) float32
}

// struct declaration
type myCalculator struct {
	firstNum, secondNum float32
}

// methods of structs declaration
func (*myCalculator) getMultiply(num1, num2 float32) float32 {
	return num1 * num2
}

func (*myCalculator) getDivision(num1, num2 float32) (float32, error) {
	if num2 == 0 {
		return 0, errorDivideByZero
	} else {
		return num1 / num2, nil
	}
}

func (*myCalculator) getAdition(num1, num2 float32) float32 {
	return num1 + num2
}

func (*myCalculator) getSubtract(num1, num2 float32) float32 {
	return num1 - num2
}

func main() {

	var num1, num2 float32
	var operation string
	flagBreak := false
	flagFirstOperation := false
	result := float32(0)
	calculate := &myCalculator{}

	for flagBreak == false {

		fmt.Println()
		fmt.Println("----------------------")
		fmt.Println("(+)   Addition")
		fmt.Println("(-)   Subtraction")
		fmt.Println("(x)   Multiplication")
		fmt.Println("(/)   Division")
		fmt.Println("(=)   Display result")
		fmt.Println("(E)   Exit")
		fmt.Println("----------------------")
		fmt.Print(">>>")
		fmt.Scan(&operation)

		switch operation {
		case "+":
			if flagFirstOperation == false {
				fmt.Println()
				fmt.Print("Number 1: ")
				fmt.Scan(&num1)
				fmt.Print("Number 2: ")
				fmt.Scan(&num2)
				result = calculate.getAdition(num1, num2)
				fmt.Printf("%.2f + %.2f = %.2f", num1, num2, result)
				flagFirstOperation = true
			} else {
				fmt.Println()
				fmt.Print("Value: ")
				fmt.Scan(&num2)
				result = calculate.getAdition(result, num2)
				fmt.Printf("%.2f + %.2f = %.2f", result, num2, result)
			}
		case "-":
			if flagFirstOperation == false {
				fmt.Println()
				fmt.Print("Number 1: ")
				fmt.Scan(&num1)
				fmt.Print("Number 2: ")
				fmt.Scan(&num2)
				result = calculate.getSubtract(num1, num2)
				fmt.Printf("%.2f - %.2f = %.2f", num1, num2, result)
				flagFirstOperation = true
			} else {
				fmt.Println()
				fmt.Print("Value: ")
				fmt.Scan(&num2)
				result = calculate.getSubtract(result, num2)
				fmt.Printf("%.2f - %.2f = %.2f", result, num2, result)
			}
		case "x":
			if flagFirstOperation == false {
				fmt.Println()
				fmt.Print("Number 1: ")
				fmt.Scan(&num1)
				fmt.Print("Number 2: ")
				fmt.Scan(&num2)
				result = calculate.getMultiply(num1, num2)
				fmt.Printf("%.2f x %.2f = %.2f", num1, num2, result)
				flagFirstOperation = true
			} else {
				fmt.Println()
				fmt.Print("Value: ")
				fmt.Scan(&num2)
				result = calculate.getMultiply(result, num2)
				fmt.Printf("%.2f x %.2f = %.2f", result, num2, result)
			}
		case "/":
			if flagFirstOperation == false {
				fmt.Println()
				fmt.Print("Number 1: ")
				fmt.Scan(&num1)
				fmt.Print("Number 2: ")
				fmt.Scan(&num2)
				result, err := calculate.getDivision(num1, num2)
				if err == nil {
					fmt.Printf("%.2f / %.2f = %.2f", num1, num2, result)
				} else {
					switch {
					case errors.Is(err, errorDivideByZero):
						fmt.Println("divide by zero error")
					default:
						fmt.Printf("unexpected division error %s", err)
					}
				}
				flagFirstOperation = true
			} else {
				fmt.Println()
				fmt.Print("Value: ")
				fmt.Scan(&num2)
				result, err := calculate.getDivision(result, num2)
				if err == nil {
					fmt.Printf("%.2f / %.2f = %.2f", result, num2, result)
				} else {
					switch {
					case errors.Is(err, errorDivideByZero):
						fmt.Println("divide by zero error")
					default:
						fmt.Printf("unexpected division error %s", err)
					}
				}
			}
		case "=":
			fmt.Println(result)
		case "E":
			fallthrough
		case "e":
			flagBreak = true
		}
	}
}
