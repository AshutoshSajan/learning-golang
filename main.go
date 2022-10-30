package main

import "fmt"

func main() {
	greet("John")
}

func greet(name string) {
	fmt.Println("hello " + name)
}
