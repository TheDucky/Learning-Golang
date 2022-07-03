/**
 * write a program to input 3 numbers and find the largest out of them.
 */

package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Println("Enter 3 numbers")

	fmt.Scan(&a)	
	fmt.Scan(&b)
	fmt.Scan(&c)
	
	if a > b && a > c {
		fmt.Println(a, "is the largest number")
	}else if b > a && b > c {
		fmt.Println(b, "is the largest number")
	}else {
		fmt.Println(c, "is the largest number")
	}
}