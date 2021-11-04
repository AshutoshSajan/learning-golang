package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	message := greetings.Hello("Hello world!")
	fmt.Println(message)
}