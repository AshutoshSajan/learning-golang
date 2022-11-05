package calculate

import (
	"fmt"
	"strconv"
)

func Calculate(num1 int, num2 int, operand string) int {
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

func StringToNumber(str string) (n int, e error) {
	integer, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Error parsing the number", err)
		return 0, err
	}

	return integer, nil
}

func GetNumberFromInput(text string) int {
	var number string
	var err error

	fmt.Printf("> Enter the %v number:", text)
	_, _ = fmt.Scanln(&number)

	num, err := StringToNumber(number)

	if err != nil {
		fmt.Println("Invalid number")
		return GetNumberFromInput(text)
	}

	return num
}

func GetOperand() string {
	var operand string
	var err error

	fmt.Print("> Enter the operand i.e + or - or * or /: ")
	_, err = fmt.Scanln(&operand)

	if operand != "+" && operand != "-" && operand != "/" && operand != "*" {
		fmt.Println("Operand must be +, -, * or /")
		return GetOperand()
	}

	if err != nil {
		fmt.Println("Something went wrong", err)
		return GetOperand()
	}

	return operand
}

func Calculator() {
	num1 := GetNumberFromInput("first")
	num2 := GetNumberFromInput("second")
	operand := GetOperand()

	result := Calculate(num1, num2, operand)
	fmt.Printf("> %v %v %v is equals to %v\n", num1, operand, num2, result)
}
