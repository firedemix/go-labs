package main

import "fmt"

func sumAndDiff(a, b float64) (float64, float64) {
	return a + b, a - b
}

func main() {
	var a, b float64
	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	s, d := sumAndDiff(a, b)
	fmt.Printf("Сумма: %.2f\n", s)
	fmt.Printf("Разность: %.2f\n", d)
}