package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func stringToNumber(str string) int {
	integer, err := strconv.Atoi(str)

	handleError(err)

	return integer
}

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

func calculator() {
	var num1 string
	var num2 string
	var operand string
	var err error

	fmt.Print("> Enter the first number: ")

	_, err = fmt.Scanln(&num1)
	handleError(err)
	n1 := stringToNumber(num1)

	fmt.Print("> Enter the second number: ")
	_, err = fmt.Scanln(&num2)
	handleError(err)
	n2 := stringToNumber(num2)

	fmt.Print("> Enter the operand i.e + or - or * or /: ")
	_, err = fmt.Scanln(&operand)
	handleError(err)

	result := calculate(n1, n2, operand)
	fmt.Printf("> %v %v %v is equals to %v\n", n1, operand, n2, result)
}

func sayHi() {
	var name string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your name? ")

	name, err := reader.ReadString('\n')

	handleError(err)

	fmt.Println("Hi", name)
}

func main() {
	sayHi()
	calculator()
}
