package main

import (
	"fmt"
)

func main() {

	foo()
	bar("jaime")
	fmt.Println(woo("Valeria"))

	x, y := saludar("Jaime", "Reyes")

	fmt.Println(x)
	fmt.Println(y)

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}

	suma := sum("hola", xi...)

	fmt.Println(suma)

	methodExample()

}

func foo() {
	fmt.Println("hola desde foo")
}

func bar(s string) {
	fmt.Printf("Hola, %v", s)
}

func woo(st string) string {
	return fmt.Sprint("Hola desde woo ", st)
}

func saludar(n string, s string) (string, bool) {

	p := fmt.Sprint(n, " ", s, ` dice "hola".`)
	q := true

	return p, q

}

//esta funcion recibe un string y un arreglo de enteros
func sum(s string, x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	suma := 0

	for i, v := range x {
		suma += v
		fmt.Println("El valor en el indice", i, "suma", v, "al total", suma)
	}

	fmt.Println("El total es : ", suma)

	return suma

}

//metodos
//func (r receptor) identificador( parametros ) (retorno(s)){ codigo }

type persona struct {
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona
	lpm bool
}

func (a agenteSecreto) presentar() {
	fmt.Println("Hola, soy", a.nombre, a.apellido, "el agente secreto se presenta")
}

func (p persona) presentar() {
	fmt.Println("Hola, soy", p.nombre, p.apellido, "la persona se presenta")
}

//Collection de metodos, nos permite definir comportamientos
type humano interface {
	presentar()
}

func callHuman(h humano) {

	//AFIRMACIONES.
	switch h.(type) {
	case persona:
		fmt.Println("(persona) ", h.(persona).nombre)
	case agenteSecreto:
		fmt.Println("(agenteSecreto) ", h.(persona).nombre)
	}

	fmt.Println("Fui pasado a la funcion callHuman", h)
}

type manzana int

func methodExample() {
	as1 := agenteSecreto{
		persona: persona{
			nombre:   "Jaime",
			apellido: "Reyes",
		},
		lpm: true,
	}
	as2 := agenteSecreto{
		persona: persona{
			nombre:   "Valeria",
			apellido: "Bastias",
		},
		lpm: true,
	}

	p := persona{
		nombre:   "Carlos",
		apellido: "Pinto",
	}

	fmt.Println(as1)
	as1.presentar()
	as2.presentar()

	callHuman(p)

	//Conversion
	var x manzana = 52
	fmt.Println(x)
	fmt.Printf("%T", x)

	var y int
	y = int(x)

	fmt.Println(x)
	fmt.Printf("%T", y)

}
