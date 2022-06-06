package main

import "fmt"

func main() {

	//GENERAL
	fmt.Printf("%v\n", "String Value")
	fmt.Printf("%T\n", 12.3)
	fmt.Printf("%T\n", true)

	//BOOLEAN
	fmt.Printf("%t\n", true)

	//INTEGERS
	var number rune = 97
	fmt.Printf("%b\n", number) //To binary
	fmt.Printf("%c\n", number) //To character (ASCII)
	fmt.Printf("%d\n", number) //To decimal
	fmt.Printf("%q\n", number) //Wraps single quotes around to converted ASCII

	//FLOAT
	var float float64 = 53.14524862452
	fmt.Printf("%f\n", float) //Print Floar Values with less precision
	fmt.Printf("%g\n", float) //Print Floar Values with more precision

	//STRING
	var str string = "I aM A StRIng"
	fmt.Printf("%s\n", str) //Prints the string
	fmt.Printf("%q\n", str) //Wraps string around double quotes

	// WP
	fmt.Printf("%15f\n", float) //Setting Width from right
	fmt.Printf("%.3f\n", float) //Setting precision of decimal point
	fmt.Printf("%08d\n", 12)    //Replaces empty padding with 0'

	// Sprintf() | SAVE FORMAT OUTPUT
	var savedOut string = fmt.Sprintf("%s %c %.0f %s %s %v%c%c %s\n", "Yoda is", 97, 900.213, "Jedi.", "He is only", 66, 99, 109, "tall")
	fmt.Print(savedOut)

}

/**
fmt.Printf()
		|
	print format
*/
