package main

import "fmt"

func main() {
	fmt.Println("hello world!")
	greet("Jhon")
}

func greet(name) {
	fmt.Println("hello " + name)
}
