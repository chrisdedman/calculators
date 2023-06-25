/*
 * Simple app calculator written in Golang
 * Author: Chris Dedman
 */
package main

import (
	"fmt"
	"math"
)

// Get user input for the numbers to be calculated
// Return the user number
func get_user_input() float64 {
	var number float64

	fmt.Printf("Enter a number: ")
	fmt.Scan(&number)

	return number
}

// Get the user input operator for our calculation
// Return the operator sign
func get_operator() string {
	var operator string
	fmt.Print("Enter an operator sign [+ - / * % ^]: ")
	fmt.Scan(&operator)

	return operator
}

// Calculate x and y according to the user choice operant
// Return the value calculated
func calculator(x float64, operant string, y float64) float64 {
	switch operant {
	case "+":
		return x + y
	case "*":
		return x * y
	case "-":
		return x - y
	case "/":
		return x / y
	case "%":
		return math.Mod(x, y)
	case "^":
		return math.Pow(x, y)
	default:
		return 404
	}
}

func main() {
	fmt.Printf("\t==========================\n")
	fmt.Printf("\t Welcome to GoCalculator! \n")
	fmt.Printf("\t==========================\n\n")

	first_number := get_user_input()
	operant := get_operator()
	second_number := get_user_input()

	result := calculator(first_number, operant, second_number)
	if result == 404 {
		fmt.Printf("\nError - Invalide Operant!\n\t Operant entered %v\n", operant)
	} else {
		fmt.Printf("%v %v %v = %v\n", first_number, operant, second_number, result)
	}
}
