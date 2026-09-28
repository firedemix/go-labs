package main

import "fmt"

func main() {
	var intVar int = 42
	var floatVar float64 = 3.14
	var stringVar string = "Привет, Go!"
	var boolVar bool = true

	fmt.Printf("int: %d\n", intVar)
	fmt.Printf("float64: %.2f\n", floatVar)
	fmt.Printf("string: %s\n", stringVar)
	fmt.Printf("bool: %t\n", boolVar)
}