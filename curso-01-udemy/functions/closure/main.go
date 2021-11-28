package main

import (
	"fmt"
)

var x int

func main() {

	fmt.Println(x)

	x++

	fmt.Println(x)

	{

		y := 42
		fmt.Println(y)

	}

	foo()

	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func foo() {
	x++
}

func incrementor() func() int {
	var x int

	return func() int {
		x++
		return x
	}
}
