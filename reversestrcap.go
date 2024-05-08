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
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) <= 1 {
		return
	}
	words := ""

	args := os.Args[1:]

	for _, word := range args {
		for i, c := range []rune(word) {
			if c == ' ' && word[i-1] >= 'a' && word[i-1] <= 'z' {
				ch := rune(word[i-1]) - 32
				words += string(ch)
			} else if c >= 'A' && c <= 'Z' {
				ch := c + 32
				words += string(ch)
			} else {
				words += string(c)
			}
		}
	}
	for _, c := range words {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
