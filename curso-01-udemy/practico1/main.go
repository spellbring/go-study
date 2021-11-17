package main

import (
	"fmt"
)

var a int
var b string
var c bool

type numero int

var w numero

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(a, b, c)

	a := 33
	b := "Jaime"
	c := false

	s := fmt.Sprintf("Agente %v\t, Edad: %v, Licencia para matar: %v", b, a, c)

	fmt.Println(s)

	x = int(w)

	fmt.Printf("%T", x)

}
