package main

import (
	"fmt"
)

func main() {
	var num int64
	i := 3 // Индекс бита, который вы хотите установить (нумерация с 0)

	// Установка i-го бита в 1
	num |= 1 << i
	fmt.Printf("После установки %d-го бита в 1: %d\n", i, num)

	// Установка i-го бита в 0
	num &^= 1 << i
	fmt.Printf("После установки %d-го бита в 0: %d\n", i, num)
}
