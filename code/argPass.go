package main

import (
	"fmt"
	"os"
)

func main() {

	var passedArgument string = os.Args[1]
	fmt.Printf("%q was the argument passed!\n", passedArgument)
}