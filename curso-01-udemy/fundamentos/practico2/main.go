package main

import (
	"fmt"
)

const (
	a     = 42
	b int = 43 //constante con tipo
)

const (
	aa = 2020 + iota
	bb = 2020 + iota
	cc = 2020 + iota
	dd = 2020 + iota
)

func main() {

	fmt.Println("Imprimir valores en binario, hexa y decimal ")

	fmt.Printf("%d\t", 42)
	fmt.Printf("%x\t", 42)
	fmt.Printf("%b\t", 42)

	fmt.Println("Operadores logicos")

	fmt.Printf("%v\t", 1 > 2)
	fmt.Printf("%v\t", 2 == 2)
	fmt.Printf("%v\t\n", 2 != 2)

	fmt.Printf("%v\t%v\n", a, b)

	fmt.Println("string literal")
	fmt.Printf(`asd
	asdasd
	
	asdasd
	`)

	fmt.Printf("%v\t%v\t%v\t%v\t", aa, bb, cc, dd)
}
