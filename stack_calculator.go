package main

import (
	"container/list"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"os/exec"
	"strconv"
)

type Stack struct {
	stack *list.List
	size  int
}

func (stack *Stack) push(element interface{}) {
	stack.size++
	stack.stack.PushFront(element)
}

func (stack *Stack) pop() (interface{}, error) {
	if stack.size > 0 {
		ele := stack.stack.Front()
		stack.size--
		return stack.stack.Remove(ele), nil
	}
	if stack.size == 0 {
		return nil, fmt.Errorf("Empty stack")
	}
	return nil, nil
}

func (stack *Stack) is_empty() bool {
	return stack.size == 0
}

func (stack *Stack) getSize() int {
	return stack.size
}

func (stack *Stack) add() float64 {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	return float64(x.(float64)) + float64(y.(float64))
}

func (stack *Stack) sub() float64 {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	return float64(x.(float64)) - float64(y.(float64))
}

func (stack *Stack) divide() (float64, error) {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	if y.(float64) == 0 {
		return 0, errors.New("Division by zero")
	}
	return float64(x.(float64)) / float64(y.(float64)), nil
}

func (stack *Stack) multiply() float64 {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	return float64(x.(float64)) * float64(y.(float64))
}

func (stack *Stack) mod() float64 {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	return math.Mod(float64(x.(float64)), float64(y.(float64)))
}

func (stack *Stack) exponent() float64 {
	x, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	y, err := stack.pop()
	if err != nil {
		log.Fatal(err)
	}
	return math.Pow(float64(x.(float64)), float64(y.(float64)))
}

func (stack *Stack) calculate(operator string) (float64, error) {
	switch operator {
	case "+":
		return stack.add(), nil
	case "*":
		return stack.multiply(), nil
	case "-":
		return stack.sub(), nil
	case "/":
		divide, err := stack.divide()
		if err != nil {
			return 0, err
		}
		return divide, nil
	case "%":
		return stack.mod(), nil
	case "^":
		return stack.exponent(), nil
	default:
		return 0, errors.New("Invalid operator")
	}
}

func (stack *Stack) printStack() {
	for e := stack.stack.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
}

func (stack *Stack) user_input() {
	var input string

	fmt.Printf("> ")
	fmt.Scan(&input)

	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		stack.push(input)
		n, _ := stack.pop()
		new_num, err := stack.calculate(n.(string))
		if err != nil {
			log.Fatal(err)
		}
		stack.push(new_num)
	} else {
		stack.push(number)
	}
}

func main() {
	log.SetPrefix("Error: ")
	log.SetFlags(0)
	stack := &Stack{
		stack: list.New(),
	}

	for stack.getSize() >= 0 {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
		fmt.Printf("\t==========================\n")
		fmt.Printf("\t   GoStack Calculator! \n")
		fmt.Printf("\t==========================\n\n")
		fmt.Print("Stack: [")
		stack.printStack()
		fmt.Print("]\n")
		stack.user_input()
	}
	fmt.Print("Stack: [")
	stack.printStack()
	fmt.Print("]\n")
}
