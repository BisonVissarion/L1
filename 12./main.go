package main

import (
	"fmt"
)

func main() {
	// Последовательность строк
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаем мапу для представления множества (используем bool в качестве значения)
	set := make(map[string]bool)

	// Добавляем строки в множество
	for _, str := range strings {
		set[str] = true
	}

	// Выводим уникальные строки из множества
	fmt.Println("Уникальные строки:")
	for str := range set {
		fmt.Println(str)
	}
}
