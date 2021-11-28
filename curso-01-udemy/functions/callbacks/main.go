package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(suma(ii...))

	s1 := sumaPares(suma, ii...)

	s2 := sumaImpares(suma, ii...)

	fmt.Println("el total del pares es : ", s1)
	fmt.Println("Suma impares : ", s2)
}

func sumaPares(f func(xi ...int) int, vi ...int) int {
	var y []int
	for _, v := range vi {
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	return f(y...)
}

func sumaImpares(f func(x1 ...int) int, vi ...int) int {
	var y []int

	for _, v := range vi {
		if v%2 != 0 {
			y = append(y, v)
		}
	}
	return f(y...)
}

func suma(xi ...int) int {

	fmt.Printf("%T\n", xi)

	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
