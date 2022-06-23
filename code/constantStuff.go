package main

import "fmt"

func main() {

	const PI float32 = 3.14 //Constants are Unchangeable and Read-only.
	// PI = 12.2 || This will throw an error as PI is a constant ^^
	fmt.Printf("Value of IP = %.2f\n", PI)
}