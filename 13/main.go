package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 10

	fmt.Printf("До обмена: a = %d, b = %d\n", a, b)

	// Меняем местами значения a и b без временной переменной
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)
}
