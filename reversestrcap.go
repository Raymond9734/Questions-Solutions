// reversestrcap
// Instructions
// Write a program that takes one or more arguments and that, for each argument, puts the last letter of each word in uppercase and the rest in lowercase. It displays the result followed by a newline ('\n').

// If there are no argument, the program displays nothing.

// Usage
// $ go run . "First SMALL TesT" | cat -e
// firsT smalL tesT$
// $ go run . "SEconD Test IS a LItTLE EasIEr" "bEwaRe IT'S NoT HARd WhEN " " Go a dernier 0123456789 for the road e" | cat -e
// seconD tesT iS A littlE easieR$
// bewarE it'S noT harD wheN $
//  gO A dernieR 0123456789 foR thE roaD E$
// $ go run .
// $

package main

import (
	"fmt"
	"os"

	// "github.com/01-edu/z01"
)

// func main() {
// 	if len(os.Args) <= 1 {
// 		return
// 	}
// 	words := ""

// 	args := os.Args[1]
// 	letters := []rune(args)

// 	for i, c := range letters {
// 		if c == ' ' && i != 0 && c >= 'a' && c <= 'z' {
// 			letters[i-1] = c - 32
// 			words += string(letters[i-1])
// 		} else if i == len(letters)-1 && c >= 'a' && c <= 'z' {
// 			c = c - 32
// 			words += string(c)
// 		} else if c >= 'A' && c <= 'Z' && i != len(letters)-1 {
// 			c = c + 32
// 			words += string(c)
// 		} else {
// 			words += string(c)
// 		}
// 	}
// 	// words = strings.Join(letters, " ")
// 	fmt.Println(words)

// }

func main() {
	wordSlice := []string{}

	if len(os.Args) != 2 {
		return
	}

	args := os.Args[1]
	var indx int

	for i, c := range args {
		if c == ' ' && indx < i{
			wordSlice = append(wordSlice, args[indx: i])
		} 
		indx = i + 1
	}

	if indx < len(args) {
		wordSlice = append(wordSlice, args[indx:])
	}
		// else if c == ' ' && indx != 0 {
		// 	wordSlice = append(wordSlice, args[indx: i])
		// 	indx = i + 1
		// } else {
		// 	wordSlice = append(wordSlice, args[indx: ])
		// }
	fmt.Println(wordSlice)
}
