package main

import "fmt"

var (
	firstName string = "Jhon"
	lastName  string = "Doe"
	age       int    = 32
)

func main() {

	firstName, lastName := "Jhon", "Doe"
	age := 32

	fmt.Println(firstName, lastName, age)

}
