/**
 * write a program to take in an int value and print out its tables from 1 to 10.
 */

package main

import "fmt"

func main() {

	var tInt int
	index := 1

	fmt.Print("Enter a number: ")
	fmt.Scan(&tInt)

	for index <= 10 {
		fmt.Printf("%d x %d = %d\n", tInt, index, (tInt*index))
		index++
	}
}