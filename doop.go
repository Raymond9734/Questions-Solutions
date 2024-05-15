package main

import (
	"os"
	"strconv"
)

//to add Atoi manual
func doop(s []string) {
	if len(s[0]) > 12 {
		return
	}
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
	//declare variables
	r := []byte{}
	temps := []byte{}
	right := 0

	//handle -ve numbers
	if n < 0 {
		r = append(r, '-')
		n *= -1
	}

	//convert each digit to a string by adding  "0"
	for {
		right = n % 10
		n = n / 10
		temps = append(temps, byte('0'+right))
		if n == 0 {
			break
		}

	}
	for i := len(temps) - 1; i >= 0; i-- {
		r = append(r, temps[i])
	}
	return string(r)
}
