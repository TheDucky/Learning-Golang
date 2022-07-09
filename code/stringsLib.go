package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "Golang is an amazing language"
	sentence2 := "Golang "
	sentence3 := []string{"Golang", "is", "POG"}

	substring := sentence[0:6]                          // Golang
	uppercase := strings.ToUpper(substring)             //GOLANG
	lowercase := strings.ToLower(substring)             //golang
	replaced := strings.Replace(sentence2, "G", "N", 1) // Nolang

	var sentenceLength int = len(sentence) // 29

	splitit := strings.Split(sentence, "a") // [Gol ng is  n  m zing l ngu ge]
	joinit := strings.Join(sentence3, "-")  // Golang-is-POG

	doesItContain := strings.Contains(sentence, sentence2) // true
	compareIt := strings.Compare(sentence, sentence2)      // 1
	compareIt2 := strings.Compare(sentence2, sentence)     // -1

	removeSpaces := strings.Trim(sentence2, "Go")    // lang
	getIndex := strings.Index(sentence, "a")         // 3
	getLastIndex := strings.LastIndex(sentence, "a") // 26

	fmt.Println(substring)
	fmt.Println(uppercase)
	fmt.Println(lowercase)
	fmt.Println(replaced)
	fmt.Println(sentenceLength)
	fmt.Println(splitit)
	fmt.Println(joinit)
	fmt.Println(doesItContain)
	fmt.Println(compareIt)
	fmt.Println(compareIt2)
	fmt.Println(removeSpaces)
	fmt.Println(getIndex)
	fmt.Println(getLastIndex)
}
