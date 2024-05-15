package main

import (
	"os"
	"strconv"
)

func doop(s []string) {
	if len(s[0])>
	num1, _ := strconv.Atoi(s[0])
	operator := s[1]
	num2, _ := strconv.Atoi(s[2])
	result := 0

	if num1 == 0 || num2 == 0 && operator == "/" {
		os.Stdout.Write([]byte("No division by 0\n\n"))
		return
	}
	if num1 == 0 || num2 == 0 && operator == "%" {
		os.Stdout.Write([]byte("No modulo by 0\n\n"))
		return
	}
	switch operator {
	case "+":
		result = num1 + num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))
	case "-":
		result = num1 - num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "*":
		result = num1 * num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "/":
		result = num1 / num2
		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))

	case "%":
		result = num1 % num2

		os.Stdout.Write([]byte(Itoa(result)))
		os.Stdout.Write([]byte("\n"))
		
	default:
		return
	}
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	doop(os.Args[1:])
}

func Itoa(n int) string {
	//to be done manually
	return strconv.Itoa(n)
}