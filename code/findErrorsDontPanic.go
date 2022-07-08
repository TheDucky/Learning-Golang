package main

import "fmt"

func main() {

	var divisor float32
	fmt.Printf("Divide 100 by a number: ")
	fmt.Scan(&divisor)

	if divisor == 0 {
		panic("YOU CANT DIVIDE A NUMBER BY ZERO")
	}

	fmt.Printf("100 / %d = %.3f\n", divisor, 100/divisor)
}
