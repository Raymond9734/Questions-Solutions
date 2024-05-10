// splitwhitespaces
// Instructions

// Write a function that separates the words of a string and puts them in a string slice.

// The separators are spaces, tabs and newlines.
// Expected function

// func SplitWhiteSpaces(s string) []string {

// }

// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Printf("%#v\n", piscine.SplitWhiteSpaces("Hello how are you?"))
// }

// And its output :

// $ go run .
// []string{"Hello", "how", "are", "you?"}
// $

package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	wordSlice := []string{}
	var indx int
	for i, c := range s {
		if c == ' ' {
			if indx < i {
				wordSlice = append(wordSlice, s[indx:i])
			}
			indx = i + 1
		}
	}
	if indx < len(s) {
		wordSlice = append(wordSlice, s[indx:])
	}

	return wordSlice
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
