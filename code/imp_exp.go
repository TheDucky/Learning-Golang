package main

import (
	"fmt"
)

func main() {

	// int declaration
	var number uint8 = 123 //This is explicit
	number2 := 123 //This is implicit
	
	fmt.Printf("%T\n", number) //uint8
	fmt.Printf("%T\n", number2) //int

	// string declaration
	var name string = "Ducky" //This is explicit
	name2 := "Daksh" //This is implicit

	fmt.Printf("%T\n", name) //string
	fmt.Printf("%T\n", name2) //string
}

/**
Explicit declaration is better
Implicit declaration lets the compiler guess which datatype to use. This might slow down the compile time.
*/
