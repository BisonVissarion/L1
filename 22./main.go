package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Создаем переменные a и b с использованием big.Int для работы с большими числами
	a := new(big.Int)
	b := new(big.Int)

	// Устанавливаем значения a и b.
	a.SetString("2000000", 10) // Установите значение a как нечто большее чем 2^20 (1 048 576)
	b.SetString("3000000", 10) // Установите значение b как нечто большее чем 2^20 (1 048 576)

	// Умножение a на b
	multiplicationResult := new(big.Int).Mul(a, b)

	// Деление a на b
	divisionResult := new(big.Int).Div(a, b)

	// Сложение a и b
	additionResult := new(big.Int).Add(a, b)

	// Вычитание b из a
	subtractionResult := new(big.Int).Sub(a, b)

	// Вывод результатов
	fmt.Printf("Умножение: %s\n", multiplicationResult.String())
	fmt.Printf("Деление: %s\n", divisionResult.String())
	fmt.Printf("Сложение: %s\n", additionResult.String())
	fmt.Printf("Вычитание: %s\n", subtractionResult.String())
}
