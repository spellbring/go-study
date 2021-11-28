package main

import (
	"fmt"
	"math"
)

//practico 4
type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {
	fmt.Println("Mi nombre es", p.nombre, "y mi edad es", p.edad)
}

//practico 5

type cuadrado struct {
	longitud float64
}

type circulo struct {
	radio float64
}

type forma interface {
	calcularArea() float64
}

func (c cuadrado) calcularArea() float64 {
	return c.longitud * c.longitud
}

func (c circulo) calcularArea() float64 {
	return math.Pi * c.radio * c.radio
}

func info(f forma) {

	switch f.(type) {
	case cuadrado:
		fmt.Println("Longitud de cuadrado: ", f.(cuadrado).longitud)
	case circulo:
		fmt.Println("Longitud de circulo: ", f.(circulo).radio)
	}

	fmt.Println(f.calcularArea())

}

func main() {

	foo()
	bar()

	xBar1 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(foo1(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(bar1(xBar1))

	fmt.Println(deferExample())

	p := persona{
		nombre:   "JAIME",
		apellido: "Reyes",
		edad:     33,
	}

	p.presentar()

	//practico 5
	circ := circulo{
		radio: 12.3,
	}

	cua := cuadrado{
		longitud: 15,
	}

	info(circ)
	info(cua)

}

//Practico 1
func foo() int {
	fmt.Println("Adios")
	return 42

}

func bar() (int, string) {
	fmt.Println("Hola")
	return 1, "hola"
}

//practico 2
func foo1(x ...int) int {
	sum := 0

	for _, v := range x {

		sum += v

	}

	return sum
}

func bar1(x []int) int {
	sum := 0

	for _, v := range x {

		sum += v

	}

	return sum
}

//practico 3

func deferExample() int {

	defer bar()

	return foo()

}
