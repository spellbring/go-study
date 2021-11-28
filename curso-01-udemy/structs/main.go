package main

import "fmt"

type persona struct { //composite literal
	first string
	last  string
	age   int
}

type agenteSecreto struct {
	persona
	lpm bool
}

func main() {

	as1 := agenteSecreto{
		persona: persona{
			first: "James",
			last:  "bond",
			age:   32,
		},
		lpm: true,
	}

	fmt.Println(as1)

	fmt.Println(as1.first, as1.last, as1.lpm, as1.age)

	//Structs Anonimos
	s1 := struct {
		name     string
		lastName string
	}{
		name:     "Jaime",
		lastName: "Reyes",
	}

	fmt.Println(s1)

	practice1()
}

func practice1() {

	type persona struct {
		nombre     string
		apellido   string
		saboresFav []string
	}

	p := persona{
		nombre:     "Jaime",
		apellido:   "Reyes",
		saboresFav: []string{"chocolate", "mantecado", "ron con pasas"},
	}
	p1 := persona{
		nombre:     "Valeria",
		apellido:   "Bastias",
		saboresFav: []string{"fresa", "vainilla", "torta suiza"},
	}

	fmt.Println(p.nombre)
	fmt.Println(p1.nombre)

}
