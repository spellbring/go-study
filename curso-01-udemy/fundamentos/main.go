package main

import (
	"fmt"
	"runtime"
)

var x bool

const aa = 34
const bb = 42.32
const cc = "Jaime Reyes"

const (
	aaa = iota
	bbb = iota
	ccc = iota
)

func main() {

	fmt.Printf("Sistema operativo: %v\n", runtime.GOOS)
	fmt.Printf("Arquitectura: %v\n", runtime.GOARCH)

	fmt.Printf("Valores booleanos\n")
	varBooleans()
	fmt.Printf("#######Fin Valores booleanos#######\n")

	fmt.Printf("#######Valores enteros#######\n")
	varIntegers()
	fmt.Printf("#######Fin Valores enteros#######\n")

	fmt.Printf("#######cadenas de bytes#######\n")
	varStrings()
	fmt.Printf("#######Fin cadena de bytes#######\n")

	fmt.Printf("#######Constants#######\n")
	constants()
	fmt.Printf("#######fin Constants#######\n")

	fmt.Printf("#######IOTA homologo a los enum####### \n")
	iotaExample()
	fmt.Printf("#######fin IOTA#######\n")

	fmt.Printf("#######bitShifting####### \n")
	bitShifting()
	fmt.Printf("#######fin bitShifting#######\n")
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

func varStrings() {

	s1 := "Hola mundo"
	s2 := `Esta es una linea
		partida`

	fmt.Printf("El tipo de s1 es: %T y su valor es %v\n", s1, s1)
	fmt.Printf("El tipo de s2 es: %T y su valor es %v\n", s2, s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println("")

	//Valores en Hexa
	for i := 0; i < len(s1); i++ {
		fmt.Printf("%#U ", s1[i])
	}

	fmt.Println("")

	for i, v := range s1 {
		fmt.Printf("El indice %d el valor es %v\n", i, v)
	}

}

func constants() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
}

func iotaExample() {

	fmt.Println(aaa)
	fmt.Println(bbb)
	fmt.Println(ccc)

	fmt.Printf("%T", aaa)

}

const (
	_   = iota
	kbb = 1 << (iota * 10)
	gbb = 1 << (iota * 10)
	tbb = 1 << (iota * 10)
)

func bitShifting() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a)

	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b)

	kb := 1024
	gb := kb * kb
	tb := gb * kb

	fmt.Printf("%d\t\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t\t%b\n", gb, gb)
	fmt.Printf("%d\t\t\t\t%b\n", tb, tb)

	fmt.Printf("%d\t\t\t\t%b\n", kbb, kbb)
	fmt.Printf("%d\t\t\t\t%b\n", gbb, gbb)
	fmt.Printf("%d\t\t\t\t%b\n", tbb, tbb)
}
