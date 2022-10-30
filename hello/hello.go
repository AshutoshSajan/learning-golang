package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Greet("Hello world!")
	fmt.Println(message)
}
