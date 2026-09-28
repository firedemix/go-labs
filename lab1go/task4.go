package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Введите первое число: ")
	fmt.Scan(&x)
	fmt.Print("Введите второе число: ")
	fmt.Scan(&y)

	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	fmt.Printf("%d - %d = %d\n", x, y, x-y)
	fmt.Printf("%d * %d = %d\n", x, y, x*y)
	if y != 0 {
		fmt.Printf("%d / %d = %d\n", x, y, x/y)
		fmt.Printf("%d %% %d = %d\n", x, y, x%y)
	} else {
		fmt.Println("Деление на ноль невозможно")
	}
}