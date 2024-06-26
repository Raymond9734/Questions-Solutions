// printbits
// Instructions

// Write a program that takes a number as argument, and prints it in binary value without a newline at the end.

//     If the the argument is not a number, the program should print 00000000.

// Usage :

// $ go run . 1
// 00000001$
// $ go run . 192
// 11000000$
// $ go run . a
// 00000000$
// $ go run . 1 1
// $ go run .
// $

package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil {
		for _, c := range "00000000" {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
		return
	}

	// Print the binary representation
	for i := 7; i >= 0; i-- {
		bit := (num >> i) & 1 // check comments below
		for _, v := range strconv.Itoa(bit) {
			z01.PrintRune(v)
		}
	}
	z01.PrintRune('\n')
}

// Let’s break down the line bit := (num >> i) & 1 from your Go program:

// num >> i: This expression performs a right shift operation on the value of num by i bits. The >> operator shifts the bits of num to the right by the specified number of positions. For example, if num is 5 (binary 101) and i is 2, the result of num >> i would be 1 (binary 001).

// & 1: This part of the expression performs a bitwise AND operation between the result of the right shift (num >> i) and the value 1. The & operator compares each corresponding bit of the two operands and sets the corresponding bit in the result to 1 if both bits are 1, otherwise it sets it to 0.
// If the rightmost bit of num >> i is 1, the result of num >> i & 1 will be 1.
// If the rightmost bit of num >> i is 0, the result of num >> i & 1 will be 0.
