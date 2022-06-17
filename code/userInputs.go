package main

import (
	"fmt"
)

func main() {

	var name string
	var bday uint16
	fmt.Print("What is your name? ")
	fmt.Scan(&name)

	fmt.Print("When were you born ", name+"? ")
	fmt.Scan(&bday)

	fmt.Printf("You are %d years old\n", 2022-bday)

}
