/**
 * write a program to take in an int value and print out the reverse of the int
 */

package main

import "fmt"

func main() {

	var input, reverse, oneDig int

	fmt.Print("Enter a number: ")
	fmt.Scan(&input)

	for input != 0 {
		oneDig = input % 10
		reverse = reverse * 10 + oneDig
		input /= 10
	}

	fmt.Println("Reverse number =", reverse)
}
