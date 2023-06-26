/*
 * Simple app calculator written in Golang
 * Author: Chris Dedman
 */
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
)

// Get user input for the numbers to be calculated
// Return the user number
func get_user_input() float64 {
	var input string

	fmt.Printf("Enter a number: ")
	fmt.Scan(&input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Printf("Invalid input. Please enter a valid number.\n")
		return get_user_input()
	}

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
func calculator(x float64, operant string, y float64) (float64, error) {
	switch operant {
	case "+":
		return x + y, nil
	case "*":
		return x * y, nil
	case "-":
		return x - y, nil
	case "/":
		if y == 0 {
			return y, errors.New("'Division by zero'")
		}
		return x / y, nil
	case "%":
		return math.Mod(x, y), nil
	case "^":
		return math.Pow(x, y), nil
	default:
		return x, errors.New("Invalid operator")
	}
}

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)

	fmt.Printf("\t==========================\n")
	fmt.Printf("\t Welcome to GoCalculator! \n")
	fmt.Printf("\t==========================\n\n")

	first_number := get_user_input()
	operant := get_operator()
	second_number := get_user_input()

	result, err := calculator(first_number, operant, second_number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nResult: %v %v %v = %v\n", first_number, operant, second_number, result)
}
