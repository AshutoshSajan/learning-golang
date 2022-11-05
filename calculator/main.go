package main

import (
	"fmt"
	"strconv"
)

func calculate(num1 int, num2 int, operand string) int {
	switch operand {
	case "+":
		return num1 + num2

	case "-":
		return num1 - num2

	case "*":
		return num1 * num2

	case "/":
		return num1 / num2

	default:
		fmt.Println("unknown operation")
		return 0
	}
}

func stringToNumber(str string) (n int, e error) {
	integer, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error parsing the number", err)
		return 0, err
	}

	return integer, nil
}

func getNumberFromInput(text string) int {
	var number string
	var err error

	fmt.Printf("> Enter the %v number:", text)
	_, err = fmt.Scanln(&number)

	num, err := stringToNumber(number)

	if err != nil {
		fmt.Println("Invalid number")
		return getNumberFromInput(text)
	}

	return num
}

func getOperand() string {
	var operand string
	var err error

	fmt.Print("> Enter the operand i.e + or - or * or /: ")
	_, err = fmt.Scanln(&operand)

	if operand != "+" && operand != "-" && operand != "/" && operand != "*" {
		fmt.Println("Operand must be +, -, * or /")
		return getOperand()
	}

	if err != nil {
		fmt.Println("Something went wrong", err)
		return getOperand()
	}

	return operand
}

func calculator() {
	num1 := getNumberFromInput("first")
	num2 := getNumberFromInput("second")
	operand := getOperand()

	result := calculate(num1, num2, operand)
	fmt.Printf("> %v %v %v is equals to %v\n", num1, operand, num2, result)
}

func main() {
	calculator()
}
