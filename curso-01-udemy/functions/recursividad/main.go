package main

import "fmt"

func main() {

	f := factorial(4)

	fmt.Println(f)
	fmt.Println(cicloFact(4))

}

func factorial(n int) int {

	//caso base
	if n == 0 {
		return 1

	}
	return n * factorial(n-1)

}

func cicloFact(n int) int {

	cont := 1

	for i := n; i > 0; i-- {

		cont *= i

	}
	return cont

}
