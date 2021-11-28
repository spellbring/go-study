package main

import (
	"fmt"
)

func main() {
	//el keyword defer, pospone la ejecución del bloque de codigo al final
	defer foo()
	bar()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
