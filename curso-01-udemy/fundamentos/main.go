package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {

	fmt.Printf("Sistema operativo: %v\n", runtime.GOOS)
	fmt.Printf("Arquitectura: %v\n", runtime.GOARCH)

	fmt.Printf("Valores booleanos\n")
	varBooleans()
	fmt.Printf("Fin Valores booleanos")

	fmt.Printf("Valores enteros\n")
	varIntegers()
	fmt.Printf("Fin Valores enteros")

}

func varBooleans() {

	fmt.Printf("El valor del boolean es: %v\n", x)
	x = true
	fmt.Printf("El valor del boolean con asignacion es: %v \n", x)

	a := 7
	b := 9

	fmt.Printf("El valor %v y el valor %v, corresponde a la igualdad:  %v\n", a, b, a == b)
}

func varIntegers() {

	x := 42
	y := 42.2
	fmt.Printf("El tipo de este valor es: %T y el valor es: %v\n", x, x)
	fmt.Printf("El tipo de este valor es: %T\n", y)

}
