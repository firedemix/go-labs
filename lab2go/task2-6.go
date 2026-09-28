package main

import "fmt"

func average(a, b int) float64 {
	return float64(a+b) / 2
}

func main() {
	var a, b int
	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&a, &b)
	fmt.Printf("Среднее значение: %.2f\n", average(a, b))
}