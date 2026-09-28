package main

import (
	"fmt"
	"unicode/utf8"
)

func strLength(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	var s string
	fmt.Print("Введите строку: ")
	fmt.Scan(&s)
	fmt.Printf("Длина строки: %d\n", strLength(s))
}