package main

import (
	"fmt"
)

// GO BY EXAMPLE, se pueden sacar ejemplos https://gobyexample.com

func main() {

	fmt.Println("#####nested loop########")
	nestedLoop()

	fmt.Println("####### declaration for")
	declarationFor()

	fmt.Println("####### break example")
	declarationFor()

	fmt.Println("#########Break and continue")
	breakAndContinue()

	fmt.Println("#########print ASCII with for")
	printAscii()

	fmt.Println("#########if clause")
	ifClause()

}

func breakAndContinue() {
	x := 1
	for {

		x++
		if x > 100 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)

	}
}

func declarationFor() {
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func breakExample() {
	i := 0
	for {
		if i > 9 {
			break
		}
		fmt.Println(i)
		i++
	}

}

func nestedLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Printf("Ciclo externo %d\n", i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t\tel ciclo interno %d\n", j)
		}

	}
}

func printAscii() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%#U\n", i, i, i)
	}
}

func ifClause() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("003")
	}
	if !false {
		fmt.Println("004")
	}

	if x := 42; x == 42 {
		fmt.Println(001)
	}

}
