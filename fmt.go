package main

import "fmt"

func main() {
	firstName := "Randi"
	lastName := "Febriadi"

	fmt.Println("Hello '", firstName, lastName, "'")
	fmt.Printf("Hello '%s %s'", firstName, lastName)
}
