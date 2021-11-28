package main

import "fmt"

func main() {

	//Literal compuesto
	x := []int{2, 3, 4, 5, 6}

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	fmt.Println(cap(x)) // IMPRIME LA CAPACIDAD DEL ARREGLO SUBYACENTE

	for i, v := range x {
		fmt.Println(i, v)
	}

	//divicion de slices

	fmt.Println(x[1:4])
	//Imprime todos los slices
	fmt.Println(x[:])

	fmt.Println("###APPEND-EXAMPLE####")
	appendExample()
	fmt.Println("###MAKE-EXAMPLE####")
	makeExample()
	fmt.Println("###MULTIDIMENSIONAL####")
	multiDimensional()
	fmt.Println("###sliceSubyacente-EXAMPLE####")
	sliceSubyacente()
	fmt.Println("###exampleMap####")
	exampleMap()

	practicos()
}

func practicos() {
	//Practico 1
	x := [5]int{1, 2, 3, 4, 5} //Composite literal

	for i, v := range x {
		fmt.Printf("indice %v, valor %v\n", i, v)
	}

	fmt.Printf("%T", x)

	y := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Ejercicio 2

	for i, v := range y {
		fmt.Printf("indice: %v, valor: %v\n", i, v)
	}

	fmt.Printf("%T\n", y)

	//Ejercicio 3
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(z[:5])
	fmt.Println(z[5:])
	fmt.Println(z[2:7])
	fmt.Println(z[1:6])
	fmt.Println(z)

	//ejercicio 7

	her := []string{"Batman", "Jefe", "Vestido de negro"}
	her1 := []string{"Robin", "Ayudante", "Vestido de colores"}

	heroes := [][]string{her, her1}

	for i, reg := range heroes {
		fmt.Println("Registro:", i)
		for j, val := range reg {
			fmt.Printf("\tIndice: %v: valor:%v\n", j, val)

		}
	}

	fmt.Println(heroes)

	//Ejercicio 8 --> Mapas
	usuarios := map[string][]string{
		`eduar_tua`:    []string{`computadoras`, `montaña`, `playa`},
		`carlos_ramon`: []string{`leer`, `comprar`, `musica`},
		`juan_bimba`:   []string{`computadoras`, `montaña`, `playa`},
	}

	//Ejercicio 9, añadir elementos a un slice
	usuarios[`luis_peres`] = []string{"comer", "dormir", "caminar"}

	for key, value := range usuarios {
		fmt.Printf("Key: %v: \n", key)
		for i, v := range value {
			fmt.Printf("\tIndice :%v, valor:%v \n", i, v)
		}
	}

	//Ejercicio de borrar elementos.
	if v, ok := usuarios["juan_bimba"]; ok {
		fmt.Println("Se borro la llave con valor", v)
		delete(usuarios, "juan_bimba")
	}

	for key, value := range usuarios {
		fmt.Printf("Key: %v: \n", key)
		for i, v := range value {
			fmt.Printf("\tIndice :%v, valor:%v \n", i, v)
		}
	}

}

func appendExample() {
	x := []int{2, 3, 4, 5, 6}
	y := []int{333, 444, 555}

	x = append(x, 44, 55, 66)

	x = append(x, y...) //AGREGAR ELEMENTOS SLICE
	fmt.Println(x)

	x = append(x[:2], x[4:]...) // ELIMINAR ELEMENTOS DEL SLICE

	fmt.Println(x)
}

func makeExample() {

	x := make([]int, 10, 100)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 1
	x[5] = 100
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

func multiDimensional() {
	et := []string{"Jaime", "Reyes", "Ciclismo", "Baseball", "Montañismo"}
	fmt.Println(et)

	jl := []string{"Jacinto", "Lopez", "Correr", "Nadar", "Bailar"}

	vp := [][]string{et, jl}

	fmt.Println(vp)
}

func sliceSubyacente() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Println(x)

	y := append(x, 6, 7, 8) //Un nuevo arreglo es asignado
	fmt.Println(y)

	z := append(y[:2], y[3:]...)

	fmt.Println(z)

}

func exampleMap() {
	m := map[string]int{
		"batman": 32,
		"robin":  27,
	}

	fmt.Println(m)
	fmt.Println(m["batman"])
	fmt.Println(m["Eduar"])

	v, ok := m["Eduar"] //Idioma, Ok -> Asi se llama en Golang
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Eduar"]; ok {
		fmt.Println("Imprimiendo desde el if", v)
	}

	m["Jaime"] = 33

	for key, value := range m { //Iterar elementos de un mapa.
		fmt.Println(key, value)
	}

	fmt.Println(m)

	if v, ok := m["robin"]; ok {
		fmt.Println("Se borro la llave con valor", v)
		delete(m, "robin")
	}

}
