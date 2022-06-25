package main

import "fmt"

func main() {

	var inputNumber rune
	fmt.Print("Enter a number: ")
	fmt.Scan(&inputNumber)

	if inputNumber%2 == 0 {
		fmt.Println(inputNumber ,"is an even number")
	} else {
		fmt.Println(inputNumber ,"is an odd number")
	}
}
