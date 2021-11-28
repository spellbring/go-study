package main

import (
	"fmt"
)

func main() {
	fmt.Println("HELLO")

	//FUNCION ANONIMA
	func() {
		fmt.Println("HELLO FROM anonymous")
	}()

	func(x int) {
		fmt.Println("HELLO FROM anonymous", x)
	}(42)
}

func foo() {
	fmt.Println("HELLO FROM FOO")
}
