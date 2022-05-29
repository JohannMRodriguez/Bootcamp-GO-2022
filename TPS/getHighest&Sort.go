package main

import "fmt"

func main() {

	lista := [5]int{4, 10, 6, 7, 8}
	mayor := [3]int{lista[0], lista[0], lista[0]}

	for times := 0; times < 3; times++ {
		for _, each := range lista {
			if times == 0 {
				if each > mayor[0] {
					mayor[0] = each
				}
			} else if times == 1 {
				if each > mayor[1] && each < mayor[0] {
					mayor[1] = each
				}
			} else {
				if each < mayor[0] && each < mayor[1] && each > mayor[2] {
					mayor[2] = each
				}
			}
		}
	}
	fmt.Println(mayor)
}
