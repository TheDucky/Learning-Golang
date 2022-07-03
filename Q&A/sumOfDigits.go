/**
 * write a program to input a number and find the sum of every didit.
 */

package main

import "fmt"

func main() {

	var inp, sum, lastDig int
	
	fmt.Print("Enter a number: ")
	fmt.Scan(&inp)

	for inp > 0 {
		lastDig = inp % 10
		sum += lastDig
		inp /= 10
	}

	fmt.Println("Sum of all digits =", sum)
}