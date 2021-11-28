package main

import (
	"fmt"
)

func main() {
	f := func() {
		fmt.Println("mi primera expresion funcion")
	}

	f()

	g := func(x int) {
		fmt.Println("Mi primera expresion funcion", x)
	}

	g(1)
}
