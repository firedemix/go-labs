package main

import "fmt"

func checkSign(n int) string {
	switch {
	case n > 0:
		return "Positive"
	case n < 0:
		return "Negative"
	default:
		return "Zero"
	}
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Println(checkSign(n))
}